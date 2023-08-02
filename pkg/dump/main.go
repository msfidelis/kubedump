package dump

import (
	"fmt"
	"kubedump/pkg/exec"
	"kubedump/pkg/files"

	"github.com/charmbracelet/log"
)

// Resource dump all named resource on informed namespace
func Resource(namespace string, resource string, kubectl string, format string, projectName string) {

	log.Info("Dumping resources", "namespace", namespace, "resource", resource)
	dumpCmd := fmt.Sprintf("%s get %s -n %s --field-selector metadata.name!=default -o %s", kubectl, resource, namespace, format)

	output, err := exec.SoExec(dumpCmd)
	if err != nil {
		fmt.Printf("Error to Dump resource %s on namespace %s: %v %v\n", resource, namespace, output, err)
		return
	}

	outputFile := fmt.Sprintf("./%s/%s/%s.%s", projectName, namespace, resource, format)
	err = files.WriteFile(outputFile, output)
	if err != nil {
		log.Error("Error to write resource file", "namespace", namespace, "resource", resource, "file", outputFile, "error", err.Error())
		return
	}
}

// Namespace dump a namespace object
func Namespace(namespace string, kubectl string, format string, projectName string) {
	log.Info("Dumping resources", "namespace", namespace, "resource", "namespace")
	dumpCmd := fmt.Sprintf("%s get ns %s -o %s", kubectl, namespace, format)
	output, err := exec.SoExec(dumpCmd)
	if err != nil {
		log.Error("Error to dump namespace", "namespace", namespace, "error", err.Error())
		return
	}
	outputFile := fmt.Sprintf("./%s/%s/00-namespace.%s", projectName, namespace, format)
	err = files.WriteFile(outputFile, output)
	if err != nil {
		log.Error("Error to write namespace resource file", "namespace", namespace, "file", outputFile, "error", err.Error())
		return
	}
}
