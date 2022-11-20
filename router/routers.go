package routers

import (
	"GoWebCode/Bubble/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//静态文件位置
	err := router.Static("/static", "static")
	if err != nil {
		fmt.Println(err)
	}
	//模板文件位置
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandler)

	v1Group := router.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return router
}
