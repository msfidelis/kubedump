package cmd

import (
	"fmt"
	"io/ioutil"
	"kubedump/pkg/config"
	"kubedump/pkg/dump"
	"kubedump/pkg/files"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// DumpFileCmd cobra command definition
var DumpFileCmd = &cobra.Command{
	Use:   "dump-file",
	Short: "dump all resources of with file custom configs",
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

		for _, v := range config.Namespaces {
			pathFolder := fmt.Sprintf("./%s/%s", config.Project, v)
			files.CreateFolder(fmt.Sprintf("./%s/%s", config.Project, v))
			dump.Namespace(v, kubectl, config.Format, config.Project)

			for _, r := range config.Resources {
				dump.Resource(v, r, kubectl, config.Format, config.Project)

			}
			log.Info("Success", "namespace", v, "output_files", pathFolder)
		}
	},
}
