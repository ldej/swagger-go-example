// Package docs A thing server.
//
// Documentation of the thing API.
//
//     	Schemes: http
//     	BasePath: /api/v1
//     	Version: 1.0.0
//     	Host: localhost:8080
//
// 	   	TermsOfService: http://swagger.io/terms/
//		Contact: Laurence de Jong<support@ldej.nl> https://ldej.nl/
//		License: MIT http://opensource.org/licenses/MIT
//
//     	Consumes:
//     	- application/json
//
//     	Produces:
//     	- application/json
//
//
// swagger:meta
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
)

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
