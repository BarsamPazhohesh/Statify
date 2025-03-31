package main

import (
	"fmt"
	"log"
	"path/filepath"
	"statfiy/Analyzer"
	"statfiy/ArgManager"
	"statfiy/FileManager"
)

// Usage examples:
// 1. `go run . -p rootpath1 -p rootpath2 -p rootpath3 ...`
// 2. To include comments, use the `-ic` flag: `go run . -ic -p rootpath1 -p rootpath2 -p rootpath3`
// 3. For Help use go run . -h
func main() {
	// Parse the arguments
	argsManager := &ArgManager.Args{}
	args, err := argsManager.ParseArgs()
	if err != nil {
		log.Fatalf("Error: Cannot parse args: %v", err)
	}

	// Check if root paths are provided
	if !args.RootPaths.IsArgProvided {
		log.Fatal("Please provide root paths using the '--paths' flag.")
	}

	// Iterate over provided paths
	for _, path := range args.RootPaths.Value {
		absPath, err := FileManager.GetAbsolutePath(path)
		if err != nil {
			log.Fatalf("error path is invalid: %v", err.Error())
		}

		outputPath := fmt.Sprintf("%v.md", filepath.Base(absPath))
		// Clear the output file
		if err := FileManager.OverwriteFile(outputPath, []byte{}); err != nil {
			log.Println("Error clearing output file:", err)
		}
		// Collect metadata from files in the specified codebase
		files, err := FileManager.CollectFilesMetadata(path)
		if err != nil {
			log.Fatal(err)
		}

		// Analyze the collected files
		analyzedFiles, err := Analyzer.AnalyzeMultipleFiles(files)
		if err != nil {
			log.Fatal(err)
		}

		// Generate and write analysis report for each file
		for _, file := range analyzedFiles {
			report := fmt.Sprintf(`## %v

| Property      | Value      |
|---------------|------------|
| File Name     | %v         |
| File Path     | %v         |
| Language      | %v         |
| Total Size    | %v bytes   |
| Code Size     | %v bytes   |
| Comment Size  | %v bytes   |
| Blank Lines   | %v         |
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

			// Write the report to the markdown file
			if err := FileManager.AppendFileString(outputPath, report); err != nil {
				log.Println("Error appending to file:", err)
			}
		}
	}
}
