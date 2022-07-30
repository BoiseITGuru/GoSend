package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/eurekadao/gosend/internal/amazon"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

var kmsConfig *jwtkms.Config
var kmsAppConfig *jwtkms.Config

type JWTClaim struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	AppVersion string `json:"app_version"`
	jwt.StandardClaims
}

func Init() {
	kmsConfig = jwtkms.NewKMSConfig(amazon.Service.KMS().Client, amazon.Service.Keys.JwtKey, false)
}

func GenerateJWT(id uint, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwtkms.SigningMethodECDSA256, claims)
	tokenString, err = token.SignedString(kmsConfig.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("can not sign JWT %s", err)
	}
	return
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return kmsConfig, nil
		},
	)
	if err != nil {
		err = errors.New("couldn't parse claims")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ValidateAppToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return kmsAppConfig, nil
		},
	)
	if err != nil {
		err = errors.New("couldn't parse claims")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
