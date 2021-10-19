# Hello Leek API 👋

Leek API 是一个使用 gin + jwt 的接口项目，[前端代码在这里。](https://github.com/feilongjump/leek)

### 目录

|  目录  |  描述  |
| ----  | ----  |
| app  | 存放业务逻辑相关代码 |
| - http  | 存放处理 HTTP 请求相关的代码 |
| - - controllers  | 处理请求并返回响应 |
| - - middlewares  | 中间件 |
| - - requests  | 请求参数验证 |
| - - resources  | 响应的资源 |
| - models  | 模型 |
| bootstrap  | 存放程序初始化相关逻辑 |
| config  | 配置 |
| pkg  | 应用程序可以使用的代码库 |
| routes  | 路由 |

### 食用

* 自行下载，然后安装依赖包吧。😈
* 别说我不告诉你，复制 `.app_env.example.toml` 文件成 `.app_env.toml` 文件
  ```
  .
  .
  .
  [database]
    connection = "mysql"
    server = "localhost"
    port = 3306
    database = "example"
    username = "username"
    password = "password"
  .
  .
  .
  ```
  好好的增加调味料（我就随便加了几个给你，剩下的你自己看着加）。
* 接下来？可以吃了啊，应该不用我来喂你吧。😏

##### 用什么吃？中国人啊，肯定用筷子啊！🥢
[Gin](https://github.com/gin-gonic/gin) + [JWT](https://github.com/dgrijalva/jwt-go) 就这双筷子吧，还有什么要特殊说明的就到时候再加上吧。

不过这个 JWT 包目前好像没有维护了，转移到了[这里](https://github.com/golang-jwt/jwt) 。以后再看看要不要换吧，先用着，还没断。

##### 什么？还有什么菜？韭菜啊！
- [x] 代码生成
- [x] JWT 身份验证
- [ ] 第三方登录
- [ ] RBAC 权限管理
- [ ] 菜单管理

### TODO
* 简单的 CURD 代码生成做好了，但是还得好好的优化，容我学习学习设计模式啥的。

### 最后
要是在食用中出现了什么问题，一定要告诉我，我一定会去解决，解决不了，那我就把它倒了。😁

好好享用！🎉