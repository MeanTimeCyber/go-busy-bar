package main

import (
	"context"
	"flag"

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
			panic(err)
		}
		status.PrettyPrint()
	default:
		println("Unknown command:", command)
		printCommands()
	}
}

func printCommands() {
	println("Available commands:")
	println("  status - Get the status of the busy bar")
}
