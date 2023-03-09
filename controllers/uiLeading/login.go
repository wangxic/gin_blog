package uiLeading

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
}

func (_ *Login) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", nil)
	return
}
