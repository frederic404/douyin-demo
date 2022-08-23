package main

import (
	"github.com/bytedance/douyin-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
