package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

type apiError struct {
	Message string `json:"error"`
}

type RequestPayload struct {
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

func BuildRequestPayload(ctx *cli.Context) RequestPayload {
	return RequestPayload{
		ProjectName:      ctx.String("name"),
		GroupId:          ctx.String("group-id"),
		Effect:           getEffect(ctx.String("effect")),
		Implementation:   getImplementation(ctx.String("server")),
		ScalaVersion:     fmt.Sprintf("Scala%d", ctx.Int("scala-version")),
		Builder:          getBuilder(ctx.String("build-tool")),
		Json:             getJson(ctx.String("json-library")),
		AddDocumentation: ctx.Bool("swagger"),
		AddMetrics:       ctx.Bool("metrics"),
	}
}

func DownloadProject(ctx *cli.Context) error {
	payload := BuildRequestPayload(ctx)

	buf, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://adopt-tapir.softwaremill.com/api/v1/starter.zip", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		} else {
			apiErr := &apiError{}
			err := json.Unmarshal(errorBody, apiErr)
			if err != nil {
				return err
			}
			return fmt.Errorf(apiErr.Message)
		}
	}

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
