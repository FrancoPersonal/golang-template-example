package filegenerator

import (
	"os"
	"path/filepath"
	"testing"
)

// estructura de configuración de prueba
type TestConfig struct {
	Name string
}

func TestGenerateFile_Errors(t *testing.T) {
	fg := FileGenerator{config: TestConfig{Name: "X"}}

	// plantilla no existe
	err := fg.GenerateFile(TemplateFile{Original: "nonexistent.tmpl", Destination: "out.tmpl"})
	if err == nil {
		t.Error("Expected error loading nonexistent template")
	}

	// crear plantilla con error de ejecución
	tmpDir := "tmpdirTestGenerateFile_Errors"
	tmplPath := filepath.Join(tmpDir, "broken.tmpl")
	destPath := filepath.Join(tmpDir, "out.txt")
	createTestTemplateFile(t, "{{.InvalidField}}", tmplPath)

	err = fg.GenerateFile(TemplateFile{Original: tmplPath, Destination: destPath})
	if err == nil {
		t.Error("Expected template execution error")
	}
	// Clean up
	os.RemoveAll(tmpDir)
}

func TestGenerateFiles_Errors(t *testing.T) {
	fg := FileGenerator{config: TestConfig{Name: "X"}}

	// plantilla no existe
	err := fg.GenerateFile(TemplateFile{Original: "nonexistent.tmpl", Destination: "out.tmpl"})
	if err == nil {
		t.Error("Expected error loading nonexistent template")
	}

	// crear plantilla con error de ejecución
	tmpDir := "tmpdirGenerateFiles_errors"
	tmplPath := filepath.Join(tmpDir, "broken.tmpl")
	createTestTemplateFile(t, "{{.InvalidField}}", tmplPath)

	err = fg.GenerateFiles()
	if err == nil {
		t.Error("Expected template execution error")
	}
	// Clean up
	os.RemoveAll(tmpDir)
}
