package ArgManager

import (
	"github.com/urfave/cli/v2"
)

type OptionalArg[T any] struct {
	Value T
	IsSet bool
}

// newOptionalArg returns an OptionalArg with the value from the flag if set, or the defaultValue otherwise.
func newOptionalArg[T any](ctx *cli.Context, flagName string, defaultValue T) OptionalArg[T] {
	val := ctx.Generic(flagName)
	if val == nil {
		return OptionalArg[T]{Value: defaultValue, IsSet: false}
	}

	typedVal, ok := val.(T)
	if !ok {
		return OptionalArg[T]{Value: defaultValue, IsSet: false}
	}

	return OptionalArg[T]{Value: typedVal, IsSet: true}
}

type Args struct {
	RootPaths      []string
	IncludeComment bool
	OutputPath     OptionalArg[string]
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
			&cli.StringFlag{
				Name:    "output-path",
				Aliases: []string{"op"},
				Usage:   "Specify output path where images and markdown file are stored",
			},
		},
		Action: func(ctx *cli.Context) error {
			args.RootPaths = ctx.StringSlice("paths")
			args.IncludeComment = ctx.Bool("include-comment")
			args.OutputPath = newOptionalArg(ctx, "output-path", "")

			return nil
		},
	}

	if err := app.Run(arguments); err != nil {
		return nil, err
	}

	return &args, nil
}
