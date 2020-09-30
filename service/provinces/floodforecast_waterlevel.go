// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package provinces

import (
	model_forecast "haii.or.th/api/thaiwater30/model/floodforecast_waterlevel"
	"time"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- struct output json
type Struct_floodForecastWaterlevel struct {
	Result string                                       `json:"result"`
	Data   *model_forecast.FloodforecastOutputWithScale `json:"data"` // data
}

// for build api cache
// struc input param
type Param_Forecast struct {
	Param *model_forecast.FloodforecastInput
}

func (s *Param_Forecast) IsValid(lastupdate time.Time) bool {
	return true
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_Forecast) GetDescription() string {
	return "refresh after 20 hour"
}

// build cache data
func (s *Param_Forecast) BuildData() (interface{}, error) {
	rs := &Struct_floodForecastWaterlevel{}
	rs_data, err := model_forecast.CpyFloodForecastLatest(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		return rs_data, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getFloodForecastWaterlevelGoCache(param *model_forecast.FloodforecastInput) ([]byte, time.Time, error) {

	cacheName := "province.floodforecast_waterlevel."

	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_Forecast{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "20,40 7 * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/floodforecast_waterlevel
// @Summary			คาดการณ์ระดับน้ำ รายจังหวัด
// @Parameter		province_code	query string example:`10,51,62` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Method			GET
// @Produces		json
// @Response		200	Struct_damUsesWater successful operation
func (srv *HttpService) getFloodForecastWaterlevel(ctx service.RequestContext) error {
	param_FloodForecast := &model_forecast.FloodforecastInput{}
	if err := ctx.GetRequestParams(param_FloodForecast); err != nil {
		return err
	}
	ctx.LogRequestParams(param_FloodForecast)
	// call data with cache
	b, t, err := getFloodForecastWaterlevelGoCache(param_FloodForecast)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
