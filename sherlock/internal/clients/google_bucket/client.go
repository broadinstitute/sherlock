package google_bucket

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

var (
	// client is what functions in this package should use whenever possible.
	client GcsClient
)

// set blob details to include list of Acl info so that we can set multiple permission sets if provided
type BlobDetails struct {
	Bucket   string
	BlobName string
	AclAttrs []storage.ACLRule
	ObjAttrs *storage.ObjectAttrsToUpdate
}

// wrapper interface for generating mocks
type GcsClient interface {
	ListBlobs(ctx context.Context, bucket string) ([]*storage.ObjectAttrs, error)
	ReadBlob(ctx context.Context, blob *storage.ObjectAttrs) ([]byte, error)
	WriteBlob(ctx context.Context, blobInfo BlobDetails, file_content []byte) error
	GetBlob(ctx context.Context, bucket_name string, blob_name string) (*storage.ObjectAttrs, error)
}

type GcsClientActual struct {
	GcsClient *storage.Client
}

func InitializeStorageClient(ctx context.Context) (GcsClient, error) {
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

func (client *GcsClientActual) WriteBlob(ctx context.Context, blobInfo BlobDetails, file_content []byte) error {
	if blobInfo.BlobName == "" || blobInfo.Bucket == "" {
		return fmt.Errorf("unable to proceed without necessary information: blob name: %v \n bucket name: %v", blobInfo.BlobName, blobInfo.Bucket)
	}
	blob := client.GcsClient.Bucket(blobInfo.Bucket).Object(blobInfo.BlobName)

	writer_client := blob.NewWriter(ctx)
	writer_client.ContentType = "application/json"

	if _, err := writer_client.Write(file_content); err != nil {
		return fmt.Errorf("createFile: unable to write data to bucket: %v", err)
	}

	if err := writer_client.Close(); err != nil {
		return fmt.Errorf("got %v, want nil", err)
	}

	if len(blobInfo.AclAttrs) != 0 {
		if err := setBlobAcl(ctx, blob, blobInfo.AclAttrs); err != nil {
			return err
		}
	}

	if blobInfo.ObjAttrs != nil {
		if err := setBlobAttrs(ctx, blob, blobInfo.ObjAttrs); err != nil {
			return err
		}
	}

	return nil
}

// internal func, set blob ACL
func setBlobAcl(ctx context.Context, blob *storage.ObjectHandle, aclList []storage.ACLRule) error {
	for _, v := range aclList {
		if err := blob.ACL().Set(ctx, v.Entity, v.Role); err != nil {
			return fmt.Errorf("issue granting read permissions to blob: %v", err)
		}
	}
	return nil
}

// interrnal func, update blob Attrs
func setBlobAttrs(ctx context.Context, blob *storage.ObjectHandle, objAttrs *storage.ObjectAttrsToUpdate) error {
	if _, err := blob.Update(ctx, *objAttrs); err != nil {
		return fmt.Errorf("issue updating object attrs %v \n blob: %v \n attrs: %v", err, blob, objAttrs)
	} else {
		return nil
	}
}

func (client *GcsClientActual) GetBlob(ctx context.Context, bucket_name string, blob_name string) (*storage.ObjectAttrs, error) {
	attrs, err := client.GcsClient.Bucket(bucket_name).Object(blob_name).Attrs(ctx)
	if errors.Is(err, storage.ErrObjectNotExist) {
		fmt.Println("The object does not exist")
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("error encountered attempting to get blob: %v", err)
	}
	return attrs, nil
}

// GetClient returns the current client (either mocked or a new real client)
func GetClient(ctx context.Context) (GcsClient, error) {
	if client != nil {
		return client, nil
	}
	return InitializeStorageClient(ctx)
}

// SetClient sets the global client (used for testing)
func SetClient(c GcsClient) {
	client = c
}

// ResetClient resets the global client to nil
func ResetClient() {
	client = nil
}
