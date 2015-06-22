package skeleton

// This file defines template file path and its output name
// which is related to cli framework.

var (
	CommonTemplates = []Template{
		{"resource/tmpl/common/CHANGELOG.md.tmpl", "CHANGELOG.md"},
		{"resource/tmpl/common/README.md.tmpl", "README.md"},
	}
)

// FrameworkTempaltes returns framework Template based on framework ID.
func FrameworkTemplates(framework int) []Template {
	switch framework {
	case Framework_go_cmd:
		return []Template{
			{"resource/tmpl/command/go_cmd/main.go.tmpl", "main.go"},
		}
	case Framework_codegangsta_cli:
		return []Template{
			{"resource/tmpl/command/codegangsta_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/command/codegangsta_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/command/codegangsta_cli/commands.go.tmpl", "commands.go"},
		}
	case Framework_mitchellh_cli:
		return []Template{
			{"resource/tmpl/command/mitchellh_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/command/mitchellh_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/command/mitchellh_cli/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/command/mitchellh_cli/commands.go.tmpl", "commands.go"},
			{"resource/tmpl/command/mitchellh_cli/command/meta.go.tmpl", "command/meta.go"},
		}
	case Framework_flag:
		return []Template{
			{"resource/tmpl/flag/flag/main.go.tmpl", "main.go"},
			{"resource/tmpl/flag/flag/version.go.tmpl", "version.go"},
			{"resource/tmpl/flag/flag/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/flag/flag/cli_test.go.tmpl", "cli_test.go"},
		}
	default:
		return []Template{}
	}
}

// CommandTempaltes returns command Tempalte based on framework ID.
// The first return value is command Tempalte.
// The second return value is command test Tempalte.
func CommandTemplates(framework int) (Template, Template) {
	switch framework {
	case Framework_go_cmd:
		return Template{"resource/tmpl/command/go_cmd/command.go.tmpl", "{{ .Name }}.go"},
			Template{"resource/tmpl/command/go_cmd/command_test.go.tmpl", "{{ .Name }}_test.go"}
	case Framework_codegangsta_cli:
		return Template{"resource/tmpl/command/codegangsta_cli/command/command.go.tmpl", "command/{{ .Name }}.go"},
			Template{"resource/tmpl/command/codegangsta_cli/command/command_test.go.tmpl", "command/{{ .Name }}_test.go"}
	case Framework_mitchellh_cli:
		return Template{"resource/tmpl/command/mitchellh_cli/command/command.go.tmpl", "command/{{ .Name }}.go"},
			Template{"resource/tmpl/command/mitchellh_cli/command/command_test.go.tmpl", "command/{{ .Name }}_test.go"}
	default:
		return Template{}, Template{}
	}
}
