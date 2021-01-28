package main

import (
	"fmt"
	"rest-api-golang/types"
)

func main() {

	article := types.NewArticle("test Header", "test Author", 4)
	fmt.Println(article.Header)
}
