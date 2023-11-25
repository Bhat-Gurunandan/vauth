package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"vauth/config"
	"vauth/mcookie"
	"vauth/mjwt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email       string
	Role        string
	AccessToken string
}

const ISSUER = "https://auth.mydomain.com"
const SUBJECT = "Authentication"

var allUsers = map[string]string{
	"user1@example.com": "password123",
	"user2@example.com": "password456",
	"user3@example.com": "password789",
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password, ok := allUsers[user.Email]
	if !ok {
		http.Error(w,
			fmt.Sprintf("cannot find user %s", user.Email),
			http.StatusUnauthorized,
		)
		return
	}

	if password != user.Password {
		http.Error(w,
			fmt.Sprintf("password for user %s does not match", user.Email),
			http.StatusUnauthorized,
		)
		return
	}

	sessionID := uuid.New().String()
	accessToken, refreshToken, err := user.makeTokens(sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookieVal := map[string]string{
		"id":           sessionID,
		"refreshToken": refreshToken,
	}
	cookie, err := mcookie.NewSecureCookie(cookieVal, time.Now().Add(time.Hour*6), false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(&LoginResponse{
		Email:       user.Email,
		Role:        "admin",
		AccessToken: accessToken,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, cookie)

	w.Write(jsonBytes)

}

func (u *User) makeTokens(cookieID string) (string, string, error) {

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Minute * 15)

	registeredClaims := jwt.RegisteredClaims{
		ID:        cookieID,
		Issuer:    ISSUER,
		Subject:   SUBJECT,
		IssuedAt:  &jwt.NumericDate{Time: issuedAt},
		ExpiresAt: &jwt.NumericDate{Time: expiresAt},
	}

	userClaims := mjwt.UserClaims{
		Role:             "admin",
		FirstName:        "Gurunandan",
		LastName:         "Bhat",
		RegisteredClaims: registeredClaims,
	}

	cfg, err := config.Configuration()
	if err != nil {
		return "", "", fmt.Errorf("error fetching configuration: %s", err)
	}

	accessToken, err := mjwt.NewAccessToken(cfg.JWTSecret, userClaims)
	if err != nil {
		return "", "", fmt.Errorf("error generating access token: %s", err)
	}

	refreshToken, err := mjwt.NewRefreshToken(cfg.JWTSecret, registeredClaims)
	if err != nil {
		return "", "", fmt.Errorf("error generating refresh token: %s", err)
	}

	return accessToken, refreshToken, nil
}
