package FileManager

import "time"

// FileMetadata represents metadata for a file
type FileMetadata struct {
	Id         int
	Name       string    // File name
	Path       string    // Full file path
	Dir        string    // Directory File
	Extension  string    // File extension
	Size       int64     // File size in bytes
	ModifiedAt time.Time // Last modification time
}
