package main

import (
	"fmt"
	"os"

	"github.com/phanikumarps/sap/cmd"
)

func main() {
	err := cmd.RunCommand(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
