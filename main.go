package main

import (
	"fmt"
	"log"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "tcp://0.0.0.0:2375"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	imgs, err := client.ListImages(true)
	if err != nil {
		log.Fatalln(err)
	}
	for _, img := range imgs {
		fmt.Println(img.ID)
	}
}
