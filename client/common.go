package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

// SuccessResponse represents a generic successful API response.
type SuccessResponse struct {
	// Result is a success status message.
	Result string `json:"result"`
}

func (s *SuccessResponse) PrettyPrint() {
	fmt.Printf("\nSuccess\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Result")
	row.Column(s.Result)

	tab.Print(os.Stdout)
	fmt.Println()
}

// ErrorResponse represents a structured API error payload.
type ErrorResponse struct {
	// Error is the error message.
	Error string `json:"error"`
	// Code is the numeric error code.
	Code int `json:"code"`
}

func (e *ErrorResponse) PrettyPrint() {
	fmt.Printf("\nError\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Error")
	row.Column(e.Error)

	row = tab.Row()
	row.Column("Code")
	row.Column(fmt.Sprintf("%d", e.Code))

	tab.Print(os.Stdout)
	fmt.Println()
}
