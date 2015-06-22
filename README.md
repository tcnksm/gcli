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

`gcli` (formerly `cli-init`) generates the codes and its directory structure you need to start building CLI tool by Golang right out of the box. All you need is to provide name, command names and [framework](#support-frameworks) you want to use. 

## Usage

To start new command line tool, run below. It generates new cli skeleton project. At least, you must provide executable name.

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

## Example

If you want to create `todo` CLI application which has `add`, `list` and `delete` command with
[mitchellh/cli](https://github.com/mitchellh/cli) framework,

```bash
$ cd $GOPATH/src/github.com/YOUR_NAME
$ gcli new -F mitchellh_cli -c add -c list -c delete todo
```

It generates below files,

```bash
$ tree todo/
todo
├── CHANGELOG.md
├── README.md
├── cli.go
├── command
│   ├── add.go
│   ├── add_test.go
│   ├── delete.go
│   ├── delete_test.go
│   ├── list.go
│   ├── list_test.go
│   └── meta.go
├── commands.go
├── main.go
└── version.go
```

You can run `go build` todo application from beginning.

## Support frameworks

`gcli` generates two types of CLI (you can choose). Flag pattern & Command pattern.

### Flag pattern

Flag pattern is the pattern which executable has only flag options (e.g., `grep`)

```bash
$ grep —i -C 4 "some string" /tmp   
    │     │              │           
    │     │               `--------- Arguments 
    │     │                          
    │      `------------------------ Option flags   
    │                                
     `------------------------------ Executable  
```

To generate this pattern, `gcli` supports,

- [flag](https://golang.org/pkg/flag/)

### Command pattern

Command pattern is the pattern which executable has command for change its behavior (e.g., `git`)

```bash
$ git --no-pager push -v origin mastter     
   │       │      │    │      │           
   │       │      │    │       `------- Arguments 
   │       │      │    │              
   │       │      │     `-------------- Command flags 
   │       │      │                   
   │       │       `------------------- Command
   │       │                          
   │        `-------------------------- Global flags
   │                                  
    `---------------------------------- Executable
```

To generate this pattern, `gcli` supports,

- [codegangsta_cli](https://github.com/codegangsta/cli)
- [mitchellh_cli](https://github.com/mitchellh/cli)
- [go_cmd](https://github.com/golang/go/blob/master/src/cmd/go/main.go#L30#L51) (No 3rd party framework, `go` command style)

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
