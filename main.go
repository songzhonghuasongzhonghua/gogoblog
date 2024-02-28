package main

import (
	"fmt"
	"github.com/songzhonghuasongzhonghua/gogoblog/core"
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"github.com/songzhonghuasongzhonghua/gogoblog/routers"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化数据库
	global.DB = core.InitGorm()
	//初始化路由
	router := routers.InitRouter()
	fmt.Println(global.Config.System.Addr())

	router.Run(global.Config.System.Addr())
}
