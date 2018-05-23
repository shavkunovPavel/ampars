package main

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

type AmazonParser struct{}

// получить адрес картинки
func (a *AmazonParser) GetImageUrl(doc *goquery.Document) (link string) {
	doc.Find("#imageBlockContainer #imgBlkFront").Each(func(i int, s *goquery.Selection) {
		li, _ := s.Attr("data-a-dynamic-image")
		re := regexp.MustCompile(`({|}|\\|"|(:\[\d+,\d+\]))`)
		link = re.ReplaceAllString(li, "")
		return
	})
	return
}

// получить заголовок
func (a *AmazonParser) GetTitle(doc *goquery.Document) (title string) {
	doc.Find("#productTitle").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
		return
	})
	return
}

// получить цену
func (a *AmazonParser) GetPrice(doc *goquery.Document) (price string) {
	doc.Find(".a-button-selected .a-size-base.a-color-price").Each(func(i int, s *goquery.Selection) {
		pr := s.Text()
		re := regexp.MustCompile(`\d+\.?\d+`)
		price = re.FindString(pr)
		return
	})
	return
}
