package api

import (
	"blog/libraries/redis"
	"blog/models"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Login struct {
}

// 登录返回数据
type Cent struct {
	Mid     uint
	Mobile  string
	Ctime   time.Time
	Account string
	Nick    string
	SignKey string
}

func (_ *Login) Index(ctx *gin.Context) {

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
		//生成简单的签名key
		sha1.Write([]byte("goLang" + string(Member.Mid)))
		cent := Cent{}
		cent.Mid = Member.Mid
		cent.Ctime = Member.Ctime
		cent.Mobile = Member.Mobile
		cent.Nick = Member.Nick
		cent.Account = Member.Account
		cent.SignKey = hex.EncodeToString(sha1.Sum(nil))
		// 结构体转成bety
		b, _ := json.Marshal(cent)
		// 存入签名key
		redis.Set(cent.SignKey, string(b), 1500)
		ctx.JSON(200, gin.H{"result": 1, "msg": "账号验证成功", "data": &cent})
		return
	}
}
