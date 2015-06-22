package command

import "strings"

// ListCommand is a Command that lists all avairable frameworks
type ListCommand struct {
	Meta
}

// Run lists all avairable frameworks.
func (c *ListCommand) Run(args []string) int {
	// TODO
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
Usage: gcli list

  Show all avairable cli frameworks. 
`
	return strings.TrimSpace(helpText)
}
