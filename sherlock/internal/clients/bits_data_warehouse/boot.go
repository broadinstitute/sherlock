package bits_data_warehouse

import (
	"cloud.google.com/go/bigquery"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"golang.org/x/net/context"
)

var client *bigquery.Client

func Init(ctx context.Context) error {
	project := bigquery.DetectProjectID
	if configProject := config.Config.String("bitsDataWarehouse.jobProject"); configProject != "" {
		project = configProject
	}

	var err error
	client, err = bigquery.NewClient(ctx, project)
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
