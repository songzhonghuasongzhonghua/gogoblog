package models

import (
	"github.com/songzhonghuasongzhonghua/gogoblog/models/ctype"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	NickName   string           `gorm:"size:32" json:"nick_name"`
	UserName   string           `gorm:"size:32" json:"user_name"`
	Password   string           `gorm:"size:128" json:"password"`
	Avatar     string           `gorm:"size:256" json:"avatar"`
	Email      string           `gorm:"size:64" json:"email"`
	Tel        string           `gorm:"size:16" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr"`
	Token      string           `gorm:"size:64" json:"token"` //其他平台的唯一id
	IP         string           `gorm:"size:20" json:"ip"`
	Role       ctype.Role       `gorm:"size:4;default:2" json:"role"`
	SignStatus ctype.SignStatus `gorm:"size:4" json:"sign_status"`
}
