package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type StorageList struct {
	// List contains storage file and directory entries.
	List []StorageListElement `json:"list"`
}

type StorageListElement struct {
	// Type is the element type (file or dir).
	Type string `json:"type"`
	// Name is the file or directory name.
	Name string `json:"name"`
	// Size is the file size in bytes and is present for file entries.
	Size int `json:"size,omitempty"`
}

type StorageStatus struct {
	// UsedBytes is used storage capacity in bytes.
	UsedBytes int `json:"used_bytes"`
	// FreeBytes is remaining storage capacity in bytes.
	FreeBytes int `json:"free_bytes"`
	// TotalBytes is total storage partition size in bytes.
	TotalBytes int `json:"total_bytes"`
}

func (s *StorageList) PrettyPrint() {
	fmt.Printf("\nStorage Files\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Type").SetAlign(tabulate.ML)
	tab.Header("Name").SetAlign(tabulate.ML)
	tab.Header("Size").SetAlign(tabulate.ML)

	for _, item := range s.List {
		row := tab.Row()
		row.Column(item.Type)
		row.Column(item.Name)
		if item.Size != 0 {
			row.Column(fmt.Sprintf("%d", item.Size))
		} else {
			row.Column("-")
		}
	}

	tab.Print(os.Stdout)
	fmt.Println()
}

func (s *StorageStatus) PrettyPrint() {
	fmt.Printf("\nStorage Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Used Bytes")
	row.Column(fmt.Sprintf("%d", s.UsedBytes))

	row = tab.Row()
	row.Column("Free Bytes")
	row.Column(fmt.Sprintf("%d", s.FreeBytes))

	row = tab.Row()
	row.Column("Total Bytes")
	row.Column(fmt.Sprintf("%d", s.TotalBytes))

	tab.Print(os.Stdout)
	fmt.Println()
}
