package command

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tcnksm/go-latest"
)

// CheckTimeout is timeout of executing go-latest.Check()
const CheckTimeout = 3 * time.Second

// VersionCommand is a Command that shows version
type VersionCommand struct {
	Meta

	Version  string
	Revision string
}

// Run shows version string and commit hash if it exists.
// It returns exit code
func (c *VersionCommand) Run(args []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "gcli version %s", c.Version)
	if c.Revision != "" {
		fmt.Fprintf(&versionString, " (%s)", c.Revision)
	}

	c.UI.Output(versionString.String())

	resCh := CheckLatest(c.Version)
	select {
	case res := <-resCh:
		if res != nil && res.Outdated {
			msg := fmt.Sprintf(
				"\nYour versin of gcli is out of date! The latest version is %s.",
				res.Current)
			c.UI.Error(msg)
		}
	case <-time.After(CheckTimeout):
		// Time out & do nothing
	}

	return 0
}

// Synopsis is a one-line, short synopsis of the command.
func (c *VersionCommand) Synopsis() string {
	return "Print the gcli version"
}

// Help is a long-form help text. In this case, help text is not  neccessary.
func (c *VersionCommand) Help() string {
	return ""
}

// CheckLatest run tcnksm/go-latest with gcli settings.
// It retuns channel of checking results. Even if something wrong happened,
// it neglects error because this is not important part of gcli execution.
func CheckLatest(version string) <-chan *latest.CheckResponse {
	// Check version is latest or not
	fix := latest.DeleteFrontV()
	github := &latest.GithubTag{
		Owner:             "tcnksm",
		Repository:        "gcli",
		FixVersionStrFunc: fix,
	}

	resCh := make(chan *latest.CheckResponse)
	go func() {
		// Ignore error because it not critical for main fucntion
		res, _ := latest.Check(github, fix(version))
		resCh <- res
	}()

	return resCh
}
