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
	APIToken    string
	ChannelName string
	Domain      string
	Bot         *Bot
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

// NewGoSlack is a method to configure a new GoSlack object
// @dev creates new GoSlack object
// @param Config object as config
// @return pointer to GoSlack struct
func NewGoSlack(config *Config) *GoSlack {
	return &GoSlack{
		Config: config,
	}
}

func (g *GoSlack) SendString(msg string) error {
	// initialize new Values type in variable payload
	payload := url.Values{}

	// set form values
	payload.Add("token", g.Config.APIToken)
	payload.Add("channel", g.Config.ChannelName)
	payload.Add("text", msg)
	payload.Add("username", g.Config.Bot.Name)
	payload.Add("icon_emoji", g.Config.Bot.Emoji)

	// generate POST request
	req, err := http.NewRequest(
		"POST",
		g.Config.Domain,
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

	// execute GET
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
