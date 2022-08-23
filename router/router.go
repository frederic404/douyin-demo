package router

import (
	action2 "github.com/bytedance/douyin-demo/action"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./web/video")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed", action2.Feed)
	apiRouter.GET("/user/", action2.UserInfo)
	apiRouter.POST("/user/register/", action2.Register)
	apiRouter.POST("/user/login/", action2.Login)
	apiRouter.POST("/publish/action/", action2.Publish)
	apiRouter.GET("/publish/list/", action2.PublishList)

	// 	// extra apis - I
	// 	apiRouter.POST("/favorite/action/", action.FavoriteAction)
	// 	apiRouter.GET("/favorite/list/", action.FavoriteList)
	// 	apiRouter.POST("/comment/action/", action.CommentAction)
	// 	apiRouter.GET("/comment/list/", action.CommentList)
	//
	// 	// extra apis - II
	// 	apiRouter.POST("/relation/action/", action.RelationAction)
	// 	apiRouter.GET("/relation/follow/list/", action.FollowList)
	// 	apiRouter.GET("/relation/follower/list/", action.FollowerList)
}
