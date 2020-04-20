package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	staticDir := flag.String("static", "./httpserver/ex4/static", "Static folder")
	templateDir := flag.String("templateIdr", "./httpserver/ex4/template", "Template folder")
	flag.Parse()

	log.Println("Static folder:", *staticDir)
	log.Println("Template folder:", *templateDir)

	template1, err := template.New("one.tpl").ParseFiles(*templateDir + "/one.tpl") // The name in the new should be the filename OR we should use `template1.ExecuteTemplate(w, "one.tpl", struct...` later at rendering
	// This is because the ParseFiles call, it's easier to use the Parse and add the raw template data in there maybe.
	if err != nil {
		log.Fatalln("template parse failed:", err)
	}

	// With multiple mux you could even listen on multiple ports with the same app and this could help a better endpoint organizing
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK!\n"))
	})

	// Server some static files from a folder
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*staticDir)))) // We should trim the prefix (the trailing `/` is important!) if we don't server the files on the root endpoint (`/`)

	mux.HandleFunc("/templated", func(w http.ResponseWriter, r *http.Request) {
		// More info: https://golang.org/pkg/html/template/
		err := template1.Execute(w, struct {
			SomeTitle  string
			SomeValues []struct {
				Key string
				Val string
			}
		}{
			SomeTitle: "Title here",
			SomeValues: []struct {
				Key string
				Val string
			}{
				{Key: "k1", Val: "v1"},
				{Key: "k2", Val: "v2"},
				{Key: "k3", Val: "v3"},
			},
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Listening:", *addr)
	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
