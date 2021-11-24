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

func TestPlugin_regexp(t *testing.T) {
	content := `### 构建信息
        > - 应用名称: [DRONE_REPO_NAME]
        > - 构建结果: 预发布成功 ✅
        > - 构建发起: [CI_COMMIT_AUTHOR_NAME]
        > - 持续时间: [CUSTOM_BUILD_TIME]s
        构建日志: [点击查看详情]([DRONE_BUILD_LINK])`
	plugin := &Plugin{Debug: true}
	content = plugin.regexp(content)
	t.Log(content)
}
