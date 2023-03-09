package routers

import (
	"blog/controllers/api"
	"blog/controllers/uiAfter"
	"blog/controllers/uiLeading"
	"blog/middlewares"
	"github.com/gin-gonic/gin"
)

// 设置跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "x-requested-with,Access-Control-Allow-Origin,Content-Type,Token")
		// 设置请求方式
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			//c.JSON(200, gin.H{"result": 0, "msg": "您输入的账号或密码为空！"})
			return
		}

		c.Next()
	}
}

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	//初始化session
	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))
	r.Use(Cors())
	//加载controllers 包中的Ui文件

	// 加载中间见
	// 前端部分
	UiLeading := new(uiLeading.Login)
	UiLeadingLogin := r.Group("/")
	{
		UiLeadingLogin.GET("/", UiLeading.Index)
		UiLeadingLogin.GET("/index", UiLeading.Index)
	}

	//后端部分
	uiAfter := new(uiAfter.After)
	uiafter := r.Group("/after")
	{
		uiafter.GET("/after", uiAfter.Index)
		uiafter.GET("/after/index", uiAfter.Index)
	}

	//文章
	Articled := new(api.Article)
	apiArticle := r.Group("/api/article", middlewares.InitMiddlewares)
	{
		apiArticle.POST("", Articled.Index)     //列表
		apiArticle.POST("/Edit", Articled.Edit) //更新
		apiArticle.POST("/look", Articled.Look) //单条
		apiArticle.POST("/del", Articled.Del)   //删除

		apiArticle.POST("/Classify", Articled.Classify)         //列表
		apiArticle.POST("/ClassifyEdit", Articled.ClassifyEdit) //更新
		apiArticle.POST("/ClassifyDel", Articled.ClassifyDel)   //单条
		apiArticle.POST("/ClassifyLook", Articled.ClassifyLook) //删除
	}

	//
	//Ui := new(api.Ui)
	//// 分组路由
	//v1 := r.Group("/ui", middlewares.InitMiddlewares)
	//{
	//	v1.GET("/", Ui.Index)
	//	v1.GET("/index", Ui.Index)
	//
	//	// v1.GET("/article/create", articles.Create)
	//	// v1.GET("/article/edit/:id", articles.Edit)
	//	// v1.GET("/article/del/:id", articles.Del)
	//	// v1.POST("/article/store", articles.Store)
	//}
	ApiLogin := new(api.Login)
	v2 := r.Group("/login")
	{

		v2.POST("", ApiLogin.Index)

	}
	return r
}
