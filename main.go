package main

import (
	"log"
	"os"

	"github.com/phanikumarps/sap/cmd"
)

func main() {
	err := cmd.RunCommand(os.Args[1:])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
