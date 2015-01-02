package obamasearch

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
)

func NewTwitterClient() *anaconda.TwitterApi {
	var (
		CONSUMER_KEY    = os.Getenv("TWITTER_CONSUMER_KEY")
		CONSUMER_SECRET = os.Getenv("TWITTER_CONSUMER_SECRET")
	)

	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)

	return anaconda.NewTwitterApi("", "")
}

func Search(query string) {
	api := NewTwitterClient()
	searchResult, error := api.GetSearch(query, nil)

	if error != nil {
		fmt.Println(error)
		return
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}
