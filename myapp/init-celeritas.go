package main

import (
	"log"
	"os"

	"github.com/polyglotdev/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := celeritas.Celeritas{}
	if err := cel.New(path); err != nil {
		log.Fatal(err)
	}

	return &application{
		App: &cel,
	}
}
