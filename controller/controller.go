package controller

import (
	"GoWebCode/Bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url -->  controller -->      logic   --> model
请求 -->  控制器      -->      业务逻辑  -->模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//获得前端请求获取参数

	//1.获得title参数
	var todo models.Todo
	c.ShouldBind(&todo)

	//2.存入数据库中
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err}) //去掉了.Error()
	} else {
		c.JSON(http.StatusOK, todo)
	}

	//3.返回响应
}

func GetTodoList(c *gin.Context) {
	//查看todo表中的所有数据
	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATode(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
