package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var uri string = "https://slack.com/api/chat.postMessage"

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
	return s.post(msg)
}

func (s *chat) post(text string) (bool, error) {
	text = text + " - cobra, " + time.Now().Format("2006-01-02 15:04:05")

	params := url.Values{}
	params.Add("token", s.token)
	params.Add("channel", s.channel)
	params.Add("text", text)

	req, err := http.NewRequest("POST", uri, nil)
	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = (params).Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var res Result
	json.Unmarshal([]byte(bodyBytes), &res)

	// fmt.Println("response Body:", string(bodyBytes))
	fmt.Printf("Result: %+v\n", res)
	return res.Ok, nil
}

type BotProfile struct {
	Id      string `json:id`
	Deleted bool   `json:deleted`
	Name    string `json:name`
	Updated int    `json:"updated"`
	AppId   string `json:"app_id"`
}

type Message struct {
	BotId      string     `json:"bot_id"`
	Type       string     `json:type`
	Text       string     `json:text`
	User       string     `json:user`
	Team       string     `json:team_id`
	BotProfile BotProfile `json:"bot_profile"`
}

type Result struct {
	Ok      bool    `json:"ok"`
	Error   string  `json:"error"`
	Channel string  `json:"channel"`
	Ts      string  `json:"ts"`
	Message Message `json:"message"`
}
