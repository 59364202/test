package shared

import (
	model "haii.or.th/api/server/model/setting"
	//model_setting "haii.or.th/api/thaiwater30/model/setting"
	"haii.or.th/api/util/service"
	//	"log"
)

func (srv *HttpService) getSystemSetting(ctx service.RequestContext, param *SharedParam) error {

	//Get Data
	objResult := &SharedParam{}
	objResult.Data = model.GetSystemSettingJson(param.Name)

	//	log.Println(param.Name)
	//Return Data
	ctx.ReplyJSON(objResult)

	return nil
}
