package FileManager

import "time"

// FileMetadata represents metadata for a file
type FileMetadata struct {
	Name       string    // File name
	Path       string    // Full file path
	Extension  string    // File extension
	Size       int64     // File size in bytes
	ModifiedAt time.Time // Last modification time
}
