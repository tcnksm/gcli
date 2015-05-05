package skeleton

import (
	"fmt"
	"strings"
)

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

// Validate validates user input.
func (f *Flag) Validate() error {
	if len(f.LongName) == 0 {
		return fmt.Errorf("LongName must be set")
	}

	if len(f.TypeString) == 0 {
		return fmt.Errorf("TypeString must be set")
	}

	return nil
}

// Fix fixed user input for templating.
func (f *Flag) Fix() error {

	// Fix Typestring
	if err := f.fixTypeString(); err != nil {
		return err
	}

	// Name must be title case of LongName
	f.Name = strings.Title(f.LongName)

	// ShortName is first character of LongName
	f.ShortName = strings.ToLower(string(f.LongName[0]))

	return nil
}

// FixTypeString fixes Type string which is provided
// by user and set Default variable.
func (f *Flag) fixTypeString() error {
	switch strings.ToLower(f.TypeString) {
	case "bool", "b":
		f.TypeString = "Bool"
		f.Default = false
	case "int", "i":
		f.TypeString = "Int"
		f.Default = 0
	case "string", "str", "s":
		f.TypeString = "String"
		f.Default = "\"\""
	default:
		return fmt.Errorf("unexpected type string: %s", f.TypeString)
	}
	return nil
}
