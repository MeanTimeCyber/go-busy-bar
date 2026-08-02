package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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

// GetSmartHomeCommissioningStatus retrieves the current smart home commissioning (pairing) status.
func (c *Client) GetSmartHomeCommissioningStatus(ctx context.Context, query url.Values) (*SmartHomePairingInfo, error) {
	return doJSON[SmartHomePairingInfo](c, ctx, http.MethodGet, "/api/smart_home/pairing", query, nil)
}

// GetSmartHomePairingPayload retrieves the smart home pairing payload (QR code and manual code).
func (c *Client) StartSmartHomePairing(ctx context.Context, query url.Values, payload any) (*SmartHomePairingPayload, error) {
	return doJSON[SmartHomePairingPayload](c, ctx, http.MethodPost, "/api/smart_home/pairing", query, payload)
}

// DeleteSmartHomePairing deletes the smart home pairing information from the device.
func (c *Client) DeleteSmartHomePairing(ctx context.Context, query url.Values) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/smart_home/pairing", query, nil)
}

// GetSmartHomeSwitch retrieves the current smart home switch state.
func (c *Client) GetSmartHomeSwitch(ctx context.Context, query url.Values) (*SmartHomeSwitchState, error) {
	return doJSON[SmartHomeSwitchState](c, ctx, http.MethodGet, "/api/smart_home/switch", query, nil)
}

// PostSmartHomeSwitch sets the smart home switch state.
func (c *Client) PostSmartHomeSwitch(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/smart_home/switch", query, payload)
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
