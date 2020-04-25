package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

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
