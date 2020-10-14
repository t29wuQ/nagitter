package main

import (
    "os"
    "github.com/urfave/cli/v2"
    "github.com/t29wuQ/nagitter/twitter"
)

func main() {
    var message string
    consumerKey := os.Getenv("NAGITTER_CK")
    consumerSecret := os.Getenv("NAGITTER_CS")
    accessToken := os.Getenv("NAGITTER_AK")
    accessTokenSecret := os.Getenv("NAGITTER_AS")

    app := &cli.App{
        Flags: []cli.Flag {
            &cli.StringFlag {
                Name: "message",
                Aliases: []string{"m"},
                Usage: "tweet text",
                Destination: &message,
            },
        },
        Action:  func(c *cli.Context) error {
            twitter.Tweet(message, consumerKey, consumerSecret, accessToken, accessTokenSecret);
            return nil
        },
    }

    app.Name = "Nagitter"
    app.Usage = "Nagitter is Tweet client"
    app.Version = "0.0.1"

    app.Run(os.Args)
}
