package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

var version = "unknown"

func main() {

	app := cli.NewApp()
	app.Name = "dingtalk-robot plugin"
	app.Usage = "Sending message to DingTalk group by robot using WebHook"
	app.Version = version
	app.Action = run
	app.Copyright = "© 2021-now hb0730"
	app.Authors = []cli.Author{
		{
			Name:  "hb0730",
			Email: "huangbing0730@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "config.debug,debug",
			Usage:  "debug mode",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.StringFlag{
			Name:   "config.dingtalk.token,access_token,token",
			Usage:  "DingTalk webhook access token",
			EnvVar: "PLUGIN_ACCESS_TOKEN,PLUGIN_TOKEN",
		},
		cli.StringFlag{
			Name:   "config.dingtalk.secret,secret",
			Usage:  "DingTalk webhok secret",
			EnvVar: "PLUGIN_SECRET",
		},
		cli.BoolFlag{
			Name:   "config.message.at_all,at_all,atAll",
			Usage:  "at all in a message",
			EnvVar: "PLUGIN_AT_ALL",
		},
		cli.StringFlag{
			Name:   "config.message.at_mobiles,at_mobiles",
			Usage:  "at someone in a DingTalk group need this guy bind's mobile",
			EnvVar: "PLUGIN_AT_MOBILES",
		},
		cli.StringFlag{
			Name:   "config.message.title,message_title",
			Usage:  "",
			EnvVar: "PLUGINS_TITLE",
		},
		cli.StringFlag{
			Name:   "config.message.content",
			Usage:  "",
			EnvVar: "PLUGIN_CONTENT",
		},
		cli.StringFlag{
			Name:   "config.message.message_type,message_type",
			Usage:  "DingTalk message type:text, markdown",
			EnvVar: "PLUGIN_MESSAGE_TYPE",
		},
	}
	if _, err := os.Stat("/run/drone/env"); err == nil {
		godotenv.Overload("/run/drone/env")
	}

	if err := app.Run(os.Args); nil != err {
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	plugin := &Plugin{
		Debug: ctx.Bool("config.debug"),
		DingTalkConfig: DingTalkConfig{
			AccessToken: ctx.String("config.dingtalk.token"),
			Secret:      ctx.String("config.dingtalk.secret"),
		},
	}
	message := Message{
		AtAll:       ctx.Bool("config.message.at_all"),
		AtMobiles:   strings.Split(ctx.String("config.message.at_mobiles"), ","),
		MessageType: ctx.String("config.message.message_type"),
		Title:       ctx.String("config.message.title"),
		Content:     ctx.String("config.message.content"),
	}
	return plugin.Exec(message)
}
