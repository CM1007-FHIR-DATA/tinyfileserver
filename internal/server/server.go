package server

import (
	"log"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)

	log.Println("Serving files from the static directory on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Server failed to start, err: ", err)
	}
}
