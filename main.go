package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getTop10(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get top 10")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/top10", getTop10)
	log.Fatal(http.ListenAndServe(":8080", router))
}