package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var version = "0.0.1"

type config struct {
	port int
	env  string
}

type application struct {
	config  config
	version string
}

func (app *application) serve() error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}

	log.Printf("server started on port: %d", app.config.port)
	return server.ListenAndServe()
}

func main() {

	cfg := config{
		port: 3000,
		env:  "development",
	}

	app := &application{
		config:  cfg,
		version: version,
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
