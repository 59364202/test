package mobile

import (
	"haii.or.th/api/thaiwater30/model/storm"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/storm
// @Summary			ข้อมูลพายุ ล่าสุด
// @Description		ข้อมูลพายุ ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	storm.Struct_storm_latest successful operation
func (srv *HttpService) handlerGet_storm(ctx service.RequestContext) error {
	// แสดงชื่อพายุ
	rs, err := storm.Get_storm_latest()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
