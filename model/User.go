package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2" label:"权限码"`
}

//查询用户是否存在
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

//  增加用户
func CreateUser(data *User) int {
	data.Password = Scrypt(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//编辑用户
func EditUser(id int, data *User) int {
	maps := make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&User{}).Where("id = ?", id).Update(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func Scrypt(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{13, 34, 15, 25, 67, 226, 97, 76}

	HashPW, err := scrypt.Key([]byte(password), salt, 1<<16, 8, 1, KeyLen)
	if err != nil {
		log.Fatalln(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if Scrypt(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrorUserNoPermission
	}
	return errmsg.SUCCESS
}
