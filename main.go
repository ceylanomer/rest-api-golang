package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api-golang/api"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", api.GetAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", api.NewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
