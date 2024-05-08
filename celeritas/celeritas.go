package celeritas

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
	return nil
}

func (c Celeritas) Init(p initPaths) error {
	//	get the root path
	root := p.rootPath
	// check if folders exist
	for _, folderPath := range p.folderNames {
		// create folderPath if it doesn't exist
		err := c.CreateDirIfNotExist(root + "/" + folderPath)
		if err != nil {
			return err
		}
	}
	return nil
}
