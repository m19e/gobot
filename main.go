package main

import (
	"bot/keys"
	"bot/media"
	"log"
	"net/url"
)

func main() {
	api := keys.GetTwitterApi()

	base64String := media.LoadEncodedMediaString("./subs/image.png")

	media, _ := api.UploadMedia(base64String)

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	text := "test tweet"
	tweet, err := api.PostTweet(text, v)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tweet.Text)
}
