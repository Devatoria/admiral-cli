package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// NewTable initializes a new table with some custom options such as row separator
func NewTable(headers []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetRowLine(true)

	return table
}
