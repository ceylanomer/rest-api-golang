package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"net/http"
	"rest-api-golang/api"
	"rest-api-golang/types"
)

func main() {
	InitializeDB()
	StartAPI()
}

func InitializeDB() {
	connString := "sqlserver://sa:Sa123456@localhost:1433?database=goDB"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&types.Article{})
	db.Create(&types.Article{
		Header:    "TestHeader",
		Author:    "TestAuthor",
		PageCount: 5,
	})

	var article types.Article
	db.First(&article, 1)
	db.First(&article, "Author = ?", "TestAuthor")
	db.Model(&article).Update("Author", "AuthorTest")

}

func StartAPI() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", api.GetAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", api.NewArticle).Methods("POST")
	fmt.Println("Now listening on: localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
