package bits_data_warehouse

import (
	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
)

var client *bigquery.Client

func Init(ctx context.Context) error {
	var err error
	client, err = bigquery.NewClient(ctx, bigquery.DetectProjectID)
	if err != nil {
		return err
	}

	calculatePersonColumns()

	err = updatePeopleCache(ctx)
	if err != nil {
		return err
	}

	go keepPeopleCacheUpdated(ctx)

	return nil
}
