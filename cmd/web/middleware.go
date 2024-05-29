package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Adds CSRF protection to all POST request
func NoSruve(next http.Handler) http.Handler {

	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{

		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfhandler

}

// Loads and save the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
