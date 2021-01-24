// Package main is the entrypoint to the RESTful api. It creates the server and hosts it with a
// goroutine. The server can be stopped with a system interrupt
package main

import (
	"fmt"
	"github.com/ferpart/pokerest/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

// routes is a structure containing all the routes used with this RESTful api.
var routes = []struct {
	path    string
	method  string
	handler http.HandlerFunc
}{
	{
		path:    "/advantage",
		method:  http.MethodGet,
		handler: handlers.Advantage,
	},
	{
		path:    "/commonMoves",
		method:  http.MethodGet,
		handler: handlers.CommonMoves,
	},
}

// main is in charge of setting the port and executing the server's goroutine.
func main() {
	port := 5000
	router := mux.NewRouter()

	for _, r := range routes {
		router.HandleFunc(fmt.Sprintf("/api%s", r.path), r.handler).Methods(r.method)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	defer srv.Close()

	// The server is ran using a gorutine so it doesn't block
	go func() {
		log.Infof("Starting server on port %d...\n", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Block execution until a sigint signal is recieved
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
