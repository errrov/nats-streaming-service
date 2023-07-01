package server

import (
	"encoding/json"
	"log"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	"nats-streaming-service/internal/model"
)

func TestRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	request, _ := http.NewRequest("GET", "/order/a", nil)
	server := NewServer(l)
	server.Cache.AddToStorage(&model.Order{OrderUID: "a"})
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
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	request, _ := http.NewRequest("GET", "/order/a", nil)
	server := NewServer(l)
	server.Srv.Handler.ServeHTTP(recorder, request)
	log.Println(recorder.Code)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("[TEST] want StatusNotFound, got: %v", recorder.Code)
	}
}

func TestRequestValid(t *testing.T) {
	recorder := httptest.NewRecorder()
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	request, _ := http.NewRequest("GET", "/order/ывы", nil)
	server := NewServer(l)
	server.Srv.Handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("[TEST] http status is not OK, %v", recorder.Code)
	}
}
