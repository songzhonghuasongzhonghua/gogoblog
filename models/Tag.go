package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	Title        string       `gorm:"size:20" json:"title"` //文章标签
	ArticleModel ArticleModel `gorm:"many2many:article_tag" json:"-"`
}
