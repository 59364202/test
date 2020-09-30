// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

//http://api2.thaiwater.net:9200/api/v1/test/thitiporn

// package for service
// training service  province
package test

import (
	"haii.or.th/api/server/model/datacache"
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	"time"

	model_a "haii.or.th/api/thaiwater30/model/a"
)

// public ใช้สำหรับเรียก service จาก main.go
type ThitipornServiceProvinceStruc struct {
}

// for build api cache
// สร้าง struc เพื่อ clone model_a.Param_handlerGetProvince เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
type ParamHandlerGetProvinceCache struct {
	Param *model_a.Param_handlerGetProvince
}

//service call model
// private

//`` ต้องการใส่ตัวอักษรหลายบรรทัด

// @DocumentName	v1.test
// @Service 		test/thitiporn
// @Summary 		training service ข้อมูลภาค จังหวัด by thitiporn

// สำหรับ ใส่รายละเอียด paramter เอง
// Parameter		region_id	query	int64	example:`5` required:true รหัสภาค
// Parameter		province_id	query	string	example : 81   รหัสจังหวัด

//แสดง document param ทั้งหมดใน struct
// Parameter		- query model_a.Param_handlerGetProvince

// แสดง document param เฉพาะ param ที่เลือก
// @Parameter		- query model_a.Param_handlerGetProvince{region_id}

// Parameter		- query {model.paramStruct}

// @Method			GET
// @Response		404	-	service not found
// @Response		422	-	invalid parameter
// @Response		200 model_a.thitipornProvinceStruc successful operation

//Service 		service url
// Response		200 {model.struct} successful operation
//func (parameter) function_name (return)
func (srv *ThitipornServiceProvinceStruc) handlerGetProvinceThitiporn(ctx service.RequestContext) error {
	// get parameter via url
	//	region_code := ctx.GetServiceParams("region_id")
	//	prov_code := ctx.GetServiceParams("province_id")

	// return error when call api with require field
	//	if prov_code == "" {
	//		return rest.NewError(422, "invalid province code", nil)
	//	}

	//get param from struct in model
	param := &model_a.Param_handlerGetProvince{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return err
	}

	// call model
	//	input each parameter
	//	province, err := model_a.Thitiporn_Province(prov_code)

	//	 input multiple parameter by interface
	//	province, err := model_a.Thitiporn_Province(param)
	//	if err != nil {
	//		ctx.ReplyError(err)
	//	} else {
	//		// reply json
	//		ctx.ReplyJSON(province)
	//	}

	// call data with cache
	b, t, err := getThitipornProvinceCache(param)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}

//------------ api cache
//s *ParamHandlerGetProvinceCache struc เพื่อ clone model_a.Param_handlerGetProvince เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
func (s *ParamHandlerGetProvinceCache) IsValid(lastupdate time.Time) bool {
	return false
}
func (s *ParamHandlerGetProvinceCache) GetDescription() string {
	return "refresh every 5 minute"
}

// build cache data
func (s *ParamHandlerGetProvinceCache) BuildData() (interface{}, error) {
	province, err := model_a.Thitiporn_Province(s.Param)
	if err != nil {
		return nil, err
	} else {
		return province, err
	}
}

func getThitipornProvinceCache(param *model_a.Param_handlerGetProvince) ([]byte, time.Time, error) {
	//	cname := "test.thitiporn.provinceId_" + param.Province_id // ชื่อของแคช
	cname := "test.thitiporn.province" // ชื่อของแคช
	if !datacache.IsRegistered(cname) {

		c := &ParamHandlerGetProvinceCache{}
		c.Param = param

		// refresh cache in every 5 minute
		datacache.RegisterDataCache(cname, c, []string{""}, c, "*/5 * * * *")
	}

	// data cache แบบ ไม่ zip json result ไม่ได้ใช้แล้ว เพราะใช้ gzip function จะทำให้ return json ได้เล็กกว่า
	//	datacache.GetJSON(cname)

	// ถ้าจะ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cname)
}
