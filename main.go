package main

import (
	"GoWebCode/Bubble/dao"
	"GoWebCode/Bubble/models"
	routers "GoWebCode/Bubble/router"
)

func main() {

	//创建数据库
	//sql:CREATE DATABASES bubble;

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	router := routers.SetupRouter()

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
