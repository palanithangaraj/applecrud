package api

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"applecrud/pkg/config"
	"applecrud/pkg/db"
	"applecrud/pkg/model"
)

const (
	Database   = "um"
	Collection = "auth"
)

func ReadAppleRouter(db db.Dao) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", GetNamedJoke(&db, config.Config{PreLog: "TEST -"})).Methods("GET")
	return r
}

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	HealthCheck(w, req)
	if w.Body.String() != "OK" || w.Code != 200 {
		t.Error("healthcheck did not return 200 OK as expected.")
	}
}

func TestRespondWithJson(t *testing.T) {
	w := httptest.NewRecorder()
	code := 200
	var name model.Name
	respondWithJSON(w, code, name)
	if w.Code != 200 {
		t.Errorf("code did not match expected result. Expected %v, got %v", code, w.Code)
	}
	if err := json.NewDecoder(w.Body).Decode(&name); err != nil {
		t.Errorf("Unable to marshal Json, expected to marshal an auth object correct. %s", err.Error())
	}
}

func TestRespondWithError(t *testing.T) {
	w := httptest.NewRecorder()
	code := 400

	respondWithError(w, code, "not found")
	if w.Code != 400 {
		t.Errorf("code did not match expected result. Expected %v, got %v", code, w.Code)
	}
	if !strings.Contains(w.Body.String(), "not found") {
		t.Error("Error message was not returned in body.")
	}
}
