package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dao = DBconn{}

//GetNumbersEndpoint ...
func GetNumbersEndpoint(w http.ResponseWriter, r *http.Request) {

	nums := []Number{}
	var err error
	nums, err = dao.GetNumbers()
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	var resp Response
	resp.CODE = http.StatusOK
	resp.DATA = nums
	resp.MSG = "Done"
	fmt.Print(resp)
	respondWithJson(w, http.StatusOK, resp)

}

//GetNumberEndpoint ...
func GetNumberEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	num := Number{params["prefix"], ""}
	nums := []Number{}
	var err error
	nums, err = dao.GetNumber(num)
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	var resp Response
	resp.CODE = http.StatusOK
	resp.DATA = nums
	resp.MSG = "Done"
	fmt.Print(resp)
	respondWithJson(w, http.StatusOK, resp)

}

//AddNumberEndpoint ...
func AddNumberEndpoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var num Number
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = dao.AddNumber(num)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var nums []Number
	nums = append(nums, num)

	var resp Response
	resp.CODE = http.StatusOK
	resp.DATA = nums
	resp.MSG = "Number added successfully"
	fmt.Print(resp)
	respondWithJson(w, http.StatusOK, resp)

}

//DeleteNumberEndpoint ...
func DeleteNumberEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var num Number
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = dao.DeleteNumber(num)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var nums []Number
	nums = append(nums, num)

	var resp Response
	resp.CODE = http.StatusOK
	resp.DATA = nums
	resp.MSG = "Number deleted successfully"
	fmt.Print(resp)
	respondWithJson(w, http.StatusOK, resp)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func initialize() {
	dao.host = "127.0.0.1"
	dao.name = "gosharing"
	dao.user = "root"
	dao.port = "3306"
	dao.pass = "password"
	err := dao.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

}

func main() {

	initialize()
	r := mux.NewRouter()
	r.HandleFunc("/", GetNumbersEndpoint).Methods("GET")
	r.HandleFunc("/numbers", GetNumbersEndpoint).Methods("GET")
	r.HandleFunc("/numbers/{prefix}", GetNumberEndpoint).Methods("GET")
	r.HandleFunc("/numbers", AddNumberEndpoint).Methods("POST")
	r.HandleFunc("/numbers", DeleteNumberEndpoint).Methods("DELETE")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal(err)
	}
}
