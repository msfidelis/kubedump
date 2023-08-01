package cmd

import (
	"fmt"
	"kubedump/pkg/restore"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore [namespace]",
	Short: "restore all resources dumped using kubedump of a Kubernetes namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespace := args[0]
		projectName, _ := cmd.Flags().GetString("project")
		kubectl, _ := cmd.Flags().GetString("kubectl-location")

		restore.RestoreNamespace(namespace, projectName, kubectl)

		resourcesPath := fmt.Sprintf("./%s/%s", projectName, namespace)
		resourcesFiles := GetFilesInFolder(resourcesPath)

		for _, v := range resourcesFiles {
			restore.RestoreResource(v, namespace, kubectl)
		}
	},
}

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
