package main

import (
	"github.com/PuerkitoBio/goquery"
)

type ParserInterface interface {
	GetImageUrl(doc *goquery.Document) (link string)
	GetTitle(doc *goquery.Document) (title string)
	GetPrice(doc *goquery.Document) (price string)
}
