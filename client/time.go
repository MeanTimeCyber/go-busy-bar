package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type TimestampInfo struct {
	// Timestamp is an ISO 8601 formatted timestamp with timezone.
	Timestamp string `json:"timestamp"`
}

type TimezoneInfo struct {
	// Name is the timezone name.
	Name string `json:"name"`
	// Offset is the timezone offset from UTC.
	Offset string `json:"offset"`
	// Abbr is the timezone abbreviation.
	Abbr string `json:"abbr"`
}

type TimezoneListResponse struct {
	// List contains supported timezone entries.
	List []TimezoneInfo `json:"list"`
}

// GetTime retrieves the current timestamp from the Busy Bar API.
func (c *Client) GetTime(ctx context.Context, query url.Values) (*TimestampInfo, error) {
	return doJSON[TimestampInfo](c, ctx, http.MethodGet, "/api/time", query, nil)
}

// SetTimeTimestamp sets the current timestamp on the Busy Bar API.
func (c *Client) SetTimeTimestamp(ctx context.Context, timestamp string, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/time/timestamp", requiredQuery(map[string]string{
		"timestamp": timestamp,
	}), payload)
}

// GetTimeTimezone retrieves the current timezone from the Busy Bar API.
func (c *Client) GetTimeTimezone(ctx context.Context, query url.Values) (*TimezoneInfo, error) {
	return doJSON[TimezoneInfo](c, ctx, http.MethodGet, "/api/time/timezone", query, nil)
}

// SetTimeTimezone sets the current timezone on the Busy Bar API.
func (c *Client) SetTimeTimezone(ctx context.Context, timezone string, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/time/timezone", requiredQuery(map[string]string{
		"timezone": timezone,
	}), payload)
}

// GetTimeTzlist retrieves the list of supported timezones from the Busy Bar API.
func (c *Client) GetTimeTzlist(ctx context.Context, query url.Values) (*TimezoneListResponse, error) {
	return doJSON[TimezoneListResponse](c, ctx, http.MethodGet, "/api/time/tzlist", query, nil)
}

func (t *TimestampInfo) PrettyPrint() {
	fmt.Printf("\nTimestamp\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Timestamp")
	row.Column(t.Timestamp)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (t *TimezoneInfo) PrettyPrint() {
	fmt.Printf("\nTimezone\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Name")
	row.Column(t.Name)

	row = tab.Row()
	row.Column("Offset")
	row.Column(t.Offset)

	row = tab.Row()
	row.Column("Abbreviation")
	row.Column(t.Abbr)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (t *TimezoneListResponse) PrettyPrint() {
	fmt.Printf("\nTimezone List\n")
	for _, tz := range t.List {
		tz.PrettyPrint()
	}
}
