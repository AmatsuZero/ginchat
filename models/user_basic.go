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
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
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
