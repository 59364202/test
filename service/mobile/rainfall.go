package mobile

import (
	"haii.or.th/api/thaiwater30/model/rainfall"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rainfall24h_latest
// @Summary			ข้อมูลปริมาณฝนสะสม 24 ชม. ล่าสุด
// @Description		ข้อมูลปริมาณฝนสะสม 24 ชม. ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	rainfall.Struct_rainfall24h_latest successful operation
func (srv *HttpService) handlerGet_Rainfall24h_latest(ctx service.RequestContext) error {
	rs, err := rainfall.Get_Rainfall24h_latest()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rainfall_latest_list
// @Summary			ข้อมูลปริมาณฝน ล่าสุด ทั้งประเทศ
// @Description		ข้อมูลปริมาณฝน ล่าสุด ทั้งประเทศ
// @Method			GET
// @Produces		json
// @Response		200	rainfall.Struct_rainfall_latest_list successful operation
func (srv *HttpService) handlerGet_Rainfall_latest_list(ctx service.RequestContext) error {
	rs, err := rainfall.Get_Rainfall_latest_list()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rainfall_latest_list/{prov_id}
// @Summary			ข้อมูลปริมาณฝน ล่าสุด จังหวัด
// @Description		ข้อมูลปริมาณฝน ล่าสุด จังหวัด
// @Method			GET
// @Parameter		prov_id	path string example:`10` รหัสจังหวัด
// @Produces		json
// @Response		200	rainfall.Struct_rainfall_latest_list successful operation
func (srv *HttpService) handlerGet_rainfall_province(ctx service.RequestContext) error {
	prov_id := ctx.GetServiceParams("prov_id")
	rs, err := rainfall.Get_rainfall_province(prov_id)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rainfall_station_graph/{station_id}
// @Summary			ข้อมูลสถานีโทรมาตรฝน และกราฟ ฝน 7 วัน, 24 ชั่วโมง
// @Description		ข้อมูลสถานีโทรมาตรฝน และกราฟ ฝน 7 วัน, 24 ชั่วโมง
// @Method			GET
// @Parameter		station_id	path string example:`13829` รหัสสถานีโทรมาตร
// @Produces		json
// @Response		200	rainfall.Struct_rainfall_station_graph successful operation
func (srv *HttpService) handlerGet_rainfall_lateest_station_graph(ctx service.RequestContext) error {
	station_id := ctx.GetServiceParams("station_id")
	rs, err := rainfall.Get_rainfall_lateest_station_graph(station_id)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
