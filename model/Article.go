package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Decs    string `gorm:"type:varchar(200)" json:"decs"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//  增加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// TODO 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Categor").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ErrorArticleNotExist
	}
	return article, errmsg.SUCCESS
}

// TODO 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

//编辑文章
func EditArticle(id int, data *Article) int {
	maps := make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Decs
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&Article{}).Where("id = ?", id).Update(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的所有文章
func GetCategoryArticle(id, pageSize, pageNum int) ([]Article, int) {
	var catArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&catArtList).Error
	if err != nil {
		return nil, errmsg.ErrorCateNotExist
	}
	return catArtList, errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int) int {
	err := db.Where("id = ?", id).Delete(&Article{}).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
