package google_bucket

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GcsClientActual struct {
	GcsClient *storage.Client
}

func InitializeStorageClient(ctx context.Context, impersonateAccount ...string) (*storage.Client, error) {
	// Client uses xml google api's by default but docs reccomend using JSON - will become default behavior in later release
	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		return nil, fmt.Errorf("Failed to create storage client: %v\n", err)
	}
	return &GcsClientActual{GcsClient: client}, err
}

func (c *GcsClientActual) ListBlobs(ctx context.Context, bucket string) ([]*storage.ObjectAttrs, error) {
	var bucket_objs []*storage.ObjectAttrs
	query := &storage.Query{Prefix: ""}
	client_bucket := c.GcsClient.Bucket(bucket)
	it := client_bucket.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Issue obtaining bucket objects: %v\n", err)
		}
		bucket_objs = append(bucket_objs, attrs)

	}
	return bucket_objs, nil
}

func (c *GcsClientActual) ReadBlob(ctx context.Context, blob *storage.ObjectAttrs) ([]byte, error) {
	reader_client, err := c.GcsClient.Bucket(blob.Bucket).Object(blob.Name).NewReader(ctx)
	if err.Is(err, storage.ErrObjectNotExist) {
		return nil, fmt.Errorf("The object does not exist")
	}
	slurp, err := io.ReadAll(reader_client)
	reader_client.Close()
	if err != nil {
		return nil, fmt.Errorf("Unable to read blob: %v", err)
	}
	// CONVERT to JSON
	// var json_data map[string]interface{}
	//if err := json.Unmarshal(slurp, &json_data); err != nil {
	//		panic(err)
	//	}
	return slurp, nil
}

func (c *GcsClientActual) WriteBlob(ctx context.Context, gcs_bucket string, blob_name string, file_content []byte) error {
	writer_client, err := c.GcsClient.Bucket(gcs_bucket).Object(blob_name).NewWriter(ctx)
	if err != nil {
		return fmt.Errorf("Error creating writer client: %v", err)
	}
	writer_client.ContentType = "application/json"
	if _, err := writer_client.Write(file_content); err != nil {
		return fmt.Errorf("createFile: unable to write data to bucket: %v", err)
	}
	return nil
}
