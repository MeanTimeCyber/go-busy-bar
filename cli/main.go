package main

import (
	"context"

	"github.com/MeanTimeCyber/go-busy-bar/client"
)

const (
	barUSBHTTPEndpoint = "http://10.0.4.20/"
)

func main() {
	// Create a new API client with the specified endpoint.
	client := client.NewClient(barUSBHTTPEndpoint)

	// Call the GetStatus method to retrieve the status of the busy bar.
	statusResp, err := client.GetStatus(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	statusResp.PrettyPrint()
}
