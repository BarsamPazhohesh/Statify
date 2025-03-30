package FileManager

import (
	"io/fs"
	"os"
	"path/filepath"
)

// GetFileMetadata retrieves metadata for a given file path.
//
// Arguments:
//   - filePath: The path to the file.
//
// Returns:
//   - FileMetadata: Contains file name, absolute path, extension, size, and last modified time.
//   - error: An error if the file cannot be accessed.
func GetFileMetadata(filePath string) (FileMetadata, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return FileMetadata{}, err
	}

	absolutePath, err := GetAbsolutePath(filePath)
	if err != nil {
		return FileMetadata{}, err
	}

	return FileMetadata{
		Name:       info.Name(),
		Path:       absolutePath,
		Extension:  filepath.Ext(absolutePath),
		Size:       info.Size(),
		ModifiedAt: info.ModTime(),
	}, nil
}

// CollectFilesMetadata walks through a directory and collects metadata for all files.
//
// Arguments:
//   - rootDir: The directory path to scan for files.
//
// Returns:
//   - []FileMetadata: A slice containing metadata of all discovered files.
//   - error: An error if directory traversal fails.
func CollectFilesMetadata(rootDir string) ([]FileMetadata, error) {
	var fileList []FileMetadata

	err := filepath.Walk(rootDir, func(filePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileMeta, err := GetFileMetadata(filePath)
		if err != nil {
			return nil
		}

		fileList = append(fileList, fileMeta)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}
