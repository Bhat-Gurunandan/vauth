package mjwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

/*
	Fields that we use:
	1. jti (Unique JWT ID)
	2. sub (Subject)
	3. iat (Issued At)
	4. exp (Expires at)
*/

func NewAccessToken(secret string, claims UserClaims) (string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(secret))
}

func NewRefreshToken(secret string, claims jwt.RegisteredClaims) (string, error) {

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString([]byte(secret))
}

func ParseAccessToken(secret string, accessToken string) *UserClaims {

	parsedAccessToken, _ := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(secret string, refreshToken string) *jwt.MapClaims {

	parsedRefreshToken, _ := jwt.ParseWithClaims(
		refreshToken,
		&jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	return parsedRefreshToken.Claims.(*jwt.MapClaims)
}
