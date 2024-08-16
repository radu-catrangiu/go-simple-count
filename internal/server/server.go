package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/radu-catrangiu/go-simple-count/internal/filewriter"
)

func getCount(w http.ResponseWriter, r *http.Request) {
	count, _ := filewriter.ReadIntFromFile()
	fmt.Fprintf(w, "%d\n", count)
}

func postCount(w http.ResponseWriter, r *http.Request) {
	count, _ := filewriter.ReadIntFromFile()
	count++
	filewriter.WriteIntToFile(count)
	fmt.Fprintf(w, "OK")
}

func Start(addr string) {
	http.HandleFunc("GET /count", getCount)
	http.HandleFunc("POST /count", postCount)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}
}
