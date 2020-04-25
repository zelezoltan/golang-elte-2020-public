package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbMysql  = "mysql"
	dbSqlite = "sqlite"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")

	mysqlConfig := flag.String("mysqldb", "root:root@tcp(127.0.0.1:3306)/testdb?parseTime=true", "Mysql database connection string")
	sqliteConfig := flag.String("sqlitedb", "file:httpserver/ex7/sqlite.db?cache=shared&parseTime=true", "SQLite database connection string")
	db := flag.String("db", "sqlite", "Supported database (sqlite, mysql)")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync() // flushes buffer, if any

	if *db != dbMysql && *db != dbSqlite {
		log.Fatalln("Invalid db selection. Supported: mysql, sqlite")
		return
	}

	// Setup the DB connection
	var sqlxConn *sqlx.DB
	if *db == dbMysql {
		sqlxConn, err = sqlx.Connect("mysql", *mysqlConfig) // Connect also sends a Ping to check the server Open won't do that!
	} else {
		sqlxConn, err = sqlx.Connect("sqlite3", *sqliteConfig) // Connect also sends a Ping to check the server Open won't do that!
	}
	if err != nil {
		logger.Fatal("database connection failure", zap.Error(err))
	}
	defer sqlxConn.Close()

	if *db == dbSqlite {
		// Limit the concurrent connections when using sqlite
		sqlxConn.SetMaxOpenConns(1)
	}

	// Setup the app part
	app := NewApp(logger, sqlxConn)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello there!")
	})

	r.Mount("/api", app.Routes(chi.NewRouter()))

	// Start the HTTP server
	logger.Info("Listening", zap.String("addr", *addr))
	err = http.ListenAndServe(*addr, r)
	if err != nil {
		logger.Fatal("server failure", zap.Error(err))
	}
}
