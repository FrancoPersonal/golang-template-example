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
type mockConfig struct {
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

func TestGenerateFile_Exception(t *testing.T) {
	tmpDir := "tmpdirTestGenerateFile_ex"
	srcPath := filepath.Join(tmpDir, "plain.txt")
	dstPath := filepath.Join(tmpDir, "out.txt")
	content := "Raw text content"

	tmpl := TemplateFile{
		Original:    srcPath,
		Destination: dstPath,
	}
	createPlainFile(t, content, srcPath)
	gen := FileGenerator{
		config:       mockConfig{},
		templatePath: tmpDir,
		outputFolder: tmpDir,
		exeptions:    []string{srcPath},
	}

	err := gen.GenerateFile(tmpl)
	if err != nil {
		t.Fatalf("GenerateFile with exception failed: %v", err)
	}

	output, _ := os.ReadFile(dstPath)
	if string(output) != content {
		t.Fatalf("Expected copied content, got: %s", string(output))
	}
	os.RemoveAll(tmpDir)
}
func TestGenerateFiles(t *testing.T) {
	tmpDir := "tmpdirTestGenerateFiles"
	srcPath := filepath.Join(tmpDir, "plain.txt")
	dstPath := filepath.Join(tmpDir, "out.txt")

	tmpl := TemplateFile{
		Original:    srcPath,
		Destination: dstPath,
	}
	createFile(t, srcPath)
	gen := FileGenerator{
		config:       mockConfig{},
		templatePath: tmpDir,
		outputFolder: tmpDir,
	}

	err := gen.GenerateFile(tmpl)
	if err != nil {
		t.Fatalf("GenerateFile with exception failed: %v", err)
	}
	os.RemoveAll(tmpDir)
}
