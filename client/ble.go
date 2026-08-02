package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type BleStatusResponse struct {
	// Status is the current BLE state.
	Status string `json:"status"`
	// Address is the remote device address when BLE status is connected.
	Address string `json:"address,omitempty"`
}

// PostBleEnable sends a request to enable BLE on the device.
func (c *Client) PostBleEnable(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/ble/enable", query, payload)
}

// PostBleDisable sends a request to disable BLE on the device.
func (c *Client) PostBleDisable(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/ble/disable", query, payload)
}

// DeleteBlePairing sends a request to delete BLE pairing on the device.
func (c *Client) DeleteBlePairing(ctx context.Context, query url.Values) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/ble/pairing", query, nil)
}

// GetBleStatus sends a request to retrieve the current BLE status on the device.
func (c *Client) GetBleStatus(ctx context.Context, query url.Values) (*BleStatusResponse, error) {
	return doJSON[BleStatusResponse](c, ctx, http.MethodGet, "/api/ble/status", query, nil)
}

func (b *BleStatusResponse) PrettyPrint() {
	fmt.Printf("\nBLE Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Status")
	row.Column(b.Status)

	if b.Address != "" {
		row = tab.Row()
		row.Column("Address")
		row.Column(b.Address)
	}

	tab.Print(os.Stdout)
	fmt.Println()
}
