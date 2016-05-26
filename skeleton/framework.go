package skeleton

import "fmt"

// Framework represents framework
type Framework struct {
	// Name is framework name
	Name string

	// AltName is alternative name which represent Framework
	AltNames []string

	// Description is description of framework
	Description string

	// URL is framework project URL
	URL string

	// BaseTemplates
	BaseTemplates []Template

	// CommandTemplate
	CommandTemplates []Template

	// If Hide is true, `list` command doesn't show
	// this framework
	Hide bool
}

// CommonTemplates is collection of templates which are used all frameworks.
var CommonTemplates = []Template{
	{"resource/tmpl/common/CHANGELOG.md.tmpl", "CHANGELOG.md"},
	{"resource/tmpl/common/README.md.tmpl", "README.md"},
	{"resource/tmpl/common/gitignore.tmpl", ".gitignore"},
}

// Frameworks is collection of Framework.
var Frameworks = []*Framework{
	{
		Name:     "mitchellh_cli",
		AltNames: []string{"mitchellh"},
		URL:      "https://github.com/mitchellh/cli",
		Description: `mitchellh/cli cli is a library for implementing powerful command-line interfaces in Go.
cli is the library that powers the CLI for Packer, Serf, and Consul.
`,
		BaseTemplates: []Template{
			{"resource/tmpl/mitchellh_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/mitchellh_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/mitchellh_cli/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/mitchellh_cli/commands.go.tmpl", "commands.go"},
			{"resource/tmpl/mitchellh_cli/command/meta.go.tmpl", "command/meta.go"},
			{"resource/tmpl/mitchellh_cli/command/version.go.tmpl", "command/version.go"},
		},
		CommandTemplates: []Template{
			{"resource/tmpl/mitchellh_cli/command/command.go.tmpl", "command/{{ .Name }}.go"},
			{"resource/tmpl/mitchellh_cli/command/command_test.go.tmpl", "command/{{ .Name }}_test.go"},
		},
	},

	{
		Name:     "codegangsta_cli",
		AltNames: []string{"codegangsta"},
		URL:      "https://github.com/codegangsta/cli",
		Description: `codegangsta/cli is simple, fast, and fun package for building command line apps in Go.
The goal is to enable developers to write fast and distributable command line applications in an expressive way.
`,
		BaseTemplates: []Template{
			{"resource/tmpl/codegangsta_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/codegangsta_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/codegangsta_cli/commands.go.tmpl", "commands.go"},
		},
		CommandTemplates: []Template{
			{"resource/tmpl/codegangsta_cli/command/command.go.tmpl", "command/{{ .Name }}.go"},
			{"resource/tmpl/codegangsta_cli/command/command_test.go.tmpl", "command/{{ .Name }}_test.go"},
		},
	},

	{
		Name:     "urfave_cli",
		AltNames: []string{"urfave"},
		URL:      "https://github.com/urfave/cli",
		Description: `This is the library formally known as codegangsta/cli. urfave/cli is simple, fast, and fun package for building command line apps in Go.
The goal is to enable developers to write fast and distributable command line applications in an expressive way.
`,
		BaseTemplates: []Template{
			{"resource/tmpl/urfave_cli/main.go.tmpl", "main.go"},
			{"resource/tmpl/urfave_cli/version.go.tmpl", "version.go"},
			{"resource/tmpl/urfave_cli/commands.go.tmpl", "commands.go"},
		},
		CommandTemplates: []Template{
			{"resource/tmpl/urfave_cli/command/command.go.tmpl", "command/{{ .Name }}.go"},
			{"resource/tmpl/urfave_cli/command/command_test.go.tmpl", "command/{{ .Name }}_test.go"},
		},
	},

	{
		Name: "go_cmd",
		URL:  "https://github.com/golang/go/tree/master/src/cmd/go",
		Description: `
`,
		BaseTemplates: []Template{
			{"resource/tmpl/go_cmd/main.go.tmpl", "main.go"},
		},
		CommandTemplates: []Template{
			{"resource/tmpl/go_cmd/command.go.tmpl", "{{ .Name }}.go"},
			{"resource/tmpl/go_cmd/command_test.go.tmpl", "{{ .Name }}_test.go"},
		},
	},

	{
		Name: "bash",
		URL:  "",
		Description: `
`,
		BaseTemplates: []Template{
			{"resource/tmpl/bash/main.sh.tmpl", "{{ .Name }}.sh"},
		},
		CommandTemplates: []Template{},
		Hide:             true,
	},

	{
		Name:        "flag",
		AltNames:    []string{},
		URL:         "https://golang.org/pkg/flag/",
		Description: `Package flag implements command-line flag parsing.`,
		BaseTemplates: []Template{
			{"resource/tmpl/flag/main.go.tmpl", "main.go"},
			{"resource/tmpl/flag/version.go.tmpl", "version.go"},
			{"resource/tmpl/flag/cli.go.tmpl", "cli.go"},
			{"resource/tmpl/flag/cli_test.go.tmpl", "cli_test.go"},
		},
		CommandTemplates: []Template{},
	},
}

// FrameworkByName retuns Framework
func FrameworkByName(name string) (*Framework, error) {
	for _, f := range Frameworks {
		if f.Name == name {
			return f, nil
		}

		for _, alt := range f.AltNames {
			if alt == name {
				return f, nil
			}
		}
	}
	return nil, fmt.Errorf("invalid framework name: %s", name)
}
