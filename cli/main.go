package main

import (
	"context"
	"flag"
	"fmt"

	barclient "github.com/MeanTimeCyber/go-busy-bar/client"
)

const defaultUSBAddress = "10.0.4.20"

func main() {
	// define and process args
	var command, ipAddress, password string
	flag.StringVar(&command, "c", "status", "Command for the busy bar")
	flag.StringVar(&ipAddress, "ip", defaultUSBAddress, "IP address of the busy bar - assumes it's over USB on the fixed address")
	flag.StringVar(&password, "p", "", "Password for the busy bar - required if on a wifi address")
	flag.Parse()

	// Check if a command was provided
	if command == "" {
		println("No command specified")
		flag.Usage()
		printCommands()
		return
	}

	// Create a new client with the specified endpoint and password
	client, err := getClient(ipAddress, password)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	// Create a context for the API calls
	ctx := context.Background()

	// TODO get network status to check we're connected to the device, and if not, print a message and exit

	// Execute the provided command
	switch command {
	// get the device status information and print it to the console
	case "status":
		status, err := client.GetStatus(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting device status: %v\n", err)
			return
		}
		status.PrettyPrint()
	// get the device name, info and account details and print them to the console
	case "info":
		deviceName, err := client.GetName(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting device name: %v\n", err)
			return
		}
		deviceName.PrettyPrint()

		deviceInfo, err := client.GetStatusDevice(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting device info: %v\n", err)
			return
		}
		deviceInfo.PrettyPrint()

		accountInfo, err := client.GetAccountInfo(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting account info: %v\n", err)
			return
		}
		accountInfo.PrettyPrint()
	// get the API version and print it to the console
	case "version":
		apiVersion, err := client.GetVersion(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting API version: %v\n", err)
			return
		}
		apiVersion.PrettyPrint()
	default:
		println("Unknown command:", command)
		printCommands()
	}
}

// getClient creates a new Busy Bar client with the specified IP address and password.
func getClient(ipAddress, password string) (*barclient.Client, error) {
	var client *barclient.Client

	// Check we have a password if on a Wi-Fi address
	if ipAddress != defaultUSBAddress {
		// If the IP address is not the default USB address, we assume it's a Wi-Fi address and require a password
		if password == "" {
			flag.Usage()
			return nil, fmt.Errorf("password is required for access over Wi-Fi address")
		}

		// Create a new client with the specified endpoint and password
		client = barclient.NewClientWithAPIKey("http://" + ipAddress +"/", password)

	} else {
		// Create a new client with the specified endpoint
		client = barclient.NewClient("http://" + ipAddress + "/")
	}

	return client, nil
}

// printCommands prints the available commands for the busy bar.
func printCommands() {
	println("Available commands:")
	println("  status - Get the status of the busy bar")
	println("  info - Get the device info of the busy bar")
	println("  version - Get the API version of the busy bar")
}
