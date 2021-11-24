package main

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

type Plugin struct {
	Debug        bool
	NoticeConfig NoticeConfig
}
type NoticeConfig struct {
	NoticeType  string
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
	var err error
	if plugin.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}
	if plugin.NoticeConfig.AccessToken == "" {
		return errors.New("missing  access token")
	}
	if message.Content == "" {
		return errors.New("missing Content")
	}
	notice, err := getSupportMessage(plugin.NoticeConfig.NoticeType, plugin.NoticeConfig.AccessToken, plugin.NoticeConfig.Secret)
	if err != nil {
		return err
	}
	content := plugin.regexp(message.Content)
	switch strings.ToLower(message.MessageType) {
	case "markdown":
		err = notice.SendMarkdown(message.Title, message.Content, message.AtAll, message.AtMobiles)
	case "text":
		err = notice.SendText(content, message.AtAll, message.AtMobiles)
	default:

		msg := "not support message type"
		err = errors.New(msg)
	}
	if err == nil {
		log.Println("send message success!")
	}
	return err
}

func (plugin *Plugin) regexp(content string) string {
	envs := plugin.getEnvs()
	// replace regex
	reg := regexp.MustCompile(`\[([^\[\]]*)]`)
	match := reg.FindAllStringSubmatch(content, -1)
	for _, m := range match {
		// from environment
		if envStr := os.Getenv(m[1]); envStr != "" {
			content = strings.ReplaceAll(content, m[0], envStr)
		}
		// check if the keyword is legal
		if _, ok := envs[m[1]]; ok {
			// replace keyword
			content = strings.ReplaceAll(content, m[0], envs[m[1]])
		}
	}
	return content
}

func (plugin *Plugin) getEnvs() map[string]string {
	envs := map[string]string{}
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}
		envs[parts[0]] = parts[1]
	}
	return envs
}
