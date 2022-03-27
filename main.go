package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main() {
	//	创建数据库
	//	sql: CREATE DATABASE bubble;
	//	连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//绑定表模型
	dao.DB.AutoMigrate(&models.Todo{})
	//	在9090端口下启动
	err = routers.SetUpRouter().Run(":9090")
	if err != nil {
		fmt.Println("run failed , err :", err)
		return
	}
}
