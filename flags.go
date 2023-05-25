package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func runCommand(args []string) error {

	if len(args) < 1 {
		return errors.New("subcommands missing")
	}

	switch os.Args[1] {
	case "server":
		return runServer()
	case "services":
		fmt.Println("services command")
	default:
		return fmt.Errorf("unknown command %s", os.Args[1])
	}

	return nil
}

func runServer() error {
	cmds := []Execute{
		newServerCommand(),
	}

	for _, cmd := range cmds {
		if cmd.Name() == os.Args[1] {
			cmd.Init(os.Args[0:])
			return cmd.Run()
		}
	}
	return nil
}

func newServerCommand() *Command {

	c := &Command{
		fs: flag.NewFlagSet("server", flag.ContinueOnError),
	}
	c.fs.StringVar(&c.name, "name", "ls", "list servers")
	return c

}

func (c *Command) Run() error {

	switch c.fs.Arg(1) {
	case "server":
		fmt.Println("command", c.fs.Arg(1))
		err := runServerCommands(c)
		if err != nil {
			return err
		}
	case "services":
		fmt.Println("command", c.fs.Arg(1))
	default:
		fmt.Println("unknown sub-command", c.fs.Arg(1))
	}
	return nil

}

func runServerCommands(c *Command) error {

	if len(c.fs.Args()) < 3 {
		return errors.New("subcommands for server missing")
	}

	switch c.fs.Arg(2) {
	case "ls":
		fmt.Printf("server subcommand %s", c.fs.Arg(2))
	case "start":
		fmt.Printf("server subcommand %s", c.fs.Arg(2))
		Server, err := startServer()
		if err != nil {
			fmt.Printf("error listening for server: %s\n", err)
		}
		fmt.Printf("started server at %s", Server.Addr)
	case "stop":
		fmt.Printf("server subcommand %s", c.fs.Arg(2))
		stopServer(SAPServer)
	default:
		return fmt.Errorf("unknown sub-command %s", c.fs.Arg(2))
	}
	return nil

}

type Command struct {
	fs   *flag.FlagSet
	name string
}

type Execute interface {
	Name() string
	Init([]string) error
	Run() error
}

func (c *Command) Name() string {
	return c.fs.Name()
}

func (c *Command) Init(args []string) error {
	return c.fs.Parse(args)
}
