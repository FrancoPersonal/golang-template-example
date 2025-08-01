package main

import (
	"github.com/FrancoPersonal/golang-template-example/src/filegenerator"
)

type Project struct {
	Name             string `json:"name"`
	GithubRepository string `json:"githubRepository"`
	Description      string `json:"descripttion"`
}

const (
	TemplatePath    = "src\\templates\\initialproyect\\"
	DestinationPath = "output"
)

func main() {
	config := &Project{}
	exeptions := []string{
		"src\\templates\\initialproyect\\files\\.github\\workflows\\release.yml",
	}
	generator, err := filegenerator.New(TemplatePath, DestinationPath, config, exeptions)
	if err != nil {
		return
	}
	err = generator.GenerateFiles()
	if err != nil {
		return
	}
}
