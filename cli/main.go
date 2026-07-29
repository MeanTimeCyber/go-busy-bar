package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/MeanTimeCyber/go-busy-bar/client"
)

func main() {
	// define and process args
	var command, ipAddress string
	flag.StringVar(&command, "c", "status", "Command for the busy bar")
	flag.StringVar(&ipAddress, "ip", "10.0.4.20", "IP address of the busy bar - assumes it's over USB on the fixed address")
	flag.Parse()

	if command == "" {
		println("No command specified")
		flag.Usage()
		printCommands()
		return
	}

	// Create a new client with the specified endpoint
	ctx := context.Background()
	c := client.NewClient("http://" + ipAddress + "/")

	switch command {
	// get the device status information and print it to the console
	case "status":
		status, err := c.GetStatus(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting device status: %v\n", err)
			return
		}
		status.PrettyPrint()
	// get the device name, info and account details and print them to the console
	case "info":
		deviceName, err := c.GetName(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting device name: %v\n", err)
			return
		}
		deviceName.PrettyPrint()

		deviceInfo, err := c.GetStatusDevice(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting device info: %v\n", err)
			return
		}
		deviceInfo.PrettyPrint()

		accountInfo, err := c.GetAccountInfo(ctx, nil)
		if err != nil {
			fmt.Printf("Error getting account info: %v\n", err)
			return
		}
		accountInfo.PrettyPrint()
	default:
		println("Unknown command:", command)
		printCommands()
	}
}

func printCommands() {
	println("Available commands:")
	println("  status - Get the status of the busy bar")
}
