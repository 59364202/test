package analyst

import (
	"haii.or.th/api/server/model/setting"
	model_sea_station "haii.or.th/api/thaiwater30/model/sea_station"
	model_sea_waterlevel "haii.or.th/api/thaiwater30/model/sea_waterlevel"
	"haii.or.th/api/thaiwater30/util/result"
	tw30setting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/service"
)

type SeaWaterlevelAnalystLatest struct {
	SeaWaterlevel  *result.Result `json:"sea_waterlevel"`
	Scale          *result.Result `json:"scale"`
	SeaAgency      *result.Result `json:"sea_agency"`
	ForecastAgency *result.Result `json:"forecast_agency"`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/sea_waterlevel_load
// @Method			GET
// @Summary			ข้อมูลระดับน้ำทะเลคาดการณ์และตรวจวัดจริง
// @Produces		json
// @Response		200		SeaWaterlevelAnalystLatestSwagger successful operation
type SeaWaterlevelLatestSwagger struct {
	Result string                                            `json:"result"` //example:`OK`
	Data   []model_sea_waterlevel.SeaWaterlevelLatestSwagger `json:"data"`
}

type SeaWaterlevelAnalystLatestSwagger struct {
	SeaWaterlevel  *SeaWaterlevelLatestSwagger `json:"sea_waterlevel"`
	Scale          *SeaSettingSwagger          `json:"scale"`
	SeaAgency      *SeaAgencySwagger           `json:"sea_agency"`
	ForecastAgency *SeaAgencySwagger           `json:"forecast_agency"`
}

type SeaSettingSwagger struct {
	Result string                            `json:"result"` //example:`OK`
	Data   []tw30setting.StructSeaWaterlevel `json:"data"`
}
type SeaAgencySwagger struct {
	Result string                        `json:"result"` //example:`OK`
	Data   []model_sea_waterlevel.Agency `json:"data"`
}

func (srv *HttpService) getSeaWaterlevelAnalyst(ctx service.RequestContext) error {
	rs := &SeaWaterlevelAnalystLatest{}
	var err error
	seaWaterlevel, err := model_sea_waterlevel.GetSeaWaterlevel()
	if err != nil {
		rs.SeaWaterlevel = result.Result0(err.Error())
	} else {
		rs.SeaWaterlevel = result.Result1(seaWaterlevel)
	}
	scale := setting.GetSystemSettingJSON("Frontend.analyst.sea_waterlevel")

	if err != nil {
		rs.Scale = result.Result0(err.Error())
	} else {
		rs.Scale = result.Result1(&scale)
	}
	seaAgency, err := model_sea_station.SeaStationByAgency("observe")
	if err != nil {
		rs.SeaAgency = result.Result0(err.Error())
	} else {
		rs.SeaAgency = result.Result1(seaAgency)
	}
	forecastAgency, err := model_sea_station.SeaStationByAgency("forecast")
	if err != nil {
		rs.ForecastAgency = result.Result0(err.Error())
	} else {
		rs.ForecastAgency = result.Result1(forecastAgency)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/sea_waterlevel_graph
// @Method			GET
// @Summary			ข้อมูลระดับน้ำทะเลตรวจวัดจริง
// @Parameter		-	query model_sea_waterlevel.SeaWaterlevelInput
// @Produces		json
// @Response		200		SeaWaterlevelOutputSwagger successful operation
type SeaWaterlevelOutputSwagger struct {
	Result string                                     `json:"result"` //example:`OK`
	Data   []model_sea_waterlevel.SeaWaterlevelOutput `json:"data"`
}

func (srv *HttpService) getSeaWaterlevelRealAnalyst(ctx service.RequestContext) error {

	p := &model_sea_waterlevel.SeaWaterlevelInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	rs, err := model_sea_waterlevel.GetSeaWaterlevelReal(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/sea_waterlevel_forecast_graph
// @Method			GET
// @Summary			ข้อมูลคาดการณ์ระดับน้ำทะเล
// @Parameter		-	query SeaWaterlevelForecastInput
// @Produces		json
// @Response		200		model_sea_waterlevel.SeaWaterlevelOutput successful operation
type SeaWaterlevelForecastInput struct {
	StationID []int64 `json:"station_id"` // รหัสสถานีระดับน้ำทะเล [23,413,243]
	StartDate string  `json:"start_date"` // วันที่เริ่มต้น 2006-01-02
	EndDate   string  `json:"end_date"`   // วันที่สิ้นสุด 2006-01-02
}

func (srv *HttpService) getSeaWaterlevelForecastAnalyst(ctx service.RequestContext) error {

	p := &model_sea_waterlevel.SeaWaterlevelInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	rs, err := model_sea_waterlevel.GetSeaWaterlevelForecast(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
