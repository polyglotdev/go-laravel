package celeritas

import "os"

// CreateDirIfNotExist is a method on the Celeritas struct.
// It checks if a directory at the provided path exists.
// If the directory does not exist, it creates the directory with the specified mode.
// The mode is set to 0755, which means the owner can read, write, and execute,
// while others can read and execute but not write.
//
// Parameters:
// path: A string representing the path of the directory to check or create.
//
// Returns:
// If the directory exists or is successfully created, it returns nil.
// If there is an error creating the directory, it returns the error.
func (c *Celeritas) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateFileIfNotExist is a method on the Celeritas struct.
// It checks if a file at the provided path exists.
// If the file does not exist, it creates the file with the specified mode.
// The mode is set to 0644, which means the owner can read and write,
// while others can only read.
//
// Parameters:
// path: A string representing the path of the file to check or create.
//
// Returns:
// If the file exists or is successfully created, it returns nil.
// If there is an error creating the file, it returns the error.
func (c *Celeritas) CreateFileIfNotExist(path string) error {
	const mode = 0644
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
	}
	return nil
}
