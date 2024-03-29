<h1 align="center">
    <br>
    <a href="https://howio.world">
        <img width="80" src="https://github.com/feilongjump/howio.world/blob/main/src/assets/logo.png?raw=1" alt="HowIO" />
    </a>
    <br>
</h1>

## [HowIO](https://howio.world) API

个人站点的 API，用 [Go](https://go.dev) 语言进行开发的接口，前端页面放在[这里](https://github.com/feilongjump/howio.world)了

### 目录结构

```
.
├── app             项目应用文件
│  └── http             HTTP 文件
│    └── controllers        Controller
│    └── middlewares        中间件
│    └── requests           请求参数校验文件
│  └── models           Model
├── bootstrap       程序初始化的代码
│  └── config.go        初始化配置
│  └── ...              ...
├── route           路由文件
│  └── api.go           API 路由
├── internal        内部引用的 Go 包
│  └── config           环境配置包
│  └── ...              ...
├── .air.toml       自动重载应用配置文件
├── .gitignore      git 忽略文件
├── env.*.toml      项目环境配置文件
├── go.mod          Go 模块文件
├── go.sum          Go 模块的依赖版本文件
├── main.go         项目应用入口文件
└── README.md       项目手册
```

不一定会提前写好有什么，但有的东西就会补充上去。

- [x] [Gin](https://github.com/gin-gonic/gin)，Web 框架
- [x] [GORM](https://github.com/go-gorm/gorm)， ORM 操作库
- [x] [Viper](https://github.com/spf13/viper)，处理应用程序配置
- [x] [JWT](https://github.com/golang-jwt/jwt)，JSON Web Token
- [ ] 待补充