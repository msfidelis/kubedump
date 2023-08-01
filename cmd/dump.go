package cmd

import (
	"fmt"
	"kubedump/pkg/dump"
	"kubedump/pkg/files"
	"strings"

	"github.com/spf13/cobra"
)

var DumpCmd = &cobra.Command{
	Use:   "dump [namespace]",
	Short: "dump all resources of a Kubernetes namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespace := args[0]
		project_name, _ := cmd.Flags().GetString("project")
		kubectl, _ := cmd.Flags().GetString("kubectl-location")
		format, _ := cmd.Flags().GetString("format")

		resourcesString, _ := cmd.Flags().GetString("resources")
		resourcesString = strings.TrimSpace(resourcesString)
		resources := strings.Split(resourcesString, ",")

		files.CreateFolder(fmt.Sprintf("./%s/%s", project_name, namespace))
		dump.DumpNamespace(namespace, kubectl, format, project_name)
		for _, v := range resources {
			dump.DumpResouce(namespace, v, kubectl, format, project_name)
		}
	},
}
