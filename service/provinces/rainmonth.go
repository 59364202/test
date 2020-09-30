// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Werawan Prongpanom <werawan@haii.or.th>

package provinces

import (
	"time"

	model_rainfallmonth "haii.or.th/api/thaiwater30/model/rainfall_month"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// struct output json
type Struct_RainfallMonth struct {
	Result string                                      `json:"result"`
	Data   []*model_rainfallmonth.Struct_RainfallMonth `json:"data"`
}

// for build api cache
// struc input param
type Param_RainfallMonth struct {
	Param *model_rainfallmonth.Param_RainfallMonth
}

func (s *Param_RainfallMonth) IsValid(lastupdate time.Time) bool {
	return true
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_RainfallMonth) GetDescription() string {
	return "refresh after 20 hour"
}

// build cache data
func (s *Param_RainfallMonth) BuildData() (interface{}, error) {
	rs := &Struct_RainfallMonth{}
	rs_data, err := model_rainfallmonth.GetRainfallMonth(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		rs.Result = "OK"
		rs.Data = rs_data
		return rs, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getRainMonthGoCache(param *model_rainfallmonth.Param_RainfallMonth) ([]byte, time.Time, error) {

	cacheName := "province.rainmonth."

	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_RainfallMonth{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "20,40 7 * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/rainfall_month
// @Summary			ฝนสะสมรายเดือนที่ผ่านมา
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Method			GET
// @Produces		json
// @Response		200	Struct_RainfallMonth successful operation
func (srv *HttpService) GetRainfallMonth(ctx service.RequestContext) error {
	param_RainfallMonth := &model_rainfallmonth.Param_RainfallMonth{}
	if err := ctx.GetRequestParams(param_RainfallMonth); err != nil {
		return err
	}
	ctx.LogRequestParams(param_RainfallMonth)
	// call data with cache
	b, t, err := getRainMonthGoCache(param_RainfallMonth)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
