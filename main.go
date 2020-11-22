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

	// TODO: ヤフオクべったりなコードになっているので他のコードが入る余地があるようにしたい
	url := fmt.Sprintf("https://auctions.yahoo.co.jp/seller/%s", userName)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		log.Panicln("http request error.")
	}
	defer res.Body.Close()

	// 出品していないユーザは404を返す(ヤグオクの仕様)
	if res.StatusCode == 404 {
		Info(fmt.Sprintf("status code is %d: %s", res.StatusCode, "出品されている商品はありません。"))
		return
	}

	// 値が返ってきたら実際のスクレイピング処理
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		Info(fmt.Sprintf("status code is %d: %s", res.StatusCode, "出品されている品物を検知しました。"))
		items, err := Scrape(res.Body)
		if err != nil {
			fmt.Println(err)
			log.Panicln("スクレイピング中にエラーが発生しました。")
		}
		for _, item := range items {
			Info(item.name)
			Info(item.href)
		}

		hoge := NotifyItems(items)
		if hoge != nil {
			fmt.Println(hoge)
			log.Panicln("Slackへのポスト中にエラーが発生しました。")
		}
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		log.Panicf("Response Bodyの読み込みに失敗しました。")
	}
	Error(fmt.Sprintf("%s", body))
	log.Panicln("予期せぬエラーが帰ってきました。")
}
