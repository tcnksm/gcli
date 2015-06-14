package skeleton

import "testing"

func TestFramework(t *testing.T) {
	tests := []struct {
		in      string
		success bool
		expt    int
	}{
		{"codegangsta_cli", true, Framework_codegangsta_cli},
		{"not_exist_cli", false, Framework_codegangsta_cli},
	}

	for i, tt := range tests {
		out, err := Framework(tt.in)
		if tt.success && err != nil {
			t.Errorf("#%d expects error not to be occurred: %s", i, err)
		}

		if !tt.success {
			continue
		}

		if out != tt.expt {
			t.Errorf("#%d expects %d to eq %d", i, out, tt.expt)
		}
	}
}
