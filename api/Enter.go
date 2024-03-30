package api

import (
	"github.com/songzhonghuasongzhonghua/gogoblog/api/images_api"
	"github.com/songzhonghuasongzhonghua/gogoblog/api/settings_api"
)

type ApiGroup struct {
	SettingsGroup settings_api.SettingsGroup
	ImagesGroup   images_api.ImagesApi
}

var ApiGroupApi = new(ApiGroup)
