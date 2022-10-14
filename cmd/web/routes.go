package main

import (
	"net/http"

	"github.com/AbahEkang/bookings/pkg/Config"
	"github.com/AbahEkang/bookings/pkg/handlers"
	//"github.com/bmizerany/pat"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *Config.AppConfig) http.Handler{

	// Using 'pat'
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Use(SessionLoad)


	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

