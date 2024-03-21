package predicates_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func TestIsFile(t *testing.T) {
	t.Parallel()

	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	// Test with a file path
	isFile := predicates.IsFile(file.Name())
	assert.True(t, isFile)

	// Test with a directory path
	isFile = predicates.IsFile(".")
	assert.False(t, isFile)
}

func TestIsDir(t *testing.T) {
	t.Parallel()

	// Test with a directory path
	isDir := predicates.IsDir(".")
	assert.True(t, isDir)

	// Test with a file path
	file, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	isDir = predicates.IsDir(file.Name())
	assert.False(t, isDir)
}

func TestPathExists(t *testing.T) {
	t.Parallel()

	file, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	exists := predicates.PathExists(file.Name())
	assert.True(t, exists)

	// Test with a non-existing file path
	exists = predicates.PathExists("non_existing_file.txt")
	assert.False(t, exists)

	// Test with an existing directory path
	exists = predicates.PathExists(".")
	assert.True(t, exists)

	// Test with a non-existing directory path
	exists = predicates.PathExists("non_existing_directory")
	assert.False(t, exists)
}
