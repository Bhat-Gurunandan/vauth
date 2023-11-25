package main

import (
	"fmt"
	"log"
	"time"
	"vauth/config"
	"vauth/mjwt"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
}

func main() {

	userClaims := mjwt.UserClaims{
		Role:      "admin",
		FirstName: "Gurunandan",
		LastName:  "Bhat",
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "Some UUID",
			Issuer:    "admin.mariodemiranda.com",
			Subject:   "Authentication",
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 15)},
		},
	}

	cfg, err := config.Configuration()
	if err != nil {
		log.Fatalf("Error fetching configuration: %s\n", err)
	}

	signedToken, err := mjwt.NewAccessToken(cfg.JWTSecret, userClaims)
	if err != nil {
		log.Fatalf("Error generating token: %s\n", err)
	}

	fmt.Println("Token: ", signedToken)

	userClaims = *mjwt.ParseAccessToken(cfg.JWTSecret, signedToken)
	fmt.Printf("Claims: %+v\n", userClaims)
}
