package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type DeviceName struct {
	// Name is the device name.
	Name string `json:"name"`
}

type NameInfo struct {
	// Name is the device name (letters, digits, spaces, and common punctuation).
	Name string `json:"name"`
}

type HttpAccessInfo struct {
	// Mode is the HTTP API access mode.
	Mode string `json:"mode"`
	// KeyValid reports whether an access key is configured and valid.
	KeyValid bool `json:"key_valid"`
}

type DisplayBrightnessInfo struct {
	// Value is the display brightness value (0-100 or auto).
	Value string `json:"value"`
}

type AudioVolumeInfo struct {
	// Volume is the current audio volume value (0-100).
	Volume float64 `json:"volume"`
}

// GetHttpAccess retrieves the HTTP access information from the Busy Bar API.
func (c *Client) GetHttpAccess(ctx context.Context, query url.Values) (*HttpAccessInfo, error) {
	return doJSON[HttpAccessInfo](c, ctx, http.MethodGet, "/api/access", query, nil)
}

// SetHttpAccess contains HTTP access control information.
func (c *Client) SetHttpAccess(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/access", query, payload)
}

// GetName retrieves the device name from the Busy Bar API.
func (c *Client) GetName(ctx context.Context, query url.Values) (*NameInfo, error) {
	return doJSON[NameInfo](c, ctx, http.MethodGet, "/api/name", query, nil)
}

// PostName sets the device name using the Busy Bar API.
func (c *Client) PostName(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/name", query, payload)
}

// GetDisplayBrightness retrieves the current display brightness from the Busy Bar API.
func (c *Client) GetDisplayBrightness(ctx context.Context, query url.Values) (*DisplayBrightnessInfo, error) {
	return doJSON[DisplayBrightnessInfo](c, ctx, http.MethodGet, "/api/display/brightness", query, nil)
}

// SetDisplayBrightness sets the display brightness using the Busy Bar API.
func (c *Client) SetDisplayBrightness(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/display/brightness", query, payload)
}

// GetAudioVolume retrieves the current audio volume from the Busy Bar API.
func (c *Client) GetAudioVolume(ctx context.Context, query url.Values) (*AudioVolumeInfo, error) {
	return doJSON[AudioVolumeInfo](c, ctx, http.MethodGet, "/api/audio/volume", query, nil)
}

// SetAudioVolume sets the audio volume using the Busy Bar API.
func (c *Client) SetAudioVolume(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/audio/volume", query, payload)
}

// ConnectWebSocket connects to the Busy Bar WebSocket API.
func (c *Client) ConnectWebSocket(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/status/ws", query, nil)
}

// GetScreen retrieves the current screen state from the Busy Bar API.
func (c *Client) GetScreen(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/screen", query, nil)
}

func (d *DeviceName) PrettyPrint() {
	(&NameInfo{Name: d.Name}).PrettyPrint()
}

func (n *NameInfo) PrettyPrint() {
	fmt.Printf("\nDevice Name\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Name")
	row.Column(n.Name)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (h *HttpAccessInfo) PrettyPrint() {
	fmt.Printf("\nHTTP Access\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Mode")
	row.Column(h.Mode)

	row = tab.Row()
	row.Column("Key Valid")
	row.Column(fmt.Sprintf("%v", h.KeyValid))

	tab.Print(os.Stdout)
	fmt.Println()
}

func (d *DisplayBrightnessInfo) PrettyPrint() {
	fmt.Printf("\nDisplay Brightness\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Value")
	row.Column(d.Value)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (a *AudioVolumeInfo) PrettyPrint() {
	fmt.Printf("\nAudio Volume\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Volume")
	row.Column(fmt.Sprintf("%.2f", a.Volume))

	tab.Print(os.Stdout)
	fmt.Println()
}
