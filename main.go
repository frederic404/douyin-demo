package main

import (
	"fmt"
	"os"

	"github.com/bytedance/douyin-demo/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	Init()
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}

func Init() {
	InitLogger()
}

func InitLogger() {
	// 设置为json格式的日志
	file, err := os.OpenFile("./gin_log.Log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("创建日志文件/打开日志文件失败")
		return
	}
	log.SetOutput(file)
	log.SetLevel(log.InfoLevel)
	return
}
