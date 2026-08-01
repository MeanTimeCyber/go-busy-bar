package main

import (
	"context"
	"flag"
	"fmt"

	barclient "github.com/MeanTimeCyber/go-busy-bar/client"
)

// checkConnection checks if the client is connected to the busy bar device.
func checkConnection(ctx context.Context, client *barclient.Client) error {
	connection, err := client.GetTransport(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Connected via: %s\n", connection.Type)

	return nil
}

// getClient creates a new Busy Bar client with the specified IP address and password.
func getClient(ipAddress, password string) (*barclient.Client, error) {
	var client *barclient.Client

	// Check we have a password if on a Wi-Fi address
	if ipAddress != defaultUSBAddress {
		// If the IP address is not the default USB address, we assume it's a Wi-Fi address and require a password
		if password == "" {
			flag.Usage()
			fmt.Println()
			return nil, fmt.Errorf("password is required for access over Wi-Fi address. See https://docs.busy.app/bar/dev/http-api#via-wi-fi for instructions")
		}

		// Create a new client with the specified endpoint and password
		client = barclient.NewClientWithAPIKey("http://"+ipAddress+"/", password)

	} else {
		// Create a new client with the specified endpoint
		client = barclient.NewClient("http://" + ipAddress + "/")
	}

	return client, nil
}
