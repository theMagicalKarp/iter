package predicates

import (
	"os"
)

// IsFile checks if the given path is a file.
// It returns true if the path is a file, false otherwise.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// IsDir checks if the given path is a directory.
// It returns true if the path is a directory, false otherwise.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

// PathExists checks if the given path exists.
// It returns true if the path exists, and false otherwise.
func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
