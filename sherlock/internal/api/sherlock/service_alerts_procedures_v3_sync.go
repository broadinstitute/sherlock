package sherlock

import (
	"encoding/json"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceAlertV3SyncRequest struct {
	OnEnvironment *string `json:"onEnvironment,omitempty" form:"onEnvironment"`
}

type ServiceAlertV3SyncResponse struct {
	AddedServiceAlerts   []string `json:"addedServiceAlerts"`
	RemovedServiceAlerts []string `json:"removedServiceAlerts"`
	UpdatedServiceAlerts []string `json:"updatedServiceAlerts"`
}

type ServiceAlertJsonData struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	Link       string `json:"link"`
	Severity   string `json:"severity"`
	IncidentID string `json:"incident_id"`
}

// serviceAlertsProceduresV3Sync godoc
//
//	@summary		Sync service alerts
//	@description	Method to get all currently active service alerts from Sherlock's DB and ensure that the service alert json files placed in Google Buckets for Terra match.
//	@tags			ServiceAlert
//	@accept			json
//	@produce		json
//	@param			environment					body		ServiceAlertV3SyncRequest	true	"Information on Service Alert environment"
//	@success		200						{array}		ServiceAlertV3SyncResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/procedures/v3/sync [post]
func syncServiceAlerts(ctx *gin.Context) {
	// set db con
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ServiceAlertV3SyncRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}
	// Get alerts from DB and gcs bucket from environment
	alerts, gcs_bucket := getAlerts(ctx, body, db)
	if len(alerts) == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) No Alerts found for this environment", errors.BadRequest))
		return
	}
	// set GCS client
	var gcsClient, gc_err = google_bucket.InitializeStorageClient(ctx)
	if gc_err != nil {
		return
	}
	// get alerts blob
	alert_json_blob, read_err := gcsClient.GetBlob(ctx, *gcs_bucket, "alerts.json")
	if read_err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("blob not found: %v", read_err))
		return
	}
	// read data from blob & parse to JSON
	json_data := getFileJson(ctx, gcsClient, alert_json_blob)
	// diff alerts from db and json data for currently live service alerts
	add, remove, update := compareAlerts(ctx, json_data, alerts)
	// if any changes then re-create file and upload to gcs bucket
	if add != nil || remove != nil || update != nil {
		// convert model data for service alerts to json formatted bytes to upload
		json_bytes, err := createSvcAlertJsonData(ctx, alerts)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("issue encountered creating json payload: %v", err))
			return
		}
		// Upload file to bucket w/ latest info
		if err = gcsClient.WriteBlob(ctx, *gcs_bucket, "alerts.json", json_bytes); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error writing updated alerts.json file: %v", err))
			return
		}
	}
}

func getAlerts(ctx *gin.Context, request ServiceAlertV3SyncRequest, db *gorm.DB) ([]models.ServiceAlert, *string) {
	var env_result models.Environment
	if request.OnEnvironment != nil {
		// match env so that we can get gcs bucket
		environmentQuery, err := environmentModelFromSelector(*request.OnEnvironment)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error parsing environment selector '%s': %w", *request.OnEnvironment, err))
			return nil, nil
		}
		if err = db.Where(&environmentQuery).Select("id").First(&env_result).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error fetching environment '%v'", err))
			return nil, nil
		}
	}
	var activeAlerts []models.ServiceAlert
	// Only return service alerts that haven't been deleted for this environment
	if err := db.Model(&models.ServiceAlert{}).Where("DeletedAt = '' AND OnEnvironmentID = '%v'", env_result.ID).Find(&activeAlerts).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error querying for Service Alerts: %w", errors.InternalServerError, err))
		return nil, nil
	}
	return activeAlerts, env_result.ServiceBannerBucket

}

func getFileJson(ctx *gin.Context, gcs_client *google_bucket.GcsClientActual, blob *storage.ObjectAttrs) []ServiceAlertJsonData {
	byte_data, read_err := gcs_client.ReadBlob(ctx, blob)
	if read_err != nil {
		// handle error
		return nil
	}
	var json_data []ServiceAlertJsonData
	err := json.Unmarshal(byte_data, &json_data)
	if err != nil {
		// handle error
		return nil
	}
	return json_data
}

func compareAlerts(ctx *gin.Context, json_slice []ServiceAlertJsonData, db_alerts []models.ServiceAlert) ([]ServiceAlertJsonData, []models.ServiceAlert, []ServiceAlertJsonData) {
	// helper func, returns true if there are any differences between alert json and info from DB
	alertsNeedUpdate := func(jsonAlert ServiceAlertJsonData, dbAlert models.ServiceAlert) bool {
		return jsonAlert.Title != *dbAlert.Title ||
			jsonAlert.Message != *dbAlert.AlertMessage ||
			jsonAlert.Link != *dbAlert.Link ||
			jsonAlert.Severity != *dbAlert.Severity
	}

	var active_alerts int = len(json_slice)
	var alerts_in_db int = len(db_alerts)

	if active_alerts == 0 && alerts_in_db == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) Nothing to do, no active service alerts and nothing to modify", errors.NotFound))
		return nil, nil, nil
	}

	jsonMap := make(map[string]ServiceAlertJsonData)
	dbMap := make(map[string]models.ServiceAlert)

	var toAdd []ServiceAlertJsonData
	var toRemove []models.ServiceAlert
	var toUpdate []ServiceAlertJsonData
	// Find alerts to add (in JSON but not in DB)
	for incidentID, jsonAlert := range jsonMap {
		if _, exists := dbMap[incidentID]; !exists {
			toAdd = append(toAdd, jsonAlert)
		}
	}
	// Find alerts to remove (in DB but not in JSON)
	for incidentID, dbAlert := range dbMap {
		if _, exists := jsonMap[incidentID]; !exists {
			toRemove = append(toRemove, dbAlert)
		}
	}

	// Find alerts to update (in both but different)
	for incidentID, jsonAlert := range jsonMap {
		if dbAlert, exists := dbMap[incidentID]; exists {
			if alertsNeedUpdate(jsonAlert, dbAlert) {
				toUpdate = append(toUpdate, jsonAlert)
			}
		}
	}

	return toAdd, toRemove, toUpdate

}

func createSvcAlertJsonData(ctx *gin.Context, active_alerts []models.ServiceAlert) ([]byte, error) {
	var alerts []ServiceAlertJsonData
	for _, v := range active_alerts {
		tmp_alert := ServiceAlertJsonData{
			Title:      *v.Title,
			Message:    *v.AlertMessage,
			Link:       *v.Link,
			Severity:   *v.Severity,
			IncidentID: uuid.UUID.String(*v.Uuid),
		}
		alerts = append(alerts, tmp_alert)
	}
	data, err := json.Marshal(alerts)
	return data, err
}
