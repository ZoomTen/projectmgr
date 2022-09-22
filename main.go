package main

import (
	"fmt"
	"log"
	"os"

	"projectmgr/commands"

	"github.com/akamensky/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("projectmgr", "A projects manager")

	subcommands := map[string]*argparse.Command{
		"init":  parser.NewCommand("init", "Initializes a project"),
		"web":   parser.NewCommand("web", "Views the project info in a web browser."),
		"info":  parser.NewCommand("info", "Views the project info in the terminal."),
		"build": parser.NewCommand("build", "Builds the project using its specified executable."),
		"validate": parser.NewCommand("validate", "Validate the project file."),
	}

	// Now parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	if subcommands["init"].Happened() {
		commands.Init()
	} else if subcommands["web"].Happened() {
		commands.Web()
	} else if subcommands["info"].Happened() {
		commands.Info()
	} else if subcommands["build"].Happened() {
		commands.Build()
	} else if subcommands["validate"].Happened() {
		commands.Validate()
	} else {
		log.Fatal("unknown action fired...")
	}
}
