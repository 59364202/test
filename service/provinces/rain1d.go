// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package provinces

import (
	"time"

	model_rainfall1d "haii.or.th/api/thaiwater30/model/rainfall_1d"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- struct output json
type Struct_rain1d struct {
	Result string                                `json:"result"`
	Data   []*model_rainfall1d.Struct_Rainfall1d `json:"data"` // data
}

// for build api cache
// struc input param
type Param_Rainfall1d struct {
	Param *model_rainfall1d.Param_Rainfall1d
}

func (s *Param_Rainfall1d) IsValid(lastupdate time.Time) bool {
	return true
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_Rainfall1d) GetDescription() string {
	return "refresh after 20 hour"
}

// build cache data
func (s *Param_Rainfall1d) BuildData() (interface{}, error) {
	rs := &Struct_rain1d{}
	rs_data, err := model_rainfall1d.GetRainfall1d(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getRain1dGoCache(param *model_rainfall1d.Param_Rainfall1d) ([]byte, time.Time, error) {

	cacheName := "province.rain1d."

	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_Rainfall1d{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "20,40 7 * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/rain1d
// @Summary			ฝนสะสมรายวัน รายจังหวัด
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Method			GET
// @Produces		json
// @Response		200	Struct_rain1d successful operation
func (srv *HttpService) getRain1d(ctx service.RequestContext) error {
	param_Rainfall1d := &model_rainfall1d.Param_Rainfall1d{}
	if err := ctx.GetRequestParams(param_Rainfall1d); err != nil {
		return err
	}
	ctx.LogRequestParams(param_Rainfall1d)
	// call data with cache
	b, t, err := getRain1dGoCache(param_Rainfall1d)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
