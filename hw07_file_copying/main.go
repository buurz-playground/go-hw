package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	from, to      string
	limit, offset uint64
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "from",
				Usage:       "file to read from",
				Destination: &from,
			},
			&cli.StringFlag{
				Name:        "to",
				Usage:       "file to write to",
				Destination: &to,
			},
			&cli.Uint64Flag{
				Name:        "limit",
				Usage:       "limit of bytes to copy",
				Destination: &limit,
			},
			&cli.Uint64Flag{
				Name:        "offset",
				Usage:       "offset in input file",
				Destination: &offset,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if err := Copy(from, to, int64(offset), int64(limit)); err != nil {
				return fmt.Errorf("failed to copy file %w", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
