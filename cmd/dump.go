package cmd

import (
	"fmt"
	"kubedump/pkg/dump"
	"kubedump/pkg/files"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// DumpCmd cobra command definition
var DumpCmd = &cobra.Command{
	Use:   "dump [namespace]",
	Short: "dump all resources of a Kubernetes namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		namespace := args[0]
		projectName, _ := cmd.Flags().GetString("project")
		kubectl, _ := cmd.Flags().GetString("kubectl-location")
		format, _ := cmd.Flags().GetString("format")

		resourcesString, _ := cmd.Flags().GetString("resources")
		resourcesString = strings.TrimSpace(resourcesString)
		resources := strings.Split(resourcesString, ",")

		projectFolder := fmt.Sprintf("./%s/%s", projectName, namespace)
		log.Info("Starting dump", "namespace", namespace)

		files.CreateFolder(projectFolder)

		dump.Namespace(namespace, kubectl, format, projectName)
		for _, v := range resources {
			dump.Resource(namespace, v, kubectl, format, projectName)
		}
		log.Info("Success", "namespace", namespace, "output_files", projectFolder)
	},
}
