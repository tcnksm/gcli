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

`gcli` generates the codes and its directory structure you need to start building CLI tool by Golang right out of the box. You can use your favorite [framework](#support-frameworks). 

## Usage

To start new command line tool, run the following command. It generates new cli skeleton project. At least, you must provide executable name. You can run `go build` application from beginning.

```bash
$ gcli new [options] NAME
```

To see available frameworks,

```bash
$ gcli list
```

See more usage,

```bash
$ gcli help
```

### Generate CLI project from `.toml` file

You can generate CLI project from design template file (`.toml`). You can define command name, its description, commands there. 
First, you can create `toml` file via `design` command,

```bash
$ gcli design <NAME>
```

Then, edit design file by your favorite `$EDITOR`. You can see sample template file [`sample.toml`](/sample.toml),

```bash
$ $EDITOR <NAME>-design.toml
```

You can validate design by `validate` command,

```bash
$ gcli validate <NAME>-design.toml
```

To generate CLI project, use `apply` command, 

```bash
$ gcli apply <NAME>-desigon.toml
```


## Support frameworks

`gcli` generates two types of CLI. flag pattern or command pattern.

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

For this flag pattern, `gcli` supports below packages, 

- [flag](https://golang.org/pkg/flag/)

### Command pattern

Command pattern is the pattern which executable has command for change its behavior. For example, `todo` CLI application which has `add` (Add new task), `list` (List all tasks) and `delete`(Delete a task) command. 

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

For this flag pattern, `gcli` supports below packages, 

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
