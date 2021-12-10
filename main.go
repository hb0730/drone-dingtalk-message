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
	app.Name = "drone-plugin-notice"
	app.Usage = "Sending message to DingTalk/FeiShu group by robot using WebHook"
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
			Name:   "config.notice.webhok",
			Usage:  "token/FeiShu webhok",
			EnvVar: "PLUGIN_NOTICE_WEB_HOK",
		},
		cli.StringFlag{
			Name:   "config.notice.secret",
			Usage:  "DingTalk/FeiShu WebHook secret for generate sign",
			EnvVar: "PLUGIN_NOTICE_SECRET,PLUGIN_SECRET",
		},
		cli.StringFlag{
			Name:   "config.notice.type",
			Usage:  "Robot type:feishu,dingtalk",
			EnvVar: "PLUGIN_NOTICE_TYPE",
		},
		cli.StringFlag{
			Name:   "config.message.type",
			Usage:  "Robot message type: text,markdown",
			EnvVar: "PLUGIN_MESSAGE_TYPE",
		},
		cli.BoolFlag{
			Name:   "config.message.at_all",
			Usage:  "at all in a message",
			EnvVar: "PLUGIN_MESSAGE_AT_ALL,PLUGIN_MESSAGE_ATALL",
		},
		cli.StringFlag{
			Name:   "config.message.at.mobiles",
			Usage:  "at someone in a DingTalk group need this guy bind's mobile(FeiShu unsupported)",
			EnvVar: "PLUGIN_MESSAGE_AT_MOBILES",
		},
		cli.BoolFlag{
			Name:   "config.message.only_failure_at",
			Usage:  "at all only in failure ",
			EnvVar: "PLUGIN_MESSAGE_O_F_AT",
		},
		cli.StringFlag{
			Name:   "config.message.title",
			Usage:  "message title(markdown supported)",
			EnvVar: "PLUGIN_MESSAGE_TITLE",
		},
		cli.StringFlag{
			Name:   "config.message.content",
			Usage:  "message content(Support placeholder[])",
			EnvVar: "PLUGIN_MESSAGE_CONTENT",
		},
		cli.StringFlag{
			Name:   "custom.started,started",
			Usage:  "started custom env name, eg., BUILD_STARTED",
			EnvVar: "PLUGIN_CUSTOM_STARTED",
		},
		cli.StringFlag{
			Name:   "custom.finished,finished",
			Usage:  "finished custom env name, eg., BUILD_FINISHED",
			EnvVar: "PLUGIN_CUSTOM_FINISHED",
		},

		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
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
		Build: Build{
			Status: ctx.String("build.status"),
		},
		NoticeConfig: NoticeConfig{
			NoticeType: ctx.String("config.notice.type"),
			WebHok:     ctx.String("config.notice.webhok"),
			Secret:     ctx.String("config.notice.secret"),
		},
		Custom: Custom{
			Consuming: Consuming{
				StartedEnv:  ctx.String("custom.started"),
				FinishedEnv: ctx.String("custom.finished"),
			},
		},
	}
	message := Message{
		OnlyFailureAt: ctx.Bool("config.message.only_failure_at"),
		AtAll:         ctx.Bool("config.message.at_all"),
		AtMobiles:     strings.Split(ctx.String("config.message.at.mobiles"), ","),
		MessageType:   ctx.String("config.message.type"),
		Title:         ctx.String("config.message.title"),
		Content:       ctx.String("config.message.content"),
	}
	return plugin.Exec(message)
}
