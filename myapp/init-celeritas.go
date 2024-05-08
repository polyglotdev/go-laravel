package main

import (
	"log"
	"os"

	"github.com/polyglotdev/celeritas"
)

// initApplication is a function that initializes the application.
func initApplication() *application {
	// Get the current working directory.
	path, err := os.Getwd()
	// If there's an error, log it and stop the program.
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a new Celeritas object.
	cel := &celeritas.Celeritas{}
	// Call the New method on the Celeritas object, passing in the current working directory.
	// The New method is responsible for setting up the initial directory structure for the application.
	if err := cel.New(path); err != nil {
		// If there's an error, log it and stop the program.
		log.Fatal(err)
	}

	cel.AppName = "myapp"
	cel.Debug = true

	// Return a new application object, with the Celeritas object embedded in it.
	return &application{
		App: cel,
	}
}
