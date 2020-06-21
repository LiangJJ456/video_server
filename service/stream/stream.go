package stream

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"log"
	"video_server/constants"
)

func GetVideo(c *gin.Context){
	vid := c.Param("vid")
	vl := constants.VIDEO_DIR+vid+".mp4"
	video,err := os.Open(vl)
	defer video.Close()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error" : "内部错误",
			})
	}

	c.Writer.Header().Set("Content-Type","video/mp4")

	http.ServeContent(c.Writer,c.Request,"",time.Now(),video)//返回视频

}
