package client

import (
	"fmt"
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
