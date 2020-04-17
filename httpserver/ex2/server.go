package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	flag.Parse()

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "use POST", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		password := r.FormValue("password")
		fmt.Println("Name:", name, "Pass", password)

		if name == "" || password == "" {
			http.Error(w, "empty name or password", http.StatusBadRequest)
			return
		}

		if name == "test" && password == "testPass" {
			_, _ = w.Write([]byte("OK!"))
			return
		}
		http.Error(w, "invalid password", http.StatusUnauthorized)
	})

	log.Println("Listening:", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
