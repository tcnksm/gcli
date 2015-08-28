package skeleton

const (
	// DefaultSynopsis is default synopsis message.
	DefaultSynopsis = ""

	// DefaultHelp is default help message.
	DefaultHelp = ""
)

// Command store command meta information.
type Command struct {
	// Name is command name.
	Name string

	// Flags are flag for the command.
	Flags []Flag

	// Synopsis is short help message of the command.
	Synopsis string

	// Help is long help message of the command.
	Help string

	// debugOutput is injected to command function
	// and generate for debugging purpose.
	// TODO: https://github.com/BurntSushi/toml/pull/90
	DebugOutput string // `toml:",omitempty"`
}
