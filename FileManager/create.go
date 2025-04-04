package FileManager

import (
	"fmt"
	"os"
)

// CreateDirectories recursively creates all directories in the given path string.
// Args:
//   - path: The string representing the full path to the directory structure you want to create.
//
// Returns:
//   - error: An error if any directory in the path cannot be created, nil otherwise.
func CreateDirectories(path string) error {
	err := os.MkdirAll(path, 0755) // 0755 permissions: owner rwx, others rx
	if err != nil {
		return fmt.Errorf("failed to create directories at '%s': %w", path, err)
	}
	return nil
}
