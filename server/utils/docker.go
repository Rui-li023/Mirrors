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

type DockerImage struct {
	Containers   string `json:"Containers"`
	CreatedAt    string `json:"CreatedAt"`
	CreatedSince string `json:"CreatedSince"`
	Digest       string `json:"Digest"`
	ID           string `json:"ID"`
	Repository   string `json:"Repository"`
	SharedSize   string `json:"SharedSize"`
	Size         string `json:"Size"`
	Tag          string `json:"Tag"`
	UniqueSize   string `json:"UniqueSize"`
	VirtualSize  string `json:"VirtualSize"`
}

// GetDockerImages returns a list of DockerImage with details of all images
func GetDockerImages() ([]environment.Images, error) {
	cmd := exec.Command("docker", "images", "--format", "json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	images := make([]environment.Images, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		var image DockerImage
		err = json.Unmarshal([]byte(line), &image)
		if err != nil {
			fmt.Println(err)
			continue
		}

		image_ := environment.Images{
			Image_Id:     image.ID,
			Repository:   image.Repository,
			Tag:          image.Tag,
			Size:         image.Size,
			CreatedSince: image.CreatedSince,
		}
		images = append(images, image_)
	}
	return images, nil
}

// GetDockerContainers returns a list of ContainerInfo with details of all running containers
func GetDockerContainers() ([]environment.Container, error) {

	// Run the docker ps command
	cmd := exec.Command("docker", "ps", "-a", "--format", "json")

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
