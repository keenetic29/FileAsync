package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileRepository interface {
	GetAllFiles(dir string) ([]string, error)
	ReadFileContent(path string) (string, error)
}

type fileRepo struct{}

func NewFileRepository() FileRepository {
	return &fileRepo{}
}

func (r *fileRepo) GetAllFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("не могу прочитать директорию: %v", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	return files, nil
}

func (r *fileRepo) ReadFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("не могу прочитать файл %s: %v", path, err)
	}
	return string(content), nil
}