package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// Testing with testify and asserts
func TestRoutes(t *testing.T) {
	app := &App{}

	r := chi.NewRouter()
	app.Routes(r)

	require.Len(t, r.Routes(), 3)
	assert.True(t, hasRoutePatter(r.Routes(), "/"))
	assert.True(t, hasRoutePatter(r.Routes(), "/add"))
	assert.True(t, hasRoutePatter(r.Routes(), "/list"))
}

func hasRoutePatter(routes []chi.Route, pattern string) bool {
	for _, r := range routes {
		if r.Pattern == pattern {
			return true
		}
	}
	return false
}

// Testing with httptest's response recorder
func TestMainHandler(t *testing.T) {
	app := &App{} // No dependencies in the handler, so this should work

	rr := httptest.NewRecorder()
	app.handleMain(rr, nil) // We don't use the request here so it could be nil. Or you could create a "fake" request: `httptest.NewRequest(http.MethodGet, "/", nil)`

	require.Equal(t, http.StatusOK, rr.Code) // It's not changed, but test it anyway here for an example
	assert.Equal(t, "main app endpoint and stuff\n", rr.Body.String())
}

func TestListHandlerWithDBError(t *testing.T) {
	dbMock := &mockDatabase{}           // No New or any init step needed
	app := NewApp(zap.NewNop(), dbMock) // We'll need the DB as a dep, so use the New call

	// Create a fake request/response for the database
	// The mock will panic if there's and unexpected call: `assert: mock: I don't know what to return because the method call was unexpected.`
	dbMock.On("Select", mock.AnythingOfType("*[]app.message"), "SELECT id, name, message, created FROM message ORDER BY id DESC").
		Return(errors.New("mocked error")) // The "error" part of the response

	rr := httptest.NewRecorder()
	app.handleList(rr, nil)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, `{"error":"mocked error","code":500}`+"\n", rr.Body.String())
}

func TestListHandlerWithEmptyResult(t *testing.T) {
	dbMock := &mockDatabase{}           // No New or any init step needed
	app := NewApp(zap.NewNop(), dbMock) // We'll need the DB as a dep, so use the New call

	// Create a fake request/response for the database
	// The mock will panic if there's and unexpected call: `assert: mock: I don't know what to return because the method call was unexpected.`
	dbMock.On("Select", mock.AnythingOfType("*[]app.message"), "SELECT id, name, message, created FROM message ORDER BY id DESC").
		Return(nil) // The "error" part of the response

	rr := httptest.NewRecorder()
	app.handleList(rr, nil)

	require.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"messages":null}`+"\n", rr.Body.String())
}

func TestListHandlerWithResults(t *testing.T) {
	dbMock := &mockDatabase{}           // No New or any init step needed
	app := NewApp(zap.NewNop(), dbMock) // We'll need the DB as a dep, so use the New call

	// Create a fake request/response for the database
	// The mock will panic if there's and unexpected call: `assert: mock: I don't know what to return because the method call was unexpected.`
	dbMock.On("Select", mock.AnythingOfType("*[]app.message"), "SELECT id, name, message, created FROM message ORDER BY id DESC").
		// Some magic here: https://github.com/vektra/mockery#return-value-provider-functions
		Return(func(dest interface{}, query string, args ...interface{}) error {
			msgs, ok := dest.(*[]message)
			if !ok {
				return errors.New("invalid type for dest")
			}

			*msgs = []message{
				{
					ID:      1,
					Name:    "testName1",
					Message: "testMessage1",
					Created: time.Date(2020, 4, 27, 10, 10, 20, 0, time.UTC),
				},
				{
					ID:      2,
					Name:    "testName2",
					Message: "testMessage2",
					Created: time.Date(2020, 4, 27, 10, 11, 20, 0, time.UTC),
				},
			}
			return nil
		})

	rr := httptest.NewRecorder()
	app.handleList(rr, nil)

	require.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"messages":[{"id":1,"name":"testName1","message":"testMessage1","created":"2020-04-27T10:10:20Z"},{"id":2,"name":"testName2","message":"testMessage2","created":"2020-04-27T10:11:20Z"}]}`+"\n", rr.Body.String())
}
