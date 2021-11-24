package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Up and running...")
		if err != nil {
			log.Fatalln("Error!")
		}
	})
	router.HandleFunc("/events", getAll).Methods("GET")
	router.HandleFunc("/event", getOne).Methods("GET")
	router.HandleFunc("/event", add).Methods("POST")
	router.HandleFunc("/event", update).Methods("PUT")
	router.HandleFunc("/event", remove).Methods("DELETE")
	log.Println(" Server is listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
