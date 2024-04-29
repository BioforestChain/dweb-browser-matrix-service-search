# support-go




## 运行环境

- 开发语言: go1.18+
- 开发框架: gin
- RPC框架: grpc
- 数据库: mysql, redis
- 消息队列: 
- 代码管理工具：

## github

### 地址

[https://github.com/JasonMetal/submodule-support-go](https://github.com/JasonMetal/submodule-support-go)

### 分支说明

- master：正式环境
- test：测试环境, 主要提供给测试同学测试使用
- bvt：预发环境, 正式环境数据库一致, 预发环境测试通过后才可以发布线上环境

## 目录结构

``` 
.
├── Dockerfile          
├── LICENSE
├── README.md
├── app                 # 应用代码
│   ├── entity            # 各种结构体定义
│   ├── grpc              # 提供GRPC服务
│   ├── http              # 提供HTTP服务
│   ├── logic             # 逻辑层
│   ├── models            # 数据模型层，数据表实体类
│   └── services          # 服务层
├── bootstrap          
│   ├── app.go
│   ├── config.go
│   ├── database.go
│   └── redis.go 
├── config              # 各个环境配置
│   ├── bvt
│   ├── local
│   ├── online
│   └── test
├── go-build.sh         # 用于ci构建
├── go.mod
├── go.sum
├── helpers             # 常用类库封装
│   ├── config
│   ├── env
│   ├── errors
│   ├── logger
│   └── strings
├── http-server
├── main.go             # 入口文件
└── routes              # 路由文件夹
    ├── base.go
    ├── other.go
    └── prize.go    
    
```

## 本地环境配置

1. 拉取代码
    ```shell
    git clone git@github.com:JasonMetal/submodule-support-go.git
    ```


2. 启动

   ```shell
   go run main.go -e=local # 本地可以省略local
   ```

## git-子模块
 - 无

## gitlab地址
- 无
 

## 相关文档

- [gin框架](https://github.com/gin-gonic/gin)
- [grpc文档](https://grpc.io/docs/)
- [git子模块文档](https://git-scm.com/book/zh/v2/Git-%E5%B7%A5%E5%85%B7-%E5%AD%90%E6%A8%A1%E5%9D%97)
