package cmd

import (
	"fmt"
	"io/ioutil"
	"kubedump/pkg/config"
	"kubedump/pkg/files"
	"kubedump/pkg/restore"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// RestoreFileCmd cobra command definition
var RestoreFileCmd = &cobra.Command{
	Use:   "restore-file",
	Short: "restore all resources of with custom configs from configuration files",
	Run: func(cmd *cobra.Command, args []string) {

		kubectl, _ := cmd.Flags().GetString("kubectl-location")
		configFile, _ := cmd.Flags().GetString("config-file")

		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			fmt.Printf("Error to read %s file: %v\n", configFile, err)
			return
		}

		var config config.Model

		// Mapear o conteúdo do arquivo YAML para a struct
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("Error to Unmarshal YAML File: %v", err)
		}

		for _, n := range config.Namespaces {

			restore.Namespace(n, config.Project, kubectl)

			resourcesPath := fmt.Sprintf("./%s/%s", config.Project, n)
			resourcesFiles := files.GetFilesInFolder(resourcesPath)

			for _, r := range resourcesFiles {
				restore.Resource(r, n, kubectl)
			}
		}
	},
}
