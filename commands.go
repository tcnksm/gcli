package main

import (
	"github.com/mitchellh/cli"
	"github.com/tcnksm/gcli/command"
)

// Commands are collections of gcli commands.
func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Meta: *meta,
			}, nil
		},

		"design": func() (cli.Command, error) {
			return &command.DesignCommand{
				Meta: *meta,
			}, nil
		},

		"apply": func() (cli.Command, error) {
			return &command.ApplyCommand{
				Meta: *meta,
			}, nil
		},

		"list": func() (cli.Command, error) {
			return &command.ListCommand{
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
