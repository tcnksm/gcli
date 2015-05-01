package command

import "strings"

// NewCommand is a Command that generates a new cli project
type NewCommand struct {
	Meta
}

func (c *NewCommand) Run(args []string) int {
	c.Ui.Output("====> Generate new project")
	return 0
}

func (c *NewCommand) Synopsis() string {
	return "Generate new cli project"
}

func (c *NewCommand) Help() string {
	helpText := `
Usage: cli-init new NAME SUB_COMMAND...

  Generate new cli project skeleton.
`
	return strings.TrimSpace(helpText)
}
