package analyst

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	//	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_media "haii.or.th/api/thaiwater30/model/media"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/radar_img
// @Summary 		ภาพเรดาร์
// @Method			GET
// @Produces		json
// @Response		200		Struct_RadarSwagger successful operation
type Struct_RadarSwagger struct {
	Result string                     `json:"result"` //example:`OK`
	Data   []model_media.Struct_Radar `json:"data"`
}

func (srv *HttpService) getRadarImg(ctx service.RequestContext) error {

	rs, err := model_latest_media.GetRadar()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Param_RadarHistory struct {
	Date      string `json:"date"`       // example:`2017-01-02` วันที่ของข้อมูล
	RadarType string `json:"radar_type"` //example:`cmp240` ประเภท Radar
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/radar_history_img
// @Summary 		ภาพเรดาร์ย้อนหลัง
// @Parameter		-	query Param_RadarHistory
// @Method			GET
// @Produces		json
// @Response		200		Result_Struct_RadarHistorySwagger successful operation
type Result_Struct_RadarHistorySwagger struct {
	Result string                                   `json:"result"` //example:`OK`
	Data   []model_media.Result_Struct_RadarHistory `json:"data"`
}

func (srv *HttpService) getRadarHistoryImg(ctx service.RequestContext) error {
	p := &Param_RadarHistory{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	rs, err := model_media.GetRadarHistory(p.RadarType, p.Date)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
