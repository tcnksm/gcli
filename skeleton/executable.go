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

func (e *Executable) Validate() (errs []error) {

	if e.Name == "" {
		errs = append(errs, fmt.Errorf("`Name` cannot be blank"))
	}

	if e.Owner == "" {
		errs = append(errs, fmt.Errorf("`Owner` cannot be blank"))
	}

	if len(e.Commands) == 0 && len(e.Flags) == 0 {
		// can be blank ?
	}

	if len(e.Commands) > 0 {
		for _, c := range e.Commands {
			if c.Name == "" {
				errs = append(errs, fmt.Errorf("`Command.Name` cannot be blank"))
			}
		}
	}

	if len(e.Flags) > 0 {
		for _, f := range e.Flags {
			if f.Name == "" {
				errs = append(errs, fmt.Errorf("`Command.Name` cannot be blank"))
			}
		}
	}

	if e.Version == "" {
		// can be blank
	}

	if e.Description == "" {
		// can be blank
	}

	if e.FrameworkStr == "" {
		// can be blank
	}

	return errs
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
