package S3__Upload

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/ank809/Image-Processor-Go/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadImage(op models.Operation, email string, s3Client *s3.Client) ([]byte, error) {

	// email
	key := email + "/" + op.Key
	bucket_name := os.Getenv("TEMP_BUCKET_NAME")
	res, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket_name),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
