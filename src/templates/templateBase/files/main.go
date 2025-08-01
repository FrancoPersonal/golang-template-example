package main

import (
	"github.com/FrancoPersonal/golang-template-example/src/filegenerator"
)

type Template struct {
	Name             string `json:"name"`
	GithubRepository string `json:"githubRepository"`
	Description      string `json:"descripttion"`
}

const (
	TemplatePath    = "src\\templates\\{{.Name}}\\"
	DestinationPath = "output"
)

func main() {
	config := &Template{}
	exeptions := []string{}
	generator, err := filegenerator.New(TemplatePath, DestinationPath, config, exeptions)
	if err != nil {
		return
	}
	err = generator.GenerateFiles()
	if err != nil {
		return
	}
}
