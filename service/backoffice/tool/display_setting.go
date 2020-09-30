package tool

import (
	model_setting "haii.or.th/api/server/model/setting"
	result "haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	//"encoding/json"
	//"log"
)

type Struct_SystemSetting_InputParam struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

type Struct_getSettingList struct {
	Result string                           `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_DamScaleColor `json:"data"`   // รายการตั้งค่า
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/setting_list
// @Summary			รายการตั้งค่าทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getSettingList successful operation
func (srv *HttpService) getSettingList(ctx service.RequestContext) error {
	resultData := model_setting.GetSystemSettingJSON("bof.Tool.DisplaySetting.SettingList")
	ctx.ReplyJSON(result.Result1(&resultData))
	return nil
}

type Struct_getDisplaySetting struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // การตั้งค่า
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/display_setting_json
// @Summary			รายละเอียดตั้งค่า
// @Description		รายละเอียดตั้งค่าจะแตกต่างกันไปข้ึนอยู้กับ รหัสรายการตั้งค่า
// @Method			GET
// @Parameter		code	query	string required:true enum:[Frontend.public.dam_scale_color,Frontend.public.rain_setting,Frontend.public.waterlevel_setting,Frontend.public.waterquality_setting,Frontend.public.wave_setting,Frontend.public.pre_rain_setting,Frontend.public.storm_setting,Frontend.public.warning_setting,Frontend.Analyst.Dam.OnLoadDamGraph]	รหัสรายการตั้งค่า
// @Produces		json
// @Response		200	Struct_getDisplaySetting successful operation
func (srv *HttpService) getDisplaySetting(ctx service.RequestContext) error {
	//Map parameters
	p := &Struct_SystemSetting_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	resultData := model_setting.GetSystemSetting(p.Code)
	ctx.ReplyJSON(result.Result1(&resultData))
	return nil
}

type Struct_putDisplaySetting struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // ตั้งค่า
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/display_setting_json
// @Summary			ตั้งค่า
// @Description		รายละเอียดตั้งค่าจะแตกต่างกันไปข้ึนอยู้กับ รหัสรายการตั้งค่า
// @Method			PUT
// @Consumes		json
// @Parameter		value	body	interface{}	 ตั้งค่า
// @Produces		json
// @Response		200	Struct_putDisplaySetting successful operation
func (srv *HttpService) putDisplaySetting(ctx service.RequestContext) error {
	//Map parameters
	p := &Struct_SystemSetting_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Check Data
	settingData := model_setting.GetSystemSetting(p.Code)
	if settingData == "" {
		return errors.Repack(rest.NewError(422, "Unknown code '"+p.Code+"'", errors.New("Unknown code '"+p.Code+"'")))
	}

	//Set Data
	objMapSetting := make(map[string]interface{}, 0)
	objMapSetting[p.Code] = p.Value
	if err := model_setting.SetSystemSetting(ctx.GetUserID(), objMapSetting, false); err != nil {
		return errors.Repack(err)
	}

	ctx.ReplyJSON(result.Result1(p))
	return nil
}
