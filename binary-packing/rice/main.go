package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

//go:generate rice embed-go

func main() {
	addr := flag.String("addr", ":8080", "HTTP Listen address")
	flag.Parse()

	box := rice.MustFindBox("static")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fp, err := box.Open("gopher.png")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer fp.Close()
		_, _ = io.Copy(w, fp)
	})

	fmt.Println("Listening", *addr)
	_ = http.ListenAndServe(*addr, nil)
}
