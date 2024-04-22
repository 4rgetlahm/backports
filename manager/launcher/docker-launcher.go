package launcher

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/4rgetlahm/backports/api/service"
	"github.com/4rgetlahm/backports/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DockerLauncher struct{}

var dockerClient *client.Client

func (launcher DockerLauncher) InitClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	return cli
}

func (launcher DockerLauncher) LaunchVolumeGenerationJob(volumeName string, cloneUrl string, credentials string, overwrite bool) (string, error) {
	if dockerClient == nil {
		dockerClient = launcher.InitClient()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if overwrite {
		log.Printf("Removing volume %s", volumeName)
		err := dockerClient.VolumeRemove(ctx, volumeName, true)

		if err != nil {
			return "", err
		}
	}

	_, err := dockerClient.VolumeCreate(ctx, volume.CreateOptions{
		Name: volumeName,
	})

	if err != nil {
		return "", err
	}

	service.UpdateVolumeStatus(volumeName, types.VolumeStatusCreating)

	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image: "4rgetlahm/repo-cloner:1.0",
		Env: []string{
			"CLONE_URL=" + cloneUrl,
			"CREDENTIALS=" + credentials,
		},
	},
		&container.HostConfig{
			RestartPolicy: container.RestartPolicy{
				Name:              container.RestartPolicyDisabled,
				MaximumRetryCount: 0,
			},
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeVolume,
					Source: volumeName,
					Target: "/repo",
				},
			},
		}, nil, nil, "volume-generator-"+volumeName)

	if err != nil {
		return "", err
	}

	err = dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{})

	if err != nil {
		return "", err
	}

	go launcher.UpdateVolumeStateAndRemoveContainerPostContainerExit(resp.ID, volumeName)

	return resp.ID, nil
}

func (launcher DockerLauncher) LaunchBackportJob(volume string, reference primitive.ObjectID, newBranchName string, targetBranchName string, commits []string) error {
	if dockerClient == nil {
		dockerClient = launcher.InitClient()
	}

	service.AddEvent(reference, types.ActionVirtualMachinePreparing, "")

	go launcher.launchBackportJob(volume, reference, newBranchName, targetBranchName, commits)

	return nil
}

func (launcher DockerLauncher) launchBackportJob(volume string, reference primitive.ObjectID, newBranchName string, targetBranchName string, commits []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newVolumeName := "backport-automation-volume-" + reference.Hex()
	_, err := launcher.DuplicateVolume(volume, newVolumeName)

	if err != nil {
		service.AddEvent(reference, types.ActionVirtualMachineError, err.Error())
		return err
	}

	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image: "4rgetlahm/backport-runner:1.0",
		Env: []string{
			"REFERENCE=" + reference.Hex(),
			"NEW_BRANCH_NAME=" + newBranchName,
			"TARGET_BRANCH_NAME=" + targetBranchName,
			"COMMITS=" + strings.Join(commits, ","),
			"SOURCE_PATH=/repo",
			"REPORTER_CONFIG=" + "",
		},
	},
		&container.HostConfig{
			RestartPolicy: container.RestartPolicy{
				Name:              container.RestartPolicyDisabled,
				MaximumRetryCount: 0,
			},
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeVolume,
					Source: newVolumeName,
					Target: "/repo",
				},
			},
		}, nil, nil, "backport-job-"+reference.Hex())

	if err != nil {
		service.AddEvent(reference, types.ActionVirtualMachineError, err.Error())
		return err
	}

	err = dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{})

	if err != nil {
		service.AddEvent(reference, types.ActionVirtualMachineError, err.Error())
		return err
	}

	go launcher.RemoveContainerAndVolumePostContainerExit(resp.ID, newVolumeName)
	service.AddEvent(reference, types.ActionVirtualMachineCreated, resp.ID)

	return nil
}

func (Launcher DockerLauncher) DuplicateVolume(volumeName string, newVolumeName string) (string, error) {
	if dockerClient == nil {
		dockerClient = Launcher.InitClient()
	}

	log.Printf("Duplicating volume %s to %s", volumeName, newVolumeName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dockerClient.VolumeInspect(ctx, volumeName)

	if err != nil {
		return "", err
	}

	_, err = dockerClient.VolumeCreate(ctx, volume.CreateOptions{
		Name: newVolumeName,
	})

	if err != nil {
		return "", err
	}

	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image: "4rgetlahm/volume-duplicator:1.0",
	},
		&container.HostConfig{
			RestartPolicy: container.RestartPolicy{
				Name:              container.RestartPolicyDisabled,
				MaximumRetryCount: 0,
			},
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeVolume,
					Source: volumeName,
					Target: "/source",
				},
				{
					Type:   mount.TypeVolume,
					Source: newVolumeName,
					Target: "/target",
				},
			},
		}, nil, nil, "volume-duplicator-"+volumeName+"-"+newVolumeName)

	if err != nil {
		return "", err
	}

	err = dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		return "", err
	}

	exitCh, _ := dockerClient.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)

	select {
	case <-exitCh:
		log.Println("Volume maker finished: " + resp.ID)
		log.Println("Removing volume maker container: " + resp.ID)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dockerClient.ContainerRemove(ctx, resp.ID, container.RemoveOptions{
			Force: true,
		})
		break
	case <-time.After(10 * time.Minute):
		break
	}

	log.Printf("Volume %s duplicated to %s", volumeName, newVolumeName)

	return resp.ID, nil
}

func (launcher DockerLauncher) UpdateVolumeStateAndRemoveContainerPostContainerExit(containerID string, volumeName string) {
	log.Println("Awaiting container to finish: " + containerID)
	exitCh, err := dockerClient.ContainerWait(context.Background(), containerID, container.WaitConditionNotRunning)

	select {
	case <-exitCh:
		log.Println("Container finished: " + containerID)
		log.Println("Removing container: " + containerID)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dockerClient.ContainerRemove(ctx, containerID, container.RemoveOptions{
			Force: true,
		})

		service.UpdateVolumeStatus(volumeName, types.VolumeStatusReady)
		break
	case <-time.After(10 * time.Minute):
		break
	}

	if err != nil {
		return
	}
}

func (launcher DockerLauncher) RemoveContainerAndVolumePostContainerExit(containerID string, volumeName string) {
	log.Println("Awaiting container to finish: " + containerID)
	exitCh, err := dockerClient.ContainerWait(context.Background(), containerID, container.WaitConditionNotRunning)

	select {
	case <-exitCh:
		log.Println("Container finished: " + containerID)
		log.Println("Removing container: " + containerID)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dockerClient.ContainerRemove(ctx, containerID, container.RemoveOptions{
			Force: true,
		})

		log.Println("Removing volume: " + volumeName)
		err := dockerClient.VolumeRemove(ctx, volumeName, true)

		if err != nil {
			log.Println("Failed to remove volume: " + volumeName)
		}
		break
	case <-time.After(10 * time.Minute):
		break
	}

	if err != nil {
		return
	}
}
