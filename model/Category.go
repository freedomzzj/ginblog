package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ErrorCateNameUsed
	}
	return errmsg.SUCCESS
}

//  增加分类
func CreateCate(data *Category) int {
	//data.Password = Scrypt(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	var total int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

//编辑分类
func EditCate(id int, data *Category) int {
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&Category{}).Where("id = ?", id).Update(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCate(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
