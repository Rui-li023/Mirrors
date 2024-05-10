package initialize

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
)

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

	containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		fmt.Printf("Error getting container information: %s\n", err)
		return
	}

	// Print more detailed information
	fmt.Printf("Detailed Container Info - Name: %s, Image: %s, State: %s, StartedAt: %s \n",
		containerInfo.Name, containerInfo.Config.Image, containerInfo.State.Status, containerInfo.State.StartedAt)

	// 遍历打印 map 中的键值对
	for key, value := range eventDetails {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

// handleImageEvent handles events related to Docker images.
func handleImageEvent(eventDetails map[string]string) {
	imageID := eventDetails["id"]
	action := eventDetails["action"]
	for key, value := range eventDetails {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	fmt.Printf("Image event - ID: %s, Action: %s\n", imageID, action)
	// Here you could update the database with image information based on ID and action
}
