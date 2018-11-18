package keys

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_API_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}
