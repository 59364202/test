package setting

import (
	model_setting "haii.or.th/api/server/model/setting"
)

func GetCronSetting(){
	model_setting.GetSystemSetting("")
}