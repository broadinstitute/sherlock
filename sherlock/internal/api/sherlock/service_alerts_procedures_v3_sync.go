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

// usersProceduresV3Deactivate godoc
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
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ServiceAlertV3SyncRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}
	var alerts, gcs_bucket = getAlerts(ctx, body, db)
	if len(alerts) == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) No Alerts found for this environment", errors.BadRequest))
		return
	}
	var gcsClient, gc_err = google_bucket.InitializeStorageClient(ctx)
	if gc_err != nil {
		return
	}
	alert_json_blob, read_err := gcsClient.GetBlob(ctx, *gcs_bucket, "alerts.json")
	if read_err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("Blob not found: %v", read_err))
		return
	}
	json_data := getFileJson(ctx, gcsClient, alert_json_blob)

}

func getAlerts(ctx *gin.Context, request ServiceAlertV3SyncRequest, db *gorm.DB) ([]models.ServiceAlert, *string) {
	var env_result models.Environment
	if request.OnEnvironment != nil {
		environmentQuery, err := environmentModelFromSelector(*request.OnEnvironment)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error parsing environment selector '%s': %w", *request.OnEnvironment, err))
			return nil
		}
		if err = db.Where(&environmentQuery).Select("id").First(&env_result).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error fetching environment '%v'", err))
			return nil
		}
	}
	var activeAlerts []models.ServiceAlert
	// Only return service alerts that haven't been deleted
	if err := db.Model(&models.ServiceAlert{}).Where("DeletedAt = '' AND OnEnvironmentID = '%v'", env_result.ID).Find(&activeAlerts).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error querying for Service Alerts: %w", errors.InternalServerError, err))
		return nil
	}
	return activeAlerts, env_result.ServiceBannerBucket

}

func getFileJson(ctx *gin.Context, gcs_client *google_bucket.GcsClientActual, blob *storage.ObjectAttrs) []interface{} {
	byte_data, read_err := gcs_client.ReadBlob(ctx, blob)
	if read_err != nil {
		// handle error
		return nil
	}
	var json_data []interface{}
	err := json.Unmarshal(byte_data, &json_data)
	if err != nil {
		// handle error
		return nil
	}
	return json_data
}
