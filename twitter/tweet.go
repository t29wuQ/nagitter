package twitter

import (
    "github.com/ChimeraCoder/anaconda"
)

func Tweet(text string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) {
    anaconda.SetConsumerKey(consumerKey)
    anaconda.SetConsumerSecret(consumerSecret)
    api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

    // ツイート実行
    _, err := api.PostTweet(text, nil)
    if(err != nil){
        panic(err)
    }
}
