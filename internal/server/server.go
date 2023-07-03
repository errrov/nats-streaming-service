package server

import (
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
	getOrderByUID.HandleFunc("/order/{orderUID:[a-zA-Z0-9-]+}", s.HandleUIDSearch)
	return sm
}

func (s *Server) HandleUIDSearch(rw http.ResponseWriter, r *http.Request) {
	s.logger.Println("inside UID search")
	view := template.Must(template.ParseFiles("./web/static/foundorder.html"))
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
		notFound := template.Must(template.ParseFiles("./web/static/notfound.html"))
		notFound.Execute(rw, order)
		return
	}
	if err != nil {
		http.Error(rw, "Error finding Order", http.StatusInternalServerError)
		return
	}
	view.Execute(rw, order)
}
