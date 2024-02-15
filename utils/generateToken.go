package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string
	Role  string
	Token string
	jwt.StandardClaims
}

var jwtKey = []byte("SECRETKEY")

func GenerateToken(email, role string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 4).Unix()
	claims := &Claims{
		Email: email,
		Role:  role,
		Token: "Access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(jwtKey))
	if err != nil {
		log.Printf("unable to generate jwt token for user %v, err: %v", email, err.Error())
		return "", err
	}

	return signedToken, nil
}
