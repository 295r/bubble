package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	//	寻找gin框架模板文件引用的静态文件
	r.Static("/static", "./static")
	//	寻找模板文件
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandler)

	//	v1
	v1Group := r.Group("/v1")
	{
		//	待办事项
		//	添加
		v1Group.POST("/todo", controller.CreateTodo)
		//	查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodo)
		//	修改某一待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//	删除某一待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
