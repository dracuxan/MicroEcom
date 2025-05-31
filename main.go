package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"MicroEcom/handlers"
)

const PORT = ":9090"

func main() {
	l := log.New(os.Stdout, "api: ", log.LstdFlags)
	h := handlers.NewHello(l)
	b := handlers.NewBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", h)
	sm.Handle("/bye", b)

	s := &http.Server{
		Addr:         PORT,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
