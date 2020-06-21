package main

import (
	"github.com/gin-gonic/gin"
	"video_server/service/stream"
)

func Init(r *gin.Engine){
	//流数据部分
	streamGroup := r.Group("/video")
	{
		streamGroup.GET("/get/:vid",stream.GetVideo)//查看视频
	}
}
