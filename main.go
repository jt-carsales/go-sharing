package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dao = DBconn{}

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
	respondWithJson(w, http.StatusOK, nums)

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

	var resp Response
	resp.CODE = http.StatusOK
	resp.DATA = num.DID
	resp.MSG = "Number Added with Sucess"
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
	var resp Response
	resp.CODE = http.StatusOK
	resp.DATA = num.DID
	resp.MSG = "Number Deleted with Sucess"
	fmt.Print(resp)
	respondWithJson(w, http.StatusOK, resp)
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

	//num := Number{"1234", "AU"}
	//nums := []Number{}
	//fmt.Println(num.DID)
	//fmt.Println(num.ISOCC)
	//err := dao.AddNumber(num)
	//fmt.Print(err)

	//nums, err = dao.GetNumber(num)
	//fmt.Print(nums)
	//fmt.Print(err)

	//err = dao.DeleteNumber(num)
	//fmt.Println(err)

	r := mux.NewRouter()
	r.HandleFunc("/numbers/{prefix}", GetNumberEndpoint).Methods("GET")
	r.HandleFunc("/numbers", AddNumberEndpoint).Methods("POST")
	r.HandleFunc("/numbers", DeleteNumberEndpoint).Methods("DELETE")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal(err)
	}
}
