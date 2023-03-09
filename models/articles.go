package models

import (
	"time"
)

type Articles struct {
	ID         uint `gorm:"primary_key"`
	Classify   int
	Title      string
	Cent       string
	Keyword    string
	CreateTime time.Time
	UpdateTime time.Time
}

// 用id查询一条记录
func (article *Articles) First(id uint) *Articles {
	orm.Where(&Articles{ID: id}).First(article)
	return article
}

// 获取文章列表
func (_ *Articles) List(like string, page int) ([]Articles, int) {
	var pageSize = 20
	var articles []Articles
	DB := orm.Select("id,keyword,title,classify,create_time").Order("id desc")
	if like != "" {
		DB = DB.Where("keyword=?", like)
	}
	var total int = 0
	DB.Limit(pageSize).Offset((page - 1) * pageSize).Order("create_time DESC").Find(&articles)
	DBCount := orm.Model(&articles)
	if like != "" {
		DBCount = DBCount.Where("keyword=?", like)
	}
	DBCount.Count(&total)
	return articles, total
}

// 返回数据插入成功后的ID
func (_ *Articles) Insert(classify int, title, cent, keyword string) uint {
	now := time.Now()
	article := &Articles{Title: title, Classify: classify, Cent: cent, Keyword: keyword, CreateTime: now, UpdateTime: now}
	orm.Create(article)
	//fmt.Println(res)
	return article.ID
}

// 返回受影响行数
func (article *Articles) Edit(id uint, classify int, title, cent, keyword string) int64 {
	ret := article.First(id)
	// 查无结果 ret为空的Article
	if ret.ID == 0 {
		return 0
	}
	updateTime := time.Now()
	rowsAffected := orm.Model(ret).Updates(map[string]interface{}{"title": title, "classify": classify, "cent": cent, "Keyword": keyword, "update_time": updateTime}).RowsAffected
	return rowsAffected
}

// 返回受影响行数
func (article *Articles) Del(id uint) int64 {
	ret := article.First(id)
	if ret.ID == 0 {
		return 0
	}

	rowsAffected := orm.Delete(ret).RowsAffected
	return rowsAffected
}
