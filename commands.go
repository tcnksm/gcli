package main

import (
	"github.com/mitchellh/cli"
	"github.com/tcnksm/cli-init/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
			}, nil
		},
	}
}
