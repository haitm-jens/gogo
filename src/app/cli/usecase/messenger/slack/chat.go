package slack

import (
	"fmt"

	"gogo/infra/messenger/slack"
)

type chat struct {
	token   string
	channel string
}

func NewChat(token string, channel string) *chat {
	return &chat{
		token:   token,
		channel: channel,
	}
}

func (s *chat) Send(msg string) (bool, error) {
	fmt.Println("send: ", msg, s.token, s.channel)

	c := slack.NewChat(s.token, s.channel)
	c.Send(msg)
	return true, nil
}
