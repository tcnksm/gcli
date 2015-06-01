package skeleton

// This file defines template file path and its output name
// which is related to cli framework.

var (
	CommonTemplates = []Template{
		{"resource/tmpl/common/CHANGELOG.md.tmpl", "CHANGELOG.md"},
		{"resource/tmpl/common/README.md.tmpl", "README.md"},
	}
)

// FrameworkTempaltes returns framewrok Template based on framework ID.
func FrameworkTemplates(framework int) []Template {
	switch framework {
	case Framework_go_cmd:
		return []Template{
			{"resource/tmpl/go_cmd/main.go.tmpl", "main.go"},
		}
	case Framework_codegangsta_cli:
		return []Template{
			{"resource/tmpl/codegangsta_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/codegangsta_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/codegangsta_cli/commands.go.tmpl", "commands.go"},
		}
	case Framework_mitchellh_cli:
		return []Template{
			{"resource/tmpl/mitchellh_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/mitchellh_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/mitchellh_cli/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/mitchellh_cli/commands.go.tmpl", "commands.go"},
			{"resource/tmpl/mitchellh_cli/command/meta.go.tmpl", "command/meta.go"},
		}
	case Framework_flag:
		return []Template{
			{"resource/tmpl/flag/main.go.tmpl", "main.go"},
			{"resource/tmpl/flag/version.go.tmpl", "version.go"},
			{"resource/tmpl/flag/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/flag/cli_test.go.tmpl", "cli_test.go"},
		}
	case Framework_tcnksm_mflag:
		return []Template{
			{"resource/tmpl/tcnksm_mflag/main.go.tmpl", "main.go"},
			{"resource/tmpl/tcnksm_mflag/version.go.tmpl", "version.go"},
			{"resource/tmpl/tcnksm_mflag/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/tcnksm_mflag/cli_test.go.tmpl", "cli_test.go"},
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
		return Template{"resource/tmpl/go_cmd/command.go.tmpl", "{{ .Name }}.go"}, Template{"", ""}
	case Framework_codegangsta_cli:
		return Template{"resource/tmpl/codegangsta_cli/command.go.tmpl", "command/{{ .Name }}.go"},
			Template{"resource/tmpl/codegangsta_cli/command_test.go.tmpl", "command/{{ .Name }}_test.go"}
	case Framework_mitchellh_cli:
		return Template{"resource/tmpl/mitchellh_cli/command/command.go.tmpl", "command/{{ .Name }}.go"},
			Template{"resource/tmpl/mitchellh_cli/command/command_test.go.tmpl", "command/{{ .Name }}_test.go"}
	default:
		return Template{}, Template{}
	}
}
