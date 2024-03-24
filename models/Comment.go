package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentId    *uint           `json:"parent_comment_id"`
	Content            string          `gorm:"size:256" json:"content"`
	DiggCount          int             `gorm:"default:0;size:8" json:"digg_count"` //点赞数
	CommentCount       int             `gorm:"default:0;size:8" json:"comment_count"`
	ArticleModel       ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID          uint            `json:"article_id"`
	UserModel          UserModel       `gorm:"foreignKey:UserID" json:"-"`
	UserID             uint            `json:"user_id"`
}
