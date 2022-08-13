package main

import (
	"errors"
	dingtalkRobot "github.com/group-robot/dingtalk-robot/v2"
	feishuRobot "github.com/group-robot/feishu-robot/v2"
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
	client *feishuRobot.Client
}

func NewFeiShuMessage(webhook, secret string) *FeiShuMessage {
	client := feishuRobot.NewClient()
	client.Webhook = webhook
	client.Secret = secret
	return &FeiShuMessage{
		client: client,
	}
}

func (message *FeiShuMessage) SendText(content string, isAll bool, _ []string) (string, error) {
	textMessage := feishuRobot.NewTextMessage(content, isAll)
	rep, err := message.client.SendMessage(textMessage)
	if err != nil {
		return "", err
	}
	return rep.Msg, nil
}
func (message *FeiShuMessage) SendMarkdown(title string, content string, isAll bool, _ []string) (string, error) {
	mdContent := content
	if isAll {
		mdContent = mdContent + "\n<at id=all></at>"
	}
	interactiveMessage := feishuRobot.NewInteractiveMessage()
	interactiveMessage.SetHeader(
		feishuRobot.NewCardHeader(feishuRobot.NewCardTitle(title, nil)),
	).SetConfig(
		feishuRobot.NewCardConfig().SetWideScreenMode(true),
	).AddElements(
		feishuRobot.NewCardMarkdown(mdContent),
	)
	rep, err := message.client.SendMessage(interactiveMessage)
	if err != nil {
		return "", err
	}
	return rep.Msg, nil
}
