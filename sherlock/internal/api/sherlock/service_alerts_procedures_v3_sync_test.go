package sherlock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket/google_bucket_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/mock"
)

// UseMockedClient temporarily replaces the global GCS client with a mock for testing
func UseMockedClient(t *testing.T, config func(mock_client *google_bucket_mocks.MockgcsClient), callback func()) {
	if config == nil {
		callback()
		return
	}
	mock_client := google_bucket_mocks.NewMockgcsClient(t)
	config(mock_client)

	// Set the mock as the global client
	google_bucket.SetClient(mock_client)

	// Ensure cleanup happens even if test panics
	defer google_bucket.ResetClient()

	// Run the test
	callback()

	// Verify expectations
	mock_client.AssertExpectations(t)
}

func (test_handler *handlerSuite) TestSyncServiceAlerts_Success() {
	env := test_handler.TestData.Environment_Dev()
	alert1 := test_handler.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	test_handler.NoError(test_handler.DB.Save(&alert1).Error)

	var permissionsList []storage.ACLRule
	permissionsList = append(permissionsList,
		storage.ACLRule{
			Role:   storage.RoleReader,
			Entity: storage.AllUsers,
		})
	var objAttrsToUpdate = storage.ObjectAttrsToUpdate{
		CacheControl: "no-store",
	}
	blobDetails := google_bucket.BlobDetails{
		BlobName: "alerts.json",
		Bucket:   *env.ServiceBannerBucket,
		AclAttrs: permissionsList,
		ObjAttrs: &objAttrsToUpdate,
	}

	UseMockedClient(test_handler.T(), func(gcs_mock_client *google_bucket_mocks.MockgcsClient) {
		gcs_mock_client.EXPECT().WriteBlob(mock.Anything, blobDetails, mock.Anything).Return(nil).Once()
	}, func() {
		var got []ServiceAlertV3
		code := test_handler.HandleRequest(
			test_handler.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: env.Name,
			}),
			&got)
		test_handler.Equal(http.StatusOK, code)
		test_handler.Len(got, 1)
		test_handler.Equal(*alert1.Title, *got[0].Title)
	})
}

// test to write blob
func (test_handler *handlerSuite) TestSyncServiceAlerts_WriteBlobFails() {
	env := test_handler.TestData.Environment_Dev()
	alert1 := test_handler.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	test_handler.NoError(test_handler.DB.Save(&alert1).Error)

	var permissionsList []storage.ACLRule
	permissionsList = append(permissionsList,
		storage.ACLRule{
			Role:   storage.RoleReader,
			Entity: storage.AllUsers,
		})

	var objAttrsToUpdate = storage.ObjectAttrsToUpdate{
		CacheControl: "no-store",
	}
	blobDetails := google_bucket.BlobDetails{
		BlobName: "alerts.json",
		Bucket:   *env.ServiceBannerBucket,
		AclAttrs: permissionsList,
		ObjAttrs: &objAttrsToUpdate,
	}
	UseMockedClient(test_handler.T(), func(gcs_mock_client *google_bucket_mocks.MockgcsClient) {
		gcs_mock_client.EXPECT().WriteBlob(mock.Anything, blobDetails, mock.Anything).Return(fmt.Errorf("error writing updated alerts.json file")).Once()
	}, func() {
		var got errors.ErrorResponse
		code := test_handler.HandleRequest(
			test_handler.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: env.Name,
			}),
			&got)
		test_handler.Equal(http.StatusInternalServerError, code)
		test_handler.Contains(got.Message, "error writing updated alerts.json file")
	})
}

func (test_handler *handlerSuite) TestSyncServiceAlerts_NoAlerts() {
	env := test_handler.TestData.Environment_Dev()

	// Ensure there are no alerts for this environment by clearing any existing ones
	//err := test_handler.DB.Where("on_environment_id = ?", env.ID).Delete(&models.ServiceAlert{}).Error
	//test_handler.NoError(err)

	var permissionsList []storage.ACLRule
	permissionsList = append(permissionsList,
		storage.ACLRule{
			Role:   storage.RoleReader,
			Entity: storage.AllUsers,
		})
	var objAttrsToUpdate = storage.ObjectAttrsToUpdate{
		CacheControl: "no-store",
	}
	blobDetails := google_bucket.BlobDetails{
		BlobName: "alerts.json",
		Bucket:   *env.ServiceBannerBucket,
		AclAttrs: permissionsList,
		ObjAttrs: &objAttrsToUpdate,
	}

	UseMockedClient(test_handler.T(), func(gcs_mock_client *google_bucket_mocks.MockgcsClient) {
		// Expect WriteBlob to be called with empty JSON array
		gcs_mock_client.EXPECT().WriteBlob(
			mock.Anything,
			blobDetails,
			mock.MatchedBy(func(content []byte) bool {
				// Verify it's an empty JSON array
				var result []interface{}
				return json.Unmarshal(content, &result) == nil && len(result) == 0
			}),
		).Return(nil).Once()
	}, func() {
		var got []ServiceAlertV3
		code := test_handler.HandleRequest(
			test_handler.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: env.Name,
			}),
			&got)
		test_handler.Equal(http.StatusOK, code)
		test_handler.Len(got, 0)
	})
}

// Ensure service alert model data can be transformed into json.
func (test_handler *handlerSuite) TestCreateServiceAlertJsonData() {
	alert1 := test_handler.TestData.ServiceAlert_1()
	alert2 := test_handler.TestData.ServiceAlert_Prod()
	alerts := []models.ServiceAlert{alert1, alert2}
	jsonData, err := createServiceAlertJsonData(alerts)
	test_handler.NoError(err)

	var got []ServiceAlertJsonData
	err = json.Unmarshal(jsonData, &got)
	test_handler.NoError(err)

	test_handler.Len(got, 2)
	test_handler.Equal(*alert1.Title, got[0].Title)
	test_handler.Equal(alert1.Uuid.String(), got[0].IncidentID)
	test_handler.Equal(*alert2.Title, got[1].Title)
	test_handler.Equal(alert2.Uuid.String(), got[1].IncidentID)
}

// Test with detailed JSON content validation
func (test_handler *handlerSuite) TestSyncServiceAlerts_ValidateJSONContent() {
	env := test_handler.TestData.Environment_Dev()
	alert1 := test_handler.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	test_handler.NoError(test_handler.DB.Save(&alert1).Error)

	var permissionsList []storage.ACLRule
	permissionsList = append(permissionsList,
		storage.ACLRule{
			Role:   storage.RoleReader,
			Entity: storage.AllUsers,
		})
	var objAttrsToUpdate = storage.ObjectAttrsToUpdate{
		CacheControl: "no-store",
	}
	blobDetails := google_bucket.BlobDetails{
		BlobName: "alerts.json",
		Bucket:   *env.ServiceBannerBucket,
		AclAttrs: permissionsList,
		ObjAttrs: &objAttrsToUpdate,
	}

	UseMockedClient(test_handler.T(), func(gcs_mock_client *google_bucket_mocks.MockgcsClient) {
		gcs_mock_client.EXPECT().WriteBlob(
			mock.Anything,
			blobDetails,
			mock.MatchedBy(func(content []byte) bool {
				// Parse and validate the JSON content
				var alerts []ServiceAlertJsonData
				if err := json.Unmarshal(content, &alerts); err != nil {
					return false
				}

				// Validate the content matches expected alert data
				return len(alerts) == 1 &&
					alerts[0].Title == *alert1.Title &&
					alerts[0].Message == *alert1.AlertMessage &&
					alerts[0].Link == *alert1.Link &&
					alerts[0].Severity == *alert1.Severity &&
					alerts[0].IncidentID == alert1.Uuid.String()
			}),
		).Return(nil).Once()
	}, func() {
		var got []ServiceAlertV3
		code := test_handler.HandleRequest(
			test_handler.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: env.Name,
			}),
			&got)
		test_handler.Equal(http.StatusOK, code)
		test_handler.Len(got, 1)
	})
}

// Test invalid environment
func (test_handler *handlerSuite) TestSyncServiceAlerts_InvalidEnvironment() {

	UseMockedClient(test_handler.T(), func(gcs_mock_client *google_bucket_mocks.MockgcsClient) {
		// No expectations set - WriteBlob should not be called for invalid environment
	}, func() {
		var got errors.ErrorResponse
		code := test_handler.HandleRequest(
			test_handler.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: "nonexistent-environment",
			}),
			&got)

		test_handler.Equal(http.StatusBadRequest, code)
		test_handler.Contains(got.Message, "error fetching environment")
	})
}
