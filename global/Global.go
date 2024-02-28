package global

import (
	"github.com/songzhonghuasongzhonghua/gogoblog/config"
	"gorm.io/gorm"
)

// 全局变量
var (
	Config *config.Config
	DB     *gorm.DB
)
