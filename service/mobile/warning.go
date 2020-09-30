package mobile

import (
	"haii.or.th/api/thaiwater30/model/rainfall24hr"
	"haii.or.th/api/thaiwater30/model/rainfall_1h"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/warning/bkk/rainfall24h
// @Summary			ข้อมูลตรวจวัด ฝนสะสม 24 ชม.
// @Description		ข้อมูลตรวจวัด ฝนสะสม 24 ชม.
// @Method			GET
// @Produces		json
// @Response		200	rainfall24hr.Struct_rainfall24h successful operation
func (srv *HttpService) handlerGetBKK_Rainfall24h(ctx service.RequestContext) error {
	rs, err := rainfall24hr.Get_BKK_Rainfall24h()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/warning/bkk/rainfall6h
// @Summary			ข้อมูลตรวจวัด ฝนรายชั่วโมง
// @Description		ข้อมูลตรวจวัด ฝนรายชั่วโมง
// @Method			GET
// @Produces		json
// @Response		200	rainfall_1h.Struct_bkk_rainfall1h successful operation
func (srv *HttpService) handlerGet_bkk_Rainfall6h(ctx service.RequestContext) error {
	rs, err := rainfall_1h.Get_bkk_rainfall1h()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
