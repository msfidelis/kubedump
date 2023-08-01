package cmd

import (
	"fmt"
	"kubedump/pkg/files"
	"kubedump/pkg/restore"

	"github.com/spf13/cobra"
)

// RestoreCmd cobra command definition
var RestoreCmd = &cobra.Command{
	Use:   "restore [namespace]",
	Short: "restore all resources dumped using kubedump of a Kubernetes namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespace := args[0]
		projectName, _ := cmd.Flags().GetString("project")
		kubectl, _ := cmd.Flags().GetString("kubectl-location")

		restore.Namespace(namespace, projectName, kubectl)

		resourcesPath := fmt.Sprintf("./%s/%s", projectName, namespace)
		resourcesFiles := files.GetFilesInFolder(resourcesPath)

		for _, v := range resourcesFiles {
			restore.Resource(v, namespace, kubectl)
		}
	},
}
