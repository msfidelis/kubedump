package restore

import (
	"fmt"
	"kubedump/pkg/exec"
)

func RestoreNamespace(namespace string, project_name string, kubectl string) {
	fmt.Printf("Restoring namespace '%s'\n", namespace)
	resource_path := fmt.Sprintf("./%s/%s/00*", project_name, namespace)
	restore_cmd := fmt.Sprintf("%s apply -f %s -n %s", kubectl, resource_path, namespace)

	output, err := exec.SoExec(restore_cmd)
	if err != nil {
		fmt.Printf("Error to restore namespace %s: %v %v\n", namespace, output, err)
		return
	}
}

func RestoreResource(resource_path string, namespace string, kubectl string) {
	fmt.Printf("Restoring %s on namespace '%s'\n", resource_path, namespace)

	restore_cmd := fmt.Sprintf("%s apply -f %s -n %s", kubectl, resource_path, namespace)

	output, err := exec.SoExec(restore_cmd)
	if err != nil {
		fmt.Printf("Error to restore resource on namespace %s: %v %v\n", namespace, output, err)
		return
	}
}
