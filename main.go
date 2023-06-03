package main

import (
	"log"
	"os"

	"github.com/phanikumarps/sap/cmd"
)

func main() {

	if err := cmd.RunCommand(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
