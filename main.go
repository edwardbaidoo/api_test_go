package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"A GET request made"}`))
	log.Output(1, "Logger: a GET method called")
}

func post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"A POST request made"}`))
	log.Output(1, "Logger: a POST method called")
}

func put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"A PUT request made"}`))
	log.Output(1, "Logger: a PUT method called")
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"A DELETE request made"}`))
	log.Output(1, "Logger: a DELETE method called")
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"An unsupported request made"}`))
	log.Output(1, "Logger: an unsupported method called")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notfound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
