package goslack

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Config struct holds config info
type Config struct {
	Token string
	Bot   *Bot
}

// GoSlack struct holds pointer to config info
type GoSlack struct {
	Config *Config
}

// Bot struct handles Slack Bot info
type Bot struct {
	Name  string
	Emoji string
}

// NewGoSlack
// @notice NewGoSlack is a method to configure a new GoSlack object
// @dev creates new GoSlack object
// @param config is a pointer to Config object
// @return pointer to GoSlack struct
func NewGoSlack(config *Config) *GoSlack {
	return &GoSlack{
		Config: config,
	}
}

// SendString
// @notice is a method to send a message to a Slack channel
// @dev sends message (string) to channel specified in parameters
// @param msg is a message (string) to be sent
// @param chnl is the channel name (string) to deliver the message to
// @return error
func (g *GoSlack) SendString(msg, chnl string) error {
	// initialize new Values map in payload variable
	payload := url.Values{}

	// set form values
	payload.Add("token", g.Config.Token)
	payload.Add("channel", chnl)
	payload.Add("text", msg)
	payload.Add("username", g.Config.Bot.Name)
	payload.Add("icon_emoji", g.Config.Bot.Emoji)

	// generate POST request
	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		strings.NewReader(payload.Encode()),
	)

	// init new client with custom timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// set headers
	req.Header.Add(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	req.Header.Add(
		"Content-Length",
		strconv.Itoa(
			len(
				payload.Encode(),
			),
		),
	)

	// execute POST
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
