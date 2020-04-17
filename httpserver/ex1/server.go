package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Print some informations here
		log.Println("Method:", r.Method)
		log.Println("URL:", r.URL)
		log.Println("Host:", r.Host)
		log.Println("Headers:", r.Header)

		//log.Println("Parse form err:", r.ParseForm())
		log.Println("Name parsed like a form:", r.Form.Get("name"))
		log.Println("Name parsed like a query value:", r.URL.Query().Get("name"))
		fmt.Println("")

		// Use the query
		queryVals := r.URL.Query()
		name := "anonymous"
		if queryVals.Get("name") != "" {
			name = queryVals.Get("name")
		}
		_, _ = fmt.Fprintf(w, "Hello %s!", name)
	})

	http.HandleFunc("/other", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello2!")
	})

	http.HandleFunc("/another/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Others with trailing `/`")
		_, _ = fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
		_, _ = fmt.Fprintf(w, "Prefix trimmed: %s\n", strings.TrimPrefix(r.URL.Path, "/another"))
	})

	log.Println("Listening:", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
