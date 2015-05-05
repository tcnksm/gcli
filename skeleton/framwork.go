package skeleton

const (
	// CLI with commands
	Framework_codegangsta_cli = 100 + iota
	Framework_mitchellh_cli

	// CLI with flag
	Framework_flag = 200 + iota
	Framework_tcnksm_flag
)

func CommonTemplates() []string {
	return []string{
		"resource/tmpl/common/CHANGELOG.md.tmpl",
		"resource/tmpl/common/README.md.tmpl",
		"resource/tmpl/common/version.go.tmpl",
	}
}

func FrameworkTemplates(framework int) []string {
	switch framework {
	case Framework_codegangsta_cli:
		return []string{}
	case Framework_mitchellh_cli:
		return []string{}
	case Framework_flag:
		return []string{
			"resource/tmpl/flag/main.go.tmpl",
			"resource/tmpl/flag/cli.go.tmpl",
		}
	case Framework_tcnksm_flag:
		return []string{
			"resource/tmpl/tcnksm_mflag/main.go.tmpl",
			"resource/tmpl/tcnksm_mflag/cli.go.tmpl",
		}
	default:
		panic("invalid framework is provided.")
	}
}
