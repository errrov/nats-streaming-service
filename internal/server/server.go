package server

import (
	"log"
	"net/http"
	"os"
	inMemoryStorage "wildberries_L0/internal/storage/in_memory"
)

type Server struct {
	Server *http.Server
	MemCache *inMemoryStorage.MemoryStorage
}

func NewServer() *Server {
	return &Server{
		Server: &http.Server{
			ErrorLog: log.New(os.Stdout, "subscriber part ", log.LstdFlags),
			Addr: ":8080",
		},
		MemCache: inMemoryStorage.NewInMemory(),
	}
}