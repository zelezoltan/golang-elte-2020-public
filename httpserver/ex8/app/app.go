package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// App holds the common context for the handlers and stuff
// Preferably this should be in a different package and have a simple New() only instead of NewApp()
// You could also use a "config struct" instead of multiple parameters
type App struct {
	logger *zap.Logger
	db     database
}

type errorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type addJSON struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type addResponse struct {
	ID int `json:"id"`
}

type message struct {
	ID int `json:"id"`

	// Right now it has the same fields  as addJSON, but they have a different meaning and could change, so I won't use that here
	Name    string `json:"name"`
	Message string `json:"message"`

	Created time.Time `json:"created"`
}
type listResponse struct {
	Messages []message `json:"messages"`
}

type database interface {
	Select(dest interface{}, query string, args ...interface{}) error

	sqlx.Execer // This is already defined inside the sqlx package because it's internal testing parts for Select and Get use them.
	// But instead using the `Execer` you could also define the `Exec` by yourself:
	//Exec(query string, args ...interface{}) (sql.Result, error)
}

func NewApp(logger *zap.Logger, db database) *App {
	return &App{
		logger: logger,
		db:     db,
	}
}

// Routes return the endpoint handled by the app
// It receives a router which could help you inject custom middleware or helpers for easier testing
func (a *App) Routes(r *chi.Mux) *chi.Mux {
	r.Get("/", a.handleMain)
	r.Post("/add", a.handleAdd)
	r.Get("/list", a.handleList)
	return r
}

func (a *App) handleMain(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "main app endpoint and stuff")
}

func (a *App) handleAdd(w http.ResponseWriter, r *http.Request) {
	var input addJSON
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		// Should we log all the invalid incoming format errors?
		sendError(a.logger, w, err, http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Message == "" {
		sendError(a.logger, w, errors.New("name or message could not be empty"), http.StatusBadRequest)
		return
	}

	res, err := a.db.Exec("INSERT INTO message (`name`, `message`) VALUES (?,?)", input.Name, input.Message)
	if err != nil {
		sendError(a.logger, w, err, http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	if err := json.NewEncoder(w).Encode(addResponse{ID: int(id)}); err != nil {
		a.logger.Error("add response send failed", zap.Error(err))
		return
	}
}

func (a *App) handleList(w http.ResponseWriter, r *http.Request) {
	var messages []message
	err := a.db.Select(&messages, "SELECT id, name, message, created FROM message ORDER BY id DESC")
	if err != nil {
		sendError(a.logger, w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(listResponse{Messages: messages}); err != nil {
		a.logger.Error("list response send failed", zap.Error(err))
		return
	}
}

// sendError will send a JSON instead of a simple string and do some logging.
// Ideally this should have it's own package with other HTTP related helpers.
// And of course it could have it's own `logger` injected as a context at init etc...
func sendError(logger *zap.Logger, w http.ResponseWriter, err error, code int) {
	logger.Error("error on backend", zap.Error(err), zap.Int("code", code))

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(errorResponse{
		Error: err.Error(),
		Code:  code,
	}); err != nil {
		logger.Error("json error response send failed", zap.Error(err))
	}
}
