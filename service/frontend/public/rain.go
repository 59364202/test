package public

import (
	"encoding/json"
	"time"

	"haii.or.th/api/server/model/setting"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_rainfall "haii.or.th/api/thaiwater30/model/rainfall"
	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_rainfall_1h "haii.or.th/api/thaiwater30/model/rainfall_1h"
	model_rainfall_daily "haii.or.th/api/thaiwater30/model/rainfall_daily"
	model_rainfall_monthly "haii.or.th/api/thaiwater30/model/rainfall_monthly"
	model_rainfall_today "haii.or.th/api/thaiwater30/model/rainfall_today"
	model_rainfall_yearly "haii.or.th/api/thaiwater30/model/rainfall_yearly"
	model_subbasin "haii.or.th/api/thaiwater30/model/subbasin"
	//	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
	"haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

type Struct_getRainOnLoad struct {
	Province *Struct_getRainOnLoad_Province `json:"province"` // สถานีกรุ๊ปตามจังหวัด
	Scale    []*uSetting.Struct_RainSetting `json:"scale"`    // เกณฑ์
}
type Struct_getRainOnLoad_Province struct {
	Result string                       `json:"result"` // example:`OK`
	Data   []*model_lt_geocode.Province `json:"data"`   // สถานีกรุ๊ปตามจังหวัด
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_load
// @Summary			เกณฑ์ฝน, ข้อมูลสถานีกรุ๊ปตามจังหวัด
// @Description		เกณฑ์ฝน, ข้อมูลสถานีกรุ๊ปตามจังหวัด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getRainOnLoad successful operation
func (srv *HttpService) getRainOnLoad(ctx service.RequestContext) error {
	type s struct {
		Province *result.Result  `json:"province"` // สถานีกรุ๊ปตามจังหวัด
		Scale    json.RawMessage `json:"scale"`    // setting
	}
	rs := &s{}

	rs_province, err := model_lt_geocode.GetProvinceByStation()
	if err != nil {
		rs.Province = result.Result0(err)
	} else {
		rs.Province = result.Result1(rs_province)
	}
	rs.Scale = setting.GetSystemSettingJson("Frontend.public.rain_setting")

	ctx.ReplyJSON(rs)

	return nil
}

type Struct_getRain24Hr struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_rainfall24hr.Rainfall24HrStruct `json:"data"`   // ฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_24h
// @Summary			ฝนย้อนหลัง 24 ชม.
// @Method			GET
// @Produces		json
// @Response		200	Struct_getRain24Hr successful operation
func (srv *HttpService) getRain24Hr(ctx service.RequestContext) error {
	p := &model_rainfall24hr.Param_Rainfall24{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	rs, err := model_rainfall24hr.GetRainfallThailandDataCache(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRain24HrGraph struct {
	Result string                                       `json:"result"` // example:`OK`
	Data   []*model_rainfall_1h.Struct_Rainfall1h_Graph `json:"data"`   // กราฟฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_24h_graph
// @Summary			กราฟฝนปริมาณย้อนหลัง 24 ชม.
// @Method			GET
// @Parameter		station_id	query int64 example: 1051 รหัสสถานี
// @Produces		json
// @Response		200	Struct_getRain24HrGraph successful operation
func (srv *HttpService) getRain24HrGraph(ctx service.RequestContext) error {
	p := &model_rainfall_1h.Param_Rainfall1h_Graph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	p.Is24 = true
	rs, err := model_rainfall_1h.GetRainGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_hour_graph
// @Summary			กราฟฝนปริมาณรายชั่วโมงย้อนหลัง
// @Method			GET
// @Parameter		station_id	query int64 example: 1051 รหัสสถานี
// @Produces		json
// @Response		200	Struct_getRain24HrGraph successful operation
func (srv *HttpService) getRainHourGraph(ctx service.RequestContext) error {
	p := &model_rainfall_1h.Param_Rainfall1h_Graph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	now := time.Now()
	if p.DateStart == "" || p.DateEnd == "" {
		p.DateStart = now.AddDate(0, 0, -3).Format("2006-01-02 15:00")
		p.DateEnd = now.Format("2006-01-02 15:00")
	}

	rs, err := model_rainfall_1h.GetRainGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainToday struct {
	Result string                                       `json:"result"` // example:`OK`
	Data   []*model_rainfall_today.Struct_RainfallToday `json:"data"`   // ฝน
}

// rain today
// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_today
// @Summary			ฝนวันนี้
// @Method			GET
// @Produces		json
// @Response		200	Struct_getRainToday successful operation
func (srv *HttpService) getRainToday(ctx service.RequestContext) error {
	p := &model_rainfall_today.Param_RainfallToday{}
	rs, err := model_rainfall_today.GetRain(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainTodayGraph struct {
	Result string                                             `json:"result"` // example:`OK`
	Data   []*model_rainfall_today.Struct_RainfallToday_Graph `json:"data"`   // กราฟฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_today_graph
// @Summary			กราฟฝนวันนี้
// @Description		กราฟฝนตั้งแต่ 7.01 - ปัจจุบัน แต่ถ้า ปัจจุบันยังไม่ถึง 8 โมง จะได้เป็นกราฟฝนย้อนหลัง  24 ชม.
// @Method			GET
// @Parameter		station_id	query int64 example: 1051 รหัสสถานี
// @Produces		json
// @Response		200	Struct_getRainTodayGraph successful operation
func (srv *HttpService) getRainTodayGraph(ctx service.RequestContext) error {
	p := &model_rainfall_today.Param_RainfallToday{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	var rs interface{}
	var err error
	p.IsToday = true
	if time.Now().Hour() < 8 {
		rs, err = model_rainfall_1h.GetRainGraph(&model_rainfall_1h.Param_Rainfall1h_Graph{StationId: p.StationId, Is24: true})
	} else {
		rs, err = model_rainfall_today.GetRainGraph(p)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainDaily struct {
	Result string                                       `json:"result"` // example:`OK`
	Data   []*model_rainfall_daily.Struct_RainfallDaily `json:"data"`   // ฝน
}

// rain daily
// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_yesterday
// @Summary			ฝนวานนี้
// @Method			GET
// @Produces		json
// @Response		200 Struct_getRainDaily successful operation
func (srv *HttpService) getRainDaily(ctx service.RequestContext) error {
	p := &model_rainfall_daily.Param_RainfallDaily{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	p.IsDaily = true
	rs, err := model_rainfall_daily.GetRain(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainDailyGraph struct {
	Result string                                             `json:"result"` // example:`OK`
	Data   []*model_rainfall_daily.Struct_RainfallDaily_Graph `json:"data"`   // กราฟฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_yesterday_graph
// @Summary			กราฟฝนรายวัน
// @Method			GET
// @Parameter		station_id	query int64 example: 1051 รหัสสถานี
// @Parameter		start_date	query	string	example: 2006-01-02 วันที่เริ่มต้น
// @Parameter		end_date	query	string	example: 2006-01-02 วันที่สิ้นสุด
// @Produces		json
// @Response		200	Struct_getRainDailyGraph successful operation
func (srv *HttpService) getRainDailyGraph(ctx service.RequestContext) error {
	p := &model_rainfall_daily.Param_RainfallDaily{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	p.IsDaily = true
	rs, err := model_rainfall_daily.GetRainGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainMonthly struct {
	Result string                                           `json:"result"` // example:`OK`
	Data   []*model_rainfall_monthly.Struct_RainfallMonthly `json:"data"`   // ฝน
}

// rain monthly
// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_monthly
// @Summary			ฝนรายเดือน
// @Method			GET
// @Produces		json
// @Response		200	Struct_getRainMonthly successful operation
func (srv *HttpService) getRainMonthly(ctx service.RequestContext) error {
	p := &model_rainfall_monthly.Param_RainfallMonthly{}
	ctx.LogRequestParams(p)
	rs, err := model_rainfall_monthly.GetRain(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainMonthlyGraph struct {
	Result string                                             `json:"result"` // example:`OK`
	Data   []*model_rainfall_daily.Struct_RainfallDaily_Graph `json:"data"`   // ฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_monthly_graph
// @Summary			กราฟฝนรายเดือน
// @Method			GET
// @Parameter		station_id	query int64 example: 1051 รหัสสถานี
// @Parameter		month	query int example:1 เดือน
// @Parameter		year	query int example:2006 ปี
// @Produces		json
// @Response		200	Struct_getRainMonthlyGraph successful operation
func (srv *HttpService) getRainMonthlyGraph(ctx service.RequestContext) error {
	p := &model_rainfall_daily.Param_RainfallDaily{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	p.IsMonthly = true
	rs, err := model_rainfall_daily.GetRainGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainYearly struct {
	Result string                                         `json:"result"` // example:`OK`
	Data   []*model_rainfall_yearly.Struct_RainfallYearly `json:"data"`   // ฝน
}

// rain yearly
// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_yearly
// @Summary			ฝนรายปี
// @Method			GET
// @Produces		json
// @Response		200	Struct_getRainYearly successful operation
func (srv *HttpService) getRainYearly(ctx service.RequestContext) error {
	p := &model_rainfall_yearly.Param_RainfallYearly{}
	rs, err := model_rainfall_yearly.GetRain(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_getRainYearlyGraph struct {
	Result string                                                 `json:"result"` // example:`OK`
	Data   []*model_rainfall_monthly.Struct_RainfallMonthly_Graph `json:"data"`   // ฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain_yearly_graph
// @Summary			กราฟฝนรายปี
// @Method			GET
// @Parameter		station_id	query int64 example: 1051 รหัสสถานี
// @Parameter		year	query int example:2006 ปี
// @Produces		json
// @Response		200	Struct_getRainYearlyGraph successful operation
func (srv *HttpService) getRainYearlyGraph(ctx service.RequestContext) error {
	p := &model_rainfall_monthly.Param_RainfallMonthly{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	//	p.IsYearly = true
	rs, err := model_rainfall_monthly.GetRainGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_rain_monthly_station_graph
// @Method			GET
// @Summary			ข้อมูลฝนรายเดือน รายสถานีสำหรับกราฟ
// @Parameter		-	query model_rainfall.GetAdvRainMonthStationGraphInput
// @Produces		json
// @Response		200		StructRainMonthlyStationGraph successful operation
type StructRainMonthlyStationGraph struct {
	Result string                   `json:"result"` // example:`OK`
	Data   model_rainfall.GraphData `json:"data"`   // ฝน
}

func (srv *HttpService) getAdvRainMonthStationGraph(ctx service.RequestContext) error {
	p := &model_rainfall.GetAdvRainMonthStationGraphInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_rainfall.GetAdvRainMonthlyStationGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_rain_monthly_graph
// @Method			GET
// @Summary			ข้อมูลฝนรายเดือนสำหรับกราฟ
// @Parameter		-	query model_rainfall.GetAdvRainMonthGraphInput
// @Produces		json
// @Response		200		StructAdvRainMonthGraph successful operation
type StructAdvRainMonthGraph struct {
	Result string                   `json:"result"` // example:`OK`
	Data   model_rainfall.GraphData `json:"data"`   // ฝน
}

func (srv *HttpService) getAdvRainMonthGraph(ctx service.RequestContext) error {
	p := &model_rainfall.GetAdvRainMonthGraphInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_rainfall.GetAdvRainMonthlyGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_rain_yearly_graph
// @Method			GET
// @Summary			ข้อมูลฝนรายปีสำหรับกราฟ
// @Parameter		-	query model_rainfall.GetAdvRainYearlyGraphInput
// @Produces		json
// @Response		200		StructAdvRainYearlyGraph successful operation
type StructAdvRainYearlyGraph struct {
	Result string                         `json:"result"` // example:`OK`
	Data   model_rainfall.GraphDataYearly `json:"data"`   // ฝน
}

func (srv *HttpService) getAdvRainYearGraph(ctx service.RequestContext) error {
	p := &model_rainfall.GetAdvRainYearlyGraphInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_rainfall.GetAdvRainYearlyGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_rain_monthly_region_graph
// @Method			GET
// @Summary			ข้อมูลฝนรายเดือน รายภาค
// @Parameter		-	query model_rainfall.GetAdvRainMonthAreaGraphInput
// @Produces		json
// @Response		200		StructGetAdvRainAreaOutput successful operation
type StructGetAdvRainAreaOutput struct {
	Result string                              `json:"result"` // example:`OK`
	Data   model_rainfall.GetAdvRainAreaOutput `json:"data"`   // ฝน
}

func (srv *HttpService) getAdvRainMonthlyAreaGraph(ctx service.RequestContext) error {
	p := &model_rainfall.GetAdvRainMonthAreaGraphInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_rainfall.GetAdvRainMonthlyAreaGraph(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Param_AdvRainDiagram struct {
	AgencyId []int64 `json:"agency_id"` // example: [1,2] รหัสหน่วยงาน
	Date     string  `json:"date"`      // example: 2016-01-02 วันเวลา
}
type Struct_getAdvRainDiagram struct {
	Result string                                      `json:"result"` // example:`OK`
	Data   []*model_rainfall24hr.Struct_AdvRainDiagram `json:"data"`   // การกระจายฝนล่วงหน้า
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_rain_distribution
// @Method			GET
// @Summary			การกระจายฝนล่วงหน้า
// @Parameter		-	query Param_AdvRainDiagram
// @Produces		json
// @Response		200	Struct_getAdvRainDiagram successful operation
func (srv *HttpService) getAdvRainDiagram(ctx service.RequestContext) error {
	p := &Param_AdvRainDiagram{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	if len(p.AgencyId) <= 0 {
		return rest.NewError(422, "no agency_id", nil)
	}
	if p.Date == "" {
		return rest.NewError(422, "no date", nil)
	}

	rs, err := model_rainfall24hr.GetAdvRainDiagram(p.AgencyId, p.Date)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type AdvRainGraphOnload struct {
	Area     []*model_lt_geocode.Struct_Region `json:"region"`   // ขอบเขต
	Province []*model_lt_geocode.Province      `json:"province"` // จังหวัด
	Basin    []*model_subbasin.Basin           `json:"basin"`    // ลุ่มน้ำ
	Baseline []int64                           `json:"normal"`   // example:`[30,48]` baseline
}
type Struct_getAdvRainGraphOnload struct {
	Result string              `json:"result"` // example:`OK`
	Data   *AdvRainGraphOnload `json:"data"`   // ขอบเขต, จังหวัด, ลุ่มน้ำ, baseline
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_graph_load
// @Summary			เริ่มต้นกราฟของผู้ใช้งานขั้นสูง
// @Description		ขอบเขต, จังหวัด, ลุ่มน้ำ, baseline สำหรับทำ dropdown
// @Method			GET
// @Produces		json
// @Response		200	AdvRainGraphOnload successful operation
func (srv *HttpService) getAdvRainGraphOnload(ctx service.RequestContext) error {
	rs := &AdvRainGraphOnload{}

	rs.Area, _ = model_lt_geocode.GetAllArea()
	rs.Province, _ = model_lt_geocode.GetProvinceByStation()
	rs.Basin, _ = model_subbasin.GetSubbasin()
	rs.Baseline = []int64{30, 48}

	ctx.ReplyJSON(result.Result1(rs))
	//	rs.Baseline
	return nil
}

type Struct_AdvLoad struct {
	Agency []*model_rainfall24hr.Struct_AdvOnload_Agency `json:"agency"` // หน่วยงาน
	Scope  json.RawMessage                               `json:"scope"`  // example:`[{"name":"Thailand","value":1}]` ขอบเขต
	Scale  json.RawMessage                               `json:"scale"`  // example:`{ "rule": { "01": [ { "operator": ">=", "rain1h": "30", "rain24h": "80", "rain3d": "160", "level": "4" }, { "operator": ">=","rain1h": "25","rain24h": "65","rain3d": "95","level": "3"}]}}` เกณฑ์
}
type Struct_getAdvLoad struct {
	Result string                  `json:"result"` // example:`OK`
	Data   *Struct_getAdvLoad_Data `json:"data"`   // หน่วยงาน, ขอบเขต, เกณฑ์
}
type Struct_getAdvLoad_Data struct {
	Agency []*model_rainfall24hr.Struct_AdvOnload_Agency `json:"agency"` // หน่วยงาน
	Scope  *uSetting.Struct_AdvanceLoad_Scope            `json:"scope"`  // ขอบเขต
	Scale  *uSetting.Struct_RainSetting                  `json:"scale"`  // เกณฑ์
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_load
// @Summary			เริ่มต้นหน้าผู้ใช้งานขั้นสูง
// @Description		หน่วยงาน, ขอบเขต, เกณฑ์
// @Method			GET
// @Produces		json
// @Response		200	Struct_getAdvLoad successful operation
func (srv *HttpService) getAdvLoad(ctx service.RequestContext) error {
	rs := &Struct_AdvLoad{}

	rs_agency, err := model_rainfall24hr.GetAdvOnload_Agency()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
		return nil
	}
	rs.Agency = rs_agency
	rs.Scope = setting.GetSystemSettingJSON("bof.analyst.advance_load.scope")
	rs.Scale = setting.GetSystemSettingJSON("Frontend.public.rain_setting")

	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

type Struct_getAdvRainSum struct {
	Result string                                  `json:"result"` // example:`OK`
	Data   []*model_rainfall24hr.Struct_AdvRainSum `json:"data"`   // ข้อมูลการกระจายตัวของฝน
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_rain_sum
// @Summary			ข้อมูลการกระจายตัวของฝน
// @Description		ข้อมูลการกระจายตัวของฝน นำไปใช้ทำแผนภาพการกระจายตัวของฝน
// @Method			GET
// @Parameter		-	query model_rainfall24hr.Param_AdvRainSum
// @Produces		json
// @Response		200	Struct_getAdvRainSum successful operation
func (srv *HttpService) getAdvRainSum(ctx service.RequestContext) error {
	p := &model_rainfall24hr.Param_AdvRainSum{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_rainfall24hr.GetAdvRainSum(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
		return nil
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}
