package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

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

type requestPayload struct {
	ProjectName      string `json:"projectName"`
	GroupId          string `json:"groupId"`
	Effect           string `json:"effect"`
	Implementation   string `json:"implementation"`
	ScalaVersion     string `json:"scalaVersion"`
	Builder          string `json:"builder"`
	Json             string `json:"json"`
	AddDocumentation bool   `json:"addDocumentation"`
	AddMetrics       bool   `json:"addMetrics"`
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

			name := getName(ctx.Args().Get(0))
			groupId := getGroupId(ctx.Args().Get(1))

			return downloadProject(options, name, groupId)
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

func getName(name string) string {
	if len(name) == 0 {
		return generateName()
	}
	return name
}

func getGroupId(groupId string) string {
	if len(groupId) == 0 {
		return "com.example"
	}
	return groupId
}

func generateName() string {
	rand.Seed(time.Now().UnixMilli())

	adjectives := []string{"accurate", "dull", "intelligent", "shiny", "dark", "light", "level", "inaccurate", "stern", "determined", "speedy", "frosty", "arbitrary", "united"}
	nouns := []string{"pelican", "turtle", "zebra", "elephant", "dog", "wolf", "caterpillar", "octopus", "tarsier", "snake", "monkey", "sloth", "spider", "koala", "panda", "seahorse"}

	return fmt.Sprintf("%s-%s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))])
}

func downloadProject(options options, name string, groupId string) error {
	payload := requestPayload{
		ProjectName:      name,
		GroupId:          groupId,
		Effect:           getEffect(options.effect),
		Implementation:   getImplementation(options.server),
		ScalaVersion:     fmt.Sprintf("Scala%d", options.scalaVersion),
		Builder:          getBuilder(options.buildTool),
		Json:             getJson(options.jsonLibrary),
		AddDocumentation: options.includeSwagger,
		AddMetrics:       options.includeMetrics,
	}

	buf, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://adopt-tapir.softwaremill.com/api/v1/starter.zip", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s-tapir-starter.zip", payload.ProjectName)

	err = os.WriteFile(fileName, body, 0644)

	if err == nil {
		fmt.Printf("Wrote %s\n", fileName)
	}

	return err
}

func getEffect(effect string) string {
	switch effect {
	case "cats":
		return "IOEffect"
	case "zio":
		return "ZIOEffect"
	case "future":
		return "FutureEffect"
	}
	return effect
}

func getImplementation(server string) string {
	switch server {
	case "netty":
		return "Netty"
	case "vertx":
		return "VertX"
	case "zio-http":
		return "ZIOHttp"
	case "http4s":
		return "Http4s"
	}
	return server
}

func getBuilder(buildTool string) string {
	switch buildTool {
	case "sbt":
		return "Sbt"
	case "scala-cli":
		return "ScalaCli"
	}
	return buildTool
}

func getJson(jsonLibrary string) string {
	switch jsonLibrary {
	case "upickle":
		return "UPickle"
	case "jsoniter":
		return "Jsoniter"
	case "circe":
		return "Circe"
	case "zio-json":
		return "ZIOJson"
	case "no":
		return "no"
	}
	return jsonLibrary
}
