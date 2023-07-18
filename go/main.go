package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	portEnvKey  = "PORT"
	defaultPort = 8080
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	port := strconv.Itoa(defaultPort)
	if len(os.Getenv(portEnvKey)) != 0 {
		port = os.Getenv(portEnvKey)
	}
	addr := ":" + port

	log.Printf("starting HTTP server on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
