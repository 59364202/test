// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package provinces

import (
	"encoding/json"
	"strconv"
	"time"

	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- rain24 output struc
type Struct_rain24Hr struct {
	Result string                                   `json:"result"`
	Data   []*model_rainfall24hr.Struct_Rainfall24H `json:"data"`  // data
	Scale  json.RawMessage                          `json:"scale"` // setting
}

// for build api cache
// สร้าง struc เพื่อ clone model_rainfall24hr.Param_Rainfall24 เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
type Param_Rainfall24 struct {
	Param *model_rainfall24hr.Param_Rainfall24
}

type hourlyCacheValidator struct {
}

//สร้าง func s *Param_Rainfall24 struc เพื่อ clone model_rainfall24hr.Param_Rainfall24 เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
func (s *Param_Rainfall24) IsValid(lastupdate time.Time) bool {
	return false
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_Rainfall24) GetDescription() string {
	return "refresh every 1 hour"
}

// build cache data
func (s *Param_Rainfall24) BuildData() (interface{}, error) {
	rs := &Struct_rain24Hr{}
	rs_data, err := model_rainfall24hr.GetRainfallThailandDataCache(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		rs.Scale = setting.GetSystemSettingJson("Frontend.public.rain_setting")
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getRain24GoCache(param *model_rainfall24hr.Param_Rainfall24) ([]byte, time.Time, error) {

	cacheName := "province.rain24."
	if param.Region_Code != "" {
		cacheName += "regionCode=" + param.Region_Code
	}
	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}
	if param.Include_zero != "" {
		cacheName += "includeZero=" + param.Include_zero
	}
	if param.Data_Limit > 0 {
		cacheName += "datalimit=" + strconv.Itoa(param.Data_Limit)
	}
	if param.Basin_id != "" {
		cacheName += "basinid=" + param.Basin_id
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_Rainfall24{}
		c.Param = param

		// refresh cache in every 1 hour
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "16 * * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/rain24
// @Summary			ฝน24ชม. รายจังหวัด
// @Parameter		region_code	query string example:`1` รหัสภาค ไม่ใส่ = ทุกภาค ,เลือกได้ทีละภาค
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ
// @Parameter		include_zero	query string example:`1` ต้องการให้แสดงข้อมูลสถานนีที่มีค่าฝนเป็น 0 ด้วยหรือไม่  include_zero=1 แสดง
// @Method			GET
// @Produces		json
// @Response		200	Struct_rain24Hr successful operation
func (srv *HttpService) getRain24(ctx service.RequestContext) error {
	param_Rainfall24 := &model_rainfall24hr.Param_Rainfall24{}
	if err := ctx.GetRequestParams(param_Rainfall24); err != nil {
		return err
	}
	ctx.LogRequestParams(param_Rainfall24)
	// call data with cache
	b, t, err := getRain24GoCache(param_Rainfall24)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
