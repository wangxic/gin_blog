package middlewares

import (
	"blog/libraries/redis"
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(ctx *gin.Context) {
	Token := ctx.Request.Header.Get("Token")
	if Token == "" {
		ctx.JSON(200, gin.H{"result": 90001, "msg": "非法请求！"})
		ctx.Abort()
		return

	}
	UserInfo := redis.Get(Token)
	if UserInfo == "" {
		ctx.JSON(200, gin.H{"result": 90001, "msg": "非法请求！"})
		ctx.Abort()
		return

	}
	//更新状态
	redis.Set(Token, UserInfo, 1500)
	ctx.Next()
}
