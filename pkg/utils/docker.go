package utils

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mattn/go-shellwords"
)

func PullDockerImage(image string) error {
	isexist, err := ImageExists(image)
	if err != nil {
		return err
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	// Pull Image if not exist

	if !isexist {
		visuals.PrintState("Task-Info", "Initialize", "Pulling Docker Image: "+visuals.PrintRandomColor(image))
		out1, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
		if err != nil {
			return fmt.Errorf("failed to pull Docker image '%s': %v", image, err)
		}
		defer out1.Close()
		_, err = io.Copy(io.Discard, out1)
		if err != nil {
			return fmt.Errorf("failed to copy Docker image pull data '%s': %v", image, err)
		}

	}
	return nil
}

func ImageExists(imageName string) (bool, error) {
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return false, fmt.Errorf("failed to create Docker client: %v", err)
	}

	// Check if the image exists
	_, _, err = cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		if client.IsErrNotFound(err) {
			// Image not found
			return false, nil
		}
		// Other error occurred
		return false, fmt.Errorf("failed to inspect Docker image '%s': %v", imageName, err)
	}

	// Image exists
	return true, nil
}

func InspectImageEntrypoint(imageName string, defaultCmd string) ([]string, error) {
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %v", err)
	}

	// Inspect the Docker image
	imgInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect Docker image '%s': %v", imageName, err)
	}

	// Use the image's entrypoint as the Cmd
	var cmd []string
	if len(imgInspect.Config.Entrypoint) == 0 {
		// Use /bin/sh -c as the default command when no entrypoint is specified
		cmd = []string{"/bin/sh", "-c", defaultCmd}
	} else {
		// Parse the command considering shell syntax
		cmd, err = shellwords.Parse(defaultCmd)
		if err != nil {
			return nil, fmt.Errorf("failed to parse command '%s': %v", defaultCmd, err)
		}
	}

	return cmd, nil
}
