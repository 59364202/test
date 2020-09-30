package analyst

import (
	model_cctv "haii.or.th/api/thaiwater30/model/cctv"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/cctv
// @Summary			Get URL CCTV
// @Description		Return URL CCTV
// @Method			GET
// @Produces		json
// @Response		200		CCTVSwagger successful operation
type CCTVSwagger struct {
	Result string       `json:"result"` //example:`OK`
	Data   []CctvOutput `json:"data"`
}

type CctvOutput struct {
	DamID       string `json:"dam_id"`            // example:`ldamrid0001` รหัสเขื่อน
	TeleStation string `json:"tele_station_name"` // example:`stationid` รหัสสถานีโทรมาตร
	BasinName   string `json:"basin_name"`        // example:`10`รหัสลุ่มน้ำ
	Lat         string `json:"lat"`               // example:`17.244115`พิกัดละติจูด
	Long        string `json:"long"`              // example:`98.972687`พิกัดลองติจูด
	Title       string `json:"title"`             // example:`เขื่อนภูมิพล`ชื่อสถานที่
	Description string `json:"description"`       // example:`description`คำอธิบายสถานที่
	MediaType	string `json:"media_type"`		  // example:`img`ชนิดข้อมูลของวิดีโอ
	URL         string `json:"cctv_url"`          // example:`http://cctv1.bhumiboldam.egat.com/` ที่อยู่วิดีโอ
}

func (srv *HttpService) getUrlCCTV(ctx service.RequestContext) error {

	rs, err := model_cctv.GetDetailsCCTV()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
