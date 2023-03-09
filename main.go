package main

import (
	"blog/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"os"
)

func main() {
	//加载配置
	cfg, err := ini.Load("./conf/api.ini")
	if err != nil {
		//fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()

	if mode == "develop" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册路由
	r := routers.Register()

	http_port := cfg.Section("").Key("http_port").String()
	//r := gin.Default()

	r.Run(http_port)
}
