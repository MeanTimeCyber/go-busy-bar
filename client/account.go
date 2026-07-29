package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type AccountInfo struct {
	Linked bool   `json:"linked"`
	ID     string `json:"id"`
	Email  string `json:"email"`
	UserID string `json:"user_id"`
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
