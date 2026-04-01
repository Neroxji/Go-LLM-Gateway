package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	
	// 用 GORM 连上数据库
	var err error
	DB,err=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatalf("连接 MySQL 失败: %v",err)
	}
	log.Println("成功连接 MySQL 数据库!")

	// 拿到底层的对象
	sqlDB,err:=DB.DB()
	if err!=nil{
		log.Fatalf("获取底层 sqlDB 失败: %v", err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移
	err=DB.AutoMigrate(&User{},&RequestLog{},&Token{})
	if err!=nil{
		log.Fatalf("自动建表失败: %v", err)
	}
	log.Println("数据库表结构同步完成！")

}