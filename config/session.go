package config

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func SayHelloToSession() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * 90 * time.Hour
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.Partitioned = true
}

func GetSessionManager()*scs.SessionManager{
	return sessionManager
}
