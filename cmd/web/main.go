package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/handlers"
	"github.com/sanjay-xdr/cmd/internals/models"
	"github.com/sanjay-xdr/cmd/internals/render"
)

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	err := run()

	if err != nil {
		log.Fatal("Unable to run the project", err)
	}

	// render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{

		Addr:    ":8080",
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("SOmethign now working ", err)
	}
}

func run() error {

	gob.Register(models.Reservation{})

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	return nil
}
