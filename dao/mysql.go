package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//注意不是：=
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
		return
	}
	DB.DB().Ping()
	return
}

func Close() {
	DB.Close()
}
