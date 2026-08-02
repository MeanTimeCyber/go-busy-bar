package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type AccountInfo struct {
	// Linked indicates whether the device is linked to an account.
	Linked bool `json:"linked"`
	// ID is the linked account identifier.
	ID string `json:"id"`
	// Email is the linked account email.
	Email string `json:"email"`
	// UserID is the linked user identifier.
	UserID string `json:"user_id"`
}

type AccountLink struct {
	// Code is the account-linking PIN code.
	Code string `json:"code"`
	// ExpiresAt is the Unix timestamp when the link code expires.
	ExpiresAt int `json:"expires_at"`
}

type AccountStatus struct {
	// Status is the current MQTT account-link status.
	Status string `json:"status"`
}

type AccountBackend struct {
	// ServerURL is the MQTT server URL to connect to.
	ServerURL string `json:"server_url"`
	// ClientCertType is the client certificate type to use.
	ClientCertType string `json:"client_cert_type"`
	// IgnoreServerCert indicates whether server certificate validation is skipped.
	IgnoreServerCert bool `json:"ignore_server_cert"`
}

// UnlinkAccount unlinks the device from the account.
func (c *Client) UnlinkAccount(ctx context.Context, query url.Values) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/account", query, nil)
}

// LinkAccount links the device to an account using the provided payload.
func (c *Client) LinkAccount(ctx context.Context, query url.Values, payload any) (*AccountLink, error) {
	return doJSON[AccountLink](c, ctx, http.MethodPost, "/api/account/link", query, payload)
}

// GetAccountInfo retrieves the account information from the Busy Bar API.
func (c *Client) GetAccountInfo(ctx context.Context, query url.Values) (*AccountInfo, error) {
	return doJSON[AccountInfo](c, ctx, http.MethodGet, "/api/account/info", query, nil)
}

// GetAccountStatus retrieves the account status information from the Busy Bar API.
func (c *Client) GetAccountStatus(ctx context.Context, query url.Values) (*AccountStatus, error) {
	return doJSON[AccountStatus](c, ctx, http.MethodGet, "/api/account/status", query, nil)
}

// GetAccountBackend retrieves the account backend information from the Busy Bar API.
func (c *Client) GetAccountBackend(ctx context.Context, query url.Values) (*AccountBackend, error) {
	return doJSON[AccountBackend](c, ctx, http.MethodGet, "/api/account/backend", query, nil)
}

// SetAccountBackend sets the account backend information in the Busy Bar API.
func (c *Client) SetAccountBackend(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPut, "/api/account/backend", query, payload)
}

func (a *AccountInfo) PrettyPrint() {
	fmt.Printf("\nAccount Info\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Linked")
	row.Column(fmt.Sprintf("%v", a.Linked))

	row = tab.Row()
	row.Column("ID")
	row.Column(a.ID)

	row = tab.Row()
	row.Column("Email")
	row.Column(a.Email)

	row = tab.Row()
	row.Column("User ID")
	row.Column(a.UserID)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (a *AccountLink) PrettyPrint() {
	fmt.Printf("\nAccount Link\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Code")
	row.Column(a.Code)

	row = tab.Row()
	row.Column("Expires At")
	row.Column(fmt.Sprintf("%d", a.ExpiresAt))

	tab.Print(os.Stdout)
	fmt.Println()
}

func (a *AccountStatus) PrettyPrint() {
	fmt.Printf("\nAccount Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Status")
	row.Column(a.Status)

	tab.Print(os.Stdout)
	fmt.Println()
}

func (a *AccountBackend) PrettyPrint() {
	fmt.Printf("\nAccount Backend\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Server URL")
	row.Column(a.ServerURL)

	row = tab.Row()
	row.Column("Client Cert Type")
	row.Column(a.ClientCertType)

	row = tab.Row()
	row.Column("Ignore Server Cert")
	row.Column(fmt.Sprintf("%v", a.IgnoreServerCert))

	tab.Print(os.Stdout)
	fmt.Println()
}
