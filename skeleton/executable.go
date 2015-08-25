package skeleton

import (
	"fmt"
	"reflect"
)

const (
	// DefaultVersion is default appliaction version
	DefaultVersion = "0.1.0"

	// defaultVersion is default application description
	DefaultDescription = ""
)

// Executable store executable meta information
type Executable struct {
	// Name is executable name
	Name string

	// Owner is owner of the executable
	Owner string

	// Commands are commands of the executable
	Commands []Command

	// Flags are flags of the executable
	Flags []Flag

	// Version is initial version
	Version string

	// Description is description of the executable
	Description string

	// FrameworkStr is framework name to use
	FrameworkStr string `toml:"Framework"`
}

func NewExecutable() *Executable {
	return &Executable{
		Version:     DefaultVersion,
		Description: DefaultDescription,
	}
}

// Overwrite overwrites provided value with default value
func (e *Executable) Overwrite(key string, v interface{}) error {
	// Check
	switch v.(type) {
	case string, []Command, []Flag:
	default:
		return fmt.Errorf("unexpected value: %#v", v)
	}

	rve := reflect.ValueOf(e)
	rve.Elem().FieldByName(key).Set(reflect.ValueOf(v))

	return nil
}
