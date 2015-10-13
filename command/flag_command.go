package command

import (
	"fmt"
	"strings"

	"github.com/tcnksm/gcli/skeleton"
)

// CommandFlag implements the flag.Value interface and allows multiple
// calls to the same variable to append a list. It parses string and set them
// as skeleton.Command.
type CommandFlag []*skeleton.Command

// String
func (c *CommandFlag) String() string {
	return ""
}

// Set parses input string and appends it on CommandFlags.
// Input format must be NAME[:SYNOPSIS] format.x
func (c *CommandFlag) Set(v string) error {
	cmdStrs := strings.Split(v, ",")
	for _, cmdStrs := range cmdStrs {

		parsedCmdStr := strings.Split(cmdStrs, ":")
		if len(parsedCmdStr) > 2 {
			return fmt.Errorf("command flag must be specified by NAME:SYNOPSIS format")
		}

		name := parsedCmdStr[0]

		// synopsis is optional
		synopsis := ""
		if len(parsedCmdStr) == 2 {
			synopsis = parsedCmdStr[1]

			// Delete unnessary characters
			// TODO, this should not here..? or extract this as other function
			synopsis = strings.Trim(synopsis, "\"")
			synopsis = strings.Trim(synopsis, "'")
		}

		command := &skeleton.Command{
			Name:     name,
			Synopsis: synopsis,
		}

		if err := command.Fix(); err != nil {
			return err
		}

		*c = append(*c, command)
	}

	return nil
}
