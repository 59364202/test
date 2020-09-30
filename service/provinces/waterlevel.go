// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package provinces

import (
	"encoding/json"
	"strconv"
	"time"

	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

type Struct_waterlevel struct {
	Result string                                     `json:"result"`
	Data   []*model_tele_waterlevel.Struct_Waterlevel `json:"data"`  // data
	Scale  json.RawMessage                            `json:"scale"` // setting
}

// for build api cache
// สร้าง struc เพื่อ clone model_tele_waterlevel.Param_waterlevel เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
type Param_waterlevel struct {
	Param *model_tele_waterlevel.Waterlevel_InputParam
}

//สร้าง func s *Param_waterlevel struc เพื่อ clone model_tele_waterlevel.Param_waterlevel เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
func (s *Param_waterlevel) IsValid(lastupdate time.Time) bool {
	return false
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_waterlevel) GetDescription() string {
	return "refresh every 1 hour"
}

// build cache data
func (s *Param_waterlevel) BuildData() (interface{}, error) {
	rs := &Struct_waterlevel{}
	rs_data, err := model_tele_waterlevel.GetWaterLevelThailandDataCache(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		rs.Scale = setting.GetSystemSettingJson("Frontend.public.waterlevel_setting")
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getWaterlevelGoCache(param *model_tele_waterlevel.Waterlevel_InputParam) ([]byte, time.Time, error) {

	cacheName := "province.waterlevel."
	if param.Region_Code != "" {
		cacheName += "regionCode=" + param.Region_Code
	}
	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}
	if param.Data_Limit > 0 {
		cacheName += "datalimit=" + strconv.Itoa(param.Data_Limit)
	}
	if param.Basin_id != "" {
		cacheName += "basinid=" + param.Basin_id
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_waterlevel{}
		c.Param = param

		// refresh cache in every 1 hour
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "14 * * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/waterlevel
// @Summary			ระดับน้ำ  รายจังหวัด
// @Parameter		region_code	query string example:`1` รหัสภาค ไม่ใส่ = ทุกภาค ,เลือกได้ทีละภาค
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		order	query string example:`asc` การเรียงข้อมูล storage_percent asc= เรียงจากค่าน้อยไปหามาก desc = เรียงจากค่ามากไปหาน้อย
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ
// @Method			GET
// @Produces		json
// @Response		200	Struct_waterlevel successful operation
func (srv *HttpService) getWaterlevel(ctx service.RequestContext) error {
	Param_waterlevel := &model_tele_waterlevel.Waterlevel_InputParam{}
	if err := ctx.GetRequestParams(Param_waterlevel); err != nil {
		return err
	}
	ctx.LogRequestParams(Param_waterlevel)
	// call data with cache
	b, t, err := getWaterlevelGoCache(Param_waterlevel)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
