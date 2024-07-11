package main

import (
	"log"
	"net/http"

	S3__Upload "github.com/ank809/Image-Processor-Go/S3"
	"github.com/ank809/Image-Processor-Go/authentication"
	"github.com/ank809/Image-Processor-Go/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/signup", authentication.SignUp)
	r.GET("/login", authentication.LoginUser)
	r.POST("/upload", S3__Upload.UploadToS3).Use(middlewares.AuthMiddleware())
	r.POST("/processimage", S3__Upload.ProcessImage)
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Println(err)
		return
	}
	log.Println("Listening on Port 8081")
}
