package main

import (
	"fmt"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

const (
	imageName = "furikuri/chain-reaction"
)

func main() {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}

	if cleanUp() {
		fmt.Println("Do cleanup")
		removeAllImageContainer(cli)
	} else {

		time.Sleep(5 * time.Second)
		stopPrevContainer(counter(), cli)
		time.Sleep(5 * time.Second)
		startChainContainer(counter(), cli)

		time.Sleep(100 * time.Second)
	}

}

func startChainContainer(counter int, cli *client.Client) {
	binds := []string{
		"/var/run/docker.sock:/var/run/docker.sock",
	}
	c, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		Cmd:   []string{"--counter", strconv.Itoa(counter - 1)},
	}, &container.HostConfig{
		Binds: binds,
	}, nil, "/chain-reaction-"+strconv.Itoa(counter-1))

	if err != nil {
		panic(err)
	}

	err2 := cli.ContainerStart(context.Background(), c.ID, types.ContainerStartOptions{})

	if err2 != nil {
		panic(err2)
	}
}

func removeAllImageContainer(cli *client.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var cleanUpContainer types.Container
	for _, c := range containers {
		if c.Image == imageName {
			if c.Command != "--cleanup" {
				cleanUpContainer = c
			} else {
				removeContainer(c, cli)

			}
		}
	}
	removeContainer(cleanUpContainer, cli)

}

func stopPrevContainer(counter int, cli *client.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	count := 0
	var imageContainers [2]types.Container
	for _, c := range containers {
		if c.Image == imageName {
			if count < 2 {
				imageContainers[count] = c
			}
			count++

		}
	}

	if count == 1 {
		// do nothing
	} else if count == 2 {
		for _, c := range imageContainers {
			if c.Names[0] != "/chain-reaction-"+strconv.Itoa(counter) {
				fmt.Println("remove " + c.Names[0] + ". Counter was " + strconv.Itoa(counter))
				removeContainer(c, cli)
			}
		}
	} else {
		panic("Something went wrong")
	}
}

func removeContainer(container types.Container, cli *client.Client) {
	cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
		Force: true,
	})
}
