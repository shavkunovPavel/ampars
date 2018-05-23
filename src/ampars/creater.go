package main

import (
	"github.com/PuerkitoBio/goquery"
)

// структуры для ответа на запрос
type ResultType struct {
	Url  string    `json:"url"`
	Meta *MetaType `json:"meta"`
}

type MetaType struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Image string `json:"image"`
}

// создает обект для ответа. parser_hadler должен реализовать ParserInterface
func Create_object(parser_hadler ParserInterface, url_in string) (*ResultType, error) {

	doc, err := goquery.NewDocument(url_in)

	if err != nil {
		return nil, err
	}
	res := new(ResultType)
	res.Meta = new(MetaType)

	res.Url = url_in
	res.Meta.Image = parser_hadler.GetImageUrl(doc)
	res.Meta.Title = parser_hadler.GetTitle(doc)
	res.Meta.Price = parser_hadler.GetPrice(doc)
	return res, nil
}
