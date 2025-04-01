package FileManager

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
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
		Dir:        filepath.Dir(absolutePath),
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

	err := filepath.Walk(rootDir,
		func(filePath string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fileMeta, err := GetFileMetadata(filePath)
			if err != nil {
				return err
			}

			fileList = append(fileList, fileMeta)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

// CollectFileMetadataByExtension scans a directory and collects metadata for files with specific extensions.
//
// Arguments:
//   - rootDir: The directory path to scan for files.
//   - extensions: A list of file extensions to filter.
//
// Returns:
//   - []FileMetadata: A slice containing metadata of filtered files.
//   - error: An error if directory traversal fails.
func CollectFileMetadataByExtension(rootDir string, extensions []string) ([]FileMetadata, error) {
	var fileList []FileMetadata

	err := filepath.Walk(rootDir,
		func(filePath string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			ext := filepath.Ext(filePath)
			if slices.Contains(extensions, ext) {
				if fileMeta, err := GetFileMetadata(filePath); err == nil {
					fileList = append(fileList, fileMeta)
				}
			}
			return nil
		})

	if err != nil {
		return nil, err
	}
	return fileList, nil
}
