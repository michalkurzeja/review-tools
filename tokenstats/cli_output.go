package tokenstats

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type CliOutput struct {
}

func NewCliOutput() *CliOutput {
	return &CliOutput{}
}

func (c *CliOutput) Output(stats TokenStats) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"User", "Label", "Count"})
	table.AppendBulk(stats.AsTable())
	table.Render()
}
