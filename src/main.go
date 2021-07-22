package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grosgrain/zipcodes/src/zipCodesService"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)

	//Freight Matching route
	freightMatchingRouter := router.PathPrefix("/freightMatching").Subrouter()
	freightMatchingRouter.HandleFunc("/zipCodes", zipCodesService.GetAllZipCodesData).Methods("GET")
	freightMatchingRouter.HandleFunc("/zipCodeLookup", zipCodesService.Lookup).Methods("POST")
	log.Fatal(http.ListenAndServe(":4001", router))

}
