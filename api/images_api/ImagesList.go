package images_api

import (
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogoblog/models"
	"github.com/songzhonghuasongzhonghua/gogoblog/service"
	"github.com/songzhonghuasongzhonghua/gogoblog/utils"
)

func (ImagesApi) ImagesListView(c *gin.Context) {
	var imagesQuery models.PageInfo
	err := c.ShouldBindQuery(&imagesQuery)
	if err != nil {
		utils.Failed(c, gin.H{
			"msg": "参数不正确",
		})
		return
	}

	list, total, err := service.CommonList(&models.BannerModel{}, service.Options{
		PageInfo: imagesQuery,
		Debug:    true,
	})

	utils.Success(c, gin.H{
		"total": total,
		"list":  list,
	})
}
