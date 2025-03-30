package FileManager

import (
	"fmt"
	"os"
)

var errorMessages map[string]string

func init() {
	errorMessages = make(map[string]string)
	errorMessages["OpenToAppend"] = "failed to open file for appending"
	errorMessages["Append"] = "failed to append data to file"
	errorMessages["OpenToOverwrite"] = "failed to open file for overwriting"
	errorMessages["Write"] = "failed to write data to file"
}

// OverwriteFile writes data to a file, overwriting it if it exists.
//
// Arguments:
// - filePath: The name (path) of the file to write to.
// - data: The data to write into the file.
//
// Returns:
// - nil if the file is successfully written to.
// - An error if the file creation or writing fails.
func OverwriteFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for overwriting: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

// OverwriteFileString writes string data to a file, overwriting it if it exists.
//
// Arguments:
// - filePath: The name (path) of the file to write to.
// - data: The string data to write into the file.
//
// Returns:
// - nil if the file is successfully written to.
// - An error if the file creation or writing fails.
func OverwriteFileString(filePath string, data string) error {
	return OverwriteFile(filePath, []byte(data))
}

// AppendFile writes data to a file, appending it if the file already exists.
//
// Arguments:
// - filePath: The path of the file to write to.
// - data: The data to append to the file.
//
// Returns:
// - nil if the file is successfully written to.
// - An error if the file creation or writing fails.
func AppendFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for appending: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to append data to file: %w", err)
	}

	return nil
}

// AppendFileString writes string data to a file, appending it if the file already exists.
//
// Arguments:
// - filePath: The path of the file to write to.
// - data: The string data to append to the file.
//
// Returns:
// - nil if the file is successfully written to.
// - An error if the file creation or writing fails.
func AppendFileString(filePath string, data string) error {
	return AppendFile(filePath, []byte(data))
}

func ErrorHandler(err error, errorMessage string) error {
	if err != nil {
		return fmt.Errorf("%v: %w", errorMessage, err)
	}
	return nil
}
