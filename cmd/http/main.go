package main

import (
	"github.com/gin-gonic/gin"
	"video_server/cmd/http/middleware"
)

func init(){
	//cl
	middleware.CLimiter = middleware.GetConnLimiter()
}
func main(){
	r := gin.Default()
	Init(r)
	r.Run(":80")
}
