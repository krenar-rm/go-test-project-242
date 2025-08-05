package main

import (
	"code"
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:        "hexlet-path-size",
		Usage:       "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		UsageText:   "hexlet-path-size [options] <path>",
		Description: "A tool to calculate size of files and directories. Supports recursive calculation, human-readable output and hidden files.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "calculate size recursively (for directories)",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "show sizes in human-readable format (B, KB, MB, GB)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories (starting with .)",
			},
		},
		Action: runCommand,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// runCommand выполняет основную логику программы
func runCommand(ctx context.Context, c *cli.Command) error {
	if c.NArg() != 1 {
		if err := cli.ShowAppHelp(c); err != nil {
			return cli.Exit(fmt.Sprintf("failed to show help: %v", err), 1)
		}
		return cli.Exit("path argument is required", 1)
	}

	path := c.Args().Get(0)
	recursive := c.Bool("recursive")
	human := c.Bool("human")
	all := c.Bool("all")

	disp, err := code.GetPathSize(path, recursive, human, all)
	if err != nil {
		return cli.Exit(fmt.Sprintf("failed to get path size: %v", err), 1)
	}

	fmt.Printf("%-10s\t%s\n", disp, path)
	return nil
}
