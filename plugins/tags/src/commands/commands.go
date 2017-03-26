package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	columnize "github.com/ryanuber/columnize"
)

const (
	helpHeader = `Usage: dokku tags[:COMMAND]

Manage docker image tags.

Additional commands:`

	helpContent = `
    tags <app>, List all app image tags
    tags:create <app> <tag>, Add tag to latest running app image
    tags:deploy <app> <tag>, Deploy tagged app image
    tags:destroy <app> <tag>, Remove app image tag
`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "tags:help":
		usage()
	case "help":
		fmt.Print(helpContent)
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	config := columnize.DefaultConfig()
	config.Delim = ","
	config.Prefix = "\t"
	config.Empty = ""
	content := strings.Split(helpContent, "\n")[1:]
	fmt.Println(helpHeader)
	fmt.Println(columnize.Format(content, config))
}
