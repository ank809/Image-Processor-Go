package helpers

import (
	"math/rand"
	"time"
	"unicode"

	emailVerifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailVerifier.NewVerifier()
)

func CheckName(name string) (bool, string) {
	if name == "" {
		return false, "Name cannot be empty"
	}
	if len(name) < 3 {
		return false, "Length of name should be greater than 3"
	}
	return true, ""
}

func CheckPassword(password string) (bool, string) {
	if password == "" {
		return false, "Password cannot be empty"
	}
	if len(password) < 6 {
		return false, "Length of password should be greater than 6"
	}

	containsUpper := false
	containsLower := false
	containsDigits := false
	containsSpecialCharacters := false

	for _, ch := range password {
		if unicode.IsDigit(ch) {
			containsDigits = true
		} else if unicode.IsLower(ch) {
			containsLower = true
		} else if unicode.IsUpper(ch) {
			containsUpper = true
		} else {
			containsSpecialCharacters = true
		}
	}
	if containsDigits && containsLower && containsSpecialCharacters && containsUpper {
		return true, "Password is valid"
	} else {
		return false, "Password should contains uppercase, lowercase, digits and special characters"
	}
}
func VerifyEmail(email string) (bool, string) {
	res, err := verifier.Verify(email)
	if err != nil {
		return false, "verify email address failed"
	}
	if !res.Syntax.Valid {
		return false, "Invalid email syntax"
	}
	return true, "Valid email"
}

func GetUniqueKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
