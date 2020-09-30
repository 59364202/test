package iframe

import (
	model_canal_waterlevel "haii.or.th/api/thaiwater30/model/canal_waterlevel"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"time"
)

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/waterlevel
// @Summary			รายชื่อจังหวัด
// @Description		เริ่มต้นหน้า iframe waterlevel
// @				* รายชื่อจังหวัด
// @Method			GET
// @Produces		json
// @Response		200	Struct_Iframe_Province successful operation
func (srv *HttpService) getWaterlevel(ctx service.RequestContext) error {
	rs, err := model_lt_geocode.GetAllProvince()
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type param_waterlevel_graph struct {
	Id           string `json:"id"`           // example:395 รหัสสถานี
	Station_type string `json:"station_type"` // enum:[tele_waterlevel,canal_waterlevel] example: tele_waterlevel ประเภทสถานี
}

type Struct_Iframe_Waterlevel_Graph struct {
	Data   *model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystOutput `json:"data"`   // ข้อมูลกราฟ
	Result string                                                                 `json:"result"` // example:`OK`
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/waterlevel_graph
// @Summary			ข้อมูลกราฟระดับน้ำ เฉพาะข้อมูล 3 วันล่าสุด ไม่สามารถใส่วันที่ได้
// @Method			GET
// @Parameter		-	query	param_waterlevel_graph
// @Produces		json
// @Response		200	Struct_Iframe_Waterlevel_Graph successful operation
func (srv *HttpService) getWaterlevelGraph(ctx service.RequestContext) error {
	p := &param_waterlevel_graph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	if p.Station_type == "tele_waterlevel" {
		now := time.Now()
		param := &model_tele_waterlevel.Waterlevel_InputParam{Station_id: p.Id, Start_date: now.AddDate(0, 0, -3).Format("2006-01-02"), End_date: now.AddDate(0, 0, 1).Format("2006-01-02")}
		rs, err := model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalyst(param)
		if err != nil {
			ctx.ReplyJSON(result.Result0(err.Error()))
		} else {
			ctx.ReplyJSON(result.Result1(rs))
		}
	} else if p.Station_type == "canal_waterlevel" {
		now := time.Now()
		param := &model_canal_waterlevel.Param_CanalWaterlevel{Station_id: p.Id, Start_date: now.AddDate(0, 0, -3).Format("2006-01-02"), End_date: now.AddDate(0, 0, 1).Format("2006-01-02")}
		rs, err := model_canal_waterlevel.GetCanalWaterlevelGraphByStationAndDateAnalyst(param)
		if err != nil {
			ctx.ReplyJSON(result.Result0(err.Error()))
		} else {
			ctx.ReplyJSON(result.Result1(rs))
		}
	}

	return nil
}
