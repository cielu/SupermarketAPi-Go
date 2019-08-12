package main

import (
	"github.com/gin-gonic/gin"
	"supermarket/app/models"
	"supermarket/library/setting"
	"supermarket/routes"
)

func init() {
	// 初始化配置
	setting.Initialized()
	// 初始化数据库连接
	models.Initialized()
}

func main() {
	defer models.CloseDb()
	// 设置app启动环境
	gin.SetMode(setting.AppCfg.AppMode)
	// 初始化 gin
	router := routes.InitRouter()
	// run
	router.Run(":8000")
}
