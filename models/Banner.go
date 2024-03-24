package models

import "gorm.io/gorm"

type BannerModel struct {
	gorm.Model
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片hash值，用于判断是否为相同图片
	Name string `gorm:"size:32" json:"name"` //图片名称
}
