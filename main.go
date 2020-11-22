package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 検知する対象のユーザ名を指定する
	userName := os.Getenv("USER_NAME")

	url := fmt.Sprintf("https://auctions.yahoo.co.jp/seller/%s", userName)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// 出品していないユーザは404を返す(ヤグオクの仕様)
	if res.StatusCode == 404 {
		Info(fmt.Sprintf("status code is %d: %s", res.StatusCode, "出品されている商品はありません。"))
		return
	}

	if res.StatusCode >= 200 && res.StatusCode < 500 {
		Info(fmt.Sprintf("status code is %d: %s", res.StatusCode, "出品されている品物を検知しました。"))
		doc, _ := goquery.NewDocumentFromReader(res.Body)

		// id=AS1m3に出品商品が格納されている
		doc.Find("div#AS1m3 div.bd.cf h3 a").Each(func(index int, s *goquery.Selection) {
			// 出品情報の商品名とリンクを表示
			Info(s.Text())
			attr, _ := s.Attr("href")
			Info(attr)
		})
		return
	}

	body, _ := ioutil.ReadAll(res.Body)
	Error(fmt.Sprintf("%s", body))
	log.Panicln("予期せぬエラーが帰ってきました。")
}
