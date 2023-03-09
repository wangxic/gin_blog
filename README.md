### 1. 环境配置
```sh
#使用环境是用 go mod 方式配置的
go env -w GOROOT=/bin/go #配置 go的环境地址
# go env -w GOPATH=/bin/go #go的项目地址 不需要了
go env -w GO111MODULE=on #设置导入包的方式 off ,auto ,on 三种
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/ #设置包的代理地址

#cd到项目目录
go mod init 项目名称 #初始化项目
```

### 2. 目录结构

项目根目录的下创建项目的结果

```shell
#目录结构
--conf        #配置
--controllers #控制器
--middlewares #中间件
--models      #模型
--libraries   #库
--routers     #路由
--static      #静态文件
--templates   #模板
--Views 	 #前端部分
main.go       #入口文件
README.md     #说明文件
```
### 3.编写入口文件程序
```go
package main

import "github.com/gin-gonic/gin"

func main()  {
    // 初始化一个http服务对象
    r := gin.Default()

    // 设置一个GET请求的路由，url: '/ping'， 控制器函数： 闭包
    r.GET("/ping", func(c *gin.Context) {
        // 通过请求上下文对象Context，返回json
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

// 监听，并在 localhost:8080上启动服务 相当于 r.Run(":8080")
    r.Run()
}
```

[gin 中文文档](https://gin-gonic.com/zh-cn/ "gin 中文")
