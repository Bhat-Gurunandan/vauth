package mcookie

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/securecookie"
)

const COOKIE_NAME = "mariomiranda"

var once sync.Once
var hashKey []byte
var blockKey []byte
var mstore *securecookie.SecureCookie

func NewSecureCookie(val map[string]string, expiresAt time.Time, secure bool) (*http.Cookie, error) {

	// Generate a cookie store once
	once.Do(func() {
		hashKey = securecookie.GenerateRandomKey(64)
		blockKey = securecookie.GenerateRandomKey(32)
		mstore = securecookie.New(hashKey, blockKey)
	})

	encVal, err := mstore.Encode(COOKIE_NAME, val)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "cookie-name",
		Value:    encVal,
		Path:     "/",
		Expires:  expiresAt,
		SameSite: http.SameSiteStrictMode,
		Secure:   secure,
		HttpOnly: true,
	}

	return cookie, nil
}
