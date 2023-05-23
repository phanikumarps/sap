package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	run()
}

func run() {

	startCmd := flag.NewFlagSet("start", flag.ExitOnError)

	stopCmd := flag.NewFlagSet("stop", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'start' or 'stop' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "start":
		startCmd.Parse(os.Args[2:])
		startServer()
	case "stop":
		stopCmd.Parse(os.Args[2:])
	default:
		fmt.Println("expected 'start' or 'stop' subcommands")
		os.Exit(1)
	}

}

func startServer() {
	mux := http.NewServeMux()
	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	mux.HandleFunc("/", rootHandler)
	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()
	<-ctx.Done()
}

type serverAddr string

const keyServerAddr serverAddr = "serverAddr"

func rootHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my website!\n")
}
