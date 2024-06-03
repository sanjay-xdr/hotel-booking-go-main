package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/driver"
	"github.com/sanjay-xdr/cmd/internals/handlers"
	"github.com/sanjay-xdr/cmd/internals/models"
	"github.com/sanjay-xdr/cmd/internals/render"
)

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	dbConn, err := run()

	if err != nil {
		log.Fatal("Unable to run the project", err)
	}

	defer dbConn.SQL.Close()

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

func run() (*driver.DB, error) {

	gob.Register(models.Reservation{})

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false
	render.NewTemplates(&app)

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	log.Println("Connecting to the Database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=sanjay")

	if err != nil {
		log.Fatal("Can not connect to Database")
	}
	log.Println("Connected to Database finally.....")
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	return db, nil
}
