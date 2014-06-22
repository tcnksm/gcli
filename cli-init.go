package main

import (
	"fmt"
	flag "github.com/dotcloud/docker/pkg/mflag"
	"log"
	"os"
	"strings"
	"text/template"
)

var versionTemplate = template.Must(template.ParseFiles("templates/version.tmpl"))
var mainTemplate = template.Must(template.ParseFiles("templates/main.tmpl"))
var commandsTemplate = template.Must(template.ParseFiles("templates/commands.tmpl"))

var versionGo = GoSource{
	Name:     "version.go",
	Template: *versionTemplate,
}

var commandsGo = GoSource{
	Name:     "commands.go",
	Template: *commandsTemplate,
}

type Application struct {
	Name, Author, Email string
	HasSubCommand       bool
	SubCommands         []SubCommand
}

type SubCommand struct {
	Name, DefineName, FunctionName string
}

func defineApplication(appName string, inputSubCommands []string) Application {

	hasSubCommand := false
	if inputSubCommands[0] != "" {
		hasSubCommand = true
	}

	return Application{
		Name:          appName,
		Author:        GitConfig("user.name"),
		Email:         GitConfig("user.email"),
		HasSubCommand: hasSubCommand,
		SubCommands:   defineSubCommands(inputSubCommands),
	}
}

func defineSubCommands(inputSubCommands []string) []SubCommand {

	var subCommands []SubCommand

	if inputSubCommands[0] == "" {
		return subCommands
	}

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

func showHelp() {
	fmt.Fprintf(os.Stderr, helpText)
}

func main() {

	var (
		flVersion     = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
		flHelp        = flag.Bool([]string{"h", "-help"}, false, "Print this message and quit")
		flDebug       = flag.Bool([]string{"-debug"}, false, "Run as DEBUG mode")
		flSubCommands = flag.String([]string{"s", "-subcommands"}, "", "Conma-seplated list of sub-commands to build")
		flForce       = flag.Bool([]string{"f", "-force"}, false, "Overwrite application without prompting")
	)

	flag.Parse()

	if *flHelp {
		showHelp()
		os.Exit(0)
	}

	if *flVersion {
		showVersion()
		os.Exit(0)
	}

	if *flDebug {
		os.Setenv("DEBUG", "1")
		debug("Run as DEBUG mode")
	}

	inputSubCommands := strings.Split(*flSubCommands, ",")
	debug("inputSubCommands:", inputSubCommands)

	appName := flag.Arg(0)
	debug("appName:", appName)

	if appName == "" {
		fmt.Fprintf(os.Stderr, "Application name must not be blank\n")
		os.Exit(1)
	}

	if _, err := os.Stat(appName); err == nil && *flForce {
		err = os.RemoveAll(appName)
		assert(err)
	}

	if _, err := os.Stat(appName); err == nil {
		fmt.Fprintf(os.Stderr, "%s is already exists, overwrite it? [Y/n]: ", appName)
		var ans string
		_, err := fmt.Scanf("%s", &ans)
		assert(err)

		if ans == "Y" {
			err = os.RemoveAll(appName)
			assert(err)
		} else {
			os.Exit(0)
		}
	}

	// Create directory
	err := os.Mkdir(appName, 0766)
	assert(err)

	application := defineApplication(appName, inputSubCommands)

	// Create verion.go
	err = versionGo.generate(appName, application)
	assert(err)

	// Create <appName>.go
	mainGo := GoSource{
		Name:     appName + ".go",
		Template: *mainTemplate,
	}
	mainGo.generate(appName, application)
	assert(err)

	// Create commands.go
	if application.HasSubCommand {
		commandsGo.generate(appName, application)
	}

	err = GoFmt(appName)
	assert(err)

	os.Exit(0)
}

const helpText = `Usage: cli-init [options] [application]

cli-init is the easy way to start building command-line app.

Options:

  -s="", --subcommands=""    Comma-separated list of sub-commands to build
  -f, --force                Overwrite application without prompting 
  -h, --help                 Print this message and quit
  -v, --version              Print version information and quit
  --debug=false              Run as DEBUG mode

Example:

  $ cli-init todo
  $ cli-init -s add,list,delete todo
`
