package ArgManager

import (
	"github.com/urfave/cli/v2"
	"os"
)

// Arg represents a generic argument with a value and a flag indicating if it's provided.
type Arg[T any] struct {
	Value         T
	IsArgProvided bool
}

// Args holds the various arguments used in the application.
type Args struct {
	RootPaths      Arg[[]string]
	IncludeComment Arg[bool]
}

// ParseArgs parses the command-line arguments and populates the Args fields.
func (a *Args) ParseArgs() (*Args, error) {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "paths",
				Aliases: []string{"p"},
				Usage:   "List of root paths for analysis files",
			},
			&cli.BoolFlag{
				Name:    "include-comment",
				Aliases: []string{"ic"},
				Usage:   "Include comments in the analysis",
			},
		},
		Action: func(ctx *cli.Context) error {
			// Parsing the paths argument
			paths := ctx.StringSlice("paths")
			if len(paths) > 0 {
				a.RootPaths.Value = paths
				a.RootPaths.IsArgProvided = true
			}

			// Parsing the include-comment flag
			includeComment := ctx.Bool("include-comment")
			a.IncludeComment.Value = includeComment
			a.IncludeComment.IsArgProvided = true
			return nil
		},
	}

	// Run the app and handle errors
	err := app.Run(os.Args)
	if err != nil {
		return nil, err
	}

	return a, nil
}
