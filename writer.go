package gomarkdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	markdownSuffix = ".md"
)

type markdownWriter struct {
	file *os.File
}

// NewMarkdownWriter creates a new markdownWriter instance and opens the file.
// If the directory in the filename does not exist, it will be created.
func NewMarkdownWriter(filename string, append bool) (*markdownWriter, error) {
	// Ensure the filename ends with .md
	if !strings.HasSuffix(filename, markdownSuffix) {
		filename += markdownSuffix
	}
	// Ensure the directory exists by extracting the directory path from the filename
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create directories: %v", err)
	}

	var file *os.File

	if append {
		// Open the file in append mode
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		// Create or overwrite the file in write mode
		file, err = os.Create(filename)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	return &markdownWriter{file: file}, nil
}

// Close closes the file
func (mw *markdownWriter) Close() error {
	return mw.file.Close()
}
