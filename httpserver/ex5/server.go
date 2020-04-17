package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK!\n"))
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "id")

		_, _ = fmt.Fprintf(w, "ID: %s\n", ID)
	})

	r.Get("/{specID:[a-z-]+}", func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "specID")

		_, _ = fmt.Fprintf(w, "Special ID: %s\n", ID)
	})

	log.Println("Listening:", *addr)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
