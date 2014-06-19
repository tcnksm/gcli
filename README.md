cli-init
====

Easy to start building command-line tool with [codegangsta/cli](https://github.com/codegangsta/cli).

## Synopsis

`cli-init` generates ...

## Usage

You just specify application name and its subcommands:

```bash
$ cli-init todo -s add,delete,list
```

This will generate three files.

```bash
.
|-- todo/
    |-- todo.go
    |-- commands.go
    |-- version.go
```

## Artifacts

### version.go

`version.go` defines its version.

```go
package main

const Version string = "0.1.0"
```

### todo.go

`todo.go` defines main function.

```go
package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var mainFlags = []cli.Flag{
	cli.BoolFlag{"debug", "Run as DEBUG mode"},
}

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Version = Version
	app.Usage = ""
	app.Author = ""
	app.Email = ""
	app.Flags = mainFlags
	app.Commands = Commands

	app.Before = func(c *cli.Context) error {

		if c.GlobalBool("debug") {
			os.Setenv("DEBUG", "1")
		}

		return nil
	}

	app.Run(os.Args)
}
```

### commands.go

`commands.go` defines sub-commands.

```go
package main

import (
	"github.com/codegangsta/cli"
	"log"
)


var Commands = []cli.Command{
	commandAdd,
	commandList,
	commandDelete,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "",
	Description: `
`,
	Action: doAdd,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: ""
	Description: `
`,
	Action: doList,
}

var commandDelete = cli.Command{
	Name:  "delete",
	Usage: "",
	Description: `
`,
	Action: doDelete,
}


func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}


func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAdd(c *cli.Context) {
}

func doList(c *cli.Context) {
}

func doDelete(c *cli.Context) {
}
```
