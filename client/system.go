package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

// SetInputKey sets the input key using the Busy Bar API.
func (c *Client) SetInputKey(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/input", query, payload)
}

// HttpAccessInfo contains HTTP access control information.
func (c *Client) SetHttpAccess(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/access", query, payload)
}

// GetVersion retrieves the firmware version from the Busy Bar API.
func (c *Client) GetVersion(ctx context.Context, query url.Values) (*VersionInfo, error) {
	return doJSON[VersionInfo](c, ctx, http.MethodGet, "/api/version", query, nil)
}

// GetTransport retrieves the transport type information from the Busy Bar API.
func (c *Client) GetTransport(ctx context.Context, query url.Values) (*NetworkInterfaceInfo, error) {
	return doJSON[NetworkInterfaceInfo](c, ctx, http.MethodGet, "/api/transport", query, nil)
}

// GetStatus retrieves the overall status information from the Busy Bar API.
func (c *Client) DumpLog(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/log_dump", query, payload)
}

func (s *StatusFirmware) PrettyPrint() {
	fmt.Printf("\nFirmware Info\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Version")
	row.Column(s.Version)

	row = tab.Row()
	row.Column("Target")
	row.Column(fmt.Sprintf("%d", s.Target))

	row = tab.Row()
	row.Column("Branch")
	row.Column(s.Branch)

	row = tab.Row()
	row.Column("Build Date")
	row.Column(s.BuildDate)

	row = tab.Row()
	row.Column("Commit Hash")
	row.Column(s.CommitHash)

	row = tab.Row()
	row.Column("NWP Version")
	row.Column(s.NwpVersion)

	row = tab.Row()
	row.Column("Matter Version")
	row.Column(s.MatterVersion)

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
