package s3

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/samanshahroudi/todo-service/internal/ports"
)

type S3Adapter struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Adapter(client *s3.Client, bucketName string) ports.S3Service {
	return &S3Adapter{
		Client:     client,
		BucketName: bucketName,
	}
}

func (s *S3Adapter) UploadFile(key string, file io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return "", err
	}

	_, err = s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    types.ObjectCannedACLPrivate,
	})
	if err != nil {
		return "", err
	}

	return key, nil
}
