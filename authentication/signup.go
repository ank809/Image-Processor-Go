package authentication

import (
	"context"
	"net/http"
	"time"

	"github.com/ank809/Image-Processor-Go/database"
	"github.com/ank809/Image-Processor-Go/helpers"
	"github.com/ank809/Image-Processor-Go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, "Error in binding json")
		return
	}
	user.ID = primitive.NewObjectID()
	//Check for name
	value, msg := helpers.CheckName(user.Name)
	if !value {
		c.JSON(400, msg)
		return
	}
	// Check for Password
	isValidPassword, res := helpers.CheckPassword(user.Password)
	if !isValidPassword {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user.Password = string(hashedPassword)
	isValidEmail, res := helpers.VerifyEmail(user.Email)
	if !isValidEmail {
		c.JSON(http.StatusBadRequest, res)
		return
	}

	collection_name := "Users"

	collection := database.OpenCollection(database.Client, collection_name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "User created successfully")

}
