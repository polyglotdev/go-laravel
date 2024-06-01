package main

import (
	"github.com/polyglotdev/celeritas"

	"github.com/polyglotdev/myapp/handlers"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
