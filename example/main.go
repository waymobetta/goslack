package main

import (
	"fmt"
	"log"
	"os"

	"github.com/waymobetta/goslack"
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
			Token: os.Getenv("SLACK_API_KEY"),
			Bot: &goslack.Bot{
				Name:  "gopher_bot",
				Emoji: ":golang:",
			},
		},
	)

	err := gs.SendString(
		msg,
		os.Getenv("SLACK_CHANNEL_NAME"), // XXXXXXXXX
	)
	if err != nil {
		log.Fatal(err)
	}
}
