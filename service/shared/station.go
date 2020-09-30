package shared

import (
	model_damstation "haii.or.th/api/thaiwater30/model/dam"
	model_telestation "haii.or.th/api/thaiwater30/model/tele_station"
	"haii.or.th/api/util/service"
)

func (srv *HttpService) getTeleStationByDataType(ctx service.RequestContext, param *SharedParam) error {
	//Map parameters
	p := &model_telestation.TeleStationParam{}
	p.DataType = param.DataType

	//Get Data
	result, err := model_telestation.GetTeleStationByDataType(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) getDamStationByDataType(ctx service.RequestContext, param *SharedParam) error {
	//Map parameters
	p := &model_damstation.Struct_GetDam_InputParam{}
	p.DataType = param.DataType

	//Get Data
	result, err := model_damstation.GetDamStationByDataType(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}
