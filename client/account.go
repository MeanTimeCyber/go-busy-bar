package client

import (
	"fmt"
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
