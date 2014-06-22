cli-init
====

The easy way to start building Golang command-line application.


## Synopsis

`cli-init` is the easy way to start building Golang command-line application with [codegangsta/cli](https://github.com/codegangsta/cli). All you need to do is to set application name and its subcommand. You can forcus on core function of application.  

## Usage

You just need to set its application name. 

```bash
$ cli-init [options] [application]
```

## Example

If you want to start to building `todo` application which has subcommands `add`, `list`, `delete`:

```bash
$ cli-init -s add,list,delete todo
```

You can see sample, [tcnksm/sample-cli-init](https://github.com/tcnksm/sample-cli-init).

## Artifacts

`cli-init` generates belows:

- `<application>.go` - defines main function. It includes its application name, usage, author name and so on. 
- `commands.go` - defines sub-commands. It includes subcomand name, usage, functions. 
- `version.go` - defines application version. default value is `0.1.0`
- `README.md`
- `CHANGELOG.md`

See more details [codegangsta/cli](https://github.com/codegangsta/cli)

## Installation

To install `cli-init`, use `go get` and `make install`. We tag versions so feel free to checkout that tag and compile.

```bash
$ go get github.com/tcnksm/cli-init
$ make install 
```

## Author

[tcnksm](https://github.com/tcnksm)

