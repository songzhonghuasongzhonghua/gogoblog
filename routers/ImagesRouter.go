package routers

import "github.com/songzhonghuasongzhonghua/gogoblog/api"

func (r *RouterGroup) ImagesRouter() {
	imagesGroup := api.ApiGroupApi.ImagesGroup
	r.POST("images", imagesGroup.ImagesUpload)
}
