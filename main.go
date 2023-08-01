package main

import (
	"kubedump/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kubedump"}
	rootCmd.AddCommand(cmd.DumpCmd, cmd.RestoreCmd)

	cmd.DumpCmd.Flags().BoolP("dry-run", "d", false, "Perform a dry-run backup (no actual backup will be performed)")
	cmd.DumpCmd.Flags().String("resources", "deployment,service,hpa,ingress,serviceaccount,daemonset,statefulset,job,cronjob", "Kubernetes resources separated by comma")
	cmd.DumpCmd.Flags().String("kubectl-location", "/usr/local/bin/kubectl", "Custom kubectl binary or alias")
	cmd.DumpCmd.Flags().String("format", "yaml", "Dump output format")
	cmd.DumpCmd.Flags().String("project", "kubedump", "Project name")
	cmd.DumpCmd.Flags().String("config-file", "", "kubedump config file location")

	cmd.RestoreCmd.Flags().String("kubectl-location", "/usr/local/bin/kubectl", "Custom kubectl binary or alias")
	cmd.RestoreCmd.Flags().String("project", "kubedump", "Project name")
	cmd.RestoreCmd.Flags().String("config-file", "", "kubedump config file location")

	rootCmd.Execute()
}
