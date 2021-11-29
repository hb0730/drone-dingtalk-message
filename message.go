package main

import (
	"errors"
	robot "github.com/hb0730/dingtalk-robot"
	"github.com/hb0730/feishu-robot"
	"strings"
)

type IMessage interface {
	SendText(content string, isAll bool, mobiles []string) error
	SendMarkdown(title string, content string, isAll bool, mobiles []string) error
}

func getSupportMessage(typeStr, webhok, secret string) (IMessage, error) {
	switch strings.ToLower(typeStr) {
	case "dingtalk":
		return NewDingTalkMessage(webhok, secret), nil
	case "feishu":
		return NewFeiShuMessage(webhok, secret), nil
	default:
		return nil, errors.New("missing message")
	}
}

type DingTalkMessage struct {
	client *robot.Client
}

// NewDingTalkMessage new DingTalkMessage
func NewDingTalkMessage(webhok string, secret string) *DingTalkMessage {
	client := robot.NewClient(webhok, secret)
	return &DingTalkMessage{
		client: client,
	}
}

// SendText send text
func (message *DingTalkMessage) SendText(content string, isAll bool, mobiles []string) error {
	text := robot.NewTextMessage().SetContent(content)
	text.SetIsAtAll(isAll).SetAtMobiles(mobiles)
	_, err := message.client.Send(text)
	return err
}
func (message DingTalkMessage) SendMarkdown(title string, content string, isAll bool, mobiles []string) error {
	markdown := robot.NewMarkdownMessage().SetTitle(title).SetText(content)
	markdown.SetAtMobiles(mobiles).SetIsAtAll(isAll)
	_, err := message.client.Send(markdown)
	return err
}

type FeiShuMessage struct {
	client *feishu.Client
}

func NewFeiShuMessage(webhok, secret string) *FeiShuMessage {
	return &FeiShuMessage{
		client: feishu.NewClient(webhok, secret),
	}
}

func (message *FeiShuMessage) SendText(content string, isAll bool, _ []string) error {
	textMessage := feishu.NewTextMessage().SetContent(content).SetAtAll(isAll)
	_, err := message.client.Send(textMessage)
	return err
}
func (message *FeiShuMessage) SendMarkdown(title string, content string, isAll bool, _ []string) error {
	mdContent := content
	if isAll {
		mdContent = mdContent + "\n<at id=all></at>"
	}
	interactiveMessage := feishu.NewInteractiveMessage()
	interactiveMessage = interactiveMessage.SetHeader(feishu.NewCardHeader().SetTitle(feishu.NewCardTitle().SetContent(title)))
	interactiveMessage.SetConfig(feishu.NewCardConfig().SetEnableForward(true))
	interactiveMessage.SetElements(feishu.NewElementsContent().AddElement(feishu.NewMarkdownCardContent().SetContent(mdContent)))
	_, err := message.client.Send(interactiveMessage)
	return err
}
