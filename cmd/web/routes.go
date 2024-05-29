package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	log.Print("Routing now")
	mux := chi.NewRouter()
	// mux.Use(middleware.Logger)
	mux.Use(NoSruve)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
