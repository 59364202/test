package iframe

import (
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/service"

	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_rainfall "haii.or.th/api/thaiwater30/model/rainfall"
	model_rainfall_1h "haii.or.th/api/thaiwater30/model/rainfall_1h"
	model_rainfall_daily "haii.or.th/api/thaiwater30/model/rainfall_daily"

	"haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
)

type rain24IframeStruct struct {
	Province *result.Result     `json:"province"`
	Datatype *result.ResultJson `json:"datatype"`
}
type Struct_Iframe_Rain24 struct {
	Province *Struct_Iframe_Province       `json:"province"` // จังหวัด
	Datatype Struct_Iframe_Rain24_Datatype `json:"datatype"` // ประเภทข้อมูล
}
type Struct_Iframe_Province struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_lt_geocode.Struct_Geocode_P `json:"data"`   // จังหวัด
}
type Struct_Iframe_Rain24_Datatype struct {
	Result string                              `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_RainfallDataType `json:"data"`   // ประเภทข้อมูล
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/rain24
// @Summary			รายชื่อจังหวัด, ประเภทข้อมูลฝน
// @Description		เริ่มต้นหน้า iframe rain24
// @				* รายชื่อจังหวัด
// @				* ประเภทข้อมูลฝน
// @Method			GET
// @Produces		json
// @Response		200	Struct_Iframe_Rain24 successful operation
func (srv *HttpService) getRain24(ctx service.RequestContext) error {
	rs := &rain24IframeStruct{}
	rs_province, err := model_lt_geocode.GetAllProvince()
	if err != nil {
		rs.Province = result.Result0(err)
	} else {
		rs.Province = result.Result1(rs_province)
	}

	rs.Datatype = &result.ResultJson{Result: "OK", Data: setting.GetSystemSettingJson("Frontend.public.rainfall_data_type")}

	ctx.ReplyJSON(rs)

	return nil
}

type Struct_Iframe_Rain24_Graph struct {
	Result string                                             `json:"result"` // example:`OK`
	Data   []*model_rainfall_daily.Struct_RainfallDaily_Graph `json:"data"`   // ข้อมูลกราฟ
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/rain24_graph
// @Summary			ข้อมูลกราฟฝน
// @Method			GET
// @Parameter		id	query	int64	example: 305 รหัสสถานี
// @Parameter		datatype	query	string	enum:[1,2] example:1 ประเภทข้อมูล 1 ฝนชม., 2 ฝนวัน
// @Produces		json
// @Response		200	Struct_Iframe_Rain24_Graph successful operation
func (srv *HttpService) getRain24Graph(ctx service.RequestContext) error {
	p := &model_rainfall.GetRainGraphParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	if p.Datatype == "1" {
		rs, err := model_rainfall_1h.GetRainGraph(&model_rainfall_1h.Param_Rainfall1h_Graph{Is24: true, StationId: p.Id})
		if err != nil {
			return err
		} else {
			ctx.ReplyJSON(result.Result1(rs))
		}
	} else {
		rs, err := model_rainfall_daily.GetRainGraph(&model_rainfall_daily.Param_RainfallDaily{IsDaily: true, StationId: p.Id})
		if err != nil {
			return err
		} else {
			ctx.ReplyJSON(result.Result1(rs))
		}
	}
	return nil
}
