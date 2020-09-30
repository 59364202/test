// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

import (
	"time"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	model_humid "haii.or.th/api/thaiwater30/model/humid"
	model_pressure "haii.or.th/api/thaiwater30/model/pressure"
	model_temperature "haii.or.th/api/thaiwater30/model/temperature"
)

// @DocumentName	v1.public
// @Service			thaiwater30//public/weather_graph
// @Method			GET
// @Summary			กราฟสภาพอากาศ (อุณหภูมิ, ความชื้น, ความกดอกากา)
// @Parameter		-	query paramGraphSwagger1
// @Produces		json
// @Response		200		getWeatherGraphByStationAndDateSwagger successful operation

type paramGraphSwagger1 struct {
	Id        string `json:"station_id"` // รหัสสถานี  เช่น 100
	StartDate string `json:"start_date"` // required:false วันเริ่มต้น 2017-06-10 ถ้าไม่กำหนดวันที่เริ่มต้น สิ้นสุด จะ return ข้อมูลย้อนหลัง 3 วันจากวันปัจจุบัน
	EndDate   string `json:"end_date"`   // required:false วันสิ้นสุด 2017-06-11
}

//----------- thaiwater30//public/weather_graph
type s struct {
	Temperature *result.Result `json:"temperature"`
	Humid       *result.Result `json:"humid"`
	Pressure    *result.Result `json:"pressure"`
}

type getWeatherGraphByStationAndDateSwagger struct {
	Result string `json:"result"` //example:`OK`
	Data   *s     `json:"data"`
}

func (srv *HttpService) getWeatherGraphByStationAndDate(ctx service.RequestContext) error {

	p := &paramTemperatureGraphSwagger1{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	var err error
	ctx.LogRequestParams(p)

	now := time.Now()
	if p.StartDate == "" || p.EndDate == "" {
		p.StartDate = now.AddDate(0, 0, -3).Format("2006-01-02")
		p.EndDate = now.Format("2006-01-02")
	}

	rs := &s{}

	params := &model_temperature.Struct_TemperatureGraphParam{Station_id: p.Id, Start_date: p.StartDate, End_date: p.EndDate}
	rs_temperature, err := model_temperature.GetTemperatureGraph(params)
	if err != nil {
		rs.Temperature = result.Result0(err)
	} else {
		rs.Temperature = result.Result1(rs_temperature)
	}

	p1 := &model_humid.Struct_GraphParam{Station_id: p.Id, Start_date: p.StartDate, End_date: p.EndDate}
	rs_humid, err := model_humid.GetGraph(p1)
	if err != nil {
		rs.Humid = result.Result0(err)
	} else {
		rs.Humid = result.Result1(rs_humid)
	}

	p2 := &model_pressure.Struct_GraphParam{Station_id: p.Id, Start_date: p.StartDate, End_date: p.EndDate}
	rs_pressure, err := model_pressure.GetGraph(p2)
	if err != nil {
		rs.Pressure = result.Result0(err)
	} else {
		rs.Pressure = result.Result1(rs_pressure)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
