package main

import (
	"context"
	"fmt"
	"os"

	pathsize "hexlet-path-size"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "Print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "sum sizes recursively for directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "print sizes in human readable form (KB, MB, GB)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: run,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return fmt.Errorf("expected exactly one path argument")
	}
	path := cmd.Args().First()

	opts := pathsize.Options{
		Recursive: cmd.Bool("recursive"),
		Human:     cmd.Bool("human"),
		All:       cmd.Bool("all"),
	}

	size, err := pathsize.GetPathSize(path, opts)
	if err != nil {
		return err
	}

	fmt.Printf("%s\t%s\n", size, path)
	return nil
}
