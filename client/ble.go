package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type BleStatusResponse struct {
	// Status is the current BLE state.
	Status string `json:"status"`
	// Address is the remote device address when BLE status is connected.
	Address string `json:"address,omitempty"`
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
