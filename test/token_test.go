package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"testing"
)

func TestToken(t *testing.T) {
	type CustomClaims struct {
		UserID   int    `json:"user_id"`
		UserKey  string `json:"user_key"`
		Username string `json:"username"`
		jwt.StandardClaims
	}

	secret := []byte("@abcdefghijklmn3#opqrstuvwxyz3=2d42")

	claims := CustomClaims{
		//UserID:  96,
		UserKey:  "82d9ff7f-840f-4a39-a604-69e2bfe1f200",
		Username: "admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JWT Token:", signedToken)

}
