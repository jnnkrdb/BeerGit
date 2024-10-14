package main

import (
	"flag"

	"github.com/xaxa-io/autogit-go/git"
)

func main() {

	_LOG := flag.Bool("log", false, "log to file: Logs/log.txt")
	_CONSOLE := flag.Bool("show", false, "show console output")

	flag.Parse()

	git.Init(_LOG, _CONSOLE)

	git.GitUpdate("gitrepos.yaml")
}
