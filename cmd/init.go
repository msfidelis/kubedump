package cmd

import (
	"fmt"
	"kubedump/pkg/config"
	"kubedump/pkg/files"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new file project",
	Run: func(cmd *cobra.Command, args []string) {
		project, _ := cmd.Flags().GetString("project")
		resourcesString, _ := cmd.Flags().GetString("resources")
		resources := strings.Split(resourcesString, ",")
		log.Info("Initializing", "project", project)

		projectFolder := fmt.Sprintf("./%s/", project)
		files.CreateFolder(projectFolder)
		log.Info("Project folder created", "folder", projectFolder)

		var model config.Model

		model.Project = project
		model.Format = "yaml"
		model.Namespaces = []string{"default"}
		model.Resources = resources

		fileProject := fmt.Sprintf("./%s/project.yaml", project)

		strProject, err := yaml.Marshal(model)
		if err != nil {
			log.Fatalf("Error to Unmarshal YAML File: %v", err)
		}

		files.WriteFile(fileProject, string(strProject))

		log.Info("Project file created", "project", fileProject)
	},
}
