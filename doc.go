/*
Command gcli generates the codes and its directory structure you need to start building CLI tool by Go. https://github.com/tcnksm/gcli

You can choose your favorite CLI framework. Currently gcli supports the following packages

  - flag (Standard package)
  - codegangsta_cli (https://github.com/codegangsta/cli)
  - mitchellh_cli (https://github.com/mitchellh/cli)
  - go_cmd (go command style)

Usage:

  $ gcli [-version] [-help] <command> [<args>]

Available commands are:

    list       List available cli frameworks
    new        Generate new cli project
    version    Print the gcli version


For example, if you want to build `todo` application which has add (Add new task), list (List all tasks) and delete (Delete a task) command with mitchellh/cli framework,


  $ cd $GOPATH/src/github.com/YOUR_NAME
  $ gcli new -F mitchellh_cli -c add -c list -c delete todo


*/
package main
