package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AbahEkang/bookings/pkg/Config"
	"github.com/AbahEkang/bookings/pkg/handlers"
	"github.com/AbahEkang/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app Config.AppConfig
var session *scs.SessionManager

func main() {

	fmt.Println("Talkin' 'bout templates!")

	fmt.Println("Starting application on port", portNumber)


	// Change this to true when in production
	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil{
		log.Fatal("connot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	

	srv := &http.Server{
		Addr: portNumber,

		Handler: routes(&app),

	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}