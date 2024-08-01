package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64,error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0,errors.New("could not parse token")
	}
	tokenIsValid := token.Valid
	if !tokenIsValid {
		return 0,errors.New("invalid token")
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0,errors.New("could not parse token claims")
	}
	// email := tokenClaims["email"].(string)
	userId := int64(tokenClaims["userId"].(float64))
	
	return userId,nil
}
