package auth

import (
	"github.com/Darthex/ink-golang/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateJWT(userId int) (string, error) {
	secret := config.Envs.JWTSecretKey
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
