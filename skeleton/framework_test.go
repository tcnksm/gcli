package skeleton

import "testing"

func TestFramework(t *testing.T) {
	tests := []struct {
		in      string
		success bool
		expt    string
	}{
		{"urfave", true, "urfave_cli"},
		{"not_exist_cli", false, ""},
	}

	for i, tt := range tests {
		out, err := FrameworkByName(tt.in)
		if tt.success && err != nil {
			t.Errorf("#%d expects error not to be occurred: %s", i, err)
		}

		if !tt.success {
			continue
		}

		if out.Name != tt.expt {
			t.Errorf("#%d expects %s to eq %s", i, out.Name, tt.expt)
		}
	}
}
