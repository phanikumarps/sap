package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/phanikumarps/sap/server"
)

func RunCommand(args []string) error {

	if len(args) < 1 {
		return errors.New("subcommands missing")
	}

	switch os.Args[1] {
	case "server":
		return cmdServer()
	case "services":
		log.Println("services command")
	case "config":
		log.Println("config command")
	default:
		var cmd string
		log.Printf("unknown command %s", cmd)
	}

	return nil
}

func cmdServer() error {
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
		log.Println("sub-command", c.fs.Arg(1))
		if err := runServerCommands(c); err != nil {
			return err
		}
	case "services":
		log.Println("sub-command", c.fs.Arg(1))
	default:
		return fmt.Errorf("unknown sub-command %s", c.fs.Arg(1))
	}
	return nil

}

func runServerCommands(c *Command) error {

	if len(c.fs.Args()) < 3 {
		return errors.New("subcommands for server missing")
	}

	var s *server.SapServer

	switch c.fs.Arg(2) {
	case "ls":
		log.Printf("server subcommand %s", c.fs.Arg(2))
	case "start":
		log.Printf("server subcommand %s", c.fs.Arg(2))
		s = server.NewSapServer("3333")
		// s, err := server.StartServer()
		err := s.StartServer()
		if err != nil {
			log.Printf("error listening for server: %s\n", err)
		}
		log.Printf("started server at %s", s.Addr)
	case "stop":
		log.Printf("server subcommand %s", c.fs.Arg(2))
		if err := s.StopServer(); err != nil {
			log.Fatal(err)
			return err
		}
		// server.StopServer(s)
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
