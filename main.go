package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	const port string = ":8000"

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Response from the sever")
	})

	router.HandleFunc("/post", getPost).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")

	log.Println("Server is Running on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
