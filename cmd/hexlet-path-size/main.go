package main

import (
	"context"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		UsageText: "hexlet-path-size [options] <path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "recursive", Aliases: []string{"r"}, Usage: "recursive size of directories"},
			&cli.BoolFlag{Name: "human", Aliases: []string{"H"}, Usage: "human-readable sizes (auto-select unit)"},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Usage: "include hidden files and directories"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NArg() != 1 {
				_ = cli.ShowAppHelp(c)
				return cli.Exit("", 1)
			}
			path := c.Args().Get(0)
			disp, err := code.GetPathSize(path, c.Bool("recursive"), c.Bool("human"), c.Bool("all"))
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			fmt.Printf("%s\t%s\n", disp, path)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
