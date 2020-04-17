package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	staticDir := flag.String("static", "/Users/gergely.radics1/go/src/github.com/gerifield/golang-elte-2020-public/httpserver/ex4/static", "Static folder")
	flag.Parse()

	// With multiple mux you could even listen on multiple ports with the same app and this could help a better endpoint organizing
	mux := http.NewServeMux()

	// Server some static files from a folder
	mux.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir(*staticDir)))) // We should trim the prefix (the trailing `/` is important!) if we don't server the files on the root endpoint (`/`)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK!\n"))
	})

	log.Println("Listening:", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
