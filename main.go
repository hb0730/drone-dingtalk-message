package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

var version = "unknown"

func main() {
	app := &cli.App{
		Name:      "drone-plugin-notice",
		Usage:     "use webhook to notify build status",
		Version:   version,
		Action:    run,
		Copyright: "© 2021-now hb0730",
		Authors: []*cli.Author{
			{
				Name:  "hb0730",
				Email: "huangbing0730@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "config.debug",
				Aliases: []string{"debug"},
				Usage:   "debug mode",
				EnvVars: []string{"PLUGIN_DEBUG", "DEBUG"},
			},
			&cli.StringFlag{
				Name:    "config.notice.webhook",
				Aliases: []string{"webhook"},
				Usage:   "robot webhook url",
				EnvVars: []string{"PLUGIN_WEBHOOK", "PLUGIN_URL", "URL", "WEBHOOK"},
			},
			&cli.StringFlag{
				Name:    "config.notice.secret",
				Aliases: []string{"secret"},
				Usage:   "robot secret",
				EnvVars: []string{"PLUGIN_NOTICE_SECRET", "PLUGIN_SECRET", "SECRET"},
			},
			&cli.StringFlag{
				Name:    "config.notice.robotType",
				Aliases: []string{"robot_type"},
				Usage:   "types supported by the library: feishu/dingtalk",
				EnvVars: []string{"PLUGIN_NOTICE_ROBOT_TYPE", "PLUGIN_ROBOT_TYPE", "ROBOT_TYPE"},
			},
			&cli.StringFlag{
				Name:    "config.message.type",
				Aliases: []string{"message_type"},
				Usage:   "send content type: markdown/text",
				EnvVars: []string{"PLUGIN_MESSAGE_TYPE", "PLUGIN_CONTENT_TYPE", "MESSAGE_TYPE"},
			},
			&cli.BoolFlag{
				Name:    "config.message.at_all",
				Aliases: []string{"at_all"},
				Usage:   "at owner: true/false",
				EnvVars: []string{"PLUGIN_MESSAGE_AT_ALL", "PLUGIN_MESSAGE_ATALL", "PLUGIN_AT_ALL", "AT_ALL"},
			},
			//暂时移除
			//&cli.StringFlag{
			//	//	Name:   "config.message.at.mobiles",
			//	//	Usage:  "at someone in a DingTalk group need this guy bind's mobile(FeiShu unsupported)",
			//	//	EnvVar: "PLUGIN_MESSAGE_AT_MOBILES",
			//},
			&cli.BoolFlag{
				Name:    "config.message.only_failure_at",
				Aliases: []string{"only_failure_at"},
				Usage:   "at all only in failure ",
				EnvVars: []string{"PLUGIN_MESSAGE_AT_ONLY_FAILURE", "PLUGIN_AT_ONLY_FAILURE", "AT_ONLY_FAILURE"},
			},
			&cli.StringFlag{
				Name:    "config.message.title",
				Aliases: []string{"title"},
				Usage:   "message title(markdown supported)",
				EnvVars: []string{"PLUGIN_MESSAGE_TITLE", "MESSAGE_TITLE", "TITLE"},
			},
			&cli.StringFlag{
				Name:    "config.message.content",
				Aliases: []string{"content"},
				Usage:   "message content(Support placeholder[])",
				EnvVars: []string{"PLUGIN_MESSAGE_CONTENT", "MESSAGE_CONTENT"},
			},
			&cli.StringFlag{
				Name:    "custom.started",
				Aliases: []string{"started"},
				Usage:   "started custom env name, eg., BUILD_STARTED",
				EnvVars: []string{"PLUGIN_CUSTOM_STARTED"},
			},
			&cli.StringFlag{
				Name:    "custom.finished",
				Aliases: []string{"finished"},
				Usage:   "finished custom env name, eg., BUILD_FINISHED",
				EnvVars: []string{"PLUGIN_CUSTOM_FINISHED"},
			},
			&cli.StringFlag{
				Name:    "build.status",
				Usage:   "build status",
				Value:   "success",
				EnvVars: []string{"DRONE_BUILD_STATUS"},
			},
		},
	}
	if _, err := os.Stat("/run/drone/env"); err == nil {
		godotenv.Overload("/run/drone/env")
	}
	//
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
			RobotType: ctx.String("config.notice.robotType"),
			WebHok:    ctx.String("config.notice.webhook"),
			Secret:    ctx.String("config.notice.secret"),
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
