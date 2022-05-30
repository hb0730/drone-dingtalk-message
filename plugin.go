package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Build struct {
	Status string
}
type Plugin struct {
	Debug        bool
	Build        Build
	NoticeConfig NoticeConfig
	Custom       Custom
}
type NoticeConfig struct {
	NoticeType string
	WebHok     string
	Secret     string
}
type Custom struct {
	Consuming Consuming
}

// Consuming custom consuming env
type Consuming struct {
	StartedEnv  string
	FinishedEnv string
}

type Message struct {
	OnlyFailureAt bool
	AtAll         bool
	AtMobiles     []string
	Title         string
	Content       string
	MessageType   string
}

func (plugin *Plugin) Exec(message Message) error {
	log.Println("start  message ...")
	var err error
	if plugin.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}
	if plugin.NoticeConfig.WebHok == "" {
		return errors.New("missing  webhook url")
	}
	if message.Content == "" {
		return errors.New("missing content")
	}
	notice, err := getSupportMessage(plugin.NoticeConfig.NoticeType, plugin.NoticeConfig.WebHok, plugin.NoticeConfig.Secret)
	if err != nil {
		return err
	}
	content := plugin.regexp(message.Content)
	if plugin.Debug {
		log.Printf("content: %s \r\n", content)
	}
	return plugin.send(message, content, notice)
}

func (plugin *Plugin) send(message Message, content string, notice IMessage) error {
	//var resMessage string
	var err error
	switch strings.ToLower(message.MessageType) {
	case "markdown":
		if message.OnlyFailureAt && plugin.Build.Status == "failure" {
			resMessage, err = notice.SendMarkdown(message.Title, content, true, message.AtMobiles)
		} else {
			resMessage, err = notice.SendMarkdown(message.Title, content, message.AtAll, message.AtMobiles)
		}
	case "text":
		if message.OnlyFailureAt && plugin.Build.Status == "failure" {
			resMessage, err = notice.SendText(content, true, message.AtMobiles)
		} else {
			resMessage, err = notice.SendText(content, message.AtAll, message.AtMobiles)
		}
	default:
		msg := "not support message type"
		err = errors.New(msg)
	}
	if err == nil {
		if plugin.Debug {
			log.Printf("response content: %s \r\n", resMessage)
		}
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
		if plugin.Debug {
			log.Printf("env str: %s = ", m[0])
		}
		// from environment
		if envStr := os.Getenv(m[1]); envStr != "" {
			if plugin.Debug {
				log.Printf(" %s\n", envStr)
			}
			content = strings.ReplaceAll(content, m[0], envStr)
		}
		// check if the keyword is legal
		if _, ok := envs[m[1]]; ok {
			if plugin.Debug {
				log.Printf(" %s\n", envs[m[1]])
			}
			// replace keyword
			content = strings.ReplaceAll(content, m[0], envs[m[1]])
		}
	}
	return content
}

func (plugin *Plugin) getEnvs() map[string]string {
	envs := map[string]string{}
	//CUSTOM_BUILD_CONSUMING
	finishedEnv := plugin.Custom.Consuming.FinishedEnv
	startedEnv := plugin.Custom.Consuming.StartedEnv
	var finishedVar, startedVar string
	var consuming, finishedAt, startedAt uint64
	if finishedEnv != "" && startedEnv != "" {
		finishedVar = os.Getenv(finishedEnv)
		startedVar = os.Getenv(startedEnv)
	} else {
		finishedVar = os.Getenv("DRONE_BUILD_FINISHED")
		startedVar = os.Getenv("DRONE_BUILD_STARTED")
	}
	if plugin.Debug {
		log.Printf("BUILD_FINISHED: %s , BUILD_STARTED: %s \n", finishedVar, startedVar)
	}
	if finishedVar != "" && startedVar != "" {
		finishedAt, _ = strconv.ParseUint(finishedVar, 10, 64)
		startedAt, _ = strconv.ParseUint(startedVar, 10, 64)
		consuming = finishedAt - startedAt
	} else {
		consuming = 0
	}
	envs["CUSTOM_BUILD_CONSUMING"] = fmt.Sprintf("%v", consuming)
	return envs
}
