package api

import (
	"encoding/json"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"net/http"
	"rest-api-golang/types"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	connString := "sqlserver://sa:Sa123456@localhost:1433?database=goDB"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var article types.Article
	db.First(&article, 1)
	json.NewEncoder(w).Encode(article)
}

func NewArticle(w http.ResponseWriter, r *http.Request) {
	var p types.Article
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
