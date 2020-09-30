// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package provinces

import (
	"strconv"
	"time"

	model_temperature "haii.or.th/api/thaiwater30/model/temperature"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- struct output json
type Struct_maxMinTemperature struct {
	Result string                                        `json:"result"`
	Data   []*model_temperature.Struct_MaxMinTemperature `json:"data"` // data
}

// for build api cache
// สร้าง struc เพื่อ clone model_temperature.Param_TemperatureProvinces เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
// struc input param
type Param_TemperatureProvinces struct {
	Param *model_temperature.Param_TemperatureProvinces
}

func (s *Param_TemperatureProvinces) IsValid(lastupdate time.Time) bool {
	return false
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_TemperatureProvinces) GetDescription() string {
	return "refresh every hour"
}

// build cache data
func (s *Param_TemperatureProvinces) BuildData() (interface{}, error) {
	rs := &Struct_maxMinTemperature{}
	rs_data, err := model_temperature.GetMaxMinTemperatureThisWeek(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getMaxMinTemperatureThisWeekGoCache(param *model_temperature.Param_TemperatureProvinces) ([]byte, time.Time, error) {

	cacheName := "province.TemperatureMaxMinThisWeek."
	if param.Region_Code != "" {
		cacheName += "regionCode=" + param.Region_Code
	}
	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}
	if param.Data_Limit > 0 {
		cacheName += "datalimit=" + strconv.Itoa(param.Data_Limit)
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_TemperatureProvinces{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "7 * * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/temperature_maxmin_thisweek
// @Summary			อุณหภูมิสูงสุด ต่ำสุดในรอบสัปดาห์ รายจังหวัด
// @Parameter		region_code	query string example:`1` รหัสภาค ไม่ใส่ = ทุกภาค ,เลือกได้ทีละภาค
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ ค่าพื้นฐาน = 1
// @Method			GET
// @Produces		json
// @Response		200	Struct_maxMinTemperature successful operation
func (srv *HttpService) getMaxMinTemperatureThisWeek(ctx service.RequestContext) error {
	param_TemperatureProvinces := &model_temperature.Param_TemperatureProvinces{}
	if err := ctx.GetRequestParams(param_TemperatureProvinces); err != nil {
		return err
	}
	ctx.LogRequestParams(param_TemperatureProvinces)
	// call data with cache
	b, t, err := getMaxMinTemperatureThisWeekGoCache(param_TemperatureProvinces)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
