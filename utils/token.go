package utils

import (
	"fmt"
	"pujaprabha/internal/presenter"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func CreateToken(userID uint, username string) (string, error) {
	//Creating Access Token
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 90).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", nil
	}
	return token, err

}

func RefreshToken(userID uint, username string) (string, error) {
	//creating refresh token
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", nil
	}

	return token, err

}

func ExtractToken(c *fiber.Ctx) string {
	bToken := c.Get("Authorization")

	strArr := strings.Split(bToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func ParseToken(tokenString string, tokenType string) (*presenter.JwtCustomClaims, error) {
	var secret string
	if tokenType == "access" {
		secret = viper.GetString("jwt.secret")
	}

	token, err := jwt.ParseWithClaims(tokenString, &presenter.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*presenter.JwtCustomClaims); ok && token.Valid {

		return claims, nil
	}
	return nil, err
}
