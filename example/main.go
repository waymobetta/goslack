package main

import (
	"fmt"
	"log"
	"os"

	"github.com/waymobetta/goslack"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(
			"usage: goslack <message> <channel_name>\n\nexample: goslack 'hello world' @jon",
		)
		return
	}

	msg := os.Args[1]
	chnl := os.Args[2]

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
		chnl, // XXXXXXXXX
	)
	if err != nil {
		log.Fatal(err)
	}
}
