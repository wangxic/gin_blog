package api

import (
	"blog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//文章相关的后台接口 包含 文章相关的类别 评论等等

type Article struct {
}

//文章相关的后台接口 包含 文章相关的类别 评论等等

//列表
func (_ *Article) Index(ctx *gin.Context) {
	like := ctx.PostForm("like")
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	if page == 0 {
		page = 1
	}
	ArticleModel := new(models.Articles)
	cent, total := ArticleModel.List(like, page)
	ctx.JSON(200, gin.H{"result": 1, "msg": "成功", "data": cent, "total": total})
	return

}

//编辑
func (_ *Article) Edit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	cent := ctx.PostForm("cent")
	title, _ := ctx.GetPostForm("title")
	keyword := ctx.PostForm("keyword")
	classify, _ := strconv.Atoi(ctx.PostForm("classify"))
	fmt.Println(ctx.Request)
	fmt.Println(cent)
	fmt.Println(title)
	fmt.Println(keyword)

	if classify == 0 || title == "" || keyword == "" || cent == "" {
		ctx.JSON(200, gin.H{"result": 9001, "msg": "参数错误"})
		return
	}
	ArticleModel := new(models.Articles)
	// 新增
	if id == 0 {

		res := ArticleModel.Insert(classify, title, cent, keyword)

		if res > 1 {
			ctx.JSON(200, gin.H{"result": 1, "msg": "新增成功"})
			return
		}
	} else {
		info := ArticleModel.First(uint(id))
		if info.ID == 0 {
			ctx.JSON(200, gin.H{"result": 9002, "msg": "非法操作"})
			return
		}
		editRes := ArticleModel.Edit(uint(id), classify, title, cent, keyword)
		if editRes < 1 {
			ctx.JSON(200, gin.H{"result": 6001, "msg": "更新失败"})
			return
		}
		ctx.JSON(200, gin.H{"result": 1, "msg": "编辑成功"})
		return
	}

}

//查看
func (_ *Article) Look(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))

	if id == 0 {
		ctx.JSON(200, gin.H{"result": 9001, "msg": "参数错误"})
		return
	}
	ArticleModel := new(models.Articles)
	info := ArticleModel.First(uint(id))
	if info.ID == 0 {
		ctx.JSON(200, gin.H{"result": 9002, "msg": "非法操作"})
		return
	}
	ctx.JSON(200, gin.H{"result": 1, "msg": "成功", "data": info})
	return
}

//删除
func (_ *Article) Del(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		ctx.JSON(200, gin.H{"result": 9001, "msg": "参数错误"})
		return
	}
	ArticleModel := new(models.Articles)
	info := ArticleModel.First(uint(id))
	if info.ID == 0 {
		ctx.JSON(200, gin.H{"result": 9002, "msg": "非法操作"})
		return
	}
	res := ArticleModel.Del(uint(id))
	if res > 0 {
		ctx.JSON(200, gin.H{"result": 1, "msg": "删除成功"})
		return
	}
	ctx.JSON(200, gin.H{"result": 0, "msg": "删除失败"})
	return
}

// 类别
func (_ *Article) Classify(ctx *gin.Context) {
	like := ctx.PostForm("like")
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	if page == 0 {
		page = 1
	}
	ClassifyM := new(models.Classifys)
	cent, total := ClassifyM.List(like, page)
	ctx.JSON(200, gin.H{"result": 1, "msg": "成功", "data": cent, "total": total})
	return
}

// 类别
func (_ *Article) ClassifyEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	pid, _ := strconv.Atoi(ctx.PostForm("pid"))
	name := ctx.PostForm("name")
	sort, _ := strconv.Atoi(ctx.PostForm("sort"))
	ClassifyM := new(models.Classifys)
	if id != 0 {
		res := ClassifyM.Update(id, pid, sort, name)
		if res > 0 {
			ctx.JSON(200, gin.H{"result": 1, "msg": "修改成功"})
			return
		}
		ctx.JSON(200, gin.H{"result": 0, "msg": "修改失败"})
		return
	}
	InsertId := ClassifyM.Insert(pid, sort, name)
	if InsertId > 0 {
		ctx.JSON(200, gin.H{"result": 1, "msg": "新增成功"})
		return
	}
	ctx.JSON(200, gin.H{"result": 1, "msg": "新增成功"})
	return

}

// 类别
func (_ *Article) ClassifyDel(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 || id < 0 {
		ctx.JSON(200, gin.H{"result": 9001, "msg": "非法请求！"})
		return
	}
	ClassifyM := new(models.Classifys)

	res := ClassifyM.Del(id)
	if res > 0 {
		ctx.JSON(200, gin.H{"result": 1, "msg": "成功！"})
		return

	}
	ctx.JSON(200, gin.H{"result": 0, "msg": "失败！"})
	return

}

// 类别
func (_ *Article) ClassifyLook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 || id < 0 {
		ctx.JSON(200, gin.H{"result": 9001, "msg": "非法请求！"})
		return
	}
	ClassifyM := new(models.Classifys)

	info := ClassifyM.GetRow(id)

	ctx.JSON(200, gin.H{"result": 1, "data": info, "msg": "非法请求！"})

}
