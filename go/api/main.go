package main

import (
	"api/data"
	"api/endpoints"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"github.com/gorilla/mux"
)

func CreateRouter() (*mux.Router, error) {

	router := mux.NewRouter().StrictSlash(true)

	userRouter := router.PathPrefix("/api/user").Subrouter()
	err := endpoints.MakeUserHandlers(userRouter)
	if err != nil {
		return nil, err
	}

	return router, nil

}

func main() {

	parser := argparse.NewParser("print", "Chat app API")

	port := parser.String("p", "port", &argparse.Options{Required: false, Help: "the port to run on", Default: "5001"})

	log.Printf("Starting API on %v", *port)

	// postgres connection
	err := data.Connection()

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	// router
	router, err := CreateRouter()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	server := &http.Server{
		Addr:              ":" + *port,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

}
