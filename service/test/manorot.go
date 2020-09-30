package test

import (
	"haii.or.th/api/server/model/datacache"		// cache
	"haii.or.th/api/util/service"
	"time"
	
	model_a "haii.or.th/api/thaiwater30/model/a"
)

// public ใช้สำหรับเรียก service จาก main.go
type Manorot struct {
}

// for build api cache
// สร้าง struc เพื่อ clone model_a.Param_handlerGetProvince เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
type ParamHandlerGetManorotProvinceCache struct {
	Param *model_a.Param_handlerGetManorot
}

// @DocumentName	v1.test
// @Service 		test/manorot
// @Summary 		ทดสอบการทำ ws
// @Parameter		-	query	model_a.Param_handlerGetManorot{region_id}
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 model_a.m_province successful operation
func (srv *Manorot) handlerGetManorot(ctx service.RequestContext) error {
	// get parameter via querystring
	param := &model_a.Param_handlerGetManorot{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return err
	}

	// get parameter via url
	// prov_code := ctx.GetServiceParams("prov_code")

	// call model
	// province, err := model_a.M_Province(prov_code)
	//	province, err := model_a.M_Province(param)
	//
	//	if err != nil {
	//		return err
	//	}

	// call data with cache
	b, t, err := getManorotProvinceCache(param)
	if err != nil {
		return err
	}
	
	// first time will return code 200, second time will return 304 (not modified)
	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}

// ------------ cache ------------
//s *ParamHandlerGetProvinceCache struc เพื่อ clone model_a.Param_handlerGetProvince เพิ่ม function is_valid,description เพื่อใช้ใน function builddata
func (s *ParamHandlerGetManorotProvinceCache) IsValid(lastupdate time.Time) bool {
	return true
}
func (s *ParamHandlerGetManorotProvinceCache) GetDescription() string {
	return "refresh every 5 minutes"
}

// build cache data
func (s *ParamHandlerGetManorotProvinceCache) BuildData() (interface{}, error) {
	province, err := model_a.M_Province(s.Param)
	if err != nil {
		return nil, err
	} else {
		return province, err
	}
}

func getManorotProvinceCache(param *model_a.Param_handlerGetManorot) ([]byte, time.Time, error) {
	//	cname := "test.manorot.provinceId_" + param.Province_id // ชื่อของแคช
	cname := "test.manorot.province" // ชื่อของแคช
	
	if !datacache.IsRegistered(cname) {

		c := &ParamHandlerGetManorotProvinceCache{}
		c.Param = param

		// refresh cache in every 5 minute
		datacache.RegisterDataCache(cname, c, []string{""}, c, "*/5 * * * *")
	}

	// data cache แบบ ไม่ zip json result ไม่ได้ใช้แล้ว เพราะใช้ gzip function จะทำให้ return json ได้เล็กกว่า
	//	datacache.GetJSON(cname)

	// ถ้าจพ return datacache.GetGZJSON ต้องใส่ time มาด้วย
	return datacache.GetGZJSON(cname)
}
// ------------ end cache ------------