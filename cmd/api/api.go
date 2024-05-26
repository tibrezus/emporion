package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tibrezus/emporion/service/user"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewAPIServer(addr string, dbPool *pgxpool.Pool) (*APIServer, error) {
	return &APIServer{
		addr: addr,
		db:   dbPool,
	}, nil
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
