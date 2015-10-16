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

The following demo shows creating `todo` CLI application which has `add`, `list` and `delete` command with [mitchellh/cli](https://github.com/mitchellh/cli) (Which is used for [Hashicorp](https://hashicorp.com/) products) with **one command**. 

![gif](/doc/gif/gcli-new.gif)

The next demo shows creating same `todo` CLI application with `design` & `apply` commands. This is the other way to start building new CLI application. First, it starts with creating design file by `design` command. In this file, you can define, CLI name, description of the CLI , framework you want to use, and commands & flags with its usages. After editing, it executes `apply` command to generating a project from that design file (The complete video is available on [Vimeo](https://vimeo.com/142134929)). 

![gif](/doc/gif/gcli-design.gif)

As you can see, the generated codes are `go build`-able from beginning. 

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
You must provides project name (`NAME`), the name will be the directory name it includes all codes and be the default binary name. In option, you can set subcommand or flag it has and its description. You can alos set your favorite [CLI framework](#frameworks) there. The followings are all available opntions,

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
You must provides project name (`NAME`). In option, you can set subcommand or flag it has and its description. You can alos set your favorite [CLI framework](#frameworks) there. You can edit these values in design file later. 

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

`gcli` can generate two types of CLI pattern, 

- [Flag pattern](#flag-pattern)
- [Command pattern](#command-pattern)

### Flag pattern

Flag pattern is the pattern which executable has only flag options like below (e.g., `grep`),

```bash
$ grep —i -C 4 "some string" /tmp   
    │     │              │           
    │     │               `--------- Arguments 
    │     │                          
    │      `------------------------ Option flags   
    │                                
     `------------------------------ Executable  
```

To generate above CLI application with [flag](https://golang.org/pkg/flag/) fraemwork,
 
```bash
$ cd $GOPATH/src/github.com/YOUR_NAME
$ gcli new -F flag -flag=i:Bool -flag=C:Int grep
  Created grep/main.go
  Created grep/CHANGELOG.md
  Created grep/cli_test.go
  Created grep/README.md
  Created grep/version.go
  Created grep/cli.go
====> Successfully generated grep
```

`gcli` supports the following packages for the flag pattern:

- [flag](https://golang.org/pkg/flag/)

### Command pattern

Command pattern is the pattern which executable has command for change its behavior. For example, todo CLI application which has add (Add new task), list (List all tasks) and delete(Delete a task) command.

```bash
$ todo add 'Buy a milk' 
   │    │      │           
   │    │       `---------- Arguments 
   │    │ 
   │     `----------------- Command 
   │                                  
    `---------------------- Executable
```

To generate above CLI application with [mitchellh/cli](https://github.com/mitchellh/cli) framework,

```bash
$ cd $GOPATH/src/github.com/YOUR_NAME
$ gcli new -F mitchellh_cli -c add -c list -c delete todo
  Created todo/main.go
  Created todo/command/meta.go
  Created todo/cli.go
  Created todo/CHANGELOG.md
  Created todo/version.go
  Created todo/commands.go
  Created todo/command/add.go
  Created todo/command/list.go
  Created todo/command/delete.go
  Created todo/README.md
  Created todo/command/add_test.go
  Created todo/command/list_test.go
  Created todo/command/delete_test.go
====> Successfully generated todo
```

`gcli` supports the following packages for the command pattern:

- [codegangsta_cli](https://github.com/codegangsta/cli)
- [mitchellh_cli](https://github.com/mitchellh/cli)
- [go_cmd](https://github.com/golang/go/blob/master/src/cmd/go/main.go#L30#L51) (without 3rd party framework, same as `go` command)

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
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Taichi Nakashima](https://github.com/tcnksm)
