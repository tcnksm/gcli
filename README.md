cli-init
====

The easy way to start building Golang command-line application.


## Synopsis

`cli-init` is the easy way to start building Golang command-line application with [codegangsta/cli](https://github.com/codegangsta/cli). All you need to do is to set application name and its subcommand. You can forcus on core function of application.

## Demo

![](http://deeeet.com/writing/images/post/cli-init.gif)

## Usage

You just need to set its application name:

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

- **main.go** - defines main function. It includes application name, version, usage, author name and so on. 
- **commands.go** - defines sub-commands. It includes subcommand name, usage, function and so on. 
- **version.go** - defines application version. default value is `0.1.0`
- **README.md** - insctructs application name, synopsis, usage and installation and so on. 
- **CHANGELOG.md** - shows version release date and its updates.

See more details [codegangsta/cli](https://github.com/codegangsta/cli).

## Installation

To install, use `go get` and `make install`. We tag versions so feel free to checkout that tag and compile.

```bash
$ go get -d github.com/tcnksm/cli-init
$ cd $GOPATH/src/github.com/tcnksm/cli-init
$ make install 
```

## Author

[tcnksm](https://github.com/tcnksm)

