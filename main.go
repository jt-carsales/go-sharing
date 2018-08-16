package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetNumbersEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet")
	return
}

func GetNumberEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet")
	return
}

func AddNumberEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet")
	return
}

func DeleteNumberEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet")
	return
}

func GetResultsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet")
	return
}

func GetCarriersEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet")
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/numbers", GetNumbersEndpoint).Methods("GET")
	r.HandleFunc("/numbers/{did}", GetNumberEndpoint).Methods("GET")
	r.HandleFunc("/numbers", AddNumberEndpoint).Methods("POST")
	r.HandleFunc("/numbers", DeleteNumberEndpoint).Methods("DELETE")
	r.HandleFunc("/results/{did}", GetResultsEndpoint).Methods("GET")
	r.HandleFunc("/carriers/{did}", GetCarriersEndpoint).Methods("GET")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal(err)
	}
}
