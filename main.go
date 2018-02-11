package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/JonathonGore/go-service/handlers"
	"github.com/JonathonGore/go-service/server"
)

func main() {
	var api handlers.API

	api, err := handlers.New()
	if err != nil {
		log.Fatalf("unable to create handler: %v", err)
	}

	s, err := server.New(api)
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	srv := &http.Server{
		Addr:      ":3000",
		Handler:   s,
		TLSConfig: &tls.Config{},
	}

	log.Fatal(srv.ListenAndServe())
}
