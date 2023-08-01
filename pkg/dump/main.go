package dump

import (
	"fmt"
	"kubedump/pkg/exec"
	"kubedump/pkg/files"
)

func DumpResouce(namespace string, resource string, kubectl string, format string, project_name string) {
	fmt.Printf("Dumping '%s' of namespace '%s'\n", resource, namespace)

	dump_cmd := fmt.Sprintf("%s get %s -n %s --field-selector metadata.name!=default -o %s", kubectl, resource, namespace, format)

	output, err := exec.SoExec(dump_cmd)
	if err != nil {
		fmt.Printf("Error to Dump resource %s on namespace %s: %v %v\n", resource, namespace, output, err)
		return
	}

	output_file := fmt.Sprintf("./%s/%s/%s.%s", project_name, namespace, resource, format)
	err = files.WriteFile(output_file, output)
	if err != nil {
		fmt.Printf("Erro to write file: %v\n", err)
		return
	}
}

func DumpNamespace(namespace string, kubectl string, format string, project_name string) {
	dump_cmd := fmt.Sprintf("%s get ns %s -o %s", kubectl, namespace, format)
	output, err := exec.SoExec(dump_cmd)
	if err != nil {
		fmt.Printf("Error to Dump on namespace %s: %v\n", namespace, err)
		return
	}
	output_file := fmt.Sprintf("./%s/%s/00-namespace.%s", project_name, namespace, format)
	err = files.WriteFile(output_file, output)
	if err != nil {
		fmt.Printf("Erro to write file: %v\n", err)
		return
	}
}
