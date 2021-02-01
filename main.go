package main

import (
	"net/http"
	"os"

	log "go.lafronz.com/tools/logger/stackdriver"
)

type serverDetails struct {
	port      string
	directory string
}

var s serverDetails

func init() {
	// use PORT environment variable, or default to 8080
	s.port = "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		s.port = fromEnv
	} else {
		log.Warning("No Port Provided, using default: %s", s.port)
	}

	// use DIR environment variable, or default to 8080
	s.directory = "./static"
	if fromEnv := os.Getenv("DIR"); fromEnv != "" {
		s.directory = fromEnv
	} else {
		log.Warning("No DIR Provided, using default: %s", s.directory)
	}
}

func main() {
	fs := http.FileServer(http.Dir(s.directory))
	http.Handle("/", fs)

	log.Info("Listening on :" + s.port + "...")
	err := http.ListenAndServe(":"+s.port, nil)
	if err != nil {
		log.Critical("%s", err)
	}
}
