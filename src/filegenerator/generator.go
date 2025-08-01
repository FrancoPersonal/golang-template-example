package filegenerator

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type IFileGenerator interface {
	GenerateFiles() error
}

type FileGenerator struct {
	config       any
	templatePath string
	outputFolder string
	exeptions    []string
}

type TemplateFile struct {
	Original    string
	Destination string
}

func New(templatePath string, outputFolder string, v any, exeptions []string) (IFileGenerator, error) {
	jsonTemplate := templatePath + "\\" + "templatejson.json"
	err := getConfig(jsonTemplate, v)
	return FileGenerator{v, templatePath, outputFolder, exeptions}, err
}

func NewFromStruct(templatePath string, outputFolder string, v any, exeptions []string) IFileGenerator {
	return FileGenerator{v, templatePath, outputFolder, exeptions}
}

func getConfig(jsonTemplate string, v any) error {
	data, err := os.ReadFile(jsonTemplate)
	if err != nil {
		log.Println("Error reading project.json:", err)
		return err
	}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Println("Error parsing JSON:", err)
		return err
	}
	return nil
}

func (service FileGenerator) GenerateFiles() error {
	files, err := getFilePathsRecursive(service.templatePath+"files\\", service.outputFolder+"\\")
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := service.GenerateFile(file); err != nil {
			return err
		}
	}
	log.Println("Total files Generated:", len(files))
	return nil
}

func (service FileGenerator) GenerateFile(templateFile TemplateFile) error {
	// Abrir la plantilla del Makefile
	log.Println("initial GenerateFile", templateFile.Original)

	dirPath := filepath.Dir(templateFile.Destination)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Println(err.Error())
	}
	outFile, err := os.Create(templateFile.Destination)
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}
	defer outFile.Close()

	if contains(service.exeptions, templateFile.Original) {
		sourceFile, err := os.Open(templateFile.Original)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		_, err = io.Copy(outFile, sourceFile)
		if err != nil {
			return err
		}
		log.Println("file generated successfully", templateFile.Destination)
		return nil
	}

	tmpl, err := template.ParseFiles(templateFile.Original)
	if err != nil {
		log.Println("Error loading template:", err)
		return err
	}

	// Ejecutar la plantilla con los datos
	if err := tmpl.Execute(outFile, service.config); err != nil {
		log.Println("Error writing Makefile:", err)
		return err
	}

	log.Println("file generated successfully", templateFile.Destination)
	return nil
}

func getFilePathsRecursive(root string, destinationPath string) ([]TemplateFile, error) {
	var filePaths []TemplateFile

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err // Handle potential errors during traversal
		}
		if !info.IsDir() { // Check if it's a file (not a directory)
			destination := strings.Replace(path, root, destinationPath, 1)
			template := TemplateFile{path, destination}
			filePaths = append(filePaths, template)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %q: %w", root, err)
	}

	return filePaths, nil
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
