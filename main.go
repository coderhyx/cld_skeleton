package main

import (
	"cld/bootstrap"
	"cld/controller"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//初始化数据库
	err := bootstrap.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	//初始化路由
	router := httprouter.New()

	//注册路由
	router.POST("/goods/add", controller.GoodsAdd)
	router.POST("/goods/list", controller.GoodsList)

	//启动并监听http服务
	http.ListenAndServe(":9002", router)
}
