package celeritas

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/polyglotdev/celeritas/render"
)

const (
	// Version is the current version of the Celeritas framework.
	Version = "1.0.0"
)

// Celeritas is the main struct for the Celeritas framework.
type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	config   config
}

type config struct {
	port     string
	renderer string
}

// New creates the initial directory structure for a new Celeritas project.
func (c *Celeritas) New(rootPath string) error {
	// create the initial directory structure
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middleware",
		},
	}
	// initialize the directory structure and check for errors.
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env file
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// create loggers
	infoLog, errLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = Version
	c.RootPath = rootPath
	c.Routes = c.routes().(*chi.Mux)

	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	c.Render = c.createRender(c)

	return nil
}

// Init initializes the directory structure for a Celeritas project.
// It takes an initPaths struct as an argument which contains the root path and the names of the folders to be created.
// It iterates over the folder names, and for each one, it calls the CreateDirIfNotExist method.
// If the CreateDirIfNotExist method returns an error, it immediately returns this error.
// If no errors occur during the folder creation, it returns nil.
func (c *Celeritas) Init(p initPaths) error {
	// Get the root path from the initPaths struct.
	root := p.rootPath
	// Iterate over the folder names in the initPaths struct.
	for _, folderPath := range p.folderNames {
		// For each folder name, call the CreateDirIfNotExist method with the full path.
		// If an error occurs, return it immediately.
		err := c.CreateDirIfNotExist(root + "/" + folderPath)
		if err != nil {
			return err
		}
	}
	// If no errors occurred during the folder creation, return nil.
	return nil
}

// ListenAndServe starts the HTTP server and listens for incoming requests.
// It configures the server with the provided settings and routes,
// and logs any errors that occur during server startup or shutdown.
func (c *Celeritas) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     c.ErrorLog,
		Handler:      c.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	c.InfoLog.Printf("Starting 🚀 on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	if err != nil {
		c.ErrorLog.Fatalf("Error starting server: %v", err)
	}

	c.InfoLog.Printf("Server stopped")
}

// checkDotEnv checks if the .env file exists in the root path of the project.
// If the file does not exist, it creates a new .env file with the default values.
func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

// startLoggers initializes and returns two loggers: an info logger and an error logger.
// The info logger is used for general logging of information, while the error logger is used for logging errors.
// Both loggers write to standard output and include the date and time in their output.
// The error logger also includes the file name and line number where the log call was made.
// The rootPath parameter is currently unused.
func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errLog
}

func (c *Celeritas) createRender(cel *Celeritas) *render.Render {
	myRenderer := render.Render{
		Renderer: cel.config.renderer,
		RootPath: cel.RootPath,
		Port:     cel.config.port,
	}

	return &myRenderer
}
