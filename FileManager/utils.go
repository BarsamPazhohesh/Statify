package FileManager

import (
	"os"
	"path/filepath"
)

// GetAbsolutePath converts a relative or absolute file path to its absolute form.
//
// Arguments:
//   - path: The file path, either relative or absolute.
//
// Returns:
//   - string: The absolute path.
//   - error: An error if conversion fails.
func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

// GetRelativePath computes the relative file path from a given root directory to the target path.
//
// Arguments:
//   - root: The root directory to calculate the relative path from.
//   - path: The target file or directory path.
//
// Returns:
//   - string: The relative path from the root directory to the target path.
//   - error: An error if the computation fails.
func GetRelativePath(root string, path string) (string, error) {
	return filepath.Rel(root, path)
}

// IsFileExists checks whether a file or directory exists at the specified path.
//
// Arguments:
//   - filename: The file or directory path to check for existence.
//
// Returns:
//   - bool: `true` if the file or directory exists, `false` otherwise.
func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
