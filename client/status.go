package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type Status struct {
	Device struct {
		SerialNumber     string `json:"serial_number"`
		UsbMac           string `json:"usb_mac"`
		WifiMac          string `json:"wifi_mac"`
		BleMac           string `json:"ble_mac"`
		OtpValid         bool   `json:"otp_valid"`
		OtpModel         string `json:"otp_model"`
		OtpTimestamp     int    `json:"otp_timestamp"`
		FirmwareSecurity string `json:"firmware_security"`
	} `json:"device"`
	Firmware struct {
		Version       string `json:"version"`
		Target        int    `json:"target"`
		Branch        string `json:"branch"`
		BuildDate     string `json:"build_date"`
		CommitHash    string `json:"commit_hash"`
		NwpVersion    string `json:"nwp_version"`
		MatterVersion string `json:"matter_version"`
	} `json:"firmware"`
	System struct {
		APISemver         string `json:"api_semver"`
		Uptime            string `json:"uptime"`
		BootTime          int    `json:"boot_time"`
		AutoUpdateEnabled bool   `json:"auto_update_enabled"`
	} `json:"system"`
	Power struct {
		State          string `json:"state"`
		BatteryCharge  int    `json:"battery_charge"`
		BatteryVoltage int    `json:"battery_voltage"`
		BatteryCurrent int    `json:"battery_current"`
		UsbVoltage     int    `json:"usb_voltage"`
	} `json:"power"`
}

func (s *Status) PrettyPrint() {
	fmt.Printf("\nDevice Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Device Serial Number")
	row.Column(s.Device.SerialNumber)

	row = tab.Row()
	row.Column("Firmware Version")
	row.Column(s.Firmware.Version)

	row = tab.Row()
	row.Column("System Uptime")
	row.Column(s.System.Uptime)

	row = tab.Row()
	row.Column("Power State")
	row.Column(s.Power.State)

	tab.Print(os.Stdout)
	fmt.Println()
}
