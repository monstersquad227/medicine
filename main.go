package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"medicine/config"
	"medicine/middleware"
	"medicine/repository"
	v1 "medicine/routes/v1"
)

func main() {
	// 配置加载
	config.LoadConfig()

	// 初始化数据库连接
	repository.InitMysql()

	// 中间件配置
	Application := gin.Default()
	Application.Use(middleware.Cors())
	Application.Use(middleware.Jwt())

	// 路由注册
	v1.UserRegister(Application.Group(config.GlobalConfig.Application.Name))
	v1.CourseRegister(Application.Group(config.GlobalConfig.Application.Name))
	v1.PlanRegister(Application.Group(config.GlobalConfig.Application.Name))
	v1.RecordRegister(Application.Group(config.GlobalConfig.Application.Name))

	// 项目启动
	err := Application.Run(":" + config.GlobalConfig.Application.Port)
	if err != nil {
		log.Fatal(config.GlobalConfig.Application.Name+" 启动失败 : ", err)
	}
}
