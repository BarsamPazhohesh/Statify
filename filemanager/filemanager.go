package filemanager

import (
	"bufio"
	"fmt"
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

// LineHandler is a callback function type for processing each line of text.
type LineHandler func(line string) error

// ProcessFileByLine reads a file line by line and applies a callback function to each line.
//
// Arguments:
//   - filePath: The path to the file to be read.
//   - handler: A callback function to process each line.
//
// Returns:
//   - error: If reading or processing fails, an error is returned. Otherwise, nil.
func ProcessFileByLine(filePath string, handler LineHandler) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := handler(scanner.Text()); err != nil {
			return fmt.Errorf("error processing line: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// ReadEntireFile reads the entire contents of a file and returns it as a string.
//
// Arguments:
//   - filePath: The path to the file to be read.
//
// Returns:
//   - string: The full content of the file.
//   - error: An error if reading the file fails.
func ReadEntireFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// CreateOrTruncateFile opens a file for reading and writing, truncating it if it exists.
//
// Arguments:
//   - filename: The name of the file to create or truncate.
//
// Returns:
//   - *os.File: A pointer to the opened file.
//   - error: An error if the file cannot be created or opened.
func CreateOrTruncateFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
}
