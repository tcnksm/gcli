package skeleton

import "testing"

func TestProcessPathTmpl(t *testing.T) {
	tests := []struct {
		Tmpl    string
		data    interface{}
		success bool
		expect  string
	}{
		{
			Tmpl:    "{{ .Name }}/README.md",
			data:    Executable{Name: "todo"},
			success: true,
			expect:  "todo/README.md",
		},

		{
			Tmpl:    "command/{{ .Name }}.go",
			data:    Command{Name: "add"},
			success: true,
			expect:  "command/add.go",
		},

		{
			Tmpl:    "{{ .NotExist }}.go",
			data:    struct{}{},
			success: false,
		},
	}

	for i, tt := range tests {
		output, err := processPathTmpl(tt.Tmpl, tt.data)
		if tt.success && err != nil {
			t.Fatalf("#%d expects error not to be occurred", i)
		}

		if output != tt.expect {
			t.Errorf("#%d expects %s to eq %s", i, output, tt.expect)
		}
	}
}
