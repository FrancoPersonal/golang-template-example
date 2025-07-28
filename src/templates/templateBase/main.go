package main

import (
	"fmt"
	"os"

	"github.com/FrancoPersonal/golang-template-example/src/filegenerator"
)

type Template struct {
	Name string `json:"name"`
}

const (
	TemplatePath = "src\\templates\\templatebase\\"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing templatename")
		os.Exit(1)
	}
	templatename := os.Args[1]
	fmt.Println(templatename)
	destinationPath := "src\\templates\\" + templatename

	config := &Template{templatename}
	generator := filegenerator.NewFromStruct(TemplatePath, destinationPath, config)
	err := generator.GenerateFiles()
	if err != nil {
		return
	}
	filePath := destinationPath + "\\files\\empyfile.txt"
	os.Remove(filePath)
}
