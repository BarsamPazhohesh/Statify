package ArgManager

import (
	"github.com/urfave/cli/v2"
)

type OptionalArg[T any] struct {
	Value T
	IsSet bool
}

type Args struct {
	RootPaths      []string
	IncludeComment bool
	OutputPaths    OptionalArg[[]string]
}

// ParseArgs parses command-line arguments and returns an Args struct.
func ParseArgs(arguments []string) (*Args, error) {
	var args Args
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "paths",
				Aliases:  []string{"p"},
				Usage:    "List of root paths for analysis files",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "include-comment",
				Aliases: []string{"ic"},
				Usage:   "Include comments in the analysis",
			},
			&cli.StringSliceFlag{
				Name:    "output-path",
				Aliases: []string{"op"},
				Usage:   "Specify output path where images and markdown file are stored",
			},
		},
		Action: func(ctx *cli.Context) error {
			args.RootPaths = ctx.StringSlice("paths")
			args.IncludeComment = ctx.Bool("include-comment")
			args.OutputPaths = parseOutputPath(ctx)

			return nil
		},
	}

	if err := app.Run(arguments); err != nil {
		return nil, err
	}

	return &args, nil
}

func parseOutputPath(ctx *cli.Context) OptionalArg[[]string] {
	values := ctx.StringSlice("output-path")
	return OptionalArg[[]string]{
		IsSet: len(values) > 0,
		Value: values,
	}
}
