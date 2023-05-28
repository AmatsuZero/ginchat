package tests

import (
	"testing"

	"github.com/amatsuzero/ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestORM(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	t.Logf("My DB %v\n", db)
	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "春菜舞"
	db.Create(user)
	t.Log(db.First(user))

	db.Model(user).Update("Password", "1234")
}
