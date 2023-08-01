package dump

import (
	"fmt"
	"kubedump/pkg/exec"
	"kubedump/pkg/files"
)

// Resource dump all named resource on informed namespace
func Resource(namespace string, resource string, kubectl string, format string, projectName string) {
	fmt.Printf("Dumping '%s' of namespace '%s'\n", resource, namespace)

	dumpCmd := fmt.Sprintf("%s get %s -n %s --field-selector metadata.name!=default -o %s", kubectl, resource, namespace, format)

	output, err := exec.SoExec(dumpCmd)
	if err != nil {
		fmt.Printf("Error to Dump resource %s on namespace %s: %v %v\n", resource, namespace, output, err)
		return
	}

	outputFile := fmt.Sprintf("./%s/%s/%s.%s", projectName, namespace, resource, format)
	err = files.WriteFile(outputFile, output)
	if err != nil {
		fmt.Printf("Erro to write file: %v\n", err)
		return
	}
}

// Namespace dump a namespace object
func Namespace(namespace string, kubectl string, format string, projectName string) {
	dumpCmd := fmt.Sprintf("%s get ns %s -o %s", kubectl, namespace, format)
	output, err := exec.SoExec(dumpCmd)
	if err != nil {
		fmt.Printf("Error to Dump on namespace %s: %v\n", namespace, err)
		return
	}
	outputFile := fmt.Sprintf("./%s/%s/00-namespace.%s", projectName, namespace, format)
	err = files.WriteFile(outputFile, output)
	if err != nil {
		fmt.Printf("Erro to write file: %v\n", err)
		return
	}
}
