gcli
====

[![GitHub release](http://img.shields.io/github/release/tcnksm/gcli.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/5587a34baf7de9c51b02e04b.svg?style=flat-square)][wercker]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[release]: https://github.com/tcnksm/gcli/releases
[wercker]: https://app.wercker.com/#applications/5587a34baf7de9c51b02e04b
[license]: https://github.com/tcnksm/gcli/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/tcnksm/gcli

`gcli` generates a skeleton (codes and its directory structure) you need to start building Command Line Interface (CLI) tool by Golang right out of the box. You can use your favorite [CLI framework](#frameworks).

## Why ?

Why you need `gcli`? Because you should focus on writing core function of CLI, not on interface. During developing CLI tool by Golang, you may find you're writing the chunk of [boilerplate code](https://en.wikipedia.org/wiki/Boilerplate_code) for interfaces. Stop writing the same codes every time. `gcli` generates them and save you a large amount of time by writing such code. This is like [Rails scaffold](http://guides.rubyonrails.org/command_line.html#rails-generate). Not only that, `gcli` know the best practices of golang CLI framework library which you want to use. Generated codes follows the most ideal way of using that framework, and you don't need to know about that. See the [frameworks](#frameworks) it supports now. 

## Demo

The following demo shows creating `todo` CLI application which has `add`, `list` and `delete` command with [mitchellh/cli](https://github.com/mitchellh/cli) (Which is used for [Hashicorp](https://hashicorp.com/) products) with **one command**. As you can see, the generated codes are `go build`-able from beginning. 

![gif](/doc/gif/gcli-new.gif)

And this [video](https://vimeo.com/142134929) shows creating same `todo` CLI application with `design` & `apply` commands. This is the other way to start building new CLI application. First, it starts with creating design file by `design` command. In this file, you can define, CLI name, description of the CLI , framework you want to use, and commands & flags with its usages. After editing, it executes `apply` command to generating a project from that design file. 

## Usage

`gcli` is single command-line application. This application then takes subcommands. To check the all available commands,

```bash
$ gcli help
```

To get help for any specific subcommand, run it with the `-h` flag.

`gcli` has 2 main subcommand to generate the project. The one is the `new` command, the other is the `design` & `apply` commands. The former is for generating the project by command line one-liner, the latter is for when you want to design it in your editor before generating (It generates design file and you can generate project based on it). The following section explain, how to use these commands.

### `new` command

The `new` command tells gcli to generate CLI project with command-line one-liner,

```bash
$ gcli new [options] NAME
```

You must provides project name (`NAME`), the name will be the directory name it includes all codes and be the default binary name. By default, `gcli` creates a project under `$GOPATH/github.com/<username>` (If you don't provide username via option, it uses `github.user` or `user.name` in `.gitconfig` file). In option, you can set subcommand or flag it has and its description. You can also set your favorite [CLI framework](#frameworks) there. The followings are all available opntions,

```bash
-command=name, -c           Command name which you want to add.
                            This is valid only when cli pacakge support commands.
                            This can be specified multiple times. Synopsis can be
                            set after ":". Namely, you can specify command by
                            -command=NAME:SYNOPSYS. Only NAME is required.
                            You can set multiple variables at same time with ","
                            separator.

-flag=name, -f              Global flag option name which you want to add.
                            This can be specified multiple times. By default, flag type
                            is string and its description is empty. You can set them,
                            with ":" separator. Namaly, you can specify flag by
                            -flag=NAME:TYPE:DESCIRPTION. Order must be flow  this and
                            TYPE must be string, bool or int. Only NAME is required.
                            You can set multiple variables at same time with ","
                            separator.

-framework=name, -F         Cli framework name. By default, gcli use "codegangsta/cli"
                            To check cli framework you can use, run 'gcli list'.
                            If you set invalid framework, it will be failed.

-owner=name, -o             Command owner (author) name. This value is also used for
                            import path name. By default, owner name is extracted from
                            ~/.gitconfig variable.

-skip-test, -T              Skip generating *_test.go file. By default, gcli generates
                            test file If you specify this flag, gcli will not generate
                            test files.
```

For example, to `todo` CLI application which has `add`, `list` and `delete` command with [mitchellh/cli](https://github.com/mitchellh/cli),

```bash
$ gcli new -F mitchellh_cli -c add -c list -c delete todo
```

### `design` & `apply` command

The `design` command tells gcli to prepare design template file ([`.toml`](https://github.com/toml-lang/toml)). The design file defines all necessary information to generate CLI application. Some fields are filled with the ideal default value, and some have empty value. You can fill that empty filed with your favorite editor with thinking like what interface that should have or description of that and so on. You can see sample template file [`sample.toml`](/sample.toml). 

After design, use `apply` command and tells gcli to generate CLI project based on the design file. The following describes this workflow. 

First, generate design template file by `design` command, 

```bash
$ gcli design [options] NAME
```
You must provides project name (`NAME`). In option, you can set subcommand or flag it has and its description. You can also set your favorite [CLI framework](#frameworks) there. You can edit these values in design file later. 

Then, edit design file by your favorite `$EDITOR`.

```bash
$ $EDITOR <NAME>-design.toml
```

After that validate design by `validate` command to check required fields are filled, 

```bash
$ gcli validate <NAME>-design.toml
```

Finnaly, generate CLI project with that design file by `apply` command, 

```bash
$ gcli apply <NAME>-desigon.toml
```

The video for this workflow is available on [Vimeo](https://vimeo.com/142134929). 

## Frameworks

There are many framework (package) for buidling command line application by golang. For example, one of the most famous frameworks is [codegangsta/cli](https://github.com/codegangsta/cli). The framework helps you not writing many codes. But still you need to write many boilerplate code for that framework. And there are different way to use that framework and learning the ideal way to use is waste of time. gcli writes out with following the best practice for that framework (learn from famous tool that is built with that framework). 

`gcli` can generate 2 types of CLI pattern. The one is [*sub-command pattern*](#sub-command), the other is [*flag pattern*](#flag). The former is flexible and you can add many behavior in one command application. The later is for simple application. You can check the all available frameworks by `list` command,

```bash
$ gcli list
```

To change framework, you can use `-framework` or `-F` option with the framework name. This option can be used for `new`, `design` and `apply` command. By default, [codegangsta_cli](https://github.com/codegangsta/cli) will be used. 

The following section will explain [*sub-command pattern*](#sub-command) and [*flag pattern*](#flag). 

### Sub-Command

*Sub-Command pattern* is the pattern that executable takes sub-command for change its behavior. `git` command is one example for this pattern. It takes `push`, `pull` subcommands. `gcli` is also this pattern. `gcli` supports the following frameworks for the command pattern.

|Name|Sample projects|
|:-:|:-:| 
|[codegangsta_cli](https://github.com/codegangsta/cli) | [docker machine](https://github.com/docker/machine) | 
|[mitchellh_cli](https://github.com/mitchellh/cli)| [consul](https://github.com/hashicorp/consul), [terraform](https://github.com/hashicorp/terraform)| 
|[go_cmd](https://github.com/golang/go/blob/master/src/cmd/go/main.go#L30#L51)| [go](https://golang.org/cmd/go/)| 

([go_cmd](https://github.com/golang/go/blob/master/src/cmd/go/main.go#L30#L51) is not framework. It only uses standard package. It generates same struct and functions that `go` command uses.)

### Flag

*Flag pattern* is the pattern that executable has flag options for changing its behavior. For example, `grep` command is this pattern. Now `gcli` only supports the official [flag](https://golang.org/pkg/flag/) package for this pattern.

For example, to create command it has `-ignore-case` option and `context` option (your own `grep`),

```bash
$ gcli new -F flag -flag=i:Bool -flag=C:Int grep
```

## Installation

To install, use `go get` and `make install`. We tag versions so feel free to checkout that tag and compile.

```bash
$ go get -d github.com/tcnksm/gcli
$ cd $GOPATH/src/github.com/tcnksm/gcli
$ make install 
```

## Contribution

1. Fork ([https://github.com/tcnksm/gcli/fork](https://github.com/tcnksm/gcli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `make test` command and confirm that it passes
1. Create a new Pull Request

## Author

[Taichi Nakashima](https://github.com/tcnksm)
