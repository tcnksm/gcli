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

type BasicInfo struct {
	Name, Author, Email string
	HasSubCommand       bool
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

func writeMain(wr io.Writer) {
	basicInfo := BasicInfo{
		Name:          "test",
		Author:        "taichi",
		Email:         "test@gmail.com",
		HasSubCommand: false,
	}

	err := mainTemplate.Execute(wr, basicInfo)
	assert(err)
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

	subCommands := strings.Split(*flSubCommands, ",")
	debug("subCommands:", subCommands)

	os.Mkdir(appName, 0766)

	versionFile, err := os.Create(strings.Join([]string{appName, "version.go"}, "/"))
	assert(err)
	defer versionFile.Close()
	writeVersion(versionFile)

	mainFile, err := os.Create(strings.Join([]string{appName, appName + ".go"}, "/"))
	assert(err)
	defer mainFile.Close()
	writeMain(mainFile)

	commandsFile, err := os.Create(strings.Join([]string{appName, "commands.go"}, "/"))
	assert(err)
	defer commandsFile.Close()
	writeVersion(commandsFile)
}
