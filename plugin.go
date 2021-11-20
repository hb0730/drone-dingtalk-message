package main

import (
	"errors"
	"github.com/CatchZeng/dingtalk"
	"log"
	"os"
)

type Plugin struct {
	Debug          bool
	DingTalkConfig DingTalkConfig
	client         *dingtalk.Client
}
type DingTalkConfig struct {
	AccessToken string
	Secret      string
}

type Message struct {
	AtAll       bool
	AtMobiles   []string
	Title       string
	Content     string
	MessageType string
}

func (plugin *Plugin) Exec(message Message) error {
	if plugin.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}
	if plugin.DingTalkConfig.AccessToken == "" {
		return errors.New("missing DingTalk access token")
	}
	if message.Content == "" {
		return errors.New("missing Content")

	}
	//create dingtalk client
	plugin.client = dingtalk.NewClient(plugin.DingTalkConfig.AccessToken, plugin.DingTalkConfig.Secret)

	switch message.MessageType {
	case string(dingtalk.MsgTypeText):
		return plugin.SendText(message.Content, message.AtMobiles, message.AtAll)
	case string(dingtalk.MsgTypeMarkdown):
		return plugin.SendMarkdown(message.Title, message.Content, message.AtMobiles, message.AtAll)
	default:
		return errors.New("not support message type")
	}

}
func (plugin *Plugin) SendMarkdown(title string, content string, atMobiles []string, atAll bool) error {
	message := dingtalk.NewMarkdownMessage().SetMarkdown(title, content).SetAt(atMobiles, atAll)
	_, err := plugin.client.Send(message)
	return err
}
func (plugin *Plugin) SendText(content string, atMobiles []string, atAll bool) error {
	message := dingtalk.NewTextMessage().
		SetContent(content).SetAt(
		atMobiles,
		atAll,
	)
	_, err := plugin.client.Send(message)

	return err
}
