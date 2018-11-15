package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/people", GetPeople).Methods("GET")
	router.HandleFunc("/api/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/api/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {

}

func GetPerson(w http.ResponseWriter, r *http.Request) {

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

}
