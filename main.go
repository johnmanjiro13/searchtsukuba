package main

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/johnmanjiro13/searchtsukuba/keys"
	"github.com/johnmanjiro13/searchtsukuba/slack"
)

func main() {
	api := keys.GetTwitterAPI()

	values := url.Values{}
	values.Set("count", "20")

	timeline, err := api.GetHomeTimeline(values)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for _, t := range timeline {
		if isContainsTsukuba(t.FullText) {
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
