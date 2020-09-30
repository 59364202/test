// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package provinces

import (
	"strconv"
	"time"

	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_rainfall3d "haii.or.th/api/thaiwater30/model/rainfall_3d"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- struct output json
type Struct_rain3d struct {
	Result string                                `json:"result"`
	Data   []*model_rainfall3d.Struct_Rainfall3d `json:"data"` // data
}

// for build api cache
// สร้าง struc เพื่อ clone model_rainfall24hr.Param_Rainfall24 เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
// struc input param
type Param_Rainfall3d struct {
	Param *model_rainfall24hr.Param_Rainfall24
}

func (s *Param_Rainfall3d) IsValid(lastupdate time.Time) bool {
	return false
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_Rainfall3d) GetDescription() string {
	return "refresh on 7.20 and 7.40 every day"
}

// build cache data
func (s *Param_Rainfall3d) BuildData() (interface{}, error) {
	rs := &Struct_rain3d{}
	rs_data, err := model_rainfall3d.GetRainfall3d(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getRain3dGoCache(param *model_rainfall24hr.Param_Rainfall24) ([]byte, time.Time, error) {

	cacheName := "province.rain3d."
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

		c := &Param_Rainfall3d{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "20,40 7 * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/rain3d
// @Summary			ฝนสะสม 3 วัน รายจังหวัด
// @Parameter		region_code	query string example:`1` รหัสภาค ไม่ใส่ = ทุกภาค ,เลือกได้ทีละภาค
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ
// @Method			GET
// @Produces		json
// @Response		200	Struct_rain3d successful operation
func (srv *HttpService) getRain3d(ctx service.RequestContext) error {
	param_Rainfall24 := &model_rainfall24hr.Param_Rainfall24{}
	if err := ctx.GetRequestParams(param_Rainfall24); err != nil {
		return err
	}
	ctx.LogRequestParams(param_Rainfall24)
	// call data with cache
	b, t, err := getRain3dGoCache(param_Rainfall24)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
