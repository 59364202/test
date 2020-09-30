package public

import (
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"

	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/service"

	result "haii.or.th/api/thaiwater30/util/result"

	"time"
)

// ----------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------- thailand main page  ----------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------
// @DocumentName	v1.public
// @Service			thaiwater30/public/thailand_main
// @Summary			หน้าหลัก
// @Description
// @				* ข้อมูลฝน 24 ชัวโมง
// @				* ภาพเรดาร์
// @				* ข้อมูลเขื่อน
// @				* ข้อมูลระดับน้ำ
// @				* ข้อมูลคุณภาพน้ำ
// @				* ภาพพายุ
// @				* คาดการณ์ฝนล่วงหน้า 7 วัน
// @				* คาดการณ์คลื่นล่วงหน้า 7 วัน
// @				* พื้นที่ประกาศภัย
// @Method			GET
// @Produces		json
// @Response		200		Struct_Main successful operation
func (srv *HttpService) getThailandMain(ctx service.RequestContext) error {
	return replyWithHourlyCache(ctx, thailandMainCacheName, getThailandMainCacheBuildData)
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/thailand_main_rain
// @Summary			ฝน24ชม. รายจังหวัด
// @Parameter		province_code	query string example:`10` รหัสจังหวัด
// @Parameter		data_limit	query string example:`20` จำนวน records ที่ต้องการ
// @Method			GET
// @Produces		json
// @Response		200		Struct_Rain_Data successful operation
func (srv *HttpService) getThailandRain(ctx service.RequestContext) error {
	p := &model_geocode.Param_Geocode{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	p1 := &model_rainfall24hr.Param_Rainfall24{}
	if err := ctx.GetRequestParams(p1); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p1)
	data, t, err := getRainCache(p.Province_code, p1.Data_Limit)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, data, t)
	ctx.Reply(r)

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/thailand_main_waterlevel
// @Summary			ระดับน้ำ รายจังหวัด
// @Method			GET
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Produces		json
// @Response		200		Struct_WaterLevel_Data successful operation
func (srv *HttpService) getThailandWaterlevel(ctx service.RequestContext) error {
	p := &model_tele_waterlevel.Waterlevel_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)
	data, t, err := getWaterLevelCache(p.Province_Code, true)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, data, t)
	ctx.Reply(r)

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/thailand_main_dam
// @Summary			เขื่อนกรมชลประทาน
// @Method			GET
// @Parameter		agency_id	query  string example:`12` รหัสหน่วยงาน ตัวอย่างกรมชลประทาน
// @Produces		json
// @Response		200		Struct_WaterLevel_Data successful operation
func (srv *HttpService) getThailandDam(ctx service.RequestContext) error {
	p := &model_dam_daily.Struct_DamDailyLastest_InputParam{Agency_id: "12"}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	data, err := model_dam_daily.GetDamDailyLastest(p)
	if err != nil {
		return errors.Repack(err)
	}
	ctx.ReplyJSON(data)
	return nil
}

// rain
type rainCache struct {
	Params model_rainfall24hr.Param_Rainfall24
}

func (rc *rainCache) BuildData() (interface{}, error) {
	data, err := model_rainfall24hr.GetRainfallThailandDataCache(&rc.Params)
	if err != nil {
		return result.Result0(err), nil
	}
	return result.Result1(data), nil
}

func getRainCache(provinceCode string, dataLimit int) ([]byte, time.Time, error) {
	cname := rainCacheName + "provinceCode" + provinceCode + "datalimit"
	if dataLimit > 0 {
		cname += string(dataLimit)
	}

	if !datacache.IsRegistered(cname) {
		c := &rainCache{}
		c.Params.Province_Code = provinceCode
		c.Params.Data_Limit = dataLimit
		datacache.RegisterDataCache(cname, c, []string{"cache.latest_rainfall24h"}, &hourlyCacheValidator{}, "1,11,21,31,41,51 * * * *")
	}

	return datacache.GetGZJSON(cname)
}

// waterlevel
type waterLevelCache struct {
	Params model_tele_waterlevel.Waterlevel_InputParam
}

func (wlc *waterLevelCache) BuildData() (interface{}, error) {
	data, err := model_tele_waterlevel.GetWaterLevelThailandDataCache(&wlc.Params)
	if err != nil {
		log.Logf("getWaterLevelCache(%s) error ...%v", wlc.Params.Province_Code, err)
		return result.Result0(err), nil
	}
	return result.Result1(data), nil
}

func getWaterLevelCache(provinceCode string, IsMain bool) ([]byte, time.Time, error) {
	cname := waterLevelCacheName + provinceCode

	if !datacache.IsRegistered(cname) {
		c := &waterLevelCache{}
		c.Params.Province_Code = provinceCode
		c.Params.IsMain = IsMain
		datacache.RegisterDataCache(cname, c, []string{"cache.latest_waterlevel"}, &hourlyCacheValidator{}, "1,11,21,31,41,51 * * * *")
	}

	return datacache.GetGZJSON(cname)
}
