package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shwethadia/BOOKINGAPI/internal/config"
	"github.com/shwethadia/BOOKINGAPI/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/laptop", handlers.Repo.Laptops)
	mux.Get("/dongle", handlers.Repo.Dongals)
	mux.Get("/allort-hardware",handlers.Repo.AllortHardware)
	mux.Post("/allort-hardware",handlers.Repo.PostAllortHardware)
	mux.Post("/allort-hardware-json",handlers.Repo.AllortHardwareJSON)
	mux.Get("/contact",handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
