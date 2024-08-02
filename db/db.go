package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo/models"
)

var Db *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(localhost:3307)/local_todo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결 실패")
	}

	err = Db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("테이블 마이그레이션 실패")
	}

	fmt.Println("DB 연결 성공")
}
