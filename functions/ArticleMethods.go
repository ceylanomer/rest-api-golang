package functions

import "rest-api-golang/types"

func GetPageCount(article *types.Article) (articlePageCount int) {
	articlePageCount = article.PageCount
	return articlePageCount
}
