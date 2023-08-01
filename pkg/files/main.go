package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Create a folders in recursive mode
func CreateFolder(location string) {
	err := os.MkdirAll(location, 0755)
	if err != nil {
		fmt.Printf("Erro to create dir: %v\n", err)
		return
	}
}

// Write a file with a string content
func WriteFile(location string, content string) error {
	err := ioutil.WriteFile(location, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// List all files in folder and return all then in slice of strings
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
		fmt.Printf("Erro to list dir %v: %v\n", path, err)
		return nil
	}
	return items[1:]
}
