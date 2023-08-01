package files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFolder(location string) {
	err := os.MkdirAll(location, 0755)
	if err != nil {
		fmt.Printf("Erro to create dir: %v\n", err)
		return
	}
}

func WriteFile(location string, content string) error {
	err := ioutil.WriteFile(location, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
