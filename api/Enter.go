package api

import "github.com/songzhonghuasongzhonghua/gogoblog/api/settings_api"

type ApiGroup struct {
	SettingsGroup settings_api.SettingsGroup
}

var ApiGroupApi = new(ApiGroup)
