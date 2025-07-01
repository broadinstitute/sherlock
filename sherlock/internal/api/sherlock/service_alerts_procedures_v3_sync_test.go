package sherlock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket/google_bucket_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/mock"
)

func (s *handlerSuite) TestSyncServiceAlerts_Success() {
	s.SetNonSuitableTestUserForDB()
	env := s.TestData.Environment_Dev()
	alert1 := s.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	s.NoError(s.DB.Save(&alert1).Error)

	google_bucket.UseMockedClient(s.T(), func(c *google_bucket_mocks.MockgcsClient) {
		c.EXPECT().WriteBlob(mock.Anything, *env.ServiceBannerBucket, "alerts.json", mock.Anything).Return(nil).Once()
	}, func() {
		var got []ServiceAlertV3
		code := s.HandleRequest(
			s.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: env.Name,
			}),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
		s.Equal(*alert1.Title, *got[0].Title)
	})
}

func (s *handlerSuite) TestSyncServiceAlerts_WriteBlobFails() {
	s.SetNonSuitableTestUserForDB()
	env := s.TestData.Environment_Dev()
	alert1 := s.TestData.ServiceAlert_1()
	alert1.OnEnvironmentID = &env.ID
	s.NoError(s.DB.Save(&alert1).Error)

	google_bucket.UseMockedClient(s.T(), func(c *google_bucket_mocks.MockgcsClient) {
		c.EXPECT().WriteBlob(mock.Anything, *env.ServiceBannerBucket, "alerts.json", mock.Anything).Return(fmt.Errorf("some GCS error")).Once()
	}, func() {
		var got errors.ErrorResponse
		code := s.HandleRequest(
			s.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: env.Name,
			}),
			&got)
		s.Equal(http.StatusInternalServerError, code)
		s.Contains(got.Message, "error writing updated alerts.json file")
	})
}

func (s *handlerSuite) TestSyncServiceAlerts_NoAlerts() {
	s.SetNonSuitableTestUserForDB()
	env := s.TestData.Environment_Dev()

	// Ensure there are no alerts for this environment by clearing any existing ones
	s.NoError(s.DB.Where("on_environment_id = ?", env.ID).Delete(&models.ServiceAlert{}).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
			OnEnvironment: env.Name,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Contains(got.Message, "No Alerts found for this environment")
}

func (s *handlerSuite) TestCreateServiceAlertJsonData() {
	alert1 := s.TestData.ServiceAlert_1()
	alert2 := s.TestData.ServiceAlert_2()
	alerts := []models.ServiceAlert{alert1, alert2}
	jsonData, err := createServiceAlertJsonData(alerts)
	s.NoError(err)

	var got []ServiceAlertJsonData
	err = json.Unmarshal(jsonData, &got)
	s.NoError(err)

	s.Len(got, 2)
	s.Equal(*alert1.Title, got[0].Title)
	s.Equal(alert1.Uuid.String(), got[0].IncidentID)
	s.Equal(*alert2.Title, got[1].Title)
	s.Equal(alert2.Uuid.String(), got[1].IncidentID)
}

func createFakeObjectAttrs(bucket, name string) *storage.ObjectAttrs {
	return &storage.ObjectAttrs{
		Bucket:      bucket,
		Name:        name,
		Size:        1024,
		ContentType: "application/json",
	}
}

func createMockAlertJsonData() []byte {
	alerts := []ServiceAlertJsonData{
		{
			Title:      "Production Issue",
			Message:    "Service experiencing high latency",
			Link:       "https://status.example.com/incident/123",
			Severity:   "high",
			IncidentID: "incident-123",
		},
		{
			Title:      "Maintenance Window",
			Message:    "Scheduled maintenance tonight",
			Link:       "https://status.example.com/maintenance/456",
			Severity:   "medium",
			IncidentID: "maintenance-456",
		},
	}
	data, err := json.Marshal(alerts)
	if err != nil {
		return nil
	}
	// GCS client ensures a newline, so we add one here for comparisons
	return append(data, '\n')
}
