package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"github.com/msegeya56/booking_platform/internal/config"
	"github.com/msegeya56/booking_platform/internal/handlers"
	"github.com/msegeya56/booking_platform/internal/render"
)

const portNumber = ":9998"

var app config.AppConfig
var session *scs.SessionManager


// NoSurf add CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
    csrfHandler := nosurf.New(next)

    csrfHandler.SetBaseCookie(http.Cookie{
        HttpOnly: true,
        Path:     "/",
        Secure:   app.InProduction,
        SameSite: http.SameSiteLaxMode,
    })
    return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
    return session.LoadAndSave(next)
}





func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
