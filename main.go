package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"statfiy/Analyzer"
	"statfiy/ArgManager"
	"statfiy/FileManager"
)

// Usage examples:
//  1. To specify multiple root paths for analysis, use the `-p` flag:
//     `go run . -p /path/to/root1 -p /path/to/root2 -p /path/to/root3 ...`
//  2. To include comments in the analysis, use the `-ic` flag:
//     `go run . -ic -p /path/to/root1 -p /path/to/root2`
//  3. To specify an output path for the results (images and markdown files), use the `-op` flag:
//     `go run . -op /path/to/output -p /path/to/root1 -p /path/to/root2`
//  4. To view the help message, use the `-h` flag:
//     `go run . -h`
func main() {
	args, err := ArgManager.ParseArgs(os.Args)
	if err != nil {
		log.Fatalf("Error: Cannot parse args: %v", err)
	}

	if !args.OutputPath.IsSet {
		args.OutputPath.Value = os.Args[0]
	}

	// Iterate over provided paths
	for _, path := range args.RootPaths {
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
