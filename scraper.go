package main

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Item is result of listing items in yahoo auction.
type Item struct {
	name string
	href string
}

// scrape methos is scraping html in yahoo auction.
func scrape(responseBoby io.ReadCloser) ([]Item, error) {

	items := []Item{}
	doc, err := goquery.NewDocumentFromReader(responseBoby)
	if err != nil {
		return nil, err
	}
	// id=AS1m3に出品商品が格納されている
	doc.Find("div#AS1m3 div.bd.cf h3 a").Each(func(index int, s *goquery.Selection) {
		// 出品情報の商品名とリンクを表示
		attr, result := s.Attr("href")
		if result {
			items = append(items, Item{s.Text(), attr})
		} else {
			Warn(fmt.Sprintf("リンクのない商品があります。: %s", s.Text()))
		}
	})
	return items, nil
}
