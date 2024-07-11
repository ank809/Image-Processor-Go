package S3__Upload

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/ank809/Image-Processor-Go/helpers"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToPermanentS3(s3Client *s3.Client, email string, image []byte, filename string) error {
	bucketName := os.Getenv("PERMANENT_BUCKET")
	if bucketName == "" {
		return fmt.Errorf("PERMANENT_BUCKET environment variable is not set")
	}

	uniqueKey := helpers.GetUniqueKey()
	key := email + "/" + uniqueKey + filename

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   bytes.NewReader(image),
	})
	if err != nil {
		return fmt.Errorf("error uploading image to S3: %w", err)
	}

	return nil
}
