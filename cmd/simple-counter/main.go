package main

import (
	"log"
	"os"

	"github.com/radu-catrangiu/go-simple-counter-server/internal/filewriter"
	"github.com/radu-catrangiu/go-simple-counter-server/internal/server"
)

func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if len(addr) == 0 {
		addr = "0.0.0.0:8080"
	}

	filepath := os.Getenv("FILE_PATH")
	if len(filepath) == 0 {
		filepath = "/tmp/counter-data"
	}

	filewriter.Init(filepath)

	log.Println("Starting server on", addr)
	server.Start(addr)
}
