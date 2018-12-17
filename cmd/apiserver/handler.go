package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type HttpHandler struct {
	logger   *logrus.Logger
}

//
//
//
func (this *HttpHandler) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		environment := os.Getenv("PCO_ENV")
		if err != nil {
			return
		}

		view := struct {
			Hostname    string `json:"hostname"`
			Environment string `json:"environment"`
		}{
			hostname,
			environment,
		}

		w.Header().Set("Content-Type", "application/json")
		j, cerr := json.Marshal(view)
		if cerr != nil {
			return
		}
		w.Write(j)
	}
}
