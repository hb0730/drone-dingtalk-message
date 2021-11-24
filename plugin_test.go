package main

import (
	"testing"
)

func TestPlugin_Exec(t *testing.T) {
	plugin := &Plugin{
		NoticeConfig: NoticeConfig{
			AccessToken: "",
			Secret:      "",
		},
	}
	message := Message{
		AtAll:       false,
		AtMobiles:   nil,
		MessageType: "text",
		Content:     "test",
	}
	err := plugin.Exec(message)
	if err != nil {

		t.Logf("%s", err.Error())
	}
}
