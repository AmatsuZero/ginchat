package models

import (
	"time"

	"github.com/amatsuzero/ginchat/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIP      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func init() {
	utils.DB.AutoMigrate(&UserBasic{})
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	var data []*UserBasic
	utils.DB.Find(&data)
	return data
}

func CreateUser(usr UserBasic) *gorm.DB {
	return utils.DB.Create(&usr)
}

func DeleteUser(usr UserBasic) *gorm.DB {
	return utils.DB.Delete(&usr)
}

func UpdateUser(usr UserBasic) *gorm.DB {
	return utils.DB.Model(&usr).Updates(usr)
}

func FindUserByName(name string) UserBasic {
	var another UserBasic
	utils.DB.Where("name = ?", name).First(&another)
	return another
}

func FindUserByPhone(phone string) UserBasic {
	var another UserBasic
	utils.DB.Where("phone = ?", phone).First(&another)
	return another
}

func FindUserByNameAndPassword(name, password string) UserBasic {
	usr := UserBasic{}
	utils.DB.Where("name = ? AND password = ?", name, password).First(&usr)
	return usr
}
