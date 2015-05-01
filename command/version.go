package command

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tcnksm/go-latest"
)

const CheckTimeout = 3 * time.Second

// VersionCommand is a Command that generates a new cli project
type VersionCommand struct {
	Meta

	Version  string
	Revision string
}

func (c *VersionCommand) Run(args []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "cli-init version %s", c.Version)
	if c.Revision != "" {
		fmt.Fprintf(&versionString, " (%s)", c.Revision)
	}

	c.Ui.Output(versionString.String())

	resCh := CheckLatest(c.Version)
	select {
	case res := <-resCh:
		if res != nil && res.Outdated {
			msg := fmt.Sprintf(
				"\nYour versin of cli-init is out of date! The latest version is %s.",
				res.Current)
			c.Ui.Error(msg)
		}
	case <-time.After(CheckTimeout):
		// Time out & do nothing
	}

	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Print the cli-init version"
}

func (c *VersionCommand) Help() string {
	return ""
}

func CheckLatest(version string) <-chan *latest.CheckResponse {
	// Check version is latest or not
	github := &latest.GithubTag{
		Owner:             "tcnksm",
		Repository:        "cli-init",
		FixVersionStrFunc: latest.DeleteFrontV(),
	}

	resCh := make(chan *latest.CheckResponse)
	go func() {
		// Ignore error because it not critical for main fucntion
		res, _ := latest.Check(github, version)
		resCh <- res
	}()

	return resCh
}
