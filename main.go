package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/johnmanjiro13/searchtsukuba/keys"
	"github.com/johnmanjiro13/searchtsukuba/slack"
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
		if isContainsTsukuba(t.FullText) {
			fmt.Println(t.FullText)
			slack.SendSlack(t.User.ScreenName, t.FullText)
		}
	}
}

func isContainsTsukuba(tweet string) bool {
	if strings.Contains(tweet, "つくば") || strings.Contains(tweet, "筑波") {
		return true
	}
	return false
}
