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
		project_name, _ := cmd.Flags().GetString("project")
		kubectl, _ := cmd.Flags().GetString("kubectl-location")

		restore.RestoreNamespace(namespace, project_name, kubectl)

		resources_path := fmt.Sprintf("./%s/%s", project_name, namespace)
		resource_files := GetFilesInFolder(resources_path)

		fmt.Println(resource_files)

		for _, v := range resource_files {
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
