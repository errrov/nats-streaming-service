package server

import (
	"encoding/json"
	"log"
	"nats-streaming-service/internal/model"
	"nats-streaming-service/internal/storage"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Server struct {
	Srv    *http.Server
	Cache  *storage.Storage
	logger *log.Logger
}

func NewServer(l *log.Logger) (*Server, error) {
	var err error
	server := &Server{
		Srv: &http.Server{
			ErrorLog: log.New(os.Stdout, "subscriber-nats ", log.LstdFlags),
			Addr:     ":8080",
		},
		logger: l,
	}
	server.Cache, err = storage.StorageInit(l)
	if err != nil {
		return nil, err
	}
	server.Srv.Handler = server.SetupHandlers()
	l.Println("Server created")
	return server, nil
}

func (s *Server) SetupHandlers() *mux.Router {
	sm := mux.NewRouter()
	getOrderByUID := sm.Methods(http.MethodGet).Subrouter()
	getOrderByUID.Path("/order").Queries("orderUID", "{orderUID}").HandlerFunc(s.HandleUIDSearch)
	getOrderByUID.HandleFunc("/order", s.getOrderPage)
	return sm
}

func (s *Server) getOrderPage(rw http.ResponseWriter, r *http.Request) {
	s.logger.Println("Handling orders page")
	tmp := template.Must(template.ParseFiles("./web/static/orders.html"))
	tmp.Execute(rw, "Hello data")
}

func (s *Server) HandleUIDSearch(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var uid string
	uid, ok := vars["orderUID"]
	s.logger.Println("Handling UID search:", uid)
	if !ok {
		http.Error(rw, "Error finding uid param", http.StatusBadRequest)
		return
	}
	order, err := s.Cache.FindByUID(uid)
	if err == model.ErrNotFound {
		http.Error(rw, "Error finding uid in memchache", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error finding Order", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(*order)
}
