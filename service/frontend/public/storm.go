// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

// ข้อมูลขื่อพายุล่าสุด

import (
	"encoding/json"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/service"

	model_storm "haii.or.th/api/thaiwater30/model/storm"
	result "haii.or.th/api/thaiwater30/util/result"
)

// @DocumentName	v1.public
// @Service			thaiwater30/public/storm_data
// @Summary			ข้อมูลพายุล่าสุด
// @Method			GET
// @Produces		json
// @Response		200		Struct_storm successful operation

type Struct_storm struct {
	Result string                      `json:"result"`
	Data   []*model_storm.Struct_Strom `json:"data"`    // พายุ
	Scale  json.RawMessage             `json:"setting"` // เกณฑ์
}

func (srv *HttpService) getStormData(ctx service.RequestContext) error {
	rs := &Struct_storm{}
	data, err := model_storm.GetStormCurrentDate()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		rs.Result = "OK"
		rs.Data = data
		rs.Scale = setting.GetSystemSettingJSON("Frontend.public.storm_setting")
		ctx.ReplyJSON(rs)
	}
	return nil
}
