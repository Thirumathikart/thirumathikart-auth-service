package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
	"gorm.io/gorm"
)

type CustomClaims struct {
	Email     string `json:"email"`
	ContactNo string `json:"contact"`
	jwt.StandardClaims
}

func CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetCurrentUser(c echo.Context) (models.User, error) {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.User{}, err
	}
	jsonBody := make(map[string]interface{})
	errM := json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(bodyBytes))).Decode(&jsonBody)
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if errM != nil {
		return models.User{}, fmt.Errorf(errM.Error())
	}

	token, err := jwt.ParseWithClaims(fmt.Sprint(jsonBody["user_token"]), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return models.User{}, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		var user models.User
		err = config.GetDB().First(&user, "email = ?", claims.Email).Error
		return user, err
	} else {
		return models.User{}, fmt.Errorf("invalid token")
	}
}

func GetCurrentUserFromToken(userToken string, db *gorm.DB) (models.User, error) {
	var user models.User
	token, err := jwt.ParseWithClaims(userToken, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.JwtSecret), nil
		})
	if err != nil {
		return user, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		claimsEmail := claims.Email
		log.Println("USER_EMAIL", claimsEmail)
		err = db.First(&user, "email = ?", claimsEmail).Error
		log.Println(user)
		if err != nil {
			return user, err
		}
	}
	return user, nil

}
