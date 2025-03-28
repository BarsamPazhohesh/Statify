package filemanager

import (
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
