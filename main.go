package main

import (
	"kubedump/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kubedump"}
	rootCmd.AddCommand(
		cmd.DumpCmd,
		cmd.RestoreCmd,
		cmd.DumpFileCmd,
		cmd.RestoreFileCmd,
	)

	cmd.DumpCmd.Flags().BoolP("dry-run", "d", false, "Perform a dry-run backup (no actual backup will be performed)")
	cmd.DumpCmd.Flags().String("resources", "deployment,service,hpa,ingress,serviceaccount,daemonset,statefulset,job,cronjob,configmaps,secrets", "Kubernetes resources separated by comma")
	cmd.DumpCmd.Flags().String("kubectl-location", "kubectl", "Custom kubectl binary or alias")
	cmd.DumpCmd.Flags().String("format", "yaml", "Dump output format")
	cmd.DumpCmd.Flags().String("project", "kubedump", "Project name")

	cmd.DumpFileCmd.Flags().String("kubectl-location", "/usr/local/bin/kubectl", "Custom kubectl binary or alias")
	cmd.DumpFileCmd.Flags().String("config-file", "", "kubedump config file location")

	cmd.RestoreCmd.Flags().String("kubectl-location", "/usr/local/bin/kubectl", "Custom kubectl binary or alias")
	cmd.RestoreCmd.Flags().String("project", "kubedump", "Project name")

	cmd.RestoreFileCmd.Flags().String("kubectl-location", "/usr/local/bin/kubectl", "Custom kubectl binary or alias")
	cmd.RestoreFileCmd.Flags().String("config-file", "", "kubedump config file location")

	rootCmd.Execute()
}
