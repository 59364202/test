package mobile

import (
	"haii.or.th/api/thaiwater30/model/tele_waterlevel"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/wl_latest_list
// @Summary			ข้อมูลระดับน้ำสูงสุด และต่ำสุด ล่าสุด 20 อันดับ
// @Description		ข้อมูลระดับน้ำสูงสุด และต่ำสุด ล่าสุด 20 อันดับ
// @Method			GET
// @Produces		json
// @Response		200	tele_waterlevel.Struct_wl_latest_list successful operation
func (srv *HttpService) handlerGet_wl_latest_list(ctx service.RequestContext) error {
	rs, err := tele_waterlevel.Get_waterlevel_latest_list()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/wl_latest_list_prov/{prov_id}
// @Summary			ข้อมูลระดับน้ำ จังหวัด ล่าสุด 20 อันดับ
// @Description		ข้อมูลระดับน้ำ จังหวัด ล่าสุด 20 อันดับ
// @Method			GET
// @Parameter		prov_id	path string example:`10` รหัสจังหวัด
// @Produces		json
// @Response		200	tele_waterlevel.Struct_wl_latest_list_prov successful operation
func (srv *HttpService) handlerGet_wl_latest_prov(ctx service.RequestContext) error {
	prov_id := ctx.GetServiceParams("prov_id")
	rs, err := tele_waterlevel.Get_wl_latest_list_prov(prov_id)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/wl_station_graph/{station_id}
// @Summary			ข้อมูลกราฟ ระดับน้ำ 7 วัน, 24 ชั่วโมง
// @Description		ข้อมูลกราฟ ระดับน้ำ 7 วัน, 24 ชั่วโมง
// @Method			GET
// @Parameter		station_id	path string example:`10` สถานีโทรมาตร
// @Produces		json
// @Response		200	tele_waterlevel.Struct_wl_station_graph successful operation
func (srv *HttpService) handlerGet_wl_station_graph(ctx service.RequestContext) error {
	station_id := ctx.GetServiceParams("station_id")
	rs, err := tele_waterlevel.Get_Wl_station_graph(station_id)
	if err != nil {
		return err
	}

	//ถ้าไม่มีข้อมูลให้แสดง Fail
	if len(rs) > 0 {
		ctx.ReplyJSON(rs)
	} else {
		ctx.ReplyJSON(MobileErr("fail"))
	}
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/wl_basin_latest
// @Summary			ข้อมูลระดับน้ำ ล่าสุด
// @Description		ข้อมูลระดับน้ำ ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	tele_waterlevel.Struct_wl_basiN_latest successful operation
func (srv *HttpService) handlerGet_wl_basin_data(ctx service.RequestContext) error {
	rs, err := tele_waterlevel.Get_wl_basin_data()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
