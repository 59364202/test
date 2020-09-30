package mobile

import (
	"haii.or.th/api/thaiwater30/model/tmd_data"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/temp_latest_prov/{prov_id}
// @Summary			ข้อมูลอุณหภูมิ จังหวัด ล่าสุด
// @Description		ข้อมูลอุณหภูมิ จังหวัด ล่าสุด
// @Method			GET
// @Parameter		prov_id	path string example:`10` รหัสจังหวัด
// @Produces		json
// @Response		200	tmd_data.Struct_temperature successful operation
func (srv *HttpService) handlerGet_temperature(ctx service.RequestContext) error {
	prov_id := ctx.GetServiceParams("prov_id")
	rs, err := tmd_data.Get_temperature(prov_id)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
