package command

import (
	"flag"
	"reflect"
	"testing"

	"github.com/tcnksm/cli-init/skeleton"
)

func TestCommandFlag_implements(t *testing.T) {
	var raw interface{}
	raw = new(CommandFlag)
	if _, ok := raw.(flag.Value); !ok {
		t.Fatal("CommandFlag should be flag.Value")
	}
}

func TestCommandFlag_Set(t *testing.T) {
	tests := []struct {
		arg     string
		success bool
		expect  CommandFlag
	}{
		{
			arg:     `add:"Add new task"`,
			success: true,
			expect: []skeleton.Command{
				{Name: "add", Synopsis: "Add new task"},
			},
		},
		{
			arg:     `add:"Add new task",delete:"Delete task"`,
			success: true,
			expect: []skeleton.Command{
				{Name: "add", Synopsis: "Add new task"},
				{Name: "delete", Synopsis: "Delete task"},
			},
		},
		{
			arg:     `add,delete,list`,
			success: true,
			expect: []skeleton.Command{
				{Name: "add"},
				{Name: "delete"},
				{Name: "list"},
			},
		},
		{
			arg:     `include:"Include " character inside"`,
			success: true,
			expect: []skeleton.Command{
				{Name: "include", Synopsis: "Include \" character inside"},
			},
		},
	}

	for i, tt := range tests {
		c := new(CommandFlag)
		err := c.Set(tt.arg)
		if tt.success && err != nil {
			t.Fatalf("#%d Set(%q) expects not to happen error: %s", i, tt.arg, err)
		}

		if !reflect.DeepEqual(*c, tt.expect) {
			t.Errorf("#%d expects %v to be eq %v", i, *c, tt.expect)
		}
	}
}
