package command

import (
	"flag"
	"reflect"
	"testing"

	"github.com/tcnksm/gcli/skeleton"
)

func TestFlagFlag_implements(t *testing.T) {
	var raw interface{}
	raw = new(FlagFlag)
	if _, ok := raw.(flag.Value); !ok {
		t.Fatal("FlagFlag should be flag.Value")
	}
}

func TestFlagFlag_Set(t *testing.T) {
	tests := []struct {
		arg     string
		success bool
		expect  FlagFlag
	}{
		{
			arg:     `debug:string:"Run as debug mode"`,
			success: true,
			expect: []skeleton.Flag{
				{
					Name:        "debug",
					LongName:    "debug",
					ShortName:   "d",
					TypeString:  skeleton.TypeStringString,
					Default:     "",
					Description: "Run as debug mode",
				},
			},
		},
		{
			arg:     `debug,help,test`,
			success: true,
			expect: []skeleton.Flag{
				{Name: "debug", LongName: "debug", ShortName: "d", TypeString: skeleton.TypeStringString, Default: ""},
				{Name: "help", LongName: "help", ShortName: "h", TypeString: skeleton.TypeStringString, Default: ""},
				{Name: "test", LongName: "test", ShortName: "t", TypeString: skeleton.TypeStringString, Default: ""},
			},
		},
	}

	for i, tt := range tests {
		c := new(FlagFlag)
		err := c.Set(tt.arg)
		if tt.success && err != nil {
			t.Fatalf("#%d Set(%q) expects not to happen error: %s", i, tt.arg, err)
		}

		if !reflect.DeepEqual(*c, tt.expect) {
			t.Errorf("#%d expects %v to be eq %v", i, *c, tt.expect)
		}
	}
}
