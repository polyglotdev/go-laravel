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
func (c Celeritas) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}
