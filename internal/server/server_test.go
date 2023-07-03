package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNotFoundRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	request, _ := http.NewRequest("GET", "/order/a", nil)
	server, err := NewServer(l)
	if err != nil {
		t.Fatalf("[TEST] error creating server")
	}
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
	server, err := NewServer(l)
	if err != nil {
		t.Fatalf("[TEST] error creating server")
	}
	server.Srv.Handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("[TEST] http status is not OK, %v", recorder.Code)
	}
}
