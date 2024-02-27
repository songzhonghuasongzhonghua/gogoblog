package core

import (
	"fmt"
	"github.com/songzhonghuasongzhonghua/gogoblog/config"
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func InitConfig() {
	var path = "settings.yaml"
	conf := new(config.Config)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("文件读取失败")
		return
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		fmt.Println("文件解析失败")
		return
	}
	global.Config = conf
}
