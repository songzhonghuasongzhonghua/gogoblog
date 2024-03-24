package models

import (
	"github.com/songzhonghuasongzhonghua/gogoblog/models/ctype"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title         string         `gorm:"size:32" json:"title"`
	Abstract      string         `json:"abstract"` //文章简介
	Content       string         `json:"content"`
	LookCount     int            `json:"look_count"`    //文章浏览量
	CommentCount  int            `json:"comment_count"` //文章评论量
	DiggCount     int            `json:"digg_count"`    //点赞数
	CollectsCount int            `json:"collects_count"`
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"-"` //关联tag表
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`
	UserID        uint           `json:"user_id"`
	Category      string         `gorm:"size:20" json:"category"`
	Source        string         `json:"source"`
	Link          string         `json:"link"`
	BannerModel   BannerModel    `gorm:"foreignKey:BannerID" json:"-"` //关联文章封面表
	BannerID      uint           `json:"banner_id"`                    //封面id
	BannerPath    string         `json:"banner_path"`                  //封面路径
	NickName      string         `gorm:"size:32" json:"nick_name"`

	Tags ctype.Array `gorm:"type:string;size:64" json:"Tags"` //文章标签
}
