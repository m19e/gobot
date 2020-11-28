package main

import (
	"bot/keys"
	"bot/media"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const subsDir = "subs/"

var subs []string
var shuffledSubs []string

func init() {
	filepath.Walk(subsDir, appendSubs)
	shuffledSubs = shuffled(subs)
}

func appendSubs(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	subs = append(subs, strings.Join(strings.Split(path, subsDir), ""))
	return nil
}

func shuffled(data []string) []string {
	n := len(data)
	r := make([]string, n)
	copy(r, data)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { r[i], r[j] = r[j], r[i] })
	return r
}

func main() {
	chStop := make(chan int, 1)
	TimerFunc(chStop)

	time.Sleep(time.Minute * 2)
	chStop <- 0

	close(chStop)

	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	time.Sleep(time.Second * 1)
	log.Println("Application End.")
}

func TimerFunc(stopTimer chan int) {
	go func() {
		ticker := time.NewTicker(1 * time.Hour)

		for {
			select {
			case <-ticker.C:
				if len(shuffledSubs) == 0 {
					shuffledSubs = shuffled(subs)
				}
				log.Println(shuffledSubs[0])
				post(subsDir, shuffledSubs[0])
				shuffledSubs = shuffledSubs[1:]

			case <-stopTimer:
				log.Println("Timer stop.")
				ticker.Stop()
				break
			}
		}
	}()
}

func post(dir, fp string) {
	api := keys.GetTwitterApi()

	base64String := media.LoadEncodedMediaString(fmt.Sprintf("%s%s", dir, fp))

	media, _ := api.UploadMedia(base64String)

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	text := fmt.Sprintf("#%s", strings.Split(fp, "/")[0])
	tweet, err := api.PostTweet(text, v)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tweet.Text)
}
