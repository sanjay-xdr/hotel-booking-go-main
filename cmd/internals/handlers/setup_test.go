package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/justinas/nosurf"
	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/models"
	"github.com/sanjay-xdr/cmd/internals/render"
)

var app config.AppConfig
var sessionManager *scs.SessionManager
var pathToTemplates="./../../templates"


func getRoutes() http.Handler{



	
	gob.Register(models.Reservation{})

	tc, err := CreateTestTemplateCache()

	if err != nil {
		log.Fatal(err)
		
	}

	app.TemplateCache = tc
	app.UseCache = true
	app.InProduction = false
	render.NewTemplate(&app)

	repo := NewRepo(&app)

	NewHandlers(repo)
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	mux := chi.NewRouter()
	// mux.Use(middleware.Logger)
	mux.Use(NoSruve)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)
	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

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


func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html",pathToTemplates))

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// fmt.Println(page)
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// fmt.Println(ts)

		matches, err :=  filepath.Glob(fmt.Sprintf("%s/*.page.html",pathToTemplates)) //checking for layout
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html",pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}

	return myCache, nil

}
