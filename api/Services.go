package api

import (
	"encoding/json"
	"net/http"
	"rest-api-golang/types"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := types.Articles{
		types.NewArticle("TestHeader", "TestAuthor", 5),
		types.NewArticle("TestHeader", "TestAuthor", 4),
	}
	json.NewEncoder(w).Encode(articles)
}

func NewArticle(w http.ResponseWriter, r *http.Request) {
	var p types.Article
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
