package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// NotifyItems method is notifing items to slack.
func NotifyItems(items []Item) error {
	channel := os.Getenv("SLACK_CHANNEL")
	err := postMessage("商品の投稿を検知しました。", channel)
	if err != nil {
		return err
	}
	for _, item := range items {
		err := postMessage(fmt.Sprintf("%s\n%s", item.name, item.href), channel)
		if err != nil {
			Error("送信途中でエラーが発生しました。")
			return err
		}
	}
	return nil
}

// postMessage method is sending message to slack.
// ref: https://api.slack.com/methods/chat.postMessage#arg_as_user
func postMessage(msg string, channel string) error {
	token := os.Getenv("SLACK_TOKEN")
	urlValue := url.Values{}
	urlValue.Set("token", token)
	urlValue.Set("channel", channel)
	urlValue.Set("text", msg)
	// 基本的にbotのtokenを利用することを想定している
	urlValue.Set("as_user", "true")
	res, err := http.PostForm("https://slack.com/api/chat.postMessage", urlValue)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	Info(fmt.Sprintf("%s", body))

	return err
}
