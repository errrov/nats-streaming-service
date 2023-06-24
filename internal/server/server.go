package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	inMemoryStorage "wildberries_L0/internal/storage/in_memory"

	"github.com/gorilla/mux"
)

type Server struct {
	Srv      *http.Server
	MemCache inMemoryStorage.MemoryStorage
}

func NewServer() *Server {
	server := &Server{
		Srv: &http.Server{
			ErrorLog: log.New(os.Stdout, "subscriber part ", log.LstdFlags),
			Addr:     ":8080",
		},
		MemCache: *inMemoryStorage.NewInMemory(),
	}
	server.Srv.Handler = server.SetupHandlers()
	return server
}

func (s *Server) SetupHandlers() *mux.Router {
	sm := mux.NewRouter()
	getOrderByUID := sm.Methods(http.MethodGet).Subrouter()
	getOrderByUID.HandleFunc("/order/{uid:[a-zA-Z0-9]+}", s.HandleUIDSearch)
	getOrderByUID.HandleFunc("/order", s.getOrder)
	return sm
}

func (s *Server) getOrder(rw http.ResponseWriter, r *http.Request) {

}

func (s *Server) HandleUIDSearch(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var uid string
	uid, ok := vars["uid"]
	if !ok {
		http.Error(rw, "Error finding uid param", http.StatusBadRequest)
		return
	}
	order, err := s.MemCache.FindByUID(uid)
	if err != nil {
		http.Error(rw, "Error finding uid in memchache", http.StatusNotFound)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(*order)
}
