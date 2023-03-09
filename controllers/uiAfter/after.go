package uiAfter

import (
	"blog/models"
	"crypto/sha1"
	"encoding/hex"
	// "github.com/gin-contrib/sessions/cookie"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type After struct {
}

func (_ *After) Index(ctx *gin.Context) {

	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "login.html", nil)
		return
	}
	if ctx.Request.Method == "POST" {
		account := ctx.PostForm("account")
		password := ctx.PostForm("password")

		if account == "" || password == "" {
			ctx.JSON(200, gin.H{"result": 0, "msg": "您输入的账号或密码为空！"})
			return
		}
		sha1 := sha1.New()
		sha1.Write([]byte(password))
		//sha1.Sum(nil) 返回16进制的bety
		// hex.EncodeToString  16进制的bety 转字符串
		passSha1 := hex.EncodeToString(sha1.Sum(nil))
		MemberModel := new(models.Member)
		Member := MemberModel.GetFirst(account, passSha1)
		if Member.Mid == 0 {
			ctx.JSON(200, gin.H{"result": 0, "msg": "用户不存在或密码错误"})
			return
		}

		sha1.Write([]byte(string(Member.Mid)))

		fmt.Println(sha1.Sum(nil))
		ctx.JSON(200, gin.H{"result": 1, "msg": "账号验证成功"})
		return
	}

}
