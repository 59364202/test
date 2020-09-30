// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

import (
	"time"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	model_temperature "haii.or.th/api/thaiwater30/model/temperature"
)

// @DocumentName	v1.public
// @Service			thaiwater30//public/temperature_graph
// @Method			GET
// @Summary			กราฟระดับอุณหภูมิ
// @Parameter		-	query paramTemperatureGraphSwagger1
// @Produces		json
// @Response		200		GetTemperatureGraphByStationAndDateSwagger successful operation

type paramTemperatureGraphSwagger1 struct {
	Id        string `json:"station_id"` // รหัสสถานี  เช่น 100
	StartDate string `json:"start_date"` // required:false วันเริ่มต้น 2017-06-10 ถ้าไม่กำหนดวันที่เริ่มต้น สิ้นสุด จะ return ข้อมูลย้อนหลัง 3 วันจากวันปัจจุบัน
	EndDate   string `json:"end_date"`   // required:false วันสิ้นสุด 2017-06-11
}

type GetTemperatureGraphByStationAndDateSwagger struct {
	Result string                                      `json:"result"` //example:`OK`
	Data   []model_temperature.Struct_TemperatureGraph `json:"data"`
}

func (srv *HttpService) getTemperatureGraphByStationAndDate(ctx service.RequestContext) error {

	p := &paramTemperatureGraphSwagger1{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	var err error
	var rs *model_temperature.Struct_TemperatureGraph

	ctx.LogRequestParams(p)
	now := time.Now()
	if p.StartDate == "" || p.EndDate == "" {
		p.StartDate = now.AddDate(0, 0, -3).Format("2006-01-02")
		p.EndDate = now.Format("2006-01-02")
	}

	params := &model_temperature.Struct_TemperatureGraphParam{Station_id: p.Id, Start_date: p.StartDate, End_date: p.EndDate}
	rs, err = model_temperature.GetTemperatureGraph(params)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
