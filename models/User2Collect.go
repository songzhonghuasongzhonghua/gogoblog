package models

import "time"

// User2CollectModel 自定义第三张表 记录用户在什么时候收藏了哪些文章
type User2CollectModel struct {
	UserID       uint         `gorm:"primaryKey" json:"user_id"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID    uint         `json:"article_id"`
	CreatedAt    time.Time
}
