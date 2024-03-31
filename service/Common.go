package service

import (
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"github.com/songzhonghuasongzhonghua/gogoblog/models"
)

type Options struct {
	models.PageInfo
	Debug bool
}

func CommonList[T any](model T, options Options) (list []T, total int64, err error) {

	offset := (options.Page - 1) * options.PageSize

	total = global.DB.Select("id").Find(&list).RowsAffected
	err = global.DB.Limit(options.PageSize).Offset(offset).Find(&list).Error
	return list, total, err

}
