package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		items := scrape(res.Body)
		for _, item := range items {
			Info(item.name)
			Info(item.href)
		}
		return
	}

	body, _ := ioutil.ReadAll(res.Body)
	Error(fmt.Sprintf("%s", body))
	log.Panicln("予期せぬエラーが帰ってきました。")
}
