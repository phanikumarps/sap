package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func StartServer() (*UmcServer, error) {
	log.Print("StartServer: start")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	u := newUmcServer(ctx)
	u.routes = u.registerRoutes()
	if err := u.start(ctx); err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := u.Errgroup.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
		return nil, err
	}
	return u, nil
}
func StopServer(server *UmcServer) error {
	//TODO
	// 1. Graceful shutdown
	// 2. Zero downtime
	return nil
}

type UmcServer struct {
	config     *config
	httpServer *http.Server
	router     *http.ServeMux
	routes     []route
	Errgroup   *errgroup.Group
	ErrCtx     context.Context
}

type config struct {
	port string
}
type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type sapServer interface {
	start(ctx context.Context) error
	stop(ctx context.Context) error
	registerRoutes() []route
}

func loadUmcConfig() *config {
	log.Print("loadUmcConfig: start")
	return &config{
		port: os.Getenv("SAP_UMC_PORT"),
	}
}
func newUmcServer(ctx context.Context) *UmcServer {
	log.Print("newUmcServer: start")
	s := &UmcServer{}
	log.Print("newUmcServer: new server assignment done")
	s.config = loadUmcConfig()
	log.Print("loadUmcConfig: ends")
	s.router = http.NewServeMux()
	log.Print("mux created")
	s.httpServer = &http.Server{
		Addr:    ":" + s.config.port,
		Handler: s.router,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
	log.Print("httpServer created")
	return nil
}

func (s *UmcServer) start(ctx context.Context) error {

	s.Errgroup, s.ErrCtx = errgroup.WithContext(ctx)
	s.Errgroup.Go(func() error {
		return s.httpServer.ListenAndServe()
	})

	return nil
}
func (s *UmcServer) stop(ctx context.Context) error {

	s.Errgroup.Go(func() error {
		<-s.ErrCtx.Done()
		return s.httpServer.Shutdown(ctx)
	})

	return nil
}

func (s *UmcServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
func (s *UmcServer) handleNotFound() http.Handler {
	return http.NotFoundHandler()
}
func (s *UmcServer) handleRootGet() http.HandlerFunc {
	log.Print("handleRootGet: start")
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodGet {
		// 	s.handleNotFound()
		// 	return
		// }
		fmt.Fprint(w, "Hello SAP!")
		log.Print("handleRootGet: ends")
	}
}
func (s *UmcServer) handleGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "World")
		log.Print("handleGreeting: ends")
	}
}
func (r *UmcServer) Get(pattern string, handler http.HandlerFunc) {

}
func (s *UmcServer) registerRoutes() []route {

	//TODO
	log.Print("registerRoutes: start")
	s.router.HandleFunc("/", s.handleGreeting())
	log.Print("registerRoutes: ends")
	return nil
}
