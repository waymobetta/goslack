package main

import (
	"fmt"
	"log"
	"os"

	goslack "github.com/waymobetta/goslack"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(
			"usage: goslack <message>\n\nexample: goslack \"hello world\"",
		)
		return
	}

	msg := os.Args[1]

	gs := goslack.NewGoSlack(
		&goslack.Config{
			APIToken:    os.Getenv("SLACK_API_KEY"),
			ChannelName: os.Getenv("SLACK_CHANNEL_NAME"),
			Domain:      os.Getenv("SLACK_DOMAIN_NAME"),
			Bot: &goslack.Bot{
				Name:  "gopher_bot",
				Emoji: ":golang:",
			},
		},
	)

	err := gs.SendString(
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}
}
