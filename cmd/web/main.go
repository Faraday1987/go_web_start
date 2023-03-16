package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Faraday1987/fgoutils/pkg/config"
	"github.com/Faraday1987/fgoutils/pkg/handlers"
	"github.com/Faraday1987/fgoutils/pkg/render"
)

const PORT = ":8888"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Server run on port: %v ", PORT)

	_ = http.ListenAndServe(PORT, nil)
}
