package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type contextKey string

func (c contextKey) String() string {
	return "test.package.context." + string(c)
}

var (
	nameContextKey      = contextKey("name")
	requestIDContextKey = contextKey("requestID")
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	flag.Parse()

	r := chi.NewRouter()

	// Group to use a middleware in front of multiple calls
	r.Group(func(r chi.Router) {
		r.Use(middlewareOne)

		r.Get("/endpoint1", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintln(w, "Hello, this endpoint is protected!")
		})

		r.Get("/endpoint2", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintln(w, "Hello, this endpoint2 is protected!")
		})
	})

	// Similar to a group, but this time specify a pattern "prefix" too
	r.Route("/api", func(r chi.Router) {
		r.Use(middlewareTwo("token2"))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`Imagine some custom "api" stuff here...\n`))
		})

		// Inline middleware inject is also possible
		r.With(middlewareThree).Get("/something", contextTesterHandler)
		r.Get("/something2", contextTesterHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(loggerMiddleware)
		r.Get("/longtime", func(w http.ResponseWriter, r *http.Request) {
			rnd := rand.Intn(5)
			fmt.Println("Sleep time:", rnd)
			time.Sleep(time.Duration(rnd) * time.Second)
			_, _ = fmt.Fprintf(w, "Sleep time: %d\n", rnd)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(betterLoggerMiddleware)
		r.Get("/longtime2", func(w http.ResponseWriter, r *http.Request) {
			requestID, _ := r.Context().Value(requestIDContextKey).(string)

			rnd := rand.Intn(5)
			fmt.Println("Sleep time:", rnd)
			time.Sleep(time.Duration(rnd) * time.Second)
			_, _ = fmt.Fprintf(w, "RequestID: %s, sleep time: %d\n", requestID, rnd)
		})
	})

	log.Println("Listening:", *addr)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func middlewareOne(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "someToken" {
			h.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func middlewareTwo(token string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") == token {
				next.ServeHTTP(w, r)
				return
			}
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
	}
}

func contextTesterHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`Imagine something more here!` + "\n"))
	name, err := getNameContext(r.Context())
	if err != nil {
		_, _ = fmt.Fprintf(w, "Sorry, no context key, error: %v\n", err)
	} else {
		_, _ = fmt.Fprintf(w, "Name from context: %s\n", name)
	}
}

func middlewareThree(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		// Useful article about the context keys: https://medium.com/@matryer/context-keys-in-go-5312346a868d
		ctx := context.WithValue(r.Context(), nameContextKey, name)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getNameContext(ctx context.Context) (string, error) {
	if name, ok := ctx.Value(nameContextKey).(string); !ok {
		return "", errors.New("missing context key")
	} else {
		return name, nil
	}
}

func loggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Println("Incoming request:", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		fmt.Println("Duration:", time.Now().Sub(now))
	})
}

func betterLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		requestID := uuid.New().String()
		fmt.Println("Incoming request:", r.Method, r.URL.Path, requestID)
		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), requestIDContextKey, requestID)))
		fmt.Println(requestID, "duration:", time.Now().Sub(now))
	})
}
