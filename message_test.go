package main

import (
	"os"
	"testing"
)

func TestDingTalkMessage_SendText(t *testing.T) {
	accessToken := os.Getenv("dingtalk_accessToken")
	secret := os.Getenv("dingtalk_secret")
	message := NewDingTalkMessage(accessToken, secret)
	_, err := message.SendText("test", true, nil)
	if err != nil {
		t.Logf("%s", err.Error())
	}
}
func TestDingTalkMessage_SendMarkdown(t *testing.T) {
	accessToken := os.Getenv("dingtalk_accessToken")
	secret := os.Getenv("dingtalk_secret")
	message := NewDingTalkMessage(accessToken, secret)
	_, err := message.SendMarkdown("test", `
### test

> - test
> - test2
`, true, nil)
	if err != nil {
		t.Logf("%s", err.Error())
	}
}

func TestFeiShuMessage_SendText(t *testing.T) {
	webhok := os.Getenv("feishu_webhok")
	secret := os.Getenv("feishu_secret")

	message := NewFeiShuMessage(webhok, secret)
	_, err := message.SendText("test", true, nil)
	if err != nil {
		t.Logf("%s", err.Error())
	}
}

func TestFeiShuMessage_SendMarkdown(t *testing.T) {
	webhok := os.Getenv("feishu_webhok")
	secret := os.Getenv("feishu_secret")
	message := NewFeiShuMessage(webhok, secret)
	_, err := message.SendMarkdown("测试", `
### test

> - test
> - test2
`, false, nil)
	if err != nil {
		t.Logf("%s", err.Error())
	}

}
