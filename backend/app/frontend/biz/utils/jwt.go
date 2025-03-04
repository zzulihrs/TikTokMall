package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserId     uint32 `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	MerchantId int64  `json:"merchant_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId int64, email string, username string, merchantId int64) (tokenString string, err error) {
	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &JWTClaim{
		UserId:     uint32(userId),
		Email:      email,
		Username:   username,
		MerchantId: merchantId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
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
