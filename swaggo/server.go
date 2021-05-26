package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"

	"github.com/ldej/swagger-go-example/swaggo/swagger"
)

// @title A thing server
// @version 1.0
// @description A thing server
// @termsOfService http://swagger.io/terms/

// @contact.name Laurence de Jong
// @contact.url https://ldej.nl/
// @contact.email support@ldej.nl

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

type Server struct {
	router *mux.Router
	log    *log.Logger
	stopCh chan os.Signal
}

func NewServer() *Server {
	s := &Server{
		log:    log.Default(),
		router: mux.NewRouter(),
		stopCh: make(chan os.Signal, 1),
	}

	swagger.SwaggerInfo.Title = "A thing server"
	swagger.SwaggerInfo.Description = "This is a dynamically set description."
	swagger.SwaggerInfo.Version = "1.0"
	swagger.SwaggerInfo.Host = "localhost:8080"
	swagger.SwaggerInfo.BasePath = "/api/v1"
	swagger.SwaggerInfo.Schemes = []string{"http", "https"}
	s.router.PathPrefix("/swaggerd/").Handler(httpSwagger.WrapHandler)

	s.router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("swagger"))))

	s.router.HandleFunc("/api/v1/thing", s.ListThings).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/thing/new", s.CreateThing).Methods(http.MethodPost)
	s.router.HandleFunc("/api/v1/thing/{uuid}", s.GetThing).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/thing/{uuid}", s.UpdateThing).Methods(http.MethodPut)
	s.router.HandleFunc("/api/v1/thing/{uuid}", s.DeleteThing).Methods(http.MethodDelete)
	return s
}

func (s *Server) ListenAndServe(addr string) {
	ctx := context.Background()

	signal.Notify(s.stopCh, syscall.SIGINT, syscall.SIGTERM)
	hs := &http.Server{Addr: addr, Handler: s.router}

	go func() {
		s.log.Println(ctx, fmt.Sprintf("Listening on: %s", addr))

		if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Fatal(ctx, err)
		}
	}()

	<-s.stopCh
	s.log.Println(ctx, "Shutting down the server...")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	if err := hs.Shutdown(ctxTimeout); err != nil {
		s.log.Fatal(ctx, err)
	}
}

func (s *Server) Shutdown() {
	s.stopCh <- os.Interrupt
}
