package S3__Upload

import (
	"context"
	"fmt"
	"os"

	aws_conf "github.com/ank809/Image-Processor-Go/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteObjectFromBucket(key string) (bool, error) {
	s3Client, err := aws_conf.GetS3Client()
	if err != nil {
		return false, fmt.Errorf("error getting S3 client: %w", err)
	}

	bucketName := os.Getenv("TEMP_BUCKET_NAME")

	_, err = s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return false, fmt.Errorf("error deleting object from S3: %w", err)
	}

	return true, nil
}
