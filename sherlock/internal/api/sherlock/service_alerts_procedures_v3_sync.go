package sherlock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceAlertV3SyncRequest struct {
	OnEnvironment string `json:"onEnvironment,omitempty" form:"onEnvironment"`
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
//	@param			environment				body		ServiceAlertV3SyncRequest	true	"Information on Service Alert environment"
//	@success		200						{array}		ServiceAlertV3
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
	}
	// Get alerts from DB and gcs bucket from environment
	alerts, gcsBucket := getAlerts(ctx, body, db)

	if alerts == nil || gcsBucket == nil {
		errors.AbortRequest(ctx, fmt.Errorf("issue obtaining alert or bucket information. alerts: %v \n bucket: %v", alerts, gcsBucket))
		return
	}
	// set GCS client
	gcsClient, googleClientError := google_bucket.GetClient(ctx)
	if googleClientError != nil {
		errors.AbortRequest(ctx, fmt.Errorf("blob not found: %v", googleClientError))
		return
	}

	jsonBytes, err := createServiceAlertJsonData(alerts)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("issue encountered creating json payload: %v", err))
		return
	}
	var permissionsList []storage.ACLRule
	permissionsList = append(permissionsList,
		storage.ACLRule{
			Role:   storage.RoleReader,
			Entity: storage.AllUsers,
		})

	var objAttrsToUpdate = storage.ObjectAttrsToUpdate{
		CacheControl: "no-store",
	}

	// ObjAttrs being set as pointer so we can check null value
	blobDetails := google_bucket.BlobDetails{
		Bucket:   *gcsBucket,
		BlobName: "alerts.json",
		AclAttrs: permissionsList,
		ObjAttrs: &objAttrsToUpdate,
	}
	// Upload file to bucket w/ latest info
	if err = gcsClient.WriteBlob(ctx, blobDetails, jsonBytes); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error writing updated alerts.json file: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.Map(alerts, ServiceAlertFromModel))

}

// Return non-deleted service alerts matching specified env
func getAlerts(ctx *gin.Context, request ServiceAlertV3SyncRequest, db *gorm.DB) ([]models.ServiceAlert, *string) {
	var envResult models.Environment
	if request.OnEnvironment != "" {
		// match env so that we can get gcs bucket
		environmentQuery, err := environmentModelFromSelector(request.OnEnvironment)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error parsing environment selector '%s': %w", request.OnEnvironment, err))
			return nil, nil
		}
		if err = db.Where(&environmentQuery).First(&envResult).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) error fetching environment '%v'", errors.BadRequest, err))
			return nil, nil
		}
	}
	var activeAlerts []models.ServiceAlert
	// Only return service alerts that haven't been deleted for this environment
	if err := db.Where(&models.ServiceAlert{OnEnvironmentID: &envResult.ID}).Find(&activeAlerts).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error querying for Service Alerts: %w", errors.BadRequest, err))
		return nil, nil
	}
	return activeAlerts, envResult.ServiceBannerBucket

}

// Transform service alerts struct to json formatted byte data to write to GCS blob.
func createServiceAlertJsonData(activeAlerts []models.ServiceAlert) ([]byte, error) {
	var alerts []ServiceAlertJsonData

	// exit early w/ empty array if no activeAlerts exist
	if len(activeAlerts) == 0 {
		// creating zero len non-null slice to set empty service alert json to [] instead of null
		emptySlice := []string{}
		data, err := json.Marshal(emptySlice)
		return data, err
	}

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
