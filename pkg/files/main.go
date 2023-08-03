package files

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

// CreateFolder a folders in recursive mode
func CreateFolder(location string) {
	err := os.MkdirAll(location, 0755)
	if err != nil {
		log.Error("Error to create dir:", "path", location, "error", err.Error())
		return
	}
}

// WriteFile a file with a string content
func WriteFile(location string, content string) error {
	err := ioutil.WriteFile(location, []byte(content), 0644)
	if err != nil {
		log.Error("Error to write a file content:", "path", location, "error", err.Error())
		return err
	}
	return nil
}

// GetFilesInFolder List all files in folder and return all then in slice of strings
func GetFilesInFolder(path string) []string {
	var items []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		items = append(items, path)
		return nil
	})
	if err != nil {
		log.Error("Error to list dir:", "path", path, "error", err.Error())
		return nil
	}
	return items[1:]
}
