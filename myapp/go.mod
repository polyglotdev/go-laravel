module github.com/polyglotdev/myapp

go 1.22.2

replace github.com/polyglotdev/celeritas => /Users/domhallan/learning/udemy/celeritas

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/polyglotdev/celeritas v1.0.7
)

require github.com/joho/godotenv v1.5.1 // indirect
