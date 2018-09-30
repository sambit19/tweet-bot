package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	//Please enter the details below by signing up for the twitter developer api
	consumerKey := ""
	consumerSecret := ""
	accessToken := ""
	accessSecret := ""

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	//Oauth1 http.Client with automatically authorize requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	//Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name: %v \n", user.ScreenName)

	// Search tweets to retweet
	searchParams := &twitter.SearchTweetParams{
		Query:      "wef",
		Count:      5,
		ResultType: "recent",
		Lang:       "en",
	}

	searchResult, _, _ := client.Search.Tweets(searchParams)

	// Retweet
	for _, tweet := range searchResult.Statuses {
		tweet_id := tweet.ID
		client.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})

		fmt.Printf("RETWEETED: %+v\n", tweet.Text)
	}
}
