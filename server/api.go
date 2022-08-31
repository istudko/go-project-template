package server

import (
	"context"
	"github.com/istudko/go-project-template/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
)

// Start is a function for starting a server
func Start() {
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	handler := c.Handler(Router())
	startServer(handler)
}

func startServer(handler http.Handler) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("starting %s server on port: %s", config.GetAppName(), config.GetPort())

	server := &http.Server{
		Addr:    ":" + config.GetPort(),
		Handler: handler,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-stop
	log.Println("stopping server.")
	server.Shutdown(context.Background())
}
