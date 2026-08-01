package main

import (
	"context"
	"fmt"

	barclient "github.com/MeanTimeCyber/go-busy-bar/client"
)

// printCommands prints the available commands for the busy bar.
func printCommands() {
	println("Available commands:")
	println("  status - Get the status of the busy bar")
	println("  info - Get the device info of the busy bar")
	println("  version - Get the API version of the busy bar")
}

// printStatus retrieves and prints the device status.
func printStatus(client *barclient.Client, ctx context.Context) {
	status, err := client.GetStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting device status: %v\n", err)
		return
	}
	status.PrettyPrint()
}

// printInfo retrieves and prints the device name, info, network status, and account details.
func printInfo(client *barclient.Client, ctx context.Context) {
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

	networkStatus, err := client.GetWifiStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting network status: %v\n", err)
		return
	}
	networkStatus.PrettyPrint()

	accountInfo, err := client.GetAccountInfo(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting account info: %v\n", err)
		return
	}
	accountInfo.PrettyPrint()
}

// printVersion retrieves the API version from the client and prints it to the console.
func printVersion(client *barclient.Client, ctx context.Context) {
	apiVersion, err := client.GetVersion(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting API version: %v\n", err)
		return
	}
	apiVersion.PrettyPrint()
}
