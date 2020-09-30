package mobile

import (
	"haii.or.th/api/thaiwater30/model/tmd_weather_report"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/weather_today/{year}/{month}/{day}
// @Summary			ข้อมูลสภาพอากาศ รายวัน
// @Description		ข้อมูลสภาพอากาศ รายวัน
// @Method			GET
// @Parameter		year	path string example:`2018` ปี
// @Parameter		month	path string example:`03` เดือน
// @Parameter		day	path string example:`15` วัน
// @Produces		json
// @Response		200	tmd_weather_report.Struct_Weather_today successful operation
func (srv *HttpService) handlerGet_weather_today(ctx service.RequestContext) error {
	year := ctx.GetServiceParams("year")
	month := ctx.GetServiceParams("month")
	day := ctx.GetServiceParams("day")

	rs, err := tmd_weather_report.Get_weather_today(year, month, day)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
