# Drone plugin notice

## example

```yaml
- name: Dingtalk Notice Failure
  pull: if-not-exists
  image: hb0730/drone-plugin-notice:1.0.2
  settings:
    debug: true
    webhook:
      from_secret: feishu-robot-webhok
    secret:
      from_secret: feishu-robot-secret
    robot_type: feishu
    content_type: markdown
    at_only_failure: true
    message_title: Drone 构建通知
    message_content: |
      ### 构建信息
      > - 应用名称: [DRONE_REPO_NAME]
      > - 构建结果: 预发布成功 ✅
      > - 构建发起: [CI_COMMIT_AUTHOR_NAME]
      > - 持续时间: [CUSTOM_BUILD_CONSUMING]s

      构建日志: [点击查看详情]([DRONE_BUILD_LINK])        
  when:
    status: 
      - success
      - failure
```

## 插件参数 plugin params
* `debug` : debug模式，打印`env`等信息
* `webhook`: 机器人的`webhook`地址
* `secret`: 机器人的`secret`
* `robot_type`: 机器人类型: `feishu`,`dingtalk` (忽略大小写)
* `content_type`: 内容的形式: `markdown`,`text` (忽略大小写)
* `at_all`: 是否`AT`所有人
* `at_only_failure`: 是否只在`failure`时`at`所有人,
* `message_title`: 标题,(目前只支持`mardkown`模式)
* `message_content` : 内容,支持占位符`[]` 替换，支持当前所有环境变量
* `custom_started` 开始时间环境变量,如:`DRONE_BUILD_STARTED`
* `custom_finished` 完成时间环境变量,如:`DRONE_BUILD_FINISHED`

---

* `CUSTOM_BUILD_CONSUMING` : 构建时间(秒)