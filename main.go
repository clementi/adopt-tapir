package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type options struct {
	effect         string
	server         string
	scalaVersion   int
	buildTool      string
	jsonLibrary    string
	includeSwagger bool
	includeMetrics bool
}

func main() {
	app := &cli.App{
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
				Value:    true,
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
				Name:     "verbose",
				Value:    true,
				Required: false,
				Usage:    "verbose output",
				Aliases:  []string{"v"},
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

			options := buildOptions(ctx)
			return downloadProject(options)
		},
		HideHelpCommand: true,
		HideVersion:     true,
		Version:         "0.1.0",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func buildOptions(ctx *cli.Context) options {
	return options{
		effect:         ctx.String("effect"),
		server:         ctx.String("server"),
		scalaVersion:   ctx.Int("scala-version"),
		buildTool:      ctx.String("build-tool"),
		jsonLibrary:    ctx.String("json-library"),
		includeSwagger: ctx.Bool("swagger"),
		includeMetrics: ctx.Bool("metrics"),
	}
}

func downloadProject(options options) error {
	return nil // TODO
}
