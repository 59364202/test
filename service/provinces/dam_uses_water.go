// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package provinces

import (
	"time"

	model_dam_uses_water "haii.or.th/api/thaiwater30/model/dam_uses_water"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// -------------------- struct output json
type Struct_damUsesWater struct {
	Result string                                                `json:"result"`
	Data   *model_dam_uses_water.Struct_DamUsesWater_OutputParam `json:"data"` // data
}

// for build api cache
// struc input param
type Param_DamUsesWater struct {
	Param *model_dam_uses_water.Struct_DamUsesWater_InputParam
}

func (s *Param_DamUsesWater) IsValid(lastupdate time.Time) bool {
	return true
}

//สร้าง func GetDescription เพื่อ ใส่ รายละเอียดการตั้งเวลา refresh cahce ที่กำหนดไว้
func (s *Param_DamUsesWater) GetDescription() string {
	return "refresh after 20 hour"
}

// build cache data
func (s *Param_DamUsesWater) BuildData() (interface{}, error) {
	//rs := &Struct_damUsesWater{}
	rs_data, err := model_dam_uses_water.GetDamUsesWater(s.Param)
	if err != nil {
		return result.Result0(err), nil
	} else {
		//		rs.Result = "OK"
		//		rs.Data = rs_data
		return rs_data, err
	}
}

//สร้าง function กำหนดชื่อ cache และตั้งเวลา
func getDamUsesWaterGoCache(param *model_dam_uses_water.Struct_DamUsesWater_InputParam) ([]byte, time.Time, error) {

	cacheName := "province.damUsesWater."

	if param.Province_Code != "" {
		cacheName += "provinceCode=" + param.Province_Code
	}

	if param.Dam_id != "" {
		cacheName += "damId=" + param.Dam_id
	}

	if param.Start_date != "" {
		cacheName += "startDate=" + param.Start_date
	}

	if param.End_date != "" {
		cacheName += "endDate=" + param.End_date
	}

	if !datacache.IsRegistered(cacheName) {

		c := &Param_DamUsesWater{}
		c.Param = param

		// refresh cache after calculate rainfall daily
		datacache.RegisterDataCache(cacheName, c, []string{""}, c, "20,40 8 * * *")
	}

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cacheName)
}

// add comment below for api-doc
// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/dam_uses_water
// @Summary			ปริมาณน้ำใช้งานได้รวมทุกเขื่อน รายจังหวัด/รายเขื่อน
// @Parameter		dam_id	query string example:`1,2,3,4` รหัสเขื่อน ไม่ใส่ = ทุกเขื่อน ,เลือกได้หลายเขื่อน เช่น 1,2,3,4
// @Parameter		province_code	query string example:`10,51,62` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		start_date	query string example:`2018-01-01` วันที่เริ่มเต้น ไม่ใส่ = ปีที่ผ่านมา
// @Parameter		end_date	query string example:`2018-01-01` วันที่สิ้นสุด ไม่ใส่ = ปีที่ผ่านมา
// @Method			GET
// @Produces		json
// @Response		200	Struct_damUsesWater successful operation
func (srv *HttpService) getDamUsesWater(ctx service.RequestContext) error {
	param_DamUsesWater := &model_dam_uses_water.Struct_DamUsesWater_InputParam{}
	if err := ctx.GetRequestParams(param_DamUsesWater); err != nil {
		return err
	}
	ctx.LogRequestParams(param_DamUsesWater)
	// call data with cache
	b, t, err := getDamUsesWaterGoCache(param_DamUsesWater)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
