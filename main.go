package main

import (
	"bot/keys"
	"bot/media"
	"log"
	"net/url"
	"time"
)

func main() {
	// chStop := make(chan int, 1)
	// TimerFunc(chStop)

	// time.Sleep(time.Hour * 24 * 365)
	// chStop <- 0

	// close(chStop)

	// time.Sleep(time.Second * 1)

	// log.Println("Application End.")
}

func TimerFunc(stopTimer chan int) {
	go func() {
		ticker := time.NewTicker(30 * time.Minute)

		for {
			select {
			case <-ticker.C:
				post()

			case <-stopTimer:
				log.Println("Timer stop.")
				ticker.Stop()
				break
			}
		}
	}()
}

func post() {
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
