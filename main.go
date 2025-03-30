package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"statfiy/analyzer"
	"statfiy/filemanager"
)

// Runs the program like: go run . --path sourceCode
// If no path is provided, it defaults to the executable's directory.
func main() {
	defaultPath := os.Args[0]

	// Define a flag for specifying the codebase path
	codebasePath := flag.String("path", defaultPath, "Path to the codebase for analysis")
	flag.Parse()

	// Collect metadata from files in the specified codebase
	files, err := filemanager.CollectFilesMetadata(*codebasePath)
	if err != nil {
		log.Fatal(err)
	}

	// Analyze the collected files
	analyzedFiles, err := analyzer.AnalyzeMultipleFiles(files)
	if err != nil {
		log.Fatal(err)
	}

	outputPath := "Info.md"

	// overwrite file with empty value
	if err := filemanager.OverwriteFile(outputPath, []byte{}); err != nil {
		log.Println("Error writing to file:", err)
	}
	// Generate and write analysis report for each file
	for _, file := range analyzedFiles {
		report := fmt.Sprintf(`## %v

| Property      | Value      |
|--------------|-----------|
| File Name    | %v        |
| File Path    | %v        |
| Language     | %v        |
| Total Size   | %v bytes  |
| Code Size    | %v bytes  |
| Comment Size | %v bytes  |
| Blank Lines  | %v        |

`,
			file.FileMetadata.Name,
			file.FileMetadata.Name,
			file.FileMetadata.Path,
			file.Language,
			file.TotalSize,
			file.CodeSize,
			file.CommentSize,
			file.BlankLines,
		)

		// Write report to the markdown file
		if err := filemanager.AppendFileString(outputPath, report); err != nil {
			log.Println("Error writing to file:", err)
		}
	}
}
