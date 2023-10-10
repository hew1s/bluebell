package main

import (
	"bulebell/controller"
	"bulebell/dao/mysql"
	"bulebell/dao/redis"
	"bulebell/logger"
	"bulebell/pkg/snowflake"
	"bulebell/router"
	"bulebell/settings"
	"fmt"
	"go.uber.org/zap"
)

// Go web 项目较为通用的脚手架模板

func main() {
	// 1.加载配置
	if err := settings.Init();err !=nil{
		fmt.Println("init settings failed err:",err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig,settings.Conf.Mode);err != nil{
		fmt.Println("init logger failed err:",err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	// 3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig);err != nil{
		fmt.Println("init mysql failed err:",err)
		return
	}
	defer mysql.Close()
	// 4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig);err != nil{
		fmt.Println("init redis failed err:",err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(settings.Conf.StartTime,settings.Conf.MachineID);err != nil{
		fmt.Println("init snowflake failed err:",err)
		return
	}
	// 初始化gin框架中内置的校验器使用的翻译器
	if err := controller.InitTrans("zh");err != nil{
		fmt.Println("init validator trans failed,err:",err)
		return
	}
	// 5.注册路由
	r := router.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf("127.0.0.1:%d",settings.Conf.Port))
	if err != nil {
		fmt.Println("run server failed",err)
		return
	}
	// 6.启动服务（优雅关机）
}