package _api

import (
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) RunApiServer() error {
	router := http.NewServeMux()
	rootRouter := newRootRouter(router, s.db)
	rootRouter.registerRoutes()
	server := http.Server{
		Addr:    s.addr,
		Handler: chain(router),
	}
	log.Printf("api svc started %s", s.addr)
	return server.ListenAndServe()
}
