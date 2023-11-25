package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Email    string
	Password string
}

type users map[string]string

var allUsers = users{
	"gbhat@pobox.com": "password123",
	"anita@pobox.com": "password456",
	"suman@pobox.com": "password789",
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

	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonBytes)

}
