package restore

import (
	"fmt"
	"kubedump/pkg/exec"
)

// Restore a namespace dumped by kubedump
func Namespace(namespace string, projectName string, kubectl string) {
	fmt.Printf("Restoring namespace '%s'\n", namespace)
	resourcePath := fmt.Sprintf("./%s/%s/00*", projectName, namespace)
	restoreCmd := fmt.Sprintf("%s apply -f %s -n %s", kubectl, resourcePath, namespace)

	output, err := exec.SoExec(restoreCmd)
	if err != nil {
		fmt.Printf("Error to restore namespace %s: %v %v\n", namespace, output, err)
		return
	}
}

// Restore all resource file exported by kubedump
func Resource(resourcePath string, namespace string, kubectl string) {
	fmt.Printf("Restoring %s on namespace '%s'\n", resourcePath, namespace)

	restoreCmd := fmt.Sprintf("%s apply -f %s -n %s", kubectl, resourcePath, namespace)

	output, err := exec.SoExec(restoreCmd)
	if err != nil {
		fmt.Printf("Error to restore resource on namespace %s: %v %v\n", namespace, output, err)
		return
	}
}
