package file

import (
	"errors"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path"
)

type ImportedFile struct {
	Path    string
	Content []byte
}

// ImportFolder Import all files from a given folder path.
// folder the folder to read
func ImportFolder(folder string) (iter.Seq[*ImportedFile], error) {
	slog.Info("Importing folder", "path", folder)

	info, err := os.Stat(folder)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		message := fmt.Sprint("Given path : ", folder, " should be a directory ")
		return nil, errors.New(message)
	}

	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	return func(yield func(*ImportedFile) bool) {
		for _, file := range files {
			filePath := path.Join(folder, file.Name())
			content, err := ImportFile(filePath)
			if err == nil {
				importedFile := &ImportedFile{
					filePath,
					content,
				}
				if !yield(importedFile) {
					break
				}
			}
		}
	}, nil
}

// ImportFile Import file content from a given file path.
// file the file to read
func ImportFile(file string) ([]byte, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		message := fmt.Sprint("Given path ", file, " should be a file")
		return nil, errors.New(message)
	}

	slog.Debug("Importing file", "path", file)
	return os.ReadFile(file)
}
