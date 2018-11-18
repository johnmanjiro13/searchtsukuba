package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/johnmanjiro13/searchtsukuba/keys"
)

func main() {
	api := keys.GetTwitterAPI()

	v := url.Values{}
	v.Set("count", "20")

	timeline, err := api.GetHomeTimeline(v)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for _, t := range timeline {
		tweet := t.FullText
		if strings.Contains(tweet, "つくば") {
			fmt.Println(tweet)
		}
	}
}
