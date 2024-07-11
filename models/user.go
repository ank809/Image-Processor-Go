package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type FileInfo struct {
	Filename string `json:"filename" binding:"required" `
	Email    string `json:"email" binding:"required"`
}

type Operation struct {
	Key           string `json:"key" binding:"required"`
	OperationName string `json:"operation_name" binding:"required"`
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	Filename      string `json:"filename" binding:"required"`
}
