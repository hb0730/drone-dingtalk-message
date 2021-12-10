# Drone plugin notice

## example

```yaml
- name: Dingtalk Notice Failure
  pull: if-not-exists
  image: hb0730/drone-plugin-notice:1.0.2
  settings:
    debug: true
    notice_web_hok:
      from_secret: feishu-robot-webhok
    notice_secret:
      from_secret: feishu-robot-secret
    notice_type: feishu
    message_type: markdown
    message_at_all: true
    message_title: Drone 构建通知
    message_content: |
      ### 构建信息
      > - 应用名称: [DRONE_REPO_NAME]
      > - 构建结果: 预发布成功 ✅
      > - 构建发起: [CI_COMMIT_AUTHOR_NAME]
      > - 持续时间: [CUSTOM_BUILD_CONSUMING]s

      构建日志: [点击查看详情]([DRONE_BUILD_LINK])        
  when:
    status: success
```

## 插件参数 plugin params

* `notice_access_token` (required) : 自定义机器人的 `webhok`
* `notice_type` (required) : 机器人类型: `dingtalk`,`feishu`
* `message_type` (required) : 消息类型: `text`,`markdown`
* `notice_secret` : 如果设置了`加签` , 可以把你的加签密钥填入此项完成加签操作。
* `message_O_F_AT` : 只有是否是 `AT`所有人
* `message_at_all` : 是否`At`所有人
* `message_at_mobiles` : 你需要@的群成员的手机号，多个时用英文逗号(`,`)分隔 , 目前只支持 `dingtalk`
* `message_title` : 标题,只支持`markdown`
* `message_content` : 内容,支持占位符`[]` 替换，支持当前所有环境变量
* `debug` : debug模式，打印`env`等信息
* `custom_started` 开始时间环境变量,如:`DRONE_BUILD_STARTED`
* `custom_finished` 完成时间环境变量,如:`DRONE_BUILD_FINISHED`

---

* `CUSTOM_BUILD_CONSUMING` : 构建时间(秒)