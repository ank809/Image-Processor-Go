package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type FileInfo struct {
	Filename string `json:"filename" binding:"required" `
	Email    string `json:"email" binding:"required"`
}
