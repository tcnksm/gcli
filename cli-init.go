package main

import (
  "fmt"
  flag "github.com/dotcloud/docker/pkg/mflag"
)

func main() {
  var (
    flVersion  = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
    flHelp     = flag.Bool([]string{"h", "-help"}, false, "Print this message")
    flDebug    = flag.Bool([]string{"-debug"}, false, "Run as DEBUG mode")
  )

  flag.Parse()

  if *flDebug {
    os.Setenv("DEBUG", "1")
  }

  if *flVersion {
    showVersion()
    os.Exit(0)
  }
  
}
