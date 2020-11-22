package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Item struct {
	name string
	href string
}

func scrape(responseBoby io.ReadCloser) []Item {

	items := []Item{}
	doc, _ := goquery.NewDocumentFromReader(responseBoby)
	// id=AS1m3に出品商品が格納されている
	doc.Find("div#AS1m3 div.bd.cf h3 a").Each(func(index int, s *goquery.Selection) {
		// 出品情報の商品名とリンクを表示
		attr, _ := s.Attr("href")
		items = append(items, Item{s.Text(), attr})
	})
	return items
}
