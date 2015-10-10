package command

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/tcnksm/gcli/skeleton"
)

// ListCommand is a Command that lists all avairable frameworks
type ListCommand struct {
	Meta
}

// Run lists all avairable frameworks.
func (c *ListCommand) Run(args []string) int {

	uflag := c.Meta.NewFlagSet("list", c.Help())
	if err := uflag.Parse(args); err != nil {
		return 1
	}

	outBuffer := new(bytes.Buffer)
	// Create a table for output
	table := tablewriter.NewWriter(outBuffer)
	header := []string{"Name", "Command", "URL"}
	table.SetHeader(header)
	for _, f := range skeleton.Frameworks {
		if f.Hide {
			continue
		}
		var cmd string
		if len(f.CommandTemplates) > 0 {
			cmd = "*"
		}
		table.Append([]string{f.Name, cmd, f.URL})
	}

	// Write a table
	table.Render()

	fmt.Fprintf(outBuffer, "COMMAND(*) means you can create command pattern CLI with that framework (you can use --command flag)")
	c.UI.Output(outBuffer.String())
	return 0
}

// Synopsis is a one-line, short synopsis of the command.
func (c *ListCommand) Synopsis() string {
	return "List available cli frameworks"
}

// Help is a long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (c *ListCommand) Help() string {
	helpText := `
Show all avairable cli frameworks.

Usage:

  gcli list
`
	return strings.TrimSpace(helpText)
}
