# File Manager Documentation
The **FileManager** is a versatile utility designed to simplify file handling operations such as reading, writing, and processing files within a project. 

## Types

### `FileMetadata`

The structure that holds metadata for a file.

#### Fields:
- `Name` (string): The name of the file.
- `Path` (string): The absolute path of the file.
- `Dir` (string): The directory containing the specified file.
- `Extension` (string): The file extension.
- `Size` (int64): The size of the file in bytes.
- `ModifiedAt` (time.Time): The last modified time of the file.

#### Example Usage:

```go
type FileMetadata struct {
    Name       string
    Path       string
    Dir        string
    Extension  string
    Size       int64
    ModifiedAt time.Time
}
```

## Metadata Functions

### GetFileMetadata
Retrieves metadata for a given file path.

#### Arguments:
- `filePath` (string): The path to the file for which metadata is being retrieved.

#### Returns:
- `FileMetadata`: Contains the file's name, absolute path, extension, size, and last modified time.
- `error`: An error if the file cannot be accessed.

#### Example Usage:

```go
fileMeta, err := GetFileMetadata("example.txt")
if err != nil {
    fmt.Println("Error retrieving file metadata:", err)
} else {
    fmt.Println("File Size:", fileMeta.Size)
    fmt.Println("File Name:", fileMeta.Name)
    fmt.Println("File Path:", fileMeta.Name)
    fmt.Println("Last Modified:", fileMeta.ModifiedAt)
}
```

---
### CollectFilesMetadata

Walks through a directory and collects metadata for all files within it.

#### Arguments:
- `rootDir` (string): The directory path to scan for files.

#### Returns:
- `[]FileMetadata`: A slice containing metadata of all discovered files.
- `error`: An error if directory traversal fails.

#### Example Usage:

```go
filesMeta, err := CollectFilesMetadata("/path/to/directory")
if err != nil {
    fmt.Println("Error collecting files metadata:", err)
} else {
    for _, fileMeta := range filesMeta {
        fmt.Println("File:", fileMeta.Name)
        fmt.Println("Path:", fileMeta.Path)
        fmt.Println("Size:", fileMeta.Size)
    }
}
```

---

### CollectFileMetadataByExtension

Walks through a directory and collects metadata for files with specific extensions.

#### Arguments:
- `rootDir` (string): The directory path to scan for files.
- `extensions` ([]string): A list of file extensions to filter.

#### Returns:
- `[]FileMetadata`: A slice containing metadata of filtered files.
- `error`: An error if directory traversal fails.

#### Example Usage:

```go
extensions := []string{".txt", ".md"}
filesMeta, err := CollectFileMetadataByExtension("/path/to/directory", extensions)
if err != nil {
    fmt.Println("Error collecting filtered files metadata:", err)
} else {
    for _, fileMeta := range filesMeta {
        fmt.Println("File:", fileMeta.Name)
        fmt.Println("Path:", fileMeta.Path)
        fmt.Println("Size:", fileMeta.Size)
    }
}
```
---


## Read Functions

### ReadLines
Reads a file line by line and applies a callback function to each line.

#### Arguments:
- `filePath` (string): The path to the file to be read.
- `handler` (LineHandler): A callback function to process each line.

#### Returns:
- `error`: If reading or processing fails, an error is returned. Otherwise, `nil`.

#### Example Usage:

```go
err := ReadLines("example.txt", func(line string) error {
    fmt.Println("Line:", line)
    return nil
})
if err != nil {
    fmt.Println("Error reading lines:", err)
}
```

---

### ReadLinesLimit
Reads a file line by line, applying a callback, up to a maximum number of lines. Stops reading when `maxLines` is reached.

#### Arguments:
- `filePath` (string): The path to the file to be read.
- `maxLines` (int): The maximum number of lines to read.
- `handler` (LineHandler): A callback function to process each line.

#### Returns:
- `error`: If reading or processing fails, an error is returned. Otherwise, `nil`.

#### Example Usage:

```go
err := ReadLinesLimit("example.txt", 5, func(line string) error {
    fmt.Println("Line:", line)
    return nil
})
if err != nil {
    fmt.Println("Error reading lines:", err)
}
```

---

### ReadFileBytes

Reads the entire contents of a file and returns it as a `[]byte`.

#### Arguments:
- `filePath` (string): The path to the file to be read.

#### Returns:
- `[]byte`: The full content of the file.
- `error`: An error if reading the file fails.

#### Example Usage:

```go
content, err := ReadFileBytes("example.txt")
if err != nil {
    fmt.Println("Error reading file:", err)
} else {
    fmt.Println("File content:", string(content))
}
```

---

### ReadFileString
Reads the entire contents of a file and returns it as a string.

#### Arguments:
- `filePath` (string): The path to the file to be read.

#### Returns:
- `string`: The full content of the file.
- `error`: An error if reading the file fails.

#### Example Usage:

```go
content, err := ReadFileString("example.txt")
if err != nil {
    fmt.Println("Error reading file:", err)
} else {
    fmt.Println("File content:", content)
}
```


## File Writing Functions 

### OverwriteFile
Writes data to a file, overwriting the file if it already exists.

#### Arguments:
- `filePath` (string): The absolute or relative path of the file to write to.
- `data` ([]byte): The byte slice containing the data to write into the file.

#### Returns:
- `nil` if the file is successfully written to.
- An error if the file creation or writing fails (e.g., permission issues, invalid path).

#### Example:
```go
data := []byte("This is the data to overwrite the file.")
err := filemanager.OverwriteFile("example.txt", data)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("File overwritten successfully.")
}
```

---

### OverwriteFileString
Writes string data to a file, overwriting the file if it already exists. This function is a variant of `OverwriteFile` that accepts a string instead of a byte slice.

#### Arguments:
- `filePath` (string): The absolute or relative path of the file to write to.
- `data` (string): The string data to write into the file.

#### Returns:
- `nil` if the file is successfully written to.
- An error if the file creation or writing fails (e.g., permission issues, invalid path).

#### Example:
```go
data := "This is the string data to overwrite the file."
err := filemanager.OverwriteFileString("example.txt", data)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("File overwritten successfully.")
}
```

---
### AppendFile
Writes data to a file, appending it to the file if the file already exists. If the file does not exist, it will be created.

#### Arguments:
- `filePath` (string): The absolute or relative path of the file to write to.
- `data` ([]byte): The byte slice containing the data to append to the file.

#### Returns:
- `nil` if the file is successfully written to.
- An error if the file creation or writing fails (e.g., permission issues, invalid path).

#### Example:
```go
data := []byte("This is the data to append to the file.")
err := filemanager.AppendFile("example.txt", data)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Data appended successfully.")
}
```

---

### AppendFileString
Writes string data to a file, appending it to the file if it already exists. This function is a variant of `AppendFile` that accepts a string instead of a byte slice.

#### Arguments:
- `filePath` (string): The absolute or relative path of the file to write to.
- `data` (string): The string data to append to the file.

#### Returns:
- `nil` if the file is successfully written to.
- An error if the file creation or writing fails (e.g., permission issues, invalid path).

#### Example:
```go
data := "This is the string data to append to the file."
err := filemanager.AppendFileString("example.txt", data)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Data appended successfully.")
}
```
## Utility Functions

### GetAbsolutePath
Converts a relative or absolute file path to its absolute form.

#### Arguments:
- `path` (string): The file path, either relative or absolute.

#### Returns:
- `string`: The absolute path.
- `error`: An error if conversion fails (e.g., invalid path).

#### Example:
```go
path := "relative/path/to/file.txt"
absPath, err := filemanager.GetAbsolutePath(path)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Absolute path:", absPath)
}
```

---

### GetRelativePath
Computes the relative file path from a given root directory to the target path.

#### Arguments:
- `root` (string): The root directory to calculate the relative path from.
- `path` (string): The target file or directory path.

#### Returns:
- `string`: The relative path from the root directory to the target path.
- `error`: An error if the computation fails (e.g., if the root is not a valid directory).

#### Example:
```go
root := "/home/user"
targetPath := "/home/user/documents/file.txt"
relPath, err := filemanager.GetRelativePath(root, targetPath)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Relative path:", relPath)
}
```

---

### IsFileExists
Checks whether a file or directory exists at the specified path.

#### Arguments:
- `filename` (string): The file or directory path to check for existence.

#### Returns:
- `bool`: `true` if the file or directory exists, `false` otherwise.

#### Example:
```go
filePath := "example.txt"
exists := filemanager.IsFileExists(filePath)
if exists {
	fmt.Println("File exists.")
} else {
	fmt.Println("File does not exist.")
}
```
