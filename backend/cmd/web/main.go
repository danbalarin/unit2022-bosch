package main

import (
	"github.com/unit2022-bosch/teapot/backend/internal/app/config"
	"log"
)

func main() {
	if err := config.LoadEnvVariables(); err != nil {
		log.Panicln(err)
	}

	app, err := BuildWebServer()
	if err != nil {
		log.Panicln(err)
	}

	if err := app.Start(); err != nil {
		log.Panicln(err)
	}
}
