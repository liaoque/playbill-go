package main

import (
	"fmt"
	"palybill/app/api"
	"palybill/app/api/playbill"
	"palybill/routers"
)

func main() {
	// 1.创建路由
	routers.Include(api.Routers, playbill.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
