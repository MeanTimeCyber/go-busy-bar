package client

import (
	"fmt"
	"os"
	"time"

	"github.com/markkurossi/tabulate"
)

type Status struct {
	// Device contains device identification and security info.
	Device StatusDevice `json:"device"`
	// Firmware contains firmware build and version information.
	Firmware StatusFirmware `json:"firmware"`
	// System contains API version and runtime state.
	System StatusSystem `json:"system"`
	// Power contains battery and USB power metrics.
	Power StatusPower `json:"power"`
}

func (s *Status) PrettyPrint() {
	fmt.Printf("\nStatus\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Firmware Version")
	row.Column(s.Firmware.Version)

	row = tab.Row()
	row.Column("System Uptime")
	row.Column(s.System.Uptime)

	row = tab.Row()
	row.Column("System Boot Time")
	row.Column(fmt.Sprintf("%s", time.Unix(int64(s.System.BootTime), 0).Format(time.RFC3339)))

	row = tab.Row()
	row.Column("Power State")
	row.Column(s.Power.State)

	row = tab.Row()
	row.Column("Battery Charge")
	row.Column(fmt.Sprintf("%d%%", s.Power.BatteryCharge))

	row = tab.Row()
	row.Column("Battery Voltage")
	row.Column(fmt.Sprintf("%d mV", s.Power.BatteryVoltage))

	row = tab.Row()
	row.Column("Battery Current")
	row.Column(fmt.Sprintf("%d mA", s.Power.BatteryCurrent))

	row = tab.Row()
	row.Column("USB Voltage")
	row.Column(fmt.Sprintf("%d mV", s.Power.UsbVoltage))

	tab.Print(os.Stdout)
	fmt.Println()
}

type StatusDevice struct {
	// SerialNumber is the device serial number.
	SerialNumber string `json:"serial_number"`
	// UsbMac is the MAC address of the USB ethernet interface.
	UsbMac string `json:"usb_mac"`
	// WifiMac is the Wi-Fi MAC address.
	WifiMac string `json:"wifi_mac"`
	// BleMac is the BLE MAC address.
	BleMac string `json:"ble_mac"`
	// OtpValid indicates whether OTP data is valid.
	OtpValid bool `json:"otp_valid"`
	// OtpModel is the device model code stored in OTP.
	OtpModel string `json:"otp_model"`
	// OtpTimestamp is the production timestamp.
	OtpTimestamp int `json:"otp_timestamp"`
	// FirmwareSecurity summarizes active firmware signature protections.
	FirmwareSecurity string `json:"firmware_security"`
}

type StatusFirmware struct {
	// Version is the firmware version.
	Version string `json:"version"`
	// Target is the firmware target code.
	Target int `json:"target"`
	// Branch is the firmware git branch name.
	Branch string `json:"branch"`
	// BuildDate is the firmware build date.
	BuildDate string `json:"build_date"`
	// CommitHash is the firmware git commit hash.
	CommitHash string `json:"commit_hash"`
	// NwpVersion is the radio firmware version.
	NwpVersion string `json:"nwp_version"`
	// MatterVersion is the Matter stack version.
	MatterVersion string `json:"matter_version"`
}

type StatusSystem struct {
	// APISemver is the HTTP API semantic version.
	APISemver string `json:"api_semver"`
	// Uptime is the formatted system uptime string.
	Uptime string `json:"uptime"`
	// BootTime is the Unix timestamp when the system booted.
	BootTime int `json:"boot_time"`
	// AutoUpdateEnabled reports whether automatic updates are enabled.
	AutoUpdateEnabled bool `json:"auto_update_enabled"`
}

type StatusPower struct {
	// State is the current power state.
	State string `json:"state"`
	// BatteryCharge is the battery charge percentage.
	BatteryCharge int `json:"battery_charge"`
	// BatteryVoltage is the battery voltage in millivolts.
	BatteryVoltage int `json:"battery_voltage"`
	// BatteryCurrent is the battery current in milliamps.
	BatteryCurrent int `json:"battery_current"`
	// UsbVoltage is the USB voltage in millivolts.
	UsbVoltage int `json:"usb_voltage"`
}

type DeviceStatus = StatusDevice

func (d *StatusDevice) PrettyPrint() {
	fmt.Printf("\nDevice Info\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Device Serial Number")
	row.Column(d.SerialNumber)

	row = tab.Row()
	row.Column("USB MAC Address")
	row.Column(d.UsbMac)

	row = tab.Row()
	row.Column("WiFi MAC Address")
	row.Column(d.WifiMac)

	row = tab.Row()
	row.Column("BLE MAC Address")
	row.Column(d.BleMac)

	row = tab.Row()
	row.Column("OTP Valid")
	row.Column(fmt.Sprintf("%v", d.OtpValid))

	tab.Print(os.Stdout)
	fmt.Println()
}

type VersionInfo struct {
	// APISemver is the API semantic version.
	APISemver string `json:"api_semver"`
}

type APIVersion = VersionInfo

func (a *VersionInfo) PrettyPrint() {
	fmt.Printf("\nAPI Version\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("API Semver")
	row.Column(a.APISemver)

	tab.Print(os.Stdout)
	fmt.Println()
}

type NetworkInterfaceInfo struct {
	// Type is the active network transport type.
	Type string `json:"type"`
}

type TransportType = NetworkInterfaceInfo

func (t *NetworkInterfaceInfo) PrettyPrint() {
	fmt.Printf("\nTransport Type\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Transport Type")
	row.Column(t.Type)

	tab.Print(os.Stdout)
	fmt.Println()
}
