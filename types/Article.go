package types

type Article struct {
	Header    string
	Author    string
	PageCount int
}

type Articles []*Article

func NewArticle(header, author string, pageCount int) *Article {
	a := new(Article)
	a.Author = author
	a.Header = header
	a.PageCount = pageCount
	return a
}
