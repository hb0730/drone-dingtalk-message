package main

import (
	"errors"
	dingtalkRobot "github.com/group-robot/dingtalk-robot/v2"
	"github.com/hb0730/feishu-robot"
	"strings"
)

type IMessage interface {
	SendText(content string, isAll bool, mobiles []string) (string, error)
	SendMarkdown(title string, content string, isAll bool, mobiles []string) (string, error)
}

func getSupportMessage(typeStr, webhook, secret string) (IMessage, error) {
	switch strings.ToLower(typeStr) {
	case "dingtalk":
		return NewDingTalkMessage(webhook, secret), nil
	case "feishu":
		return NewFeiShuMessage(webhook, secret), nil
	default:
		return nil, errors.New("unsupported webhook type")
	}
}

type DingTalkMessage struct {
	client *dingtalkRobot.Client
}

// NewDingTalkMessage new DingTalkMessage
func NewDingTalkMessage(webhook string, secret string) *DingTalkMessage {
	client := dingtalkRobot.NewClient()
	client.Webhook = webhook
	client.Secret = secret
	return &DingTalkMessage{
		client: client,
	}
}

// SendText send text
func (message *DingTalkMessage) SendText(content string, isAll bool, mobiles []string) (string, error) {
	textMsg := dingtalkRobot.NewTextMessage(content)
	textMsg.At = dingtalkRobot.NewAt(isAll).SetAtMobiles(mobiles...)
	response, err := message.client.SendMessage(textMsg)
	if err != nil {
		return "", err
	}
	return response.ErrorMessage, nil
}
func (message *DingTalkMessage) SendMarkdown(title string, content string, isAll bool, mobiles []string) (string, error) {
	markdownMsg := dingtalkRobot.NewMarkdownMessage(title, content)
	markdownMsg.At = dingtalkRobot.NewAt(isAll).SetAtMobiles(mobiles...)
	response, err := message.client.SendMessage(markdownMsg)
	if err != nil {
		return "", err
	}
	return response.ErrorMessage, nil
}

type FeiShuMessage struct {
	client *feishu.Client
}

func NewFeiShuMessage(webhook, secret string) *FeiShuMessage {
	return &FeiShuMessage{
		client: feishu.NewClient(webhook, secret),
	}
}

func (message *FeiShuMessage) SendText(content string, isAll bool, _ []string) (string, error) {
	textMessage := feishu.NewTextMessage().SetContent(content).SetAtAll(isAll)
	response, err := message.client.Send(textMessage)
	//return err
	if err != nil {
		return "", err
	}
	return response.Msg, nil
}
func (message *FeiShuMessage) SendMarkdown(title string, content string, isAll bool, _ []string) (string, error) {
	mdContent := content
	if isAll {
		mdContent = mdContent + "\n<at id=all></at>"
	}
	interactiveMessage := feishu.NewInteractiveMessage()
	interactiveMessage = interactiveMessage.SetHeader(feishu.NewCardHeader().SetTitle(feishu.NewCardTitle().SetContent(title)))
	interactiveMessage.SetConfig(feishu.NewCardConfig().SetEnableForward(true))
	interactiveMessage.SetElements(feishu.NewElementsContent().AddElement(feishu.NewMarkdownCardContent().SetContent(mdContent)))
	response, err := message.client.Send(interactiveMessage)
	if err != nil {
		return "", err
	}
	return response.Msg, nil
}
