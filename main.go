package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sunderee/thoughtster/api"
)

func main() {
	var tweetContents *string = parseTweetContentFromFlags()
	if tweetContents == nil {
		os.Exit(1)
	}

	envFile, envFileLoadingError := loadDotEnv()
	if envFileLoadingError != nil {
		panic(envFileLoadingError)
	}

	oauth1Model := &api.Oauth1Model{
		ConsumerKey:    envFile["API_KEY"],
		ConsumerSecret: envFile["API_KEY_SECRET"],
		AccessToken:    envFile["ACCESS_TOKEN"],
		TokenSecret:    envFile["ACCESS_TOKEN_SECRET"],
	}
	oauth1Header := oauth1Model.BuildOauth1AuthenticationHeader("post", "https://api.twitter.com/2/tweets")

	postTweetResponse, postTweetResponseError := api.PostTweet(*tweetContents, oauth1Header)
	if postTweetResponseError != nil {
		panic(postTweetResponseError)
	}

	fmt.Printf("Tweet %s posted successfully\n", postTweetResponse.Data.TweetID)
}

func parseTweetContentFromFlags() *string {
	twitterText := flag.String("tweet", "", "Tweet you want to post (shorter than 280 characters)")
	flag.Parse()

	if len(*twitterText) == 0 || len(*twitterText) > 280 {
		fmt.Println("Your Tweet can't be empty or longer than 280 characters")
		return nil
	}

	return twitterText
}
