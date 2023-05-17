package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MetaAuth struct {
	ID            string
	Username      string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaAuth
}

func Login(Data map[string]interface{}, SecretKeyEnv string, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()
	jwtSecretKey := GoDotEnv(SecretKeyEnv)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := to.SignedString([]byte(jwtSecretKey))
	if err != nil {
		logrus.Error(err.Error())
		return token, err
	}

	return token, nil
}

func VerifyHeaderToken(ctx *gin.Context, SecretKeyEnv string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.Split(tokenHeader, "Bearer")[1]
	jwtSecretKey := GoDotEnv(SecretKeyEnv)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return token, nil
}

func VerifyJWTToken(accessToken, SecretKeyEnv string) (*jwt.Token, error) {
	jwtSecretKey := GoDotEnv(SecretKeyEnv)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return token, nil
}

func DecodeJWTToken(token *jwt.Token) AccessToken {
	var accessToken AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &accessToken)

	return accessToken
}
