package main

import (
	"github.com/gin-gonic/gin"
	"video_server/cmd/http/middleware"
	"video_server/service/stream"
)

func Init(r *gin.Engine){
	r.Use(middleware.Cors())
	//流数据部分
	streamGroup := r.Group("/video")
	streamGroup.Use(middleware.ConnLimiterHandler())
	{
		streamGroup.GET("/get/:vid",stream.GetVideo)//查看视频
	}

	//用户数据部分
	userGroup := r.Group("/user")
	userGroup.Use(middleware.CheckLogin()) //登录前需要验证
	{
		userGroup.GET("")
	}
}
