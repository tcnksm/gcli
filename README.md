gcli
====

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[license]: https://github.com/tcnksm/gcli/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/tcnksm/gcli

`gcli` (formerly `cli-init`) generates the codes and its directory structure you need to start building CLI tool right out of the box.
All you need is to provide name, commands and [framework](#support-frameworks) you want to use. 

## Usage

To start new command line tool,

```bash
$ gcli new [options] NAME
```
It generates new cli skeleton project. At least, you must provide executable name.

See more usage,

```bash
$ gcli help
```

## Example

If you want to create `todo` CLI application which has `add`, `list` and `delete` command with
[mitchellh/cli](https://github.com/mitchellh/cli) framework,

```bash
$ cd $GOPATH/src/github.com/YOUR_NAME
$ gcli new -F mitchellh_cli -c add -c list -c delete todo
```

You can run `go build` todo application from beginning.

## Support frameworks

`gcli` supports bellow cli frameworks,

- [codegangsta_cli](https://github.com/codegangsta/cli)
- [mitchellh_cli](https://github.com/mitchellh/cli)
- [go_cmd]() (Standard `go` command style)
- [flag](https://golang.org/pkg/flag/)

`gcli` has tempaltes of these frameworks. Template file includes best practices of each frameworks like
how to separate file or how to set directory structure and so on.

In future, we will also suppport other CLI frameworks like below (Need help),

- [spf13/cobra](https://github.com/spf13/cobra)
- [docopt.go](https://github.com/docopt/docopt.go)
- [motemen/go-cli](https://github.com/motemen/go-cli)
- [mow.cli](https://github.com/jawher/mow.cli)
- [ogier/pflag](https://github.com/ogier/pflag)
- [go-flags](https://github.com/jessevdk/go-flags)

## Installation

To install, use `go get` and `make install`. We tag versions so feel free to checkout that tag and compile.

```bash
$ go get -d github.com/tcnksm/gcli
$ cd $GOPATH/src/github.com/tcnksm/gcli
$ make install 
```

`gcli` was re-written from scratch. If you prefer old version of `gcli`, checkout,

```bash
$ git checkout v0.1.0
```

## Contribution

1. Fork ([https://github.com/tcnksm/gcli/fork](https://github.com/tcnksm/gcli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `make test` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Taichi Nakashima](https://github.com/tcnksm)
