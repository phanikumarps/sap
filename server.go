package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	mux.HandleFunc("/", rootHandler)
	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

type serverAddr string

const keyServerAddr serverAddr = "serverAddr"

func rootHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my website!\n")
}
