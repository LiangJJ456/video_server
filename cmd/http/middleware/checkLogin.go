package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"video_server/constants"
)

//检查用户是否登录，使用jwt技术
func CheckLogin()gin.HandlerFunc{
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if mapClaims := ParseToken(token); mapClaims != nil {
			context.Set("mapClaims",mapClaims)
			context.Next()
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
				"error" : "token 错误",
			})
			return
		}
	}
}

//创建token
func CreatToken()(string,error){
	//设置使用的算法
	token := jwt.New(jwt.SigningMethodHS256) //HS256算法
	//设置使用的playload
	claims := make(jwt.MapClaims)
	claims["user"] = "lzj"
	claims["exp"] = time.Now().Add(time.Hour*time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()//发行时间
	token.Claims = claims
	//设置签名，根据秘钥;生成token
	tokenString,err := token.SignedString([]byte(constants.SIGN_NAME_SCERET))

	if err != nil {
		log.Printf("err : %s\n",err.Error())
		return "",err
	}

	return tokenString,nil
}

//解析token
func ParseToken(tokenString string) jwt.MapClaims{
	token,err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.SIGN_NAME_SCERET),nil//返回密钥
	})

	if err != nil {
		log.Printf("the error is %s",err)
		return nil
	}
	var claims jwt.MapClaims
	var ok bool
	if claims,ok = token.Claims.(jwt.MapClaims);ok&&token.Valid{
		return claims
	}else{
		return nil
	}
}