package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	// 读取json文件
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("读文件错误:", err)
		return
	}

	// 初始化keypool
	for i := 0; i < len(config.Providers); i++ {
		config.Providers[i].Pool = NewKeyPool(config.Providers[i].Keys)
	}

	// 启动并连接数据库
	InitDB(config.DSN)

	// 启动日志系统
	for i := 0; i < 5; i++ {
		go logworker()
	}

	// 注册gin引擎
	r := gin.Default()

	// 挂载中间件
	r.Use(CorsMiddleware())
	r.Use(AuthMiddleware())

	// 注册路由
	r.POST("/api/chat", apiChatHandler(config))

	// 启动服务
	go func() {
		log.Println("服务器启动在 :8080...")
		r.Run(":8080")
	}()

	// 等关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("检测到退出信号，正在关闭服务...")
	close(LogChan)
}
