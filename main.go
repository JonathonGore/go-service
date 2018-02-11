package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/JonathonGore/dots/yaml"
	"github.com/JonathonGore/go-service/config"
	"github.com/JonathonGore/go-service/handlers"
	"github.com/JonathonGore/go-service/server"
)

func main() {
	var api handlers.API
	var conf config.Config

	conf, err := yaml.New("config.yml")
	if err != nil {
		log.Fatalf("unable to parse configuration file", err)
	}

	api, err = handlers.New()
	if err != nil {
		log.Fatalf("unable to create handler: %v", err)
	}

	s, err := server.New(api)
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%v", conf.GetInt("server.port")),
		Handler:   s,
		TLSConfig: &tls.Config{},
	}

	log.Fatal(srv.ListenAndServe())
}
