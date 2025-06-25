package google_bucket

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GcsClientActual struct {
	GcsClient *storage.Client
}

func InitializeStorageClient(ctx context.Context, impersonateAccount ...string) (*GcsClientActual, error) {
	// Client uses xml google api's by default but docs reccomend using JSON - will become default behavior in later release
	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %v", err)
	}
	return &GcsClientActual{GcsClient: client}, err
}

func (client *GcsClientActual) ListBlobs(ctx context.Context, bucket string) ([]*storage.ObjectAttrs, error) {
	var bucket_objs []*storage.ObjectAttrs
	query := &storage.Query{Prefix: ""}
	client_bucket := client.GcsClient.Bucket(bucket)
	blobIterator := client_bucket.Objects(ctx, query)
	for {
		attrs, err := blobIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("issue obtaining bucket objects: %v", err)
		}
		bucket_objs = append(bucket_objs, attrs)

	}
	return bucket_objs, nil
}

func (client *GcsClientActual) ReadBlob(ctx context.Context, blob *storage.ObjectAttrs) ([]byte, error) {
	reader_client, err := client.GcsClient.Bucket(blob.Bucket).Object(blob.Name).NewReader(ctx)
	if errors.Is(err, storage.ErrObjectNotExist) {
		return nil, fmt.Errorf("the object does not exist")
	}
	slurp, err := io.ReadAll(reader_client)
	reader_client.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to read blob: %v", err)
	}
	return slurp, nil
}

func (client *GcsClientActual) WriteBlob(ctx context.Context, gcs_bucket string, blob_name string, file_content []byte) error {
	writer_client := client.GcsClient.Bucket(gcs_bucket).Object(blob_name).NewWriter(ctx)

	writer_client.ContentType = "application/json"
	if _, err := writer_client.Write(file_content); err != nil {
		return fmt.Errorf("createFile: unable to write data to bucket: %v", err)
	}
	return nil
}

func (client *GcsClientActual) GetBlob(ctx context.Context, bucket_name string, blob_name string) (*storage.ObjectAttrs, error) {
	attrs, err := client.GcsClient.Bucket(bucket_name).Object(blob_name).Attrs(ctx)
	if errors.Is(err, storage.ErrObjectNotExist) {
		fmt.Println("The object does not exist")
		return nil, err
	}
	return attrs, nil
}
