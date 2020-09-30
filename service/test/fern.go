package test

import (
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	"haii.or.th/api/server/model/datacache"

	"time"

	model_a "haii.or.th/api/thaiwater30/model/a"
)

type Fern struct {
}

// ตัวอย่าง json result
type TestFernResult struct {
	Province_id string    `json:"province_id"` // example:`81` required:false รหัสจังหวัดนะ
	Region_id   string    `json:"region_id"`   // example:`5` required:false มันเป็นรหัสภาคนะ
	Time        time.Time `json:"time"`
}

// ตัวอย่าง struct ของ cache
type TestFernCache struct {
	Param *model_a.Param_handlerGetFern // เพิ่ม Param มาเพื่อเก็บ param จาก url
}

//  ตัวอย่าง func เช็คความใหม่ของข้อมูล
//  บังคับชื่อ
//	*require*
func (s *TestFernCache) IsValid(lastupdate time.Time) bool {
	return true
}


//  คัวอย่าง Description
//  บังคับชื่อ
//	*require*
func (s *TestFernCache) GetDescription() string {
	return "refresh every 5 minute"
}

//	ตัวอย่าง สร้างข้อมูล
//  บังคับชื่อ
//	*require*
func (s *TestFernCache) BuildData() (interface{}, error) {
	//	province, err := model_a.Fern_Province(s.Param)
	//	if err != nil {
	//		return nil, err
	//	} else {
	//		return province, err
	//	}

	return &TestFernResult{Province_id: s.Param.Province_id, Region_id: s.Param.Region_id, Time: time.Now()}, nil
}


// @DocumentName	v1.test
// @Service			test/test_fern
// @Summary			ทดสอบ service
// @Parameter		- query model_a.Param_handlerGetFern
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 model_a.fern_province successful operation

func (srv *Fern) handlerGetFern(ctx service.RequestContext) error {
	//	province, err := model_a.Fern_Province(param)
	//	if err != nil {
	//		ctx.ReplyError(err)
	//	} else {
	//		ctx.ReplyJSON(province)
	//	}

	param := &model_a.Param_handlerGetFern{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return err
	}
	
	b, t, err := getFernCache(param)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t) // สร้าง cache result (ที่จะ return 304)
	ctx.Reply(r)

	return nil

	//prov_code := ctx.GetServiceParams("prov_code")
	//province, err := model_a.Fern_Province("prov_code")
	/*if err != nil {
		return err
	}
	ctx.ReplyJSON(province)
	return nil*/
}

func getFernCache(param *model_a.Param_handlerGetFern) ([]byte, time.Time, error) {
	cname := "test.Fern.provinceId_" + param.Province_id // ชื่อของแคช

	if !datacache.IsRegistered(cname) { // เช็ค ไม่เคยมีแคชชื่อ cname

		c := &TestFernCache{}
		c.Param = param

		//
		datacache.RegisterDataCache(cname, c, []string{"m_tele_station", "rainfall"}, c, "*/5 * * * *") // regis cache cname
	}

	//	datacache.GetJSON(cname)
	return datacache.GetGZJSON(cname) // return gzip json cname
}

