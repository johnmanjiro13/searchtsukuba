package slack

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
)

func SendSlack(userID, tweet string) error {
	channel := "search-tsukuba"
	token := os.Getenv("SLACK_TSUKUBA_TOKEN")
	endpoint := "https://slack.com/api/chat.postMessage"
	message := "@" + userID + ": " + tweet + "\nhttps://twitter.com/" + userID

	values := url.Values{}
	values.Set("token", token)
	values.Add("channel", channel)
	values.Add("text", message)

	req, err := http.NewRequest(
		"POST",
		endpoint,
		bytes.NewBufferString(values.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
