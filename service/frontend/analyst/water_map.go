package analyst

import (
	model_watermap "haii.or.th/api/thaiwater30/model/water_map"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)
// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/water_map
// @Summary			แผนผังน้ำ
// @Method			GET
// @Produces		json
// @Response		200		WaterMapOutputSwagger successful operation

type WaterMapOutputSwagger struct {
	Result string `json:"result"`  //example:`OK`
	Data []model_watermap.WaterMapOutput `json:"data"`
}
func (srv *HttpService) getWaterMap(ctx service.RequestContext) error {

	rs, err := model_watermap.GetWaterMap()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
