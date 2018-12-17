package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"gopkg.pco.dk/masterclass/demoapi"

)

func main() {
	logger := demoapi.NewLogger()
	var port uint
	var verbose bool = false
	flag.UintVar(&port, "port", 8080, "Port to run webserver on")
	flag.BoolVar(&verbose, "verbose", false, "Run in verbosemode. Prints debug statement to stdout")
	flag.Parse()
	if verbose {
		logger.Level = logrus.DebugLevel
	}

	demoapi.LoadConfig()
	logger.Infof("Stating API server on port %d\n", port)

	handler := HttpHandler{logger}
	router := mux.NewRouter()

	//Home handler
	router.HandleFunc("/", handler.HomeHandler()).Methods("GET", "HEAD", "OPTIONS")

	http.Handle("/", router)

	hostAndPort := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(hostAndPort, nil); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
			"Host":  hostAndPort,
		}).Fatal("Unable to start webserver")
	}

}
