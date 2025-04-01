## Argument Manager
The `ArgManager` package provides a structured way to handle command-line arguments in a Go application using the `urfave/cli/v2` package.

### Structs

#### `OptionalArg[T any]`
Represents an optional argument with a value and a flag indicating whether the argument was provided or not.

#### Fields:
- `Value` (T): The actual argument value.
- `IsSet` (bool): A flag indicating whether the argument was provided and is valid (not nil or incorrect type).

#### `Args`
Holds the various arguments used in the application.

#### Fields:
- `RootPaths` (`[]string`): A list of root paths for analysis.
- `IncludeComment` (bool): A flag indicating whether to include comments in the analysis.
- `OutputPath` (`OptionalArg[string]`): The output path where images and markdown files are stored. This is an optional argument.

---

### Functions

#### `ParseArgs`
Parses the command-line arguments and populates the `Args` fields.

#### Arguments:
- `arguments []string`: The list of command-line arguments passed to the application (e.g., `os.Args`).

#### Returns:
- `(*Args, error)`: A pointer to an `Args` instance containing the parsed values, or an error if parsing fails.

#### Example Usage:

```go

args, err := args.ParseArgs(os.Args)
if err != nil {
  log.Fatal("Error parsing arguments:", err)
}

// Check if output path is set
if !parsedArgs.RootPaths.IsSet {
  log.Fatal("Please provide output path using the '--output-path' flag.")
}

  fmt.Println("Root Paths:", parsedArgs.RootPaths)
  fmt.Println("Include Comments:", parsedArgs.IncludeComment)
  fmt.Println("Output Path:", parsedArgs.OutputPath.Value)
}
```

---

### Flags Supported

#### `--paths` / `-p`
**Description:** Specifies the list of root paths for analysis.

**Example:**
```sh
go run . --paths /path/to/files --paths /another/path
```
or 

```sh
go run . -p /path/to/files -p /another/path
```

#### `--include-comment` / `-ic`
**Description:** Determines whether comments should be included in the analysis.

**Example:**
```sh
go run . --include-comment
```
or 

```sh
go run . -ic
```

#### `--output-path` / `-op`
**Description:** Specifies the output path where images and markdown files are stored. This is an optional flag.

**Example:**
```sh
go run . --output-path /path/to/output
```
or 

```sh
go run . -op /path/to/output
```
