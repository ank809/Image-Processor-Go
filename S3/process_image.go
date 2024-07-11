package S3__Upload

import (
	"log"
	"net/http"

	"github.com/ank809/Image-Processor-Go/ImageOperations"
	aws_conf "github.com/ank809/Image-Processor-Go/aws"
	"github.com/ank809/Image-Processor-Go/models"
	"github.com/gin-gonic/gin"
)

func ProcessImage(c *gin.Context) {

	s3Client, err := aws_conf.GetS3Client()
	if err != nil {
		c.JSON(400, "Erorr in getting s3 client")
		return
	}
	var op models.Operation
	var processedImage []byte

	err = c.ShouldBindJSON(&op)
	if err != nil {
		c.JSON(400, "Error in binding json")
		return
	}

	userclaims, _ := c.Get("user")
	claims, _ := userclaims.(*models.Claims)

	img, err := DownloadImage(op, claims.Email, s3Client)
	if err != nil {
		c.JSON(400, err)
		return
	}
	if len(img) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Downloaded image is empty"})
		return
	}
	// Process the image and perform some operations
	switch op.OperationName {
	case "crop":
		processedImage, err = ImageOperations.CropImage(img, op.Height, op.Width)
	case "enlarge":
		processedImage, err = ImageOperations.Enlarge(img, op.Height, op.Width)
	case "smartcrop":
		processedImage, err = ImageOperations.SmartCrop(img, op.Height, op.Width)
	case "resize":
		processedImage, err = ImageOperations.Resize(img, op.Height, op.Width)
	default:
		c.JSON(400, "Image operation not allowed")
		return
	}
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = UploadToPermanentS3(s3Client, claims.Email, processedImage, op.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	key := claims.Email + "/" + op.Key
	isDeleted, err := DeleteObjectFromBucket(key)
	if !isDeleted {
		log.Printf("Error in deleteing from temporary bucket %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":         "Image downloaded successfully",
		"img_size":        len(img),
		"processed_image": len(processedImage),
		"email":           claims.Email,
	})

}
