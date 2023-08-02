package restore

import (
	"fmt"
	"kubedump/pkg/exec"
)

// Namespace restore a namespace dumped by kubedump
func Namespace(namespace string, projectName string, kubectl string) {
	fmt.Printf("Restoring namespace '%s'\n", namespace)
	resourcePath := fmt.Sprintf("./%s/%s/00*", projectName, namespace)
	restoreCmd := fmt.Sprintf("%s apply -f %s -n %s --validate=false", kubectl, resourcePath, namespace)

	output, err := exec.SoExec(restoreCmd)
	if err != nil {
		fmt.Printf("Error to restore namespace %s: %v %v\n", namespace, output, err)
		return
	}
}

// Resource restore all resource file exported by kubedump
func Resource(resourcePath string, namespace string, kubectl string) {

	restoreCmd := fmt.Sprintf("%s apply -f %s -n %s", kubectl, resourcePath, namespace)

	output, err := exec.SoExec(restoreCmd)
	if err != nil {
		fmt.Printf("Error to restore %s resource on namespace %s: %v %v\n", resourcePath, namespace, output, err)
		return
	}

	fmt.Printf("Restored %s resources on namespace '%s'\n", resourcePath, namespace)

}
