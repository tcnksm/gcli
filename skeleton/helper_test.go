package skeleton

import (
	"testing"
)

func TestCamelCase(t *testing.T) {

	tests := []struct {
		in, expt string
	}{
		{
			in:   "ignore-case",
			expt: "ignoreCase",
		},
	}

	for i, tt := range tests {
		out := camelCase(tt.in)
		for out != tt.expt {
			t.Fatalf("#%d expects %q to be eq %q", i, out, tt.expt)
		}
	}
}
