package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type UpdateStatus struct {
	// Install contains firmware installation state and progress details.
	Install UpdateInstallStatus `json:"install"`
	// Check contains firmware update availability-check status.
	Check UpdateCheckStatus `json:"check"`
}

type UpdateInstallStatus struct {
	// IsAllowed indicates whether installation is allowed (battery checks).
	IsAllowed bool `json:"is_allowed"`
	// Event is the current update lifecycle event.
	Event string `json:"event"`
	// Action is the current update action.
	Action string `json:"action"`
	// Status is the current or last update operation status.
	Status string `json:"status"`
	// Detail is an optional human-readable status detail.
	Detail string `json:"detail"`
	// Download contains live download metrics for remote install.
	Download UpdateDownloadStatus `json:"download"`
}

type UpdateDownloadStatus struct {
	// SpeedBytesPerSec is current download speed in bytes per second.
	SpeedBytesPerSec int `json:"speed_bytes_per_sec"`
	// ReceivedBytes is bytes downloaded so far.
	ReceivedBytes int `json:"received_bytes"`
	// TotalBytes is expected download size in bytes.
	TotalBytes int `json:"total_bytes"`
}

type UpdateCheckStatus struct {
	// AvailableVersion is available firmware version (empty if none).
	AvailableVersion string `json:"available_version"`
	// Event is the current update-check event.
	Event string `json:"event"`
	// Status is the update-check result status.
	Status string `json:"status"`
}

type AutoupdateSettings struct {
	// IsEnabled controls whether automatic updates are enabled.
	IsEnabled bool `json:"is_enabled"`
	// IntervalStart is the update window start time in HH:MM format.
	IntervalStart string `json:"interval_start"`
	// IntervalEnd is the update window end time in HH:MM format.
	IntervalEnd string `json:"interval_end"`
}

type UpdateChangelog struct {
	// Changelog contains release notes for a requested firmware version.
	Changelog string `json:"changelog"`
}

// UpdatePayload is the payload for firmware update requests.
func (c *Client) UpdateFirmware(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/update", query, payload)
}

// CheckFirmwareUpdate checks for available firmware updates.
func (c *Client) CheckFirmwareUpdate(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/update/check", query, payload)
}

// GetFirmwareUpdateStatus retrieves the current firmware update status.
func (c *Client) GetFirmwareUpdateStatus(ctx context.Context, query url.Values) (*UpdateStatus, error) {
	return doJSON[UpdateStatus](c, ctx, http.MethodGet, "/api/update/status", query, nil)
}

// GetFirmwareUpdateChangelog retrieves the changelog for a specific firmware version.
func (c *Client) GetUpdateChangelog(ctx context.Context, query url.Values) (*UpdateChangelog, error) {
	return doJSON[UpdateChangelog](c, ctx, http.MethodGet, "/api/update/changelog", query, nil)
}

// GetAutoupdateSettings retrieves the current automatic update settings.
func (c *Client) InstallFirmwareUpdate(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/update/install", query, payload)
}

// AbortFirmwareDownload aborts an ongoing firmware download.
func (c *Client) AbortFirmwareDownload(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/update/abort_download", query, payload)
}

// GetAutoupdateSettings retrieves the current automatic update settings.
func (c *Client) GetAutoupdateSettings(ctx context.Context, query url.Values) (*AutoupdateSettings, error) {
	return doJSON[AutoupdateSettings](c, ctx, http.MethodGet, "/api/update/autoupdate", query, nil)
}

// SetAutoupdateSettings updates the automatic update settings.
func (c *Client) SetAutoupdateSettings(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/update/autoupdate", query, payload)
}

func (u *UpdateStatus) PrettyPrint() {
	fmt.Printf("\nFirmware Update Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Install Event")
	row.Column(u.Install.Event)

	row = tab.Row()
	row.Column("Install Action")
	row.Column(u.Install.Action)

	row = tab.Row()
	row.Column("Install Status")
	row.Column(u.Install.Status)

	row = tab.Row()
	row.Column("Available Version")
	row.Column(u.Check.AvailableVersion)

	row = tab.Row()
	row.Column("Check Status")
	row.Column(u.Check.Status)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (a *AutoupdateSettings) PrettyPrint() {
	fmt.Printf("\nAutoupdate Settings\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Enabled")
	row.Column(fmt.Sprintf("%v", a.IsEnabled))

	row = tab.Row()
	row.Column("Interval Start")
	row.Column(a.IntervalStart)

	row = tab.Row()
	row.Column("Interval End")
	row.Column(a.IntervalEnd)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (u *UpdateChangelog) PrettyPrint() {
	fmt.Printf("\nUpdate Changelog\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Changelog")
	row.Column(u.Changelog)

	tab.Print(os.Stdout)
	fmt.Println()
}
