package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func init() {
	initConfig()
	initMySQL()
	initRedis()
}

func initConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("config app inited...")
}

func initMySQL() {
	// 自定义日志模板 打印 SQL 语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,
		},
	)
	dsn := viper.GetString("mysql.dsn")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	fmt.Println("MySQL inited...")
}

func initRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
		Password:     viper.GetString("redis.password"),
	})

	fmt.Println("Redis inited...")
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到 redis
func Publish(ctx context.Context, channel, msg string) (err error) {
	err = Red.Publish(ctx, channel, msg).Err()
	fmt.Println("publish...", err.Error())
	return
}

// Subscribe 订阅 redis 消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	rmsg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe ...", rmsg.Payload)
	return rmsg.Payload, err
}
