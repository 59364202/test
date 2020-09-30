package public

import (
	//model_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_canal_station "haii.or.th/api/thaiwater30/model/canal_station"
	model_canal_waterlevel "haii.or.th/api/thaiwater30/model/canal_waterlevel"
	model_forecast "haii.or.th/api/thaiwater30/model/forecast"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"

	"haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	//"strings"
	"time"
)

type Struct_OnLoadWaterlevelLastest struct {
	Data     *result.Result `json:"waterlevel_data"`
	Basin    *result.Result `json:"basin"`
	Agency   *result.Result `json:"agency"`
	Station  *result.Result `json:"station"`
	Scale    *result.Result `json:"scale,omitempty"`
	Province *result.Result `json:"province"`
}

// watergateload
type Struct_OnLoadWaterlevelInOutLastest struct {
	Data     *result.Result `json:"watergate_data"`
	Station  *result.Result `json:"station"`
	Province *result.Result `json:"province"`
}

type Struct_result struct {
	Tele  []*model_tele_station.Struct_TeleStationGroupByProvince   `json:"tele_waterlevel"`  // ระดับน้ำ
	Canal []*model_canal_station.Struct_CanalStationGroupByProvince `json:"canal_waterlevel"` // ระดับน้าในคลอง
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/flood_forecast_data
// @Method			GET
// @Summary			คาดการณ์น้ำ
// @Produces		json
// @Response		200		FloodForecastOutPutCpySwagger successful operation
type FloodForecastOutPutCpySwagger struct {
	Result string                                               `json:"result"` //example:`OK`
	Data   []model_forecast.FloodforecastOutputWithScaleSwagger `json:"data"`
}

func (srv *HttpService) getCpyForecastWaterlevel(ctx service.RequestContext) error {

	rs, err := model_forecast.CpyFloodForecastLatest()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/swan_station
// @Method			GET
// @Summary			สถานีคาดการณ์ความสูงคลื่น
// @Produces		json
// @Response		200		SwanStationOutputSwagger successful operation

type SwanStationOutputSwagger struct {
	Result string                             `json:"result"` //example:`OK`
	Data   []model_forecast.SwanStationOutput `json:"data"`
}

func (srv *HttpService) getSwanStation(ctx service.RequestContext) error {

	rs, err := model_forecast.SwanStationDetails()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/wave_forecast
// @Method			GET
// @Parameter		-	query Struct_Swan_Forecast_InputParam
// @Summary			คาดการณ์ความสูงคลื่น
// @Produces		json
// @Response		200		SwanForecastOutputSwagger successful operation

type SwanForecastOutputSwagger struct {
	Result string                              `json:"result"` //example:`OK`
	Data   []model_forecast.SwanForecastOutput `json:"data"`
}

func (srv *HttpService) getSwanForecast(ctx service.RequestContext) error {
	//Map parameters
	p := &Struct_Swan_Forecast_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	paramSwanForecast := &model_forecast.Struct_Swan_Forecast_InputParam{}
	paramSwanForecast.StationID = p.StationID

	rs, err := model_forecast.SwanForecastLatest(paramSwanForecast)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_Swan_Forecast_InputParam struct {
	StationID string `json:"station_id"` // required:false example:`1` รหัสสถานีคาดการณ์คลื่น ไม่ใส่ = ทุกสถานี เลือกได้หลายสถานี เช่น 1,2,3
}

type Struct_onLoadWaterlevel struct {
	Data     *Struct_onLoadWaterlevel_Data     `json:"waterlevel_data"` // ระดับน้ำ
	Basin    *Struct_onLoadWaterlevel_Basin    `json:"basin"`           // ลุ่มน้ำ
	Agency   *Struct_onLoadWaterlevel_Agency   `json:"agency"`          // หน่วยงาน
	Station  *Struct_onLoadWaterlevel_Station  `json:"station"`         // สถานี
	Scale    *Struct_onLoadWaterlevel_Scale    `json:"scale,omitempty"` // เกณฑ์
	Province *Struct_onLoadWaterlevel_Province `json:"province"`        // จังหวัด
}
type Struct_onLoadWaterlevel_Data struct {
	Result string                                     `json:"result"` // example:`OK`
	Data   []*model_tele_waterlevel.Struct_Waterlevel `json:"data"`   // ระดับน้ำ
}
type Struct_onLoadWaterlevel_Basin struct {
	Result string                      `json:"result"` // example:`OK`
	Data   []*model_basin.Struct_Basin `json:"data"`   // ลุ่มน้ำ
}
type Struct_onLoadWaterlevel_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}
type Struct_onLoadWaterlevel_Station struct {
	Result string         `json:"result"` // example:`OK`
	Data   *Struct_result `json:"data"`   // สถานี
}
type Struct_onLoadWaterlevel_Scale struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_WaterlevelSetting `json:"data"`   // เกณฑ์
}
type Struct_onLoadWaterlevel_Province struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model_geocode.Struct_Geocode `json:"data"`   // จังหวัด
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/waterlevel_load
// @Summary			เริ่มต้นหน้าระดับน้ำ
// @Description		ระดับน้ำ, ลุ่มน้ำ, หน่วยงาน, สถานี, เกณฑ์, จังหวัด
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadWaterlevel successful operation

func (srv *HttpService) onLoadWaterlevel(ctx service.RequestContext) error {

	objResult := &Struct_OnLoadWaterlevelLastest{}

	//=== Scale ===//
	rsScale := model_setting.GetSystemSettingJSON("Frontend.public.waterlevel_setting")
	objResult.Scale = result.Result1(&rsScale)

	//=== Basin ===//
	resultBasin, err := model_basin.GetAllBasin()
	if err != nil {
		objResult.Basin = result.Result0(err)
	} else {
		objResult.Basin = result.Result1(resultBasin)
	}

	//=== Station ===//
	var rs *Struct_result = new(Struct_result)
	rs_station, err_station := model_tele_station.GetTeleStationGroupByProvince(&model_tele_station.TeleStationParam{DataType: "waterlevel"})
	if err_station != nil {
		objResult.Station = result.Result0(err_station)
	} else {
		rs.Tele = make([]*model_tele_station.Struct_TeleStationGroupByProvince, 0)
		rs.Tele = append(rs.Tele, rs_station...)
	}
	rs_canal, err_canal := model_canal_station.GetCanalStationGroupByProvince()
	if err_canal != nil {
		objResult.Station = result.Result0(err_canal)
	} else {
		rs.Canal = make([]*model_canal_station.Struct_CanalStationGroupByProvince, 0)
		rs.Canal = append(rs.Canal, rs_canal...)
	}
	objResult.Station = result.Result1(rs)

	//=== Agency ===//
	resultAgency, err := model_agency.GetAgencyInCanalWaterlevelTable()
	if err != nil {
		objResult.Agency = result.Result0(err)
	} else {
		objResult.Agency = result.Result1(resultAgency)
	}

	//=== Data ===//
	p := &model_tele_waterlevel.Waterlevel_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		objResult.Data = result.Result0(err)
	} else {
		resultData, err := model_tele_waterlevel.GetWaterLevelThailandDataCache(p)
		if err != nil {
			objResult.Data = result.Result0(err)
		} else {
			objResult.Data = result.Result1(resultData)
		}
	}

	//=== Province ===//
	rsProvince, err := model_geocode.GetAllProvince()
	if err != nil {
		objResult.Province = result.Result0(err)
	} else {
		objResult.Province = result.Result1(rsProvince)
	}

	ctx.ReplyJSON(objResult)
	return nil
}

type Struct_getWaterlevelAnalyst struct {
	Result string                                     `json:"result"` // example:`OK`
	Data   []*model_tele_waterlevel.Struct_Waterlevel `json:"data"`   // ระดับน้ำ
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/waterlevel
// @Summary			ลุ่มน้ำ
// @Method			GET
// @Parameter		agency_id	query	string	example: 9 รหัสหน่วยงาน
// @Parameter		basin_id	query	string	example: 1 รหัสลุ่มน้ำ
// @Produces		json
// @Response		200	Struct_getWaterlevelAnalyst successful operation
func (srv *HttpService) getWaterlevelAnalyst(ctx service.RequestContext) error {

	p := &model_tele_waterlevel.Waterlevel_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	rs, err := model_tele_waterlevel.GetWaterLevelThailandDataCache(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type paramWaterlevelGraph struct {
	Id          string `json:"station_id"`   // รหัสสถานีระดับน้ำ เช่น 3425
	StationType string `json:"station_type"` // ประเภทของสถานีระดับน้ำ เช่น W,A
	Year        []int  `json:"year"`         // ปี เช่น [2015,2016]
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type paramWaterlevelGraphSwagger1 struct {
	Id          string `json:"station_id"`   // รหัสสถานีระดับน้ำ เช่น 100
	StationType string `json:"station_type"` // ประเภทของสถานีระดับน้ำ เช่น tele_waterlevel หรือ canal_waterlevel
	StartDate   string `json:"start_date"`   // required:false วันเริ่มต้น 2017-06-10 ถ้าไม่กำหนดวันที่เริ่มต้น สิ้นสุด จะ return ข้อมูลย้อนหลัง 3 วันจากวันปัจจุบัน
	EndDate     string `json:"end_date"`     // required:false วันสิ้นสุด 2017-06-11
}

type paramWaterlevelGraphSwagger2 struct {
	Id          string `json:"station_id"`   // รหัสสถานีระดับน้ำ เช่น 3425
	StationType string `json:"station_type"` // ประเภทของสถานีระดับน้ำ เช่น tele_waterlevel หรือ canal_waterlevel
	Year        []int  `json:"year"`         // ปี เช่น [2016]
}

// @DocumentName	v1.public
// @Service			thaiwater30//public/waterlevel_graph
// @Method			GET
// @Summary			กราฟระดับน้ำตามช่วงวัน
// @Parameter		-	query paramWaterlevelGraphSwagger1
// @Produces		json
// @Response		200		GetWaterlevelGraphByStationAndDateAnalystOutputSwagger successful operation

type GetWaterlevelGraphByStationAndDateAnalystOutputSwagger struct {
	Result string                                                                  `json:"result"` //example:`OK`
	Data   []model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystOutput `json:"data"`
}

func (srv *HttpService) getWaterlevelGraphByStationAndDateAnalyst(ctx service.RequestContext) error {

	p := &paramWaterlevelGraph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	var err error
	var rs *model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystOutput
	ctx.LogRequestParams(p)
	now := time.Now()
	if p.StartDate == "" || p.EndDate == "" {
		p.StartDate = now.AddDate(0, 0, -3).Format("2006-01-02")
		p.EndDate = now.Format("2006-01-02")
	}

	if p.StationType == "tele_waterlevel" {
		params := &model_tele_waterlevel.Waterlevel_InputParam{Station_id: p.Id, Start_date: p.StartDate, End_date: p.EndDate}
		rs, err = model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalyst(params)
	} else {
		params := &model_canal_waterlevel.Param_CanalWaterlevel{Station_id: p.Id, Start_date: p.StartDate, End_date: p.EndDate}
		rs, err = model_canal_waterlevel.GetCanalWaterlevelGraphByStationAndDateAnalyst(params)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/waterlevel_yearly_graph
// @Method			GET
// @Summary			กราฟระดับน้ำรายปี
// @Parameter		-	query paramWaterlevelGraphSwagger2
// @Produces		json
// @Response		200		model_tele_waterlevel.GetWaterlevelYearlyGraphAnalystOutput successful operation
func (srv *HttpService) getWaterlevelYearlyGraphAnalyst(ctx service.RequestContext) error {

	p := &paramWaterlevelGraph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	var err error
	var rs *model_tele_waterlevel.GetWaterlevelYearlyGraphAnalystOutput
	ctx.LogRequestParams(p)
	if p.StationType == "tele_waterlevel" {
		params := &model_tele_waterlevel.GetWaterlevelYearlyGraphInput{StationID: p.Id, Year: p.Year}
		rs, err = model_tele_waterlevel.GetWaterlevelYearlyGraphAnalyst(params)
	} else {
		params := &model_canal_waterlevel.GetCanalWaterlevelYearlyGraphInput{StationID: p.Id, Year: p.Year}
		rs, err = model_canal_waterlevel.GetWaterlevelYearlyGraphAnalyst(params)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/watergate_graph
// @Method			GET
// @Summary			ระดับน้ำประตูระบายน้ำเข้าและระบายน้ำออกสำหรับกราฟ
// @Parameter		-	query model_tele_waterlevel.GetWaterlevelInOutGrapthAnalystInput
// @Produces		json
// @Response		200		WaterlevelInOutGrapthSwagger successful operation
type WaterlevelInOutGrapthSwagger struct {
	Result string                                                        `json:"result"` //example:`OK`
	Data   []model_tele_waterlevel.GetWaterlevelInOutGrapthAnalystOutput `json:"data"`
}

// water gate ปตร./ฝาย
func (srv *HttpService) getWaterlevelInOutGraphAnalyst(ctx service.RequestContext) error {
	p := &model_tele_waterlevel.GetWaterlevelInOutGrapthAnalystInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	rs, err := model_tele_waterlevel.GetWaterlevelInOutGraphAnalyst(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/watergate_load
// @Summary			ระดับน้ำประตูระบายน้ำเข้าและระบายน้ำออก
// @Method			GET
// @Produces		json
// @Response		200		ResultWaterGateOutput successful operation
type ResultWaterGate struct {
	Result string                                                         `json:"result"` //example:`OK`
	Data   []*model_tele_waterlevel.GetWaterlevelInOutLatestAnalystOutput `json:"data"`
}

type ResultWaterGateStation struct {
	Result string                                                             `json:"result"` //example:`OK`
	Data   []*model_tele_station.Struct_WaterlevelCanalStationGroupByProvince `json:"data"`
}

type Struct_OnLoadWaterlevelInOutLastestSwagger struct {
	Data    *ResultWaterGate        `json:"watergate_data"` //ข้อมูลประตูระบายน้ำ
	Station *ResultWaterGateStation `json:"station"`        // สถานีประตูระบายน้ำ
}

type ResultWaterGateOutput struct {
	Result string                                      `json:"result"` //example:`OK`
	Data   *Struct_OnLoadWaterlevelInOutLastestSwagger `json:"data"`
}

// watergate ปตร./ฝาย
func (srv *HttpService) onLoadWaterlevelInOut(ctx service.RequestContext) error {

	objResult := &Struct_OnLoadWaterlevelInOutLastest{}

	//Get weir data
	resultData, err := model_tele_waterlevel.GetWaterlevelInOutLatestAnalyst()
	if err != nil {
		objResult.Data = result.Result0(err)
	} else {
		objResult.Data = result.Result1(resultData)
	}

	//Get weir station / watergate
	//	สำหรับกราฟ
	resultStation, err := model_tele_station.GetWeirStationGroupByProvince()
	if err != nil {
		objResult.Station = result.Result0(err)
	} else {
		objResult.Station = result.Result1(resultStation)
	}

	//=== Province ===//
	//	สำหรับกราฟ
	rsProvince, err := model_geocode.GetAllProvince()
	if err != nil {
		objResult.Province = result.Result0(err)
	} else {
		objResult.Province = result.Result1(rsProvince)
	}

	ctx.ReplyJSON(result.Result1(objResult))
	return nil
}

type paramAdvanceWaterlevelBasinGraph struct {
	SubbasinId int64  `json:"subbasin_id"` // example:`9` รหัสลุ่มน้ำสาขา
	Datetime   string `json:"datetime"`    // example:`2017-08-20 15:00` วันเวลา
}

type Struct_getAdvanceWaterlevelBasinGraph struct {
	Result string                                                             `json:"result"` // example:`OK`
	Data   []*model_tele_waterlevel.Struct_WaterlevelBasinGraphAnalystAdvance `json:"data"`   // ข้อมูลกราฟ
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_waterlevel_basin_graph
// @Summary			กราฟระดับน้ำตามความยาวลำน้า
// @Method			GET
// @Parameter		-	query paramAdvanceWaterlevelBasinGraph
// @Produces		json
// @Response		200	Struct_getAdvanceWaterlevelBasinGraph successful operation
func (srv *HttpService) getAdvanceWaterlevelBasinGraph(ctx service.RequestContext) error {
	p := &paramAdvanceWaterlevelBasinGraph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	if p.SubbasinId <= 0 {
		return rest.NewError(422, "no subbasin id ", nil)
	}
	if p.Datetime == "" {
		return rest.NewError(422, "no datetime", nil)
	}

	rs, err := model_tele_waterlevel.GetWaterlevelBasinGraphAnalystAdvance(p.SubbasinId, p.Datetime)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_getAdvanceWaterlevelBasin24hGraph struct {
	Result string                                                                `json:"result"` // example:`OK`
	Data   []*model_tele_waterlevel.Struct_WaterlevelBasin24HGraphAnalystAdvance `json:"data"`   // ข้อมูลกราฟ
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/advance_waterlevel_basin_24h_graph
// @Summary			กราฟระดับน้ำตามความยาวลำน้าย้อนหลัง 24 ชม.
// @Method			GET
// @Parameter		subbasin_id query int64 example:1  รหัสลุ่มน้ำสาขา
// @Produces		json
// @Response		200	Struct_getAdvanceWaterlevelBasin24hGraph successful operation
func (srv *HttpService) getAdvanceWaterlevelBasin24hGraph(ctx service.RequestContext) error {
	p := &paramAdvanceWaterlevelBasinGraph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	if p.SubbasinId <= 0 {
		return rest.NewError(422, "no subbasin id ", nil)
	}

	rs, err := model_tele_waterlevel.GetWaterlevelBasinGraph24HAnalystAdvance(p.SubbasinId)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/flood_forecast_monitoring
// @Method			GET
// @Summary			ข้อมูลคาดการณ์ระดับน้ำลุ่มน้ำเจ้าพระยา
// @Produces		json
// @Response		200		FloodForecastMonitoringSwagger successful operation
type FloodForecastMonitoringSwagger struct {
	Result string                                   `json:"result"` //example:`OK`
	Data   []model_forecast.FloodForecastMonitoring `json:"data"`
}

func (srv *HttpService) getFloodForecastMonitoring(ctx service.RequestContext) error {

	rs, err := model_forecast.FloodForecastSubbasin()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
