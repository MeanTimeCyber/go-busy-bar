package client

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
	println("Device Serial Number:", s.Device.SerialNumber)
	println("Firmware Version:", s.Firmware.Version)
	println("System Uptime:", s.System.Uptime)
	println("Power State:", s.Power.State)
}
