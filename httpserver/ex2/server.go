package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type inputJSON struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type outputJSON struct {
	hiddenResponse string `json:"hiddenResponse"`
	Response       string `json:"response"`
}

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
			_, _ = w.Write([]byte("OK!\n"))
			return
		}
		http.Error(w, "invalid password", http.StatusUnauthorized)
	})

	http.HandleFunc("/postjson", func(w http.ResponseWriter, r *http.Request) {
		var inputData inputJSON
		err := json.NewDecoder(r.Body).Decode(&inputData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// OR
		//buff, err := ioutil.ReadAll(r.Body) // Add error handling etc., but it's not so efficient to read everything into a buffer
		//_ = json.Unmarshal(buff, &inputData)

		if inputData.Name == "test" && inputData.Password == "testPass" {
			r := outputJSON{
				hiddenResponse: "notSeen",
				Response:       "OK",
			}

			_ = json.NewEncoder(w).Encode(r)
			// OR
			// b, err := json.Marshal(r)
			// w.Write(b)
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
