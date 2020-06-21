package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

//由于上传和请求流数据需要长连接，如果请求一多，会导致服务器挂掉，所以我们需要
//对连接数进行控制，这里我们用到了bucket token算法

var (
	CLimiter *ConnLimiter
	once sync.Once
)
const cc = 500 //容量
type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter{
	return &ConnLimiter{
		concurrentConn:cc,
		bucket: make(chan int ,cc),
	}
}

func (cl *ConnLimiter)getConn()bool{
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitaon.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter)releaseConn(){
	c := <- cl.bucket
	log.Printf("New connction coming %d",c)
}
//单例，全局唯一
func GetConnLimiter()*ConnLimiter{
	once.Do(func() {
		CLimiter = NewConnLimiter(cc)
	})
	return CLimiter
}
func ConnLimiterHandler() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer CLimiter.releaseConn()
		if CLimiter.getConn() {
			c.Next()
		}else {
			c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
				"Error":"the request is too many",
			})
		}
	}
}
