package restore

import (
	"fmt"
	"kubedump/pkg/exec"

	"github.com/charmbracelet/log"
)

// Namespace restore a namespace dumped by kubedump
func Namespace(namespace string, projectName string, kubectl string) {
	log.Info("Restoring resources", "namespace", namespace, "resource", "namespace")
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
	log.Debug("Restoring resources", "namespace", namespace, "resources", resourcePath)
	restoreCmd := fmt.Sprintf("%s apply -f %s -n %s", kubectl, resourcePath, namespace)

	output, err := exec.SoExec(restoreCmd)
	if err != nil {
		log.Warn("Error to restore resource on:", "namespace", namespace, "resources", resourcePath, "error", err.Error(), "output", output)
		return
	}
	log.Info("Resources restored:", "namespace", namespace, "resources", resourcePath)
}
