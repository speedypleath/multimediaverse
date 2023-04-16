package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:      "multimediaverse",
		Version:   "v0.0.1",
		Usage:     "AI powered multimedia editing tool",
		UsageText: "multimediaverse command [command options] [arguments...]",
		Commands: []*cli.Command{
			{
				Name:  "add_subtitles",
				Usage: "Add subtitles to a video",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Video: ", cCtx.Args().First())
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
