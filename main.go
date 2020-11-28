package main

import (
	"bot/keys"
	"fmt"
	"log"
)

func main() {
	api := keys.GetTwitterApi()

	text := "test tweet"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tweet.Text)
}
