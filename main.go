package main

import (
	"log"
	"net/http"
	"os"

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

	l.Printf("listening on https://localhost%s\n", PORT)
	http.ListenAndServe(PORT, sm)
}
