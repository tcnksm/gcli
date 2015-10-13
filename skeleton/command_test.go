package skeleton

import "testing"
import "reflect"

func TestCommand_Fix(t *testing.T) {
	tests := []struct {
		in, exp *Command
		success bool
	}{
		{
			in: &Command{
				Name:  "server-start",
				Flags: []Flag{},
			},

			exp: &Command{
				Name:         "server-start",
				FunctionName: "serverStart",
				Flags:        []Flag{},
			},

			success: true,
		},
	}

	for i, tt := range tests {

		err := tt.in.Fix()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect Fix to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect Fix not to fail but %q", i, err.Error())
		}

		if !reflect.DeepEqual(*tt.in, *tt.exp) {
			t.Errorf("#%d expect %v to eq %v", i, tt.in, tt.exp)
		}
	}

}
