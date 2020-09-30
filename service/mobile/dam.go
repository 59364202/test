package mobile

import (
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/model/dam_daily"

	"time"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/dam_summary_latest
// @Summary			ข้อมูลน้ำในเขื่อนหลัก ล่าสุด
// @Description		ข้อมูลน้ำในเขื่อนหลัก ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	dam_daily.Struct_WaterInformation successful operation
func (srv *HttpService) handlerGetDamSummaryLatest(ctx service.RequestContext) error {
	result, err := dam_daily.GetWaterInformationMainDam()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/dam_list
// @Summary			ข้อมูลพิกัด และปริมาณ
// @Description		ข้อมูลน้ำในเขื่อนหลัก ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	dam_daily.Struct_WaterInformation_DamList successful operation
func (srv *HttpService) handlerGetDamList(ctx service.RequestContext) error {
	result, err := dam_daily.GetWaterInformationDamList()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/dam_graph/history/{dam_id}
// @Summary			ข้อมูลกราฟเขื่อน ย้อนหลัง 2 ปี
// @Description		ข้อมูลกราฟเขื่อน ย้อนหลัง 2 ปี
// @Method			GET
// @Parameter		dam_id	path string example:`1` รหัสเขื่อน
// @Produces		json
// @Response		200	dam_daily.Struct_DamGraphHistory successful operation
func (srv *HttpService) handlerGetDamGraphHistory(ctx service.RequestContext) error {
	dam_id := ctx.GetServiceParams("dam_id")
	year := time.Now().AddDate(0, 0, -1).Year() - 2
	startYear := datatype.MakeString(year)
	endYear := datatype.MakeString(year + 1)

	result, err := dam_daily.GetDamGraphByYear(dam_id, startYear, endYear)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/dam_graph/current/{dam_id}
// @Summary			ข้อมูลกราฟเขื่อน ปีปัจจุบัน
// @Description		ข้อมูลกราฟเขื่อน ปีปัจจุบัน
// @Method			GET
// @Parameter		dam_id	path string example:`1` รหัสเขื่อน
// @Produces		json
// @Response		200	dam_daily.Struct_DamGraphHistory successful operation
func (srv *HttpService) handlerGetDamGraphCurrent(ctx service.RequestContext) error {
	dam_id := ctx.GetServiceParams("dam_id")
	year := time.Now().AddDate(0, 0, -1).Year()
	startYear := datatype.MakeString(year)

	result, err := dam_daily.GetDamGraphByYear(dam_id, startYear, startYear)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}
