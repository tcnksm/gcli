package skeleton

import (
	"reflect"
	"testing"
)

func TestOverwrite(t *testing.T) {
	tests := []struct {
		initExecutable, expt *Executable
		success              bool
		inputKey             string
		inputValue           interface{}
	}{
		{
			initExecutable: &Executable{Name: "todo"},
			inputKey:       "Name",
			inputValue:     "todo-ng",
			success:        true,
			expt:           &Executable{Name: "todo-ng"},
		},
		{
			initExecutable: &Executable{Name: "todo"},
			inputKey:       "Name",
			inputValue:     1,
			success:        false,
			expt:           &Executable{},
		},

		{
			initExecutable: &Executable{
				Name: "todo",
				Commands: []Command{
					{Name: "add"},
				},
			},
			inputKey: "Commands",
			inputValue: []Command{
				{Name: "list"},
			},
			success: true,
			expt: &Executable{
				Name: "todo",
				Commands: []Command{
					{Name: "list"},
				},
			},
		},
		{
			initExecutable: &Executable{
				Name: "todo",
				Flags: []Flag{
					{Name: "add"},
				},
			},
			inputKey: "Flags",
			inputValue: []Flag{
				{Name: "list"},
			},
			success: true,
			expt: &Executable{
				Name: "todo",
				Flags: []Flag{
					{Name: "list"},
				},
			},
		},
	}

	for i, tt := range tests {
		out := tt.initExecutable
		err := out.Overwrite(tt.inputKey, tt.inputValue)
		if !tt.success && err == nil {
			t.Fatalf("#%d expects to be error", i)
		}

		if !tt.success {
			continue
		}

		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(out, tt.expt) {
			t.Errorf("#%d expects %#v to be eq %#v", i, out, tt.expt)
		}
	}

}
