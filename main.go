package main

import (
	"fmt"
	"os"

	"github.com/AakashKath/PubSub/migrations"
	"github.com/mkideal/cli"
)

func validateArgs(ctx *cli.Context) error {
	_ = ctx.Argv().(*nodeTree)
	return nil
}

func arguments() interface{} {
	return new(nodeTree)
}

var help = cli.HelpCommand("Display help information.")

type nodeTree struct {
	cli.Helper
}

var root = &cli.Command{
	Argv: arguments,
	Fn:   validateArgs,
}

func command(commandName string, Description string) *cli.Command {
	var commandObj = &cli.Command{
		Name: commandName,
		Desc: Description,
		Argv: arguments,
		Fn:   validateArgs,
	}
	return commandObj
}

func findArgument(argToFind string, argsList []string) bool {
	for _, value := range argsList {
		if value == argToFind {
			return true
		}
	}
	return false
}

func main() {
	err := cli.Root(root, cli.Tree(help),
		cli.Tree(command("run-server", "Argument to run HTTP Server.")),
		cli.Tree(command("run-client", "Argument to run Client.")),
		cli.Tree(command("migrate", "Argument for running Migrations."))).Run(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
	if findArgument("migrate", os.Args[1:]) {
		migrations.Migrate()
		os.Exit(1)
	}
	if findArgument("run-server", os.Args[1:]) {
		RunRESTServer()
	}
	if findArgument("run-client", os.Args[1:]) {
		RunClient()
	}
}
