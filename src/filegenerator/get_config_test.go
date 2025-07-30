package filegenerator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetConfig(t *testing.T) {
	// Setup: Create a temporary directory for the test
	tmpDir := "tmpdirTestGetConfig"
	testJson := filepath.Join(tmpDir, "test.json")
	data := createFile(t, testJson)

	// Test success case
	var result map[string]string
	err := getConfig(testJson, &result)
	if err != nil {
		t.Errorf("getConfig failed: %v", err)
	}
	if diff := cmp.Diff(data, result); diff != "" {
		t.Errorf("getConfig returned unexpected result. Diff:\n%s", diff)
	}

	// Test error cases
	testJsonCorrupted := filepath.Join(tmpDir, "corrupted.json")
	err = os.WriteFile(testJsonCorrupted, []byte("invalid json"), 0644)
	if err == nil {
		err = getConfig(testJsonCorrupted, &result)
		if err == nil {
			t.Error("Expected error when reading corrupted JSON, got none.")
		}
	}
	os.RemoveAll(tmpDir)
}

func TestGetConfig_ErrorJSON(t *testing.T) {
	tmpDir := "tmpdirTestGetConfig_ErrorJSON"
	tmpFile := filepath.Join(tmpDir, "bad.json")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.WriteFile(tmpFile, []byte("{invalid json}"), 0644); err != nil {
		t.Fatal(err.Error())
	}

	var cfg TestConfig
	err := getConfig(tmpFile, &cfg)
	if err == nil {
		t.Fatal("Expected error parsing invalid JSON")
	}
	os.RemoveAll(tmpDir)
}
