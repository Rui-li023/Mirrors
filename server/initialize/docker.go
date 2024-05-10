package initialize

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"time"
)

var containerService = service.ServiceGroupApp.EnvironmentServiceGroup.ContainerService

// ListenDockerEvents handles setting up the Docker client and listening for Docker events.
func ListenDockerEvents() error {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	// Create a context that is never cancelled
	ctx := context.Background()

	// Listen to Docker events
	eventsChan, errChan := cli.Events(ctx, types.EventsOptions{})

	// Process events and errors
	for {
		select {
		case event := <-eventsChan:
			handleEvent(cli, event)
		case err := <-errChan:
			return err
		}
	}
}

// handleEvent processes Docker events based on their type.
func handleEvent(cli *client.Client, event events.Message) {
	switch event.Type {
	case "container":
		handleContainerEvent(cli, event.Actor.ID, event.Actor.Attributes)
	case "image":
		handleImageEvent(event.Actor.Attributes)
	default:
		// Handle other types of events
	}
}

// handleContainerEvent handles events related to Docker containers.
func handleContainerEvent(cli *client.Client, containerID string, eventDetails map[string]string) {

	//containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	//if err != nil {
	//	fmt.Printf("Error getting container information: %s\n", err)
	//	return
	//}
	// 获取容器信息
	//containerInfo, err := getContainerInfo(cli, containerIDOrName)
	//if err != nil {
	//	fmt.Println("Error getting container information:", err)
	//	return
	//}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		fmt.Println("Error listing containers:", err)
		return
	}

	for _, container := range containers {
		fmt.Printf("Container ID: %s\n", container.ID[:12]) // 显示容器ID的前12个字符
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Command: \"%s\"\n", container.Command)
		fmt.Printf("Created: %s\n", time.Unix(container.Created, 0).Format(time.RFC1123))
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Printf("Ports: %s\n", formatPorts(container.Ports))
		fmt.Printf("Names: %s\n", container.Names[0][1:]) // 从路径中剔除前导'/'
		fmt.Println("-----------------------------------")
	}
	// 打印容器信息
	//fmt.Printf("Container ID: %s\n", containerInfo.ID)
	//fmt.Printf("Image: %s\n", containerInfo.Image)
	//fmt.Printf("Command: %s\n", containerInfo.Command)
	//fmt.Printf("Created: %s\n", containerInfo.Created)
	//fmt.Printf("Status: %s\n", containerInfo.State.Status)

	//var container environment.Container
	//container.ContainerId = containerInfo.ID
	//container.Create = containerInfo.Created
	//container.ImageName = containerInfo.Config.Image
	//container.Status = containerInfo.State.Status
	//container.Ports = containerInfo.
	//
	//err = containerService.CreateContainer(&container)
	//if err != nil {
	//	return
	//}

	// Print more detailed information
	//fmt.Printf("Detailed Container Info - Name: %s, Image: %s, State: %s, StartedAt: %s \n",
	//	containerInfo.Name, containerInfo.Config.Image, containerInfo.State.Status, containerInfo.State.StartedAt)

	// 遍历打印 map 中的键值对
	for key, value := range eventDetails {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

// handleImageEvent handles events related to Docker images.
func handleImageEvent(eventDetails map[string]string) {
	for key, value := range eventDetails {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	// Here you could update the database with image information based on ID and action
}

// formatPorts formats the port mappings for display.
func formatPorts(ports []types.Port) string {
	var result string
	for _, port := range ports {
		result += fmt.Sprintf("%d/%s -> %d; ", port.PrivatePort, port.Type, port.PublicPort)
	}
	if len(result) > 0 {
		result = result[:len(result)-2] // Remove the last semicolon and space
	}
	return result
}
