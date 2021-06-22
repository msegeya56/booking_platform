package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wander4747/booking/pkg/config"
	"github.com/wander4747/booking/pkg/handlers"
	"github.com/wander4747/booking/pkg/render"
)

const portNumber = ":9998"

func main() {
	var app config.AppConfig

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
