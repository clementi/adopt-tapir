package main

import "github.com/urfave/cli/v2"

const version = "1.0.0"

var App = &cli.App{
	Name:  "adopt-tapir",
	Usage: "Generate a Scala Tapir project.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "effect",
			Value:    "cats",
			Required: false,
			Usage:    "effect type",
			Aliases:  []string{"e"},
		},
		&cli.StringFlag{
			Name:     "server",
			Value:    "netty",
			Required: false,
			Usage:    "server",
			Aliases:  []string{"s"},
		},
		&cli.IntFlag{
			Name:     "scala-version",
			Value:    2,
			Required: false,
			Usage:    "Scala version",
			Aliases:  []string{"r"},
		},
		&cli.StringFlag{
			Name:     "build-tool",
			Value:    "sbt",
			Required: false,
			Usage:    "build tool",
			Aliases:  []string{"b"},
		},
		&cli.StringFlag{
			Name:     "json-library",
			Value:    "circe",
			Required: false,
			Usage:    "JSON library",
			Aliases:  []string{"j"},
		},
		&cli.BoolFlag{
			Name:     "swagger",
			Value:    false,
			Required: false,
			Usage:    "include Swagger",
			Aliases:  []string{"w"},
		},
		&cli.BoolFlag{
			Name:     "metrics",
			Value:    false,
			Required: false,
			Usage:    "include metrics",
			Aliases:  []string{"m"},
		},
		&cli.BoolFlag{
			Name:     "version",
			Value:    false,
			Required: false,
			Usage:    "show version",
			Aliases:  []string{"V"},
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.Bool("version") {
			cli.ShowVersion(ctx)
			return nil
		}
		return DownloadProject(ctx)
	},
	HideHelpCommand: true,
	HideVersion:     true,
	Version:         version,
}
