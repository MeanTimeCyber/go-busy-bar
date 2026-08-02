package main

import (
	"context"
	"flag"
	"fmt"
)

// as per https://docs.busy.app/bar/dev/http-api#via-usb
const defaultUSBAddress = "10.0.4.20"

func main() {
	// define and process args
	var command, ipAddress, password string
	var help bool
	flag.BoolVar(&help, "?", false, "Show help")
	flag.StringVar(&command, "c", "status", "Command for the busy bar")
	flag.StringVar(&ipAddress, "ip", defaultUSBAddress, "IP address of the busy bar - assumes it's over USB on the fixed address")
	flag.StringVar(&password, "p", "", "Password for the busy bar - required if on a wifi address")
	flag.Parse()

	if help {
		flag.Usage()
		fmt.Println()
		printCommands()
		return
	}

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
	err = checkConnection(ctx, client)

	if err != nil {
		fmt.Printf("Error checking connection: %v\n", err)
		return
	}

	// Execute the provided command
	switch command {
	// get the device status information and print it to the console
	case "status":
		printStatus(client, ctx)
	// get the device name, info and account details and print them to the console
	case "info":
		printInfo(client, ctx)
	// get the API version and print it to the console
	case "version":
		printVersion(client, ctx)
	case "firmware":
		printFirmware(client, ctx)
	case "settings":
		printSettings(client, ctx)
	case "storage":
		printStorage(client, ctx)
	// show the help message
	case "help":
		printCommands()
	default:
		println("Unknown command:", command)
		printCommands()
	}
}
