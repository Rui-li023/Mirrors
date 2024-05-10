package utils

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/environment"
	"os/exec"
	"strings"
)

type Container struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
}

// GetDockerContainers returns a list of ContainerInfo with details of all running containers
func GetDockerContainers() ([]environment.Container, error) {

	// Run the docker ps command
	cmd := exec.Command("docker", "ps", "--format", "json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	containers := make([]environment.Container, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		var container Container
		err = json.Unmarshal([]byte(line), &container)
		if err != nil {
			fmt.Println(err)
			continue
		}

		container_ := environment.Container{
			ContainerId: container.ID,
			ImageName:   container.Image,
			Status:      container.Status,
			Create:      container.CreatedAt,
			Command:     container.Command,
			Ports:       container.Ports,
		}
		containers = append(containers, container_)
	}
	return containers, nil
}
