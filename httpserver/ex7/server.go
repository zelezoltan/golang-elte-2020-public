package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	dbConfig := flag.String("db", "root:root@tcp(127.0.0.1:3306)/testdb?parseTime=true", "Database connection string")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync() // flushes buffer, if any

	db, err := sqlx.Connect("mysql", *dbConfig) // Connect also sends a Ping to check the server Open won't do that!
	if err != nil {
		logger.Fatal("database connection failure", zap.Error(err))
	}
	defer db.Close()

	app := NewApp(logger, db)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello there!")
	})

	r.Mount("/api", app.Routes(chi.NewRouter()))

	logger.Info("Listening", zap.String("addr", *addr))
	err = http.ListenAndServe(*addr, r)
	if err != nil {
		logger.Fatal("server failure", zap.Error(err))
	}
}
