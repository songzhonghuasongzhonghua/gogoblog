package main

import (
	"fmt"
	"github.com/songzhonghuasongzhonghua/gogoblog/core"
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
)

func main() {
	fmt.Println("hello  mo")
	//读取配置文件
	core.InitConfig()
	//初始化数据库
	global.DB = core.InitGorm()
}
