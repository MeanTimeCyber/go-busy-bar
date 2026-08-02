package main

import (
	"context"
	"fmt"

	barclient "github.com/MeanTimeCyber/go-busy-bar/client"
)

// printCommands prints the available commands for the busy bar.
func printCommands() {
	fmt.Println("Available commands:")
	fmt.Println("  status - Get the status of the busy bar")
	fmt.Println("  info - Get the device info of the busy bar")
	fmt.Println("  version - Get the API version of the busy bar")
	fmt.Println("  firmware - Get the firmware version of the busy bar")
	fmt.Println("  settings - Get the settings of the busy bar")
	fmt.Println("  storage - Get the storage info of the busy bar")
	fmt.Println("  help - Show this help message")
}

// printStatus retrieves and prints the device status.
func printStatus(client *barclient.Client, ctx context.Context) {
	status, err := client.GetStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting device status: %v\n", err)
		return
	}
	status.PrettyPrint()

	networkStatus, err := client.GetWifiStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting network status: %v\n", err)
		return
	}
	networkStatus.PrettyPrint()

	bleStatus, err := client.GetBleStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting BLE status: %v\n", err)
		return
	}
	bleStatus.PrettyPrint()
}

// printSettings retrieves and prints the HTTP access info, display brightness, and audio volume settings.
func printSettings(client *barclient.Client, ctx context.Context) {
	httpAccessInfo, err := client.GetHttpAccess(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting HTTP access info: %v\n", err)
		return
	}
	httpAccessInfo.PrettyPrint()

	displayBrightnessInfo, err := client.GetDisplayBrightness(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting display brightness info: %v\n", err)
		return
	}
	displayBrightnessInfo.PrettyPrint()

	audioVolumeInfo, err := client.GetAudioVolume(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting audio volume info: %v\n", err)
		return
	}
	audioVolumeInfo.PrettyPrint()
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

// printFirmware retrieves the firmware version from the client and prints it to the console.
func printFirmware(client *barclient.Client, ctx context.Context) {
	// get the firmware status from the client
	firmwareStatus, err := client.GetStatusFirmware(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting firmware status: %v\n", err)
		return
	}
	firmwareStatus.PrettyPrint()

	// Get the firmware update information from the client
	firmwareVersion, err := client.GetFirmwareUpdateStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting firmware version: %v\n", err)
		return
	}
	firmwareVersion.PrettyPrint()

	// get the auto-update status from the client
	autoUpdateStatus, err := client.GetAutoupdateSettings(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting auto-update status: %v\n", err)
		return
	}
	autoUpdateStatus.PrettyPrint()
}

func printStorage(client *barclient.Client, ctx context.Context) {
	storageStatus, err := client.GetStorageStatus(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting storage status: %v\n", err)
		return
	}
	storageStatus.PrettyPrint()

	storageList, err := client.ListStorageFiles(ctx, "/ext")
	if err != nil {
		fmt.Printf("Error getting storage list: %v\n", err)
		return
	}
	storageList.PrettyPrint()
}
