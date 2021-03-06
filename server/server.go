package server

import (
	"net/http"

	"github.com/JonathonGore/go-service/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func New(api handlers.API) (*Server, error) {
	s := &Server{Router: mux.NewRouter()}

	s.Router.HandleFunc("/v1/{name}", api.HelloWorld).Methods(http.MethodGet)

	return s, nil
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin) // Restrict this to proper origins
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.Router.ServeHTTP(rw, req)
}
