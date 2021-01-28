package tests

import (
	"rest-api-golang/functions"
	"rest-api-golang/types"
	"testing"
)

func CreateTestArticle(header, author string, pageCount int) *types.Article {
	return types.NewArticle(header, author, pageCount)
}

func TestGetPageCount(t *testing.T) {
	if functions.GetPageCount(CreateTestArticle("testHeader","testAuthor",4)) != 4 {
		t.Error("should be 4")
	}
}
