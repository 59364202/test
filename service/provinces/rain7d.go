// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package provinces

import (
	"time"

	model_rainfall7d "haii.or.th/api/thaiwater30/model/rainfall_7d"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- struct output json
type Struct_rain7d struct {
	Result string                                `json:"result"`
	Data   []*model_rainfall7d.Struct_Rainfall7d `json:"data"` // data
}

// for build api cache
// struc input param
type Param_Rainfall7d struct {
	Param *model_rainfall7d.Param_Rainfall7d
}

func (s *Param_Rainfall7d) IsValid(lastupdate time.Time) bool {
	return true
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_Rainfall7d) GetDescription() string {
	return "refresh after 20 hour"
}

// build cache data
func (s *Param_Rainfall7d) BuildData() (interface{}, error) {
	rs := &Struct_rain7d{}
	rs_data, err := model_rainfall7d.GetRainfall7d(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getRain7dGoCache(param *model_rainfall7d.Param_Rainfall7d) ([]byte, time.Time, error) {

	cacheName := "province.rain7d."

	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_Rainfall7d{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "20,40 7 * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/rain7d
// @Summary			ฝนสะสม 7 วัน รายจังหวัด
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Method			GET
// @Produces		json
// @Response		200	Struct_rain7d successful operation
func (srv *HttpService) getRain7d(ctx service.RequestContext) error {
	param_Rainfall7d := &model_rainfall7d.Param_Rainfall7d{}
	if err := ctx.GetRequestParams(param_Rainfall7d); err != nil {
		return err
	}
	ctx.LogRequestParams(param_Rainfall7d)
	// call data with cache
	b, t, err := getRain7dGoCache(param_Rainfall7d)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
