package main

import (
	"errors"
	"github.com/CatchZeng/dingtalk"
	"github.com/hb0730/feishu-robot"
)

type IMessage interface {
	SendText(content string, isAll bool, mobiles []string) error
	SendMarkdown(title string, content string, isAll bool, mobiles []string) error
}

func getSupportMessage(typeStr, accessToken, secret string) (IMessage, error) {
	switch typeStr {
	case "dingtalk":
		return NewDingTalkMessage(accessToken, secret), nil
	case "feishu":
		return NewFeiShuMessage(accessToken, secret), nil
	default:
		return nil, errors.New("missing message")
	}
}

type DingTalkMessage struct {
	client *dingtalk.Client
}

// NewDingTalkMessage new DingTalkMessage
func NewDingTalkMessage(accessToken string, secret string) *DingTalkMessage {
	client := dingtalk.NewClient(accessToken, secret)
	return &DingTalkMessage{
		client: client,
	}
}

// SendText send text
func (message *DingTalkMessage) SendText(content string, isAll bool, mobiles []string) error {
	text := dingtalk.NewTextMessage().SetContent(content).SetAt(
		mobiles,
		isAll,
	)
	_, err := message.client.Send(text)
	return err
}
func (message DingTalkMessage) SendMarkdown(title string, content string, isAll bool, mobiles []string) error {
	markdown := dingtalk.NewMarkdownMessage().SetMarkdown(title, content).SetAt(mobiles, isAll)
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
