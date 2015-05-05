package skeleton

type Executable struct {
	// Name is executable name
	Name string

	// Owner is owner of the executable
	Owner string

	// Commands are commands of the executable
	Commands []Command

	// Flags are flags of the exexutable
	Flags []Flag

	// Version is initial version
	Version string

	// Description is description of the executable
	Description string

	// Framework is cli package
	Framework int
}
