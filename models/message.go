package models

import (
	"github.com/amatsuzero/ginchat/utils"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId         string // 发送者
	TargetId       string // 消息接受者
	Type           string // 消息类型 群聊 / 私聊 / 广播
	Media          int    // 消息嘞哦行 文字 / 图片 / 音频
	Content        string
	Pic, Url, Desc string
	Amount         int
}

func init() {
	utils.DB.AutoMigrate(&Message{})
}

func (table *Message) TableName() string {
	return "messge"
}
