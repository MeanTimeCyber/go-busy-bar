package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type DeviceName struct {
	Name string `json:"name"`
}

func (d *DeviceName) PrettyPrint() {
	fmt.Printf("\nDevice Name\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Name")
	row.Column(d.Name)

	tab.Print(os.Stdout)
	fmt.Println()
}
