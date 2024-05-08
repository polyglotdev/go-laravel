package celeritas

import (
	"fmt"
	"github.com/joho/godotenv"
)

const (
	// Version is the current version of the Celeritas framework.
	Version = "1.0.0"
)

// Celeritas is the main struct for the Celeritas framework.
type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

// New creates the initial directory structure for a new Celeritas project.
func (c Celeritas) New(rootPath string) error {
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

	return nil
}

// Init initializes the directory structure for a Celeritas project.
// It takes an initPaths struct as an argument which contains the root path and the names of the folders to be created.
// It iterates over the folder names, and for each one, it calls the CreateDirIfNotExist method.
// If the CreateDirIfNotExist method returns an error, it immediately returns this error.
// If no errors occur during the folder creation, it returns nil.
func (c Celeritas) Init(p initPaths) error {
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

// checkDotEnv checks if the .env file exists in the root path of the project.
// If the file does not exist, it creates a new .env file with the default values.
func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}
