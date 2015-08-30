package skeleton

import (
	"fmt"
	"strings"
)

// TypeString represents type as string
const (
	TypeStringInt    = "int"
	TypeStringBool   = "bool"
	TypeStringString = "string"
)

// Flag stores flag meta informations
type Flag struct {
	// Name is flag name, this is used for flag variable name in generated code.
	// Name is equal to titled LongName.
	Name string

	// LongName is long form of the flag name.
	// This must be provided by user
	LongName string

	// ShortName is short form of flag name.
	// This is generated automatically from LongName
	ShortName string

	// TypeString is flag type. This must be provided by user
	TypeString string

	// Default is default value.
	// This is automatically generated from TypeString
	Default interface{}

	// Description is help message of the flag.
	Description string
}

// Fix fixed user input for templating.
func (f *Flag) Fix() error {

	// Fix Typestring
	if err := f.fixTypeString(); err != nil {
		return err
	}

	f.LongName = strings.ToLower(f.LongName)

	// Name is same as LongName by default
	f.Name = f.LongName

	// ShortName is first character of LongName
	// TODO, when same first character is provided.
	f.ShortName = string(f.LongName[0])

	return nil
}

// FixTypeString fixes Type string which is provided
// by user and set Default variable.
func (f *Flag) fixTypeString() error {
	switch strings.ToLower(f.TypeString) {
	case "bool", "b":
		f.TypeString = TypeStringBool
		if f.Default == nil {
			f.Default = false
		}
	case "int", "i":
		f.TypeString = TypeStringInt
		if f.Default == nil {
			f.Default = 0
		}
	case "string", "str", "s":
		f.TypeString = TypeStringString
		if f.Default == nil {
			f.Default = ""
		}
	default:
		return fmt.Errorf("unexpected type is provided: %s", f.TypeString)
	}
	return nil
}
