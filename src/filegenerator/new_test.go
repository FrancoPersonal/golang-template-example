package filegenerator

import (
	"os"
	"testing"
)

func TestNew_ErrorReadingJSON(t *testing.T) {
	tmpDir := "tmpdirTestNew_ErrorReadingJSON"
	exeptions := []string{}
	_, err := New(tmpDir, "out", &TestConfig{}, exeptions)
	if err == nil {
		t.Fatal("Expected error due to missing JSON file")
	}
	os.RemoveAll(tmpDir)
}

func TestNewFromStruct(t *testing.T) {
	exeptions := []string{}
	fg := NewFromStruct("template", "out", TestConfig{Name: "fromstruct"}, exeptions)
	if fg == nil {
		t.Fatal("Expected valid FileGenerator from struct")
	}
}
