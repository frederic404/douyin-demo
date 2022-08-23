package util

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogrus() {
	// 设置为json格式的日志
	Log.Formatter = &logrus.JSONFormatter{}
	file, err := os.OpenFile("./gin_log.Log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("创建日志文件/打开日志文件失败")
		return
	}
	// 设置log默认文件输出
	Log.Out = file
	gin.SetMode(gin.ReleaseMode)
	// gin框架自己记录的日志也会输出
	gin.DefaultWriter = Log.Out
	// 设置日志级别
	Log.Level = logrus.InfoLevel
	return
}
