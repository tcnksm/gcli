package skeleton

import "fmt"

const (
	// CLI with commands
	Framework_go_cmd = 100 + iota
	Framework_codegangsta_cli
	Framework_mitchellh_cli

	// CLI with flag
	Framework_flag = 1000 + iota
	Framework_tcnksm_mflag
)

// Framework returns framework ID (unique variable in gcli)
// from name string. If not match any framework, it retuns error.
func Framework(name string) (int, error) {
	switch name {
	case "go_cmd":
		return Framework_go_cmd, nil

	case "codegangsta_cli", "codegagsta":
		return Framework_codegangsta_cli, nil

	case "mitchellh_cli", "mitchellh":
		return Framework_mitchellh_cli, nil

	case "flag":
		return Framework_flag, nil

	default:
		return -1, fmt.Errorf("invalid framework name: %s", name)
	}
}
