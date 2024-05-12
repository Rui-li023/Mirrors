package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/environment"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"go.uber.org/zap"
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
		global.GVA_LOG.Error("运行查询命令失败!", zap.Error(err))
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
			global.GVA_LOG.Error("转换JSON失败!", zap.Error(err))
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

func GetDockerContainers() ([]environment.Container, error) {

	// Run the docker ps command
	cmd := exec.Command("docker", "ps", "-a", "--format", "json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		global.GVA_LOG.Error("运行查询命令失败!", zap.Error(err))
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	containers := make([]environment.Container, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		var container_ Container
		err = json.Unmarshal([]byte(line), &container_)
		if err != nil {
			global.GVA_LOG.Error("转换JSON失败!", zap.Error(err))
			continue
		}
		container__ := environment.Container{
			ContainerId: container_.ID,
			ImageName:   container_.Image,
			Status:      container_.Status,
			Create:      container_.CreatedAt,
			Command:     container_.Command,
			Ports:       container_.Ports,
			Name:        container_.Names,
		}
		containers = append(containers, container__)
	}
	return containers, nil
}

func StartContainerByID(id string) error {
	// Create a new Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Get the container by ID
	container_, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	// Start the container
	err = cli.ContainerStart(ctx, container_.ID, container.StartOptions{})
	if err != nil {
		return err
	}
	global.GVA_LOG.Info("Container started successfully: ", zap.String("id", id))
	return nil
}

func StopContainerByID(id string) error {
	// Create a new Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Get the container by ID
	container_, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	// Stop the container
	err = cli.ContainerStop(ctx, container_.ID, container.StopOptions{})
	if err != nil {
		return err
	}
	global.GVA_LOG.Info("Container stopped successfully: ", zap.String("id", id))
	return nil
}

func RemoveContainerByID(id string) error {
	// Create a new Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Get the container by ID
	container_, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	// Remove the container
	err = cli.ContainerRemove(ctx, container_.ID, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}
	global.GVA_LOG.Info("Container removed successfully: ", zap.String("id", id))
	return nil
}

func RestartContainerByID(id string) error {
	// Create a new Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Get the container by ID
	container_, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	// Restart the container
	err = cli.ContainerRestart(ctx, container_.ID, container.StopOptions{})
	if err != nil {
		return err
	}
	global.GVA_LOG.Info("Container restarted successfully: ", zap.String("id", id))
	return nil
}

func DeleteImageByID(id string) error {
	// Create a new Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	image_, _, err := cli.ImageInspectWithRaw(ctx, id)
	if err != nil {
		return err
	}

	_, err = cli.ImageRemove(ctx, image_.ID, image.RemoveOptions{})
	if err != nil {
		return err
	}
	global.GVA_LOG.Info("Image deleted successfully: ", zap.String("id", id))
	return nil
}

func PullImage(imageName, imageTag string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	if imageTag == "" {
		_, err = cli.ImagePull(ctx, fmt.Sprintf("%s", imageName), image.PullOptions{})
	} else {
		_, err = cli.ImagePull(ctx, fmt.Sprintf("%s:%s", imageName, imageTag), image.PullOptions{})
	}

	if err != nil {
		return err
	}

	global.GVA_LOG.Info("Image pulled successfully: ", zap.String("image", fmt.Sprintf("%s:%s", imageName, imageTag)))
	return nil
}

// CreateContainer 创建容器
func CreateContainer(containerConfig environment.ContainerConfig) error {
	// 创建 Docker 客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	// 创建容器配置
	config := &container.Config{
		Image: containerConfig.ImageName,
		Cmd:   containerConfig.Cmds,
	}
	hostConfig := &container.HostConfig{}
	portBindings := nat.PortMap{}

	for _, pm := range containerConfig.PortMappings {
		containerPort := nat.Port(fmt.Sprintf("%s/%s", pm.ContainerPort, pm.ContainerType))
		fmt.Println(containerPort)
		fmt.Println(pm.HostPort)

		portBindings[containerPort] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: pm.HostPort,
			},
		}
	}
	hostConfig.PortBindings = portBindings
	// Create a new networking config (optional)
	networkingConfig := &network.NetworkingConfig{}

	// Create a new platform config (optional)
	platform := &ocispec.Platform{}

	// 创建容器
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, platform, containerConfig.ContainerName)
	if err != nil {
		return err
	}

	// 启动容器
	err = cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{})
	if err != nil {
		return err
	}

	return nil
}
