package FileManager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// LineHandler is a callback function type for processing each line of text.
type LineHandler func(line string) error

// ReadLines reads a file line by line and applies a callback function to each line.
//
// Arguments:
//   - filePath: The path to the file to be read.
//   - handler: A callback function to process each line.
//
// Returns:
//   - error: If reading or processing fails, an error is returned. Otherwise, nil.
func ReadLines(filePath string, handler LineHandler) error {
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

// MaxLinesReachedError is a custom error indicating the maximum lines were reached.
var MaxLinesReachedError = errors.New("maximum lines reached")

// ReadLinesLimit reads a file line by line, applying a callback, up to a maximum number of lines.
// It stops reading when maxLines is reached, returning nil if no other error occurred.
//
// Arguments:
//   - filePath: The path to the file to be read.
//   - maxLines: The maximum number of lines to read.
//   - handler: A callback function to process each line.
//
// Returns:
//   - error: If reading or processing fails, an error is returned. Otherwise, nil.
func ReadLinesLimit(filePath string, maxLines int, handler LineHandler) error {
	linesRead := 0
	err := ReadLines(filePath, func(line string) error {
		if linesRead >= maxLines {
			return MaxLinesReachedError
		}
		if err := handler(line); err != nil {
			return err
		}
		linesRead++
		return nil
	})

	if err == nil || errors.Is(err, MaxLinesReachedError) {
		return nil
	}

	return err
}

// ReadFileBytes reads the entire contents of a file and returns it as a []byte.
//
// Arguments:
//   - filePath: The path to the file to be read.
//
// Returns:
//   - []byte: The full content of the file.
//   - error: An error if reading the file fails.
func ReadFileBytes(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// ReadFileString reads the entire contents of a file and returns it as a string.
//
// Arguments:
//   - filePath: The path to the file to be read.
//
// Returns:
//   - string: The full content of the file.
//   - error: An error if reading the file fails.
func ReadFileString(filePath string) (string, error) {
	content, err := ReadFileBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
