package main

import (
	"fmt"
	"net/http"
	"os"
)

var SAPServer *http.Server

func main() {
	err := runCommand(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
