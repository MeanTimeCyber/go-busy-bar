package client

import (
	"fmt"
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
