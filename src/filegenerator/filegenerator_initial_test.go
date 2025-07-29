package filegenerator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func CreateFiles(t *testing.T, root string) []struct {
	path    string
	content string
} {
	// Create files in different subdirectories
	testFiles := []struct {
		path    string
		content string
	}{
		{filepath.Join(root, "file1.txt"), ""},
		{filepath.Join(root, "subdir", "file2.txt"), ""},
		{filepath.Join(root, "subdir", "nested", "file3.txt"), ""},
	}

	for _, f := range testFiles {
		if err := os.MkdirAll(filepath.Dir(f.path), 0755); err != nil {
			t.Fatalf("Error creating directory: %v", err)
		}
		if err := os.WriteFile(f.path, []byte(f.content), 0644); err != nil {
			t.Fatalf("Error writing file: %v", err)
		}
	}
	return testFiles
}
func createTestTemplateFile(t *testing.T, content string, path string) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		t.Fatalf("Error creating test dir: %v", err)
	}
	err = os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error writing test template file: %v", err)
	}
}

func createJSONFile(t *testing.T, content any, path string) {
	jsonData, _ := json.Marshal(content)
	os.MkdirAll(filepath.Dir(path), 0755)
	err := os.WriteFile(path, jsonData, 0644)
	if err != nil {
		t.Fatalf("Error writing JSON file: %v", err)
	}
}

func createFile(t *testing.T, path string) map[string]string {
	data := map[string]string{
		"key": "value",
	}
	createJSONFile(t, data, path)
	return data
}
