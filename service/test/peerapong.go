package test

import (
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	"haii.or.th/api/server/model/datacache"

	"time"

	model_a "haii.or.th/api/thaiwater30/model/a"
)

type Peerapong struct {
}

// ตัวอย่าง json result
type TestPeerapongResult struct {
	Province_id string    `json:"province_id"` // example:`81` required:false รหัสจังหวัดนะ
	Region_id   string    `json:"region_id"`   // example:`5` required:false มันเป็นรหัสภาคนะ
	Time        time.Time `json:"time"`
}

// ตัวอย่าง struct ของ cache
type TestPeerapongCache struct {
	Param *model_a.Param_handlerGetCim // เพิ่ม Param มาเพื่อเก็บ param จาก url
}

//  ตัวอย่าง func เช็คความใหม่ของข้อมูล
//  บังคับชื่อ
//	*require*
func (s *TestPeerapongCache) IsValid(lastupdate time.Time) bool {
	return true
}

//  คัวอย่าง Description
//  บังคับชื่อ
//	*require*
func (s *TestPeerapongCache) GetDescription() string {
	return "refresh every 5 minute"
}

//	ตัวอย่าง สร้างข้อมูล
//  บังคับชื่อ
//	*require*
func (s *TestPeerapongCache) BuildData() (interface{}, error) {
	//	province, err := model_a.Cim_Province(s.Param)
	//	if err != nil {
	//		return nil, err
	//	} else {
	//		return province, err
	//	}

	return &TestPeerapongResult{Province_id: s.Param.Province_id, Region_id: s.Param.Region_id, Time: time.Now()}, nil
}

// @DocumentName	v1.test
// @Module			test/peerapong
// @Description		ระบบให้บริการข้อมูล Api Services

// @DocumentName	v1.test
// @Service 		test/peerapong
// @Summary 		ทดสอบ service
// @Parameter		-	query	model_a.Param_handlerGetCim{region_id}
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 model_a.cim_province successful operation
func (srv *Peerapong) handlerGetPeerapong(ctx service.RequestContext) error {

	//	prov_code := ctx.GetServiceParams("a")
	//	if prov_code == "" {
	//		err := rest.NewError(422, "invalid prov_code", nil)
	//		ctx.ReplyError(err)
	//		return nil
	//	}

	param := &model_a.Param_handlerGetCim{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return err
	}

	b, t, err := getPeerapongCache(param)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t) // สร้าง cache result (ที่จะ return 304)
	ctx.Reply(r)

	//	province, err := model_a.Cim_Province(param)
	//	if err != nil {
	//		ctx.ReplyError(err)
	//	} else {
	//		ctx.ReplyJSON(province)
	//	}

	return nil
}

func getPeerapongCache(param *model_a.Param_handlerGetCim) ([]byte, time.Time, error) {
	cname := "test.peerapong.provinceId_" + param.Province_id // ชื่อของแคช

	if !datacache.IsRegistered(cname) { // เช็ค ไม่เคยมีแคชชื่อ cname

		c := &TestPeerapongCache{}
		c.Param = param

		//
		datacache.RegisterDataCache(cname, c, []string{"m_tele_station", "rainfall"}, c, "*/5 * * * *") // regis cache cname
	}

	//	datacache.GetJSON(cname)
	return datacache.GetGZJSON(cname) // return gzip json cname
}
