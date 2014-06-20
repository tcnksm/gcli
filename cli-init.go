package main

import (
	"fmt"
	flag "github.com/dotcloud/docker/pkg/mflag"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

var versionTemplate = template.Must(template.ParseFiles("templates/version.tmpl"))
var mainTemplate = template.Must(template.ParseFiles("templates/main.tmpl"))
var commandsTemplate = template.Must(template.ParseFiles("templates/commands.tmpl"))

type Application struct {
	Name, Author, Email string
	HasSubCommand       bool
	SubCommands         []SubCommand
}

type SubCommand struct {
	Name, DefineName, FunctionName string
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

func showVersion() {
	fmt.Fprintf(os.Stderr, "cli-init v%s\n", Version)
}

func writeVersion(wr io.Writer) {
	err := versionTemplate.Execute(wr, nil)
	assert(err)
}

func writeMain(application Application, wr io.Writer) {
	err := mainTemplate.Execute(wr, application)
	assert(err)
}

func writeCommands(application Application, wr io.Writer) {
	err := commandsTemplate.Execute(wr, application)
	assert(err)
}

func defineSubCommands(inputSubCommands []string) []SubCommand {
	var subCommands []SubCommand
	for _, name := range inputSubCommands {
		subCommand := SubCommand{
			Name:         name,
			DefineName:   "command" + ToUpperFirst(name),
			FunctionName: "do" + ToUpperFirst(name),
		}
		subCommands = append(subCommands, subCommand)
	}

	return subCommands
}

func ToUpperFirst(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func main() {

	var (
		flVersion = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
		flHelp    = flag.Bool([]string{"h", "-help"}, false, "Print this message")
	)

	flagSub := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		flDebug       = flagSub.Bool([]string{"-debug"}, false, "Run as DEBUG mode")
		flSubCommands = flagSub.String([]string{"s", "-subcommands"}, "", "Sub commands")
	)

	flag.Parse()

	if *flHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *flVersion {
		showVersion()
		os.Exit(0)
	}

	appName := flag.Arg(0)
	debug("appName:", appName)

	flagSub.Parse(os.Args[2:])
	if *flDebug {
		os.Setenv("DEBUG", "1")
		debug("Run as DEBUG mode")
	}

	inputSubCommands := strings.Split(*flSubCommands, ",")
	debug("inputSubCommands:", inputSubCommands)

	hasSubCommand := false
	if inputSubCommands[0] != "" {
		hasSubCommand = true
	}
	debug("hasSubCommand:", hasSubCommand)

	application := Application{
		Name:          appName,
		Author:        GitConfig("user.name"),
		Email:         GitConfig("user.email"),
		HasSubCommand: hasSubCommand,
		SubCommands:   defineSubCommands(inputSubCommands),
	}

	os.Mkdir(appName, 0766)

	// Create version.go
	versionFile, err := os.Create(strings.Join([]string{appName, "version.go"}, "/"))
	assert(err)
	defer versionFile.Close()
	writeVersion(versionFile)

	// Create <appName>.go
	mainFile, err := os.Create(strings.Join([]string{appName, appName + ".go"}, "/"))
	assert(err)
	defer mainFile.Close()
	writeMain(application, mainFile)

	if hasSubCommand {
		// Create commands.go
		commandsFile, err := os.Create(strings.Join([]string{appName, "commands.go"}, "/"))
		assert(err)
		defer commandsFile.Close()
		writeCommands(application, commandsFile)
	}

	err = GoFmt(appName)
	assert(err)

	os.Exit(0)
}
