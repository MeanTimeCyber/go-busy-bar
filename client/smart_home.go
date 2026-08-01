package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type SmartHomePairingInfo struct {
	// FabricCount is the number of smart homes (Matter fabrics) paired to the device.
	FabricCount int `json:"fabric_count"`
	// LatestPairingStatus is the latest smart home pairing state.
	LatestPairingStatus SmartHomePairingStatus `json:"latest_pairing_status"`
}

type SmartHomePairingStatus struct {
	// Value is the latest state of smart home pairing (commissioning).
	Value string `json:"value"`
	// Timestamp is the UTC Unix second timestamp of the latest pairing status update.
	Timestamp int `json:"timestamp,omitempty"`
}

type SmartHomePairingPayload struct {
	// AvailableUntil is the UTC Unix millisecond timestamp until pairing is available.
	AvailableUntil string `json:"available_until"`
	// QRCode is the QR payload used for smart home pairing.
	QRCode string `json:"qr_code"`
	// ManualCode is the manual smart home pairing code.
	ManualCode string `json:"manual_code"`
}

type SmartHomeSwitchState struct {
	// State is the emulated switch state.
	State bool `json:"state"`
	// Startup is the switch startup behavior value accepted by the API.
	Startup string `json:"startup,omitempty"`
}

func (s *SmartHomePairingInfo) PrettyPrint() {
	fmt.Printf("\nSmart Home Pairing\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Fabric Count")
	row.Column(fmt.Sprintf("%d", s.FabricCount))

	row = tab.Row()
	row.Column("Latest Status")
	row.Column(s.LatestPairingStatus.Value)

	if s.LatestPairingStatus.Timestamp != 0 {
		row = tab.Row()
		row.Column("Latest Timestamp")
		row.Column(fmt.Sprintf("%d", s.LatestPairingStatus.Timestamp))
	}

	tab.Print(os.Stdout)
	fmt.Println()
}

func (s *SmartHomePairingPayload) PrettyPrint() {
	fmt.Printf("\nSmart Home Pairing Payload\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Available Until")
	row.Column(s.AvailableUntil)

	row = tab.Row()
	row.Column("QR Code")
	row.Column(s.QRCode)

	row = tab.Row()
	row.Column("Manual Code")
	row.Column(s.ManualCode)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (s *SmartHomeSwitchState) PrettyPrint() {
	fmt.Printf("\nSmart Home Switch\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("State")
	row.Column(fmt.Sprintf("%v", s.State))

	if s.Startup != "" {
		row = tab.Row()
		row.Column("Startup")
		row.Column(s.Startup)
	}

	tab.Print(os.Stdout)
	fmt.Println()
}
