package command

import (
	"fmt"
	"strings"

	"github.com/tcnksm/gcli/skeleton"
)

const (
	// defaultTypeString is default flag type
	defaultTypeString = "string"
)

// FlagFlag implements the flag.Value interface and allows multiple
// calls to the same variable to append a list. It parses string and set them
// as skeleton.Flag.
type FlagFlag []*skeleton.Flag

// String
func (f *FlagFlag) String() string {
	return ""
}

// Set parses input string and appends it on FlagdFlag.
// Input format must be NAME:TYPE:SYNOPSIS format.
func (f *FlagFlag) Set(v string) error {
	flgStrs := strings.Split(v, ",")

	for _, flgStr := range flgStrs {

		parsedFlgStr := strings.Split(flgStr, ":")
		if len(parsedFlgStr) > 3 {
			return fmt.Errorf("flag must be specified by NAME:TYPE:DESCRIPTION format")
		}

		name := parsedFlgStr[0]
		typeString := defaultTypeString
		desc := ""

		if len(parsedFlgStr) > 1 {
			typeString = parsedFlgStr[1]
		}

		if len(parsedFlgStr) > 2 {
			desc = parsedFlgStr[2]

			// Delete unnessary characters
			// TODO, this should not here..? or extract this as other function
			desc = strings.Trim(desc, "\"")
			desc = strings.Trim(desc, "'")
		}

		flag := &skeleton.Flag{
			LongName:    name,
			TypeString:  typeString,
			Description: desc,
		}

		// Fix inputs string for using main processing
		if err := flag.Fix(); err != nil {
			return err
		}

		*f = append(*f, flag)
	}

	return nil
}
