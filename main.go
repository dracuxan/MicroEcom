package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"MicroEcom/handlers"
)

const PORT = ":9090"

func main() {
	l := log.New(os.Stdout, "api: ", log.LstdFlags)
	ph := handlers.NewProducts(l)

	// sm := http.NewServeMux()
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddelwarePoroductValidation)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddelwarePoroductValidation)

	s := &http.Server{
		Addr:         PORT,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Printf("listening on https://localhost%s\n", PORT)

		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Println("Got Signal:", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
