package routers

import "github.com/songzhonghuasongzhonghua/gogoblog/api"

func (r *RouterGroup) SettingsRouter() {
	settingGroup := api.ApiGroupApi.SettingsGroup

	r.GET("settings", settingGroup.SettingInfoView)
}
