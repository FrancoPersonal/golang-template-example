package filegenerator

import (
	"os"
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

func TestNewFromStruct(t *testing.T) {
	fg := NewFromStruct("template", "out", TestConfig{Name: "fromstruct"})
	if fg == nil {
		t.Fatal("Expected valid FileGenerator from struct")
	}
}
