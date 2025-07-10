package sherlock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket/google_bucket_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/mock"
)

func (test_handler *handlerSuite) TestSyncServiceAlerts_Success() {
	test_handler.SetNonSuitableTestUserForDB()
	env := test_handler.TestData.Environment_Dev()
	alert1 := test_handler.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	test_handler.NoError(test_handler.DB.Save(&alert1).Error)

	google_bucket.UseMockedClient(test_handler.T(), func(client *google_bucket_mocks.MockgcsClient) {
		client.EXPECT().WriteBlob(mock.Anything, *env.ServiceBannerBucket, "alerts.json", mock.Anything).Return(nil).Once()
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
	test_handler.SetNonSuitableTestUserForDB()
	env := test_handler.TestData.Environment_Dev()
	alert1 := test_handler.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	test_handler.NoError(test_handler.DB.Save(&alert1).Error)

	google_bucket.UseMockedClient(test_handler.T(), func(client *google_bucket_mocks.MockgcsClient) {
		client.EXPECT().WriteBlob(mock.Anything, *env.ServiceBannerBucket, "alerts.json", mock.Anything).Return(fmt.Errorf("some GCS error")).Once()
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
	test_handler.SetNonSuitableTestUserForDB()
	env := test_handler.TestData.Environment_Dev()

	// Ensure there are no alerts for this environment by clearing any existing ones
	test_handler.NoError(test_handler.DB.Where("on_environment_id = ?", env.ID).Delete(&models.ServiceAlert{}).Error)

	google_bucket.UseMockedClient(test_handler.T(), func(client *google_bucket_mocks.MockgcsClient) {
		client.EXPECT().WriteBlob(mock.Anything, *env.ServiceBannerBucket, "alerts.json", mock.Anything).Return(nil).Once()
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
	alert2 := test_handler.TestData.ServiceAlert_2()
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
