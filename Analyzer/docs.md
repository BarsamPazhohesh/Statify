# Analyzer Documentation
The **analyzer** package provides tools for analyzing source code files to extract useful insights like file language, total size, comment size, code size, and blank line count. It supports analyzing multiple files at once or individual files.

## Types

### FileMetadata
The structure that holds metadata for a file.

#### Fields:
- `Name` (string): The name of the file.
- `Path` (string): The absolute path of the file.
- `Dir` (string): The directory containing the specified file.
- `Extension` (string): The file extension.
- `Size` (int64): The size of the file in bytes.
- `ModifiedAt` (time.Time): The last modified time of the file.

### AnalyzeFileResult  
- **AnalyzeFileResult:** Represents the result of analyzing a single file, containing:

#### Fields:
- `FileMetadata`: File metadata.
- `Language`: The programming language of the file.
- `TotalSize`: The total size of the file. (count utf8 char).
- `CommentSize`: The size of the comments in the file (count utf8 char).
- `CodeSize`: The size of the actual code in the file (count utf8 char).
- `BlankLines`: The number of blank lines in the file.


## **Functions:**

#### AnalyzeSingleFile
- **AnalyzeSingleFile(metadata FileMetadata) (AnalyzeFileResult, error):**
  Analyzes a single file to determine:
  - The language of the file.
  - The total size of the file.
  - The size of the comments.
  - The size of the code (excluding comments and blank lines).
  - The number of blank lines in the file.
  
  **Arguments:**
  - `metadata`: A `FileMetadata` object that contains details about the file (path, extension).
  
  **Returns:**
  - `AnalyzeFileResult`: The result of the analysis containing details about file size, comment size, code size, and blank lines.
  - `error`: If an error occurs during file reading or analysis, an error is returned.
  
  **Example:**
  ```go
  metadata := filemanager.FileMetadata{Path: "example.go"}
  result, err := analyzer.AnalyzeSingleFile(metadata)
  if err != nil {
      fmt.Println("Error analyzing file:", err)
  } else {
      fmt.Println("Code Size:", result.CodeSize)
      fmt.Println("Comment Size:", result.CommentSize)
      fmt.Println("Blank Lines:", result.BlankLines)
  }
  ```

---
#### AnalyzeMultipleFiles
- **AnalyzeMultipleFiles(files []FileMetadata) ([]AnalyzeFileResult, error):**
  Analyzes multiple files and returns the analysis results for each valid file.

  **Arguments:**
  - `files`: A slice of `FileMetadata` representing the files to be analyzed.
  
  **Returns:**
  - `[]AnalyzeFileResult`: A slice of analysis results for each valid file.
  - `error`: If any file reading or analysis fails, an error is returned.
  
  **Example:**
  ```go
  files := []filemanager.FileMetadata{
      {Path: "example.go",extension:".go", Name:"example", ...},
      {Path: "test.py",extension:".py", Name:"test", ...},
  }
  
  results, err := analyzer.AnalyzeMultipleFiles(files)
  if err != nil {
      fmt.Println("Error analyzing files:", err)
  } else {
      for _, result := range results {
          fmt.Printf("File: %s, Code Size: %d, Comment Size: %d\n", result.FileMetadata.Path, result.CodeSize, result.CommentSize)
      }
  }
  ```
