// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

import (
	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/util/service"
	"time"

	model_humid "haii.or.th/api/thaiwater30/model/humid"
	model_pressure "haii.or.th/api/thaiwater30/model/pressure"
	model_rainfall_1h "haii.or.th/api/thaiwater30/model/rainfall_1h"
	model_temperature "haii.or.th/api/thaiwater30/model/temperature"
	model_wind "haii.or.th/api/thaiwater30/model/wind"

	result "haii.or.th/api/thaiwater30/util/result"
)

// ----------------------------------------------------------------------------------------------------------------------
// @DocumentName	v1.public
// @Service			thaiwater30/public/thaiwater_main
// @Summary			thaiwater หน้าหลัก
// @Description
// @				* thaiwater หน้าหลัก
// @Method			GET
// @Produces		json
// @Response		200		Index successful operation
func (srv *HttpService) getThaiwaterMain(ctx service.RequestContext) error {
	return replyWithHourlyCache(ctx, thaiwaterMainCacheName, getThaiwaterMainCacheBuildData)
}

// ----------------------------------------------------------------------------------------------------------------------
// @DocumentName	v1.public
// @Service			thaiwater30/public/thaiwater/weather
// @Summary			thaiwater หน้าสภาพอากาศ
// @Parameter		region_code	query string example:`1` รหัสภาค ไม่ใส่ = ทุกภาค ,เลือกได้ทีละภาค
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ
// @Parameter		order	query string example:`DESC` การจัดเรียงข้อมูล ASC น้อยไปหามาก หรือ DESC มากไปหาน้อย
// @Description
// @				* thaiwater หน้าสภาพอากาศ
// @Method			GET
// @Produces		json
// @Response		200		ThaiwaterWeatherStruct successful operation
func (srv *HttpService) getThaiwaterWeather(ctx service.RequestContext) error {
	rs := &ThaiwaterWeatherStruct{}

	// max wind
	paramWind := &model_wind.Param_Provinces{}
	if err := ctx.GetRequestParams(paramWind); err != nil {
		return err
	}
	ctx.LogRequestParams(paramWind)
	paramWind.Order = "DESC"
	paramWind.Data_Limit = 1
	rs.Wind = buildResult(model_wind.GetWind(paramWind))

	// max humid
	paramHumid := &model_humid.Param_Provinces{}
	if err := ctx.GetRequestParams(paramHumid); err != nil {
		return err
	}
	ctx.LogRequestParams(paramHumid)
	paramHumid.Order = "DESC"
	paramHumid.Data_Limit = 1
	rs.Humid = buildResult(model_humid.GetHumid(paramHumid))

	//max pressure
	paramPressure := &model_pressure.Param_Provinces{}
	if err := ctx.GetRequestParams(paramPressure); err != nil {
		return err
	}
	ctx.LogRequestParams(paramPressure)
	paramPressure.Order = "DESC"
	paramPressure.Data_Limit = 1
	rs.Pressure = buildResult(model_pressure.GetPressure(paramPressure))

	//max,min temperature
	paramTemperature := &model_temperature.Param_TemperatureProvinces{}
	if err := ctx.GetRequestParams(paramTemperature); err != nil {
		return err
	}
	ctx.LogRequestParams(paramTemperature)
	rs.Temperature = buildResult(model_temperature.GetMaxMinTemperature(paramTemperature))

	//max,min rainfall 1h
	paramRainfall := &model_rainfall_1h.Param_Provinces{}
	if err := ctx.GetRequestParams(paramRainfall); err != nil {
		return err
	}
	ctx.LogRequestParams(paramRainfall)
	rs.Rain = buildResult(model_rainfall_1h.GetMaxMinRainfall(paramRainfall))

	ctx.ReplyJSON(rs)
	return nil
}

// ----------------------------------------------------------------------------------------------------------------------
// @DocumentName	v1.public
// @Service			thaiwater30/public/thaiwater/temperature
// @Summary			thaiwater หน้าสภาพอากาศ
// @Parameter		region_code	query string example:`1` รหัสภาค ไม่ใส่ = ทุกภาค ,เลือกได้ทีละภาค
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ
// @Parameter		order	query string example:`DESC` การจัดเรียงข้อมูล ASC น้อยไปหามาก หรือ DESC มากไปหาน้อย
// @Description
// @				* thaiwater หน้าสภาพอากาศ อุณหภูมิ
// @Method			GET
// @Produces		json
// @Response		200		ThailandModuleStruct successful operation

// for build api cache
// สร้าง struc เพื่อ clone model_temperature.Param_TemperatureProvinces เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
type Param_Temperature struct {
	Param *model_temperature.Param_TemperatureProvinces
}

//สร้าง func s *Param_Temperature struc เพื่อ clone model_rainfall24hr.Param_Rainfall24 เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
func (s *Param_Temperature) IsValid(lastupdate time.Time) bool {
	return true
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_Temperature) GetDescription() string {
	return "refresh every 1 hour"
}

// build cache data
func (s *Param_Temperature) BuildData() (interface{}, error) {
	rs := &result.Result{}
	rs = buildResult(model_temperature.GetTemperatureLatest(s.Param))
	return rs, nil
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getTemperatureCache(param *model_temperature.Param_TemperatureProvinces) ([]byte, time.Time, error) {

	cacheName := "public.thaiwater.temperature"
	if param.Region_Code != "" {
		cacheName += "regionCode=" + param.Region_Code
	}
	if param.Region_Code_tmd != "" {
		cacheName += "regionCodeTmd=" + param.Region_Code_tmd
	}
	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}
	if param.Data_Limit > 0 {
		cacheName += "datalimit=" + string(param.Data_Limit)
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_Temperature{}
		c.Param = param

		// refresh cache in every 1 hour
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "11 * * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

func (srv *HttpService) getThaiwaterTemperature(ctx service.RequestContext) error {
	paramTemperature := &model_temperature.Param_TemperatureProvinces{}
	if err := ctx.GetRequestParams(paramTemperature); err != nil {
		return err
	}
	// call data with cache
	b, t, err := getTemperatureCache(paramTemperature)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
