package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type StatusResponse struct {
	// State is the current Wi-Fi state.
	State string `json:"state"`
	// SSID is the connected network name and is present only when connected.
	SSID string `json:"ssid,omitempty"`
	// BSSID is the connected network BSSID and is present only when connected.
	BSSID string `json:"bssid,omitempty"`
	// Channel is the connected Wi-Fi channel and is present only when connected.
	Channel int `json:"channel,omitempty"`
	// RSSI is the signal level and is present only when connected.
	RSSI int `json:"rssi,omitempty"`
	// Security is the Wi-Fi security method and is present only when connected.
	Security string `json:"security,omitempty"`
	// IPConfig contains IP settings and is present only when connected.
	IPConfig *WifiStatusIPCfg `json:"ip_config,omitempty"`
}

type WifiStatusIPCfg struct {
	// IPMethod is the IP configuration method.
	IPMethod string `json:"ip_method,omitempty"`
	// IPType is the active IP protocol family.
	IPType string `json:"ip_type,omitempty"`
	// Address is the current IP address.
	Address string `json:"address,omitempty"`
}

type NetworkResponse struct {
	// Count is the number of scanned Wi-Fi networks.
	Count int `json:"count"`
	// Networks contains scanned network entries.
	Networks []Network `json:"networks"`
}

type Network struct {
	// SSID is the network name.
	SSID string `json:"ssid"`
	// Security is the network security mode.
	Security string `json:"security"`
	// RSSI is the network signal level.
	RSSI int `json:"rssi"`
}

// GetWifiStatus retrieves the current Wi-Fi status from the device.
func (c *Client) GetWifiStatus(ctx context.Context, query url.Values) (*StatusResponse, error) {
	return doJSON[StatusResponse](c, ctx, http.MethodGet, "/api/wifi/status", query, nil)
}

// PostWifiConnect sends a request to connect to a Wi-Fi network with the provided payload.
func (c *Client) PostWifiConnect(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/wifi/connect", query, payload)
}

// PostWifiDisconnect sends a request to disconnect from the current Wi-Fi network.
func (c *Client) PostWifiDisconnect(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/wifi/disconnect", query, payload)
}

// GetWifiNetworks retrieves the list of available Wi-Fi networks from the device.
func (c *Client) GetWifiNetworks(ctx context.Context, query url.Values) (*NetworkResponse, error) {
	return doJSON[NetworkResponse](c, ctx, http.MethodGet, "/api/wifi/networks", query, nil)
}

func (s *StatusResponse) PrettyPrint() {
	fmt.Printf("\nWi-Fi Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("State")
	row.Column(s.State)

	if s.SSID != "" {
		row = tab.Row()
		row.Column("SSID")
		row.Column(s.SSID)
	}

	if s.BSSID != "" {
		row = tab.Row()
		row.Column("BSSID")
		row.Column(s.BSSID)
	}

	if s.Channel != 0 {
		row = tab.Row()
		row.Column("Channel")
		row.Column(fmt.Sprintf("%d", s.Channel))
	}

	if s.RSSI != 0 {
		row = tab.Row()
		row.Column("RSSI")
		row.Column(fmt.Sprintf("%d", s.RSSI))
	}

	if s.Security != "" {
		row = tab.Row()
		row.Column("Security")
		row.Column(s.Security)
	}

	if s.IPConfig != nil && s.IPConfig.Address != "" {
		row = tab.Row()
		row.Column("IP")
		row.Column(s.IPConfig.Address)
	}

	tab.Print(os.Stdout)
	fmt.Println()
}

func (n *NetworkResponse) PrettyPrint() {
	fmt.Printf("\nWi-Fi Networks\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("SSID").SetAlign(tabulate.ML)
	tab.Header("Security").SetAlign(tabulate.ML)
	tab.Header("RSSI").SetAlign(tabulate.ML)

	for _, network := range n.Networks {
		row := tab.Row()
		row.Column(network.SSID)
		row.Column(network.Security)
		row.Column(fmt.Sprintf("%d", network.RSSI))
	}

	tab.Print(os.Stdout)
	fmt.Println()
}
