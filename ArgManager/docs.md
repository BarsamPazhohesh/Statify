## Argument Manager
The `ArgManager` package provides a structured way to handle command-line arguments in a Go application using the `urfave/cli/v2` package.

### Structs

#### `Arg[T any]`
Represents a generic argument with a value and a flag indicating if it was provided.

#### Fields:
- `Value` (T): The actual argument value.
- `IsArgProvided` (bool): A flag indicating whether the argument was provided (not nil or incorrect value).

#### `Args`
Holds the various arguments used in the application.

#### Fields:
- `RootPaths` (`Arg[[]string]`): A list of root paths for analysis.
- `IncludeComment` (`Arg[bool]`): A flag indicating whether to include comments in the analysis.

---

### Functions

#### `ParseArgs`
Parses the command-line arguments and populates the `Args` fields.

#### Arguments:
- None

#### Returns:
- `(*Args, error)`: A pointer to an `Args` instance containing the parsed values, or an error if parsing fails.

#### Example Usage:

```go
package main

import (
    "fmt"
    "log"
    "arg_manager"
)

func main() {
    args := &arg_manager.Args{}
    parsedArgs, err := args.ParseArgs()
    if err != nil {
        log.Fatal("Error parsing arguments:", err)
    }

    // Check if root paths are provided
	if !args.RootPaths.IsArgProvided {
		log.Fatal("Please provide root paths using the '--paths' flag.")
	}

    fmt.Println("Root Paths:", parsedArgs.RootPaths.Value)
    fmt.Println("Include Comments:", parsedArgs.IncludeComment.Value)
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