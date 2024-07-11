package S3__Upload

import (
	"context"
	"net/http"
	"os"
	"time"

	aws_conf "github.com/ank809/Image-Processor-Go/aws"
	"github.com/ank809/Image-Processor-Go/helpers"
	"github.com/ank809/Image-Processor-Go/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func UploadToS3(c *gin.Context) {
	var fileinfo models.FileInfo
	err := c.BindJSON(&fileinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}
	s3client, err := aws_conf.GetS3Client()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting S3 client"})
		return
	}
	uniquekey := helpers.GetUniqueKey()

	key := fileinfo.Email + "/" + uniquekey + fileinfo.Filename
	bucketname := os.Getenv("BUCKET_NAME")

	presignClient := s3.NewPresignClient(s3client)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	presignedurl, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(key),
	})

	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"url":         presignedurl.URL,
		"bucket_name": bucketname,
		"key":         key,
	})
}
