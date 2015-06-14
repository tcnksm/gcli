package command

import "strings"

// ListCommand is a Command that generates a new cli project
type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	// TODO
	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List available cli frameworks"
}

func (c *ListCommand) Help() string {
	helpText := `
Usage: gcli list

  Show all avairable cli frameworks. 
`
	return strings.TrimSpace(helpText)
}
