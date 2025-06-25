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
	alerts, gcsBucket := getAlerts(ctx, body, db)
	if len(alerts) == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) No Alerts found for this environment", errors.BadRequest))
		return
	}
	// set GCS client
	var gcsClient, googleClientError = google_bucket.InitializeStorageClient(ctx)
	if googleClientError != nil {
		return
	}
	// get alerts blob
	alertJsonBlob, readErr := gcsClient.GetBlob(ctx, *gcsBucket, "alerts.json")
	if readErr != nil {
		errors.AbortRequest(ctx, fmt.Errorf("blob not found: %v", readErr))
		return
	}
	// read data from blob & parse to JSON
	jsonData := getFileJson(ctx, gcsClient, alertJsonBlob)

	if len(jsonData) == 0 && len(alerts) == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) Nothing to do, no active service alerts and nothing to modify", errors.NotFound))
		return
	}
	// convert model data for service alerts to json formatted bytes to upload
	jsonBytes, err := createServiceAlertJsonData(ctx, alerts)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("issue encountered creating json payload: %v", err))
		return
	}
	// Upload file to bucket w/ latest info
	if err = gcsClient.WriteBlob(ctx, *gcsBucket, "alerts.json", jsonBytes); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error writing updated alerts.json file: %v", err))
		return
	}

}

// Return non-deleted service alerts matching specified env
func getAlerts(ctx *gin.Context, request ServiceAlertV3SyncRequest, db *gorm.DB) ([]models.ServiceAlert, *string) {
	var envResult models.Environment
	if request.OnEnvironment != nil {
		// match env so that we can get gcs bucket
		environmentQuery, err := environmentModelFromSelector(*request.OnEnvironment)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error parsing environment selector '%s': %w", *request.OnEnvironment, err))
			return nil, nil
		}
		if err = db.Where(&environmentQuery).Select("id").First(&envResult).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error fetching environment '%v'", err))
			return nil, nil
		}
	}
	var activeAlerts []models.ServiceAlert
	// Only return service alerts that haven't been deleted for this environment
	if err := db.Model(&models.ServiceAlert{}).Where("DeletedAt = '' AND OnEnvironmentID = '%v'", envResult.ID).Find(&activeAlerts).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error querying for Service Alerts: %w", errors.InternalServerError, err))
		return nil, nil
	}
	return activeAlerts, envResult.ServiceBannerBucket

}

// Read file from GCS, then parse JSON data and return []ServiceAlertJsonData struct
func getFileJson(ctx *gin.Context, gcsClient *google_bucket.GcsClientActual, blob *storage.ObjectAttrs) []ServiceAlertJsonData {
	byteData, readErr := gcsClient.ReadBlob(ctx, blob)
	if readErr != nil {
		// handle error
		return nil
	}
	var jsonData []ServiceAlertJsonData
	err := json.Unmarshal(byteData, &jsonData)
	if err != nil {
		// handle error
		return nil
	}
	return jsonData
}

// Transform service alerts struct to json formatted byte data to write to GCS blob
func createServiceAlertJsonData(ctx *gin.Context, activeAlerts []models.ServiceAlert) ([]byte, error) {
	var alerts []ServiceAlertJsonData
	for _, v := range activeAlerts {
		alertJsonStruct := ServiceAlertJsonData{
			Title:      *v.Title,
			Message:    *v.AlertMessage,
			Link:       *v.Link,
			Severity:   *v.Severity,
			IncidentID: uuid.UUID.String(*v.Uuid),
		}
		alerts = append(alerts, alertJsonStruct)
	}
	data, err := json.Marshal(alerts)
	return data, err
}
