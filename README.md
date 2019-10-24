## 打印消息路由日志

实现方案
```text
1：将消息推送给所有的客户端
    优点：实现简单，方便协作，客户端收到相同的消息。
    缺点：用在多个项目中，客户端收到多个项目的消息，消息彼此之间混淆，不方便排查错误
2：将消息推送给发起问答的客户端
    优点：消息彼此独立，方便排错，不受多个项目的影响
    缺点：实现复杂，独立调试，不能协作Bug

方案确认:选择方案1，侧重协同排错

开发语言: golang,javascript
```
---

使用方法：
```text
1. 吾来平台填写回调消息路由地址 http://console.wul.ai/bot/message_route
2. 打开console页面，查看日志  http://console.wul.ai/
```

<img src="https://github.com/shzy2012/static/blob/master/console.png?raw=true" width="600" height="400">


---

部署注意事项:单机部署，防止代理分发混乱
