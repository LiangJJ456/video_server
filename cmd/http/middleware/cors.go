package middleware

import "github.com/gin-gonic/gin"

//解决跨域问题
func Cors()gin.HandlerFunc{
	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")
		context.Writer.Header().Set("Access-Control-Allow-Origin",origin)
		context.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		context.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type,Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if context.Request.Method == "OPTIONS" {
			context.Writer.Header().Set("Access-Control-Max-Age", "3600")
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}
