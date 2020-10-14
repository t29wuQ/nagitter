package twitter

import (
    "github.com/ChimeraCoder/anaconda"
)

type Twitter struct {
    api *anaconda.TwitterApi
}

func NewTwitter(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Twitter {
    anaconda.SetConsumerKey(consumerKey)
    anaconda.SetConsumerSecret(consumerSecret)
    api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
    twitter := Twitter{
        api,
    }
    return &twitter
}

func (twitter *Twitter) Tweet(text string) {
    _, err := twitter.api.PostTweet(text, nil)
    if(err != nil){
        panic(err)
    }
}
