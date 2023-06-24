package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"wildberries_L0/internal/model"
)

func TestRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/order/a", nil)
	server := NewServer()
	server.MemCache.Add(&model.Order{OrderUID: "a"})
	var order model.Order
	server.Srv.Handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("[TEST] http status is not OK, %v", recorder.Code)
	}
	if err := json.NewDecoder(recorder.Body).Decode(&order); err != nil {
		t.Errorf("[TEST] error parsing response %v", err)
	}
}

func TestNotFoundRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/order/a", nil)
	server := NewServer()
	server.Srv.Handler.ServeHTTP(recorder, request)
	log.Println(recorder.Code)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("[TEST] want StatusNotFound, got: %v", recorder.Code)
	}
}

func TestRequestValid(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/order/ывы", nil)
	server := NewServer()
	server.Srv.Handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("[TEST] http status is not OK, %v", recorder.Code)
	}
}
