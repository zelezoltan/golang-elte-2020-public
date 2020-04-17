package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	cert := flag.String("cert", "httpserver/ex3/server.crt", "HTTPS cert")
	privKey := flag.String("privKey", "httpserver/ex3/server.key", "HTTPS cert private key")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("It works!"))
	})

	log.Println("Listening:", *addr)
	err := http.ListenAndServeTLS(*addr, *cert, *privKey, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
