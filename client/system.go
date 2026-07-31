package client

import (
	"fmt"
	"os"
	"time"

	"github.com/markkurossi/tabulate"
)

// Status represents the status information returned by the device.
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

// DeviceStatus represents the device status information returned by the device.
type DeviceStatus struct {
	SerialNumber     string `json:"serial_number"`
	UsbMac           string `json:"usb_mac"`
	WifiMac          string `json:"wifi_mac"`
	BleMac           string `json:"ble_mac"`
	OtpValid         bool   `json:"otp_valid"`
	OtpModel         string `json:"otp_model"`
	OtpTimestamp     int    `json:"otp_timestamp"`
	FirmwareSecurity string `json:"firmware_security"`
}

func (d *DeviceStatus) PrettyPrint() {
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

// APIVersion represents the API version information returned by the device.
type APIVersion struct {
	APISemver string `json:"api_semver"`
}

func (a *APIVersion) PrettyPrint() {
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

// TransportType represents the transport type information returned by the device.
type TransportType struct {
	Type string `json:"type"`
}

func (t *TransportType) PrettyPrint() {
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
