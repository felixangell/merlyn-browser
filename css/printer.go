package css

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func TokenPrinter(tokens []*Token) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Lexeme", "Token Type"})

	for _, t := range tokens {
		table.Append([]string{
			string(t.lexeme), t.kind.String(),
		})
	}
	table.Render()
}
