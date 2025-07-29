package filegenerator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFilePathsRecursive(t *testing.T) {
	// Setup: Create a test directory structure
	tempDir := "tempDir"
	root := filepath.Join(tempDir, "testfiles")

	if err := os.MkdirAll(root, 0755); err != nil {
		t.Fatalf("Error creating directory: %v", err)
	}

	testFiles := CreateFiles(t, root)
	// Test
	paths, err := getFilePathsRecursive(root, root)
	if err != nil {
		t.Errorf("getFilePathsRecursive failed: %v", err)
	}

	// Verify all test files are present
	for _, f := range testFiles {
		path := f.path
		found := false
		for _, p := range paths {
			if p.Original == path {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found in results", path)
		}
	}
	os.RemoveAll(tempDir)
}
