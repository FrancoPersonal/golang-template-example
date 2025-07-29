package filegenerator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew_ErrorReadingJSON(t *testing.T) {
	tmpDir := "tmpdirTestNew_ErrorReadingJSON"
	_, err := New(tmpDir, "out", &TestConfig{})
	if err == nil {
		t.Fatal("Expected error due to missing JSON file")
	}
	os.RemoveAll(tmpDir)
}

func TestNew(t *testing.T) {
	tmpDir := "tmpdirNew"
	jsonPath := filepath.Join(tmpDir, "templatejson.json")
	expected := TestConfig{Name: "example"}
	createJSONFile(t, expected, jsonPath)

	fg, err := New(tmpDir, "out", &TestConfig{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if fg == nil {
		t.Fatal("Expected valid FileGenerator, got nil")
	}
	os.RemoveAll(tmpDir)
}

func TestNewFromStruct(t *testing.T) {
	fg := NewFromStruct("template", "out", TestConfig{Name: "fromstruct"})
	if fg == nil {
		t.Fatal("Expected valid FileGenerator from struct")
	}
}
