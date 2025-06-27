package sherlock

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_bucket/google_bucket_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func (s *handlerSuite) TestSyncSingleServiceAlerts() {
	serviceAlert := s.TestData.ServiceAlert_1()

	fakeObjectAttrs := &storage.ObjectAttrs{
		Bucket:      *serviceAlert.OnEnvironment.ServiceBannerBucket,
		Name:        "alerts.json",
		Size:        1024,
		ContentType: "application/json",
		Created:     time.Now(),
		Updated:     time.Now(),
	}
	// Mock JSON data that would be returned
	mockJsonData := []ServiceAlertJsonData{
		{
			Title:      *serviceAlert.Title,
			Message:    *serviceAlert.AlertMessage,
			Link:       *serviceAlert.Link,
			Severity:   *serviceAlert.Severity,
			IncidentID: uuid.UUID.String(*serviceAlert.Uuid),
		},
	}
	mockJsonBytes, _ := json.Marshal(mockJsonData)

	ctx := context.Background()
	google_bucket.UseMockedClient(s.T(), func(mock_client *google_bucket_mocks.MockgcsClient) {
		mock_client.EXPECT().GetBlob(ctx, serviceAlert.OnEnvironment.ServiceBannerBucket, "alerts.json").Return(fakeObjectAttrs, nil).Once()
		mock_client.EXPECT().ReadBlob(ctx, fakeObjectAttrs).Return(mockJsonBytes, nil).Once()
		mock_client.EXPECT().WriteBlob(ctx, serviceAlert.OnEnvironment.ServiceBannerBucket, "alerts.json", mock.Anything).Return(nil).Once()

	}, func() {
		var got []models.ServiceAlert

		code := s.HandleRequest(
			s.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: &serviceAlert.OnEnvironment.Name,
			}),
			&got)
		s.Equal(http.StatusOK, code)
	})

}

func (s *handlerSuite) TestSyncMultipleServiceAlerts() {
	serviceAlert := s.TestData.ServiceAlert_1()
	serviceAlert2 := s.TestData.ServiceAlert_2()

	fakeObjectAttrs := &storage.ObjectAttrs{
		Bucket:      *serviceAlert.OnEnvironment.ServiceBannerBucket,
		Name:        "alerts.json",
		Size:        1024,
		ContentType: "application/json",
		Created:     time.Now(),
		Updated:     time.Now(),
	}
	// Mock JSON data that would be returned
	mockJsonData := []ServiceAlertJsonData{
		{
			Title:      *serviceAlert.Title,
			Message:    *serviceAlert.AlertMessage,
			Link:       *serviceAlert.Link,
			Severity:   *serviceAlert.Severity,
			IncidentID: uuid.UUID.String(*serviceAlert.Uuid),
		},
		{
			Title:      *serviceAlert2.Title,
			Message:    *serviceAlert2.AlertMessage,
			Link:       *serviceAlert2.Link,
			Severity:   *serviceAlert2.Severity,
			IncidentID: uuid.UUID.String(*serviceAlert2.Uuid),
		},
	}
	mockJsonBytes, _ := json.Marshal(mockJsonData)

	ctx := context.Background()
	google_bucket.UseMockedClient(s.T(), func(mock_client *google_bucket_mocks.MockgcsClient) {
		mock_client.EXPECT().GetBlob(ctx, serviceAlert.OnEnvironment.ServiceBannerBucket, "alerts.json").Return(fakeObjectAttrs, nil).Once()
		mock_client.EXPECT().ReadBlob(ctx, fakeObjectAttrs).Return(mockJsonBytes, nil).Once()
		mock_client.EXPECT().WriteBlob(ctx, serviceAlert.OnEnvironment.ServiceBannerBucket, "alerts.json", mock.Anything).Return(nil).Once()

	}, func() {
		var got []models.ServiceAlert

		code := s.HandleRequest(
			s.NewRequest("POST", "/api/service-alerts/procedures/v3/sync", ServiceAlertV3SyncRequest{
				OnEnvironment: &serviceAlert.OnEnvironment.Name,
			}),
			&got)
		s.Equal(http.StatusOK, code)
		s.Assert()
	})

}
