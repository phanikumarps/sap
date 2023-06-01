package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type resource struct {
}
type SapServer struct {
	resource resource
	*http.Server
	Router *http.ServeMux
}

func NewSapServer(addr string) *SapServer {
	r := resource{}
	m := http.NewServeMux()
	s := &http.Server{
		Addr:    addr,
		Handler: m,
	}
	return &SapServer{r, s, m}
}
func (s *SapServer) StopServer() error {
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := s.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
	return nil
}

func (s *SapServer) StartServer() error {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() error {

		err := s.Server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}()
	log.Print("Server Started")
	<-done

	return nil

}

func StartServer() (*http.Server, error) {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	mux.HandleFunc("/", rootHandler)
	go func() error {
		err := server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}()
	log.Print("Server Started")
	<-done

	return server, nil

}

func StopServer(s *http.Server) {
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
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
