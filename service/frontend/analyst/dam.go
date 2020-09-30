package analyst

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_dam_daily_sum_by_region "haii.or.th/api/thaiwater30/model/dam_daily_sum_by_region"
	model_dam_hourly "haii.or.th/api/thaiwater30/model/dam_hourly"
	model_dam_yearly "haii.or.th/api/thaiwater30/model/dam_yearly"
	model_medium_dam "haii.or.th/api/thaiwater30/model/medium_dam"

	result "haii.or.th/api/thaiwater30/util/result"
	//model_hourly "haii.or.th/api/thaiwater30/model/dam_hourly"

	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	"strconv"
	// "time"
)

type Struct_DamGraph struct {
	Dam        *model_dam.Struct_GetDam         `json:"dam"`         // เขื่อน
	GraphLabel string                           `json:"graph_label"` // ชนิดข้อมูล
	GraphData  *model_dam_daily.Struct_DamGraph `json:"graph_data"`  // ข้อมูลกราฟ
}

type Struct_OnLoadDamLastest struct {
	Agency    *result.Result `json:"agency"`       // หน่วยงาน
	Data      *result.Result `json:"dam_data"`     // เขื่อนในตาราง
	Graph     *result.Result `json:"graph_data"`   // กราฟเขื่อน
	Basin     *result.Result `json:"basin"`        // ลุ่มน้ำ
	Datatype  *result.Result `json:"dam_datatype"` // ชนิดข้อมูล
	DamLarge  *result.Result `json:"dam_large"`    // เขื่อนขนาดใหญ่
	DamMedium *result.Result `json:"dam_medium"`   // เขื่อนขนาดกลาง
	DamHour   *result.Result `json:"dam_hour"`     // เขื่อนขนาดใหญ่รายชม.
	Scale     *result.Result `json:"scale"`        // เกณฑ์
}

type Struct_Dam_Inputparam struct {
	DamDate    string `json:"dam_date"`    // required:false example:`2006-01-02` วันที่ ไม่ใส่ = วันปัจจุบัน
	DamType    int    `json:"dam_size"`    // required:false enum:[1,2] example:`1` ขนาดของเขื่อน ไม่ใส่ = ทุกขนาด, 1 = ขนาดใหญ่, 2 = ขนาดกลาง
	BasinID    string `json:"basin_id"`    // required:false example:`1` รหัสลุ่มน้ำ ไม่ใส่ = ทุกลุ่มน้ำ เลือกได้หลายลุ่มน้ำ เช่น 1,2,4
	ProvinceID string `json:"province_id"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	RegionID   string `json:"region_id"`   // required:false example:`1` รหัสภาค ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	//	ไม่จำเป็นต้องใช้เงื่อนไขนี้ เมื่อเลือกเขื่อนขนาดใหญ่ หน้าแสดงผล มี เขื่อนรายชั่วดมงด้วย
	//	IsHourly bool   `json:"is_hourly"` // example:true รายชั่วโมง
}

type struct_GraphConfig struct {
	Dam_id        int64  `json:"dam_id"`
	Dam_data      string `json:"dam_data"`
	Dam_data_name string `json:"dam_data_name"`
}

type Struct_onLoadDamLastest struct {
	Agency    *Struct_onLoadDamLastest_Agency    `json:"agency"`       // หน่วยงาน
	Data      *Struct_onLoadDamLastest_Data      `json:"dam_data"`     // เขื่อนในตาราง
	Graph     *Struct_onLoadDamLastest_Graph     `json:"graph_data"`   // กราฟเขื่อน
	Basin     *Struct_onLoadDamLastest_Basin     `json:"basin"`        // ลุ่มน้ำ
	Datatype  *Struct_onLoadDamLastest_Datatype  `json:"dam_datatype"` // ชนิดข้อมูล
	DamLarge  *Struct_onLoadDamLastest_DamLarge  `json:"dam_large"`    // เขื่อนขนาดใหญ่
	DamMedium *Struct_onLoadDamLastest_DamMedium `json:"dam_medium"`   // เขื่อนขนาดกลาง
	Scale     *Struct_onLoadDamLastest_Scale     `json:"scale"`        // เกณฑ์
}
type Struct_onLoadDamLastest_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}
type Struct_onLoadDamLastest_Data struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_dam_daily.Struct_DamDaily `json:"data"`   // เขื่อนในตาราง
}
type Struct_onLoadDamLastest_Graph struct {
	Result string             `json:"result"` // example:`OK`
	Data   []*Struct_DamGraph `json:"data"`   // กราฟเขื่อน
}
type Struct_onLoadDamLastest_Basin struct {
	Result string                      `json:"result"` // example:`OK`
	Data   []*model_basin.Struct_Basin `json:"data"`   // ลุ่มน้ำ
}
type Struct_onLoadDamLastest_Datatype struct {
	Result string          `json:"result"` // example:`OK`
	Data   json.RawMessage `json:"data"`   // example:`[{"id":"1","value":"dam_storage","text":{"th":"ปริมาตรกักเก็บ","en":"dam storage"}}]` ชนิดข้อมูล
}
type Struct_onLoadDamLastest_DamLarge struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_dam.Struct_DamGroupByAgency `json:"data"`   // เขื่อนขนาดใหญ่
}
type Struct_onLoadDamLastest_DamMedium struct {
	Result string                                      `json:"result"` // example:`OK`
	Data   []*model_medium_dam.Struct_DamGroupByAgency `json:"data"`   // เขื่อนขนาดกลาง
}
type Struct_onLoadDamLastest_Scale struct {
	Result string          `json:"result"` // example:`OK`
	Data   json.RawMessage `json:"data"`   // example:`{"scale":[{"operator":">","term":"100","color":"#C70000","colorname":"min","text":">100"}]}` เกณฑ์
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_load
// @Summary			เริ่มต้นหน้าเขื่อน
// @Description		หน่วยงาน, เขื่อนในตาราง, กราฟเขื่อน, ลุ่มน้ำ, ชนิดข้อมูล, เขื่อนขนาดใหญ่, เขื่อนขนาดกลาง, เกณฑ์
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadDamLastest successful operation
func (srv *HttpService) onLoadDamLastest(ctx service.RequestContext) error {

	objResult := &Struct_OnLoadDamLastest{}

	//=== Dam Data ===//
	param := &model_dam_daily.Struct_DamDailyLastest_InputParam{}
	resultDamDailyLastest, err := model_dam_daily.GetDamDailyLastest(param)
	if err != nil {
		objResult.Data = result.Result0(err)
	} else {
		objResult.Data = result.Result1(resultDamDailyLastest)
	}

	//=== Dam Graph ===//
	arrObjGraph := make([]*Struct_DamGraph, 0)
	//Setting Dam Graph
	result_setting := model_setting.GetSystemSettingJson("Frontend.Analyst.Dam.OnLoadDamGraph") //analyst_dam_graph_config
	arrSetting := make([]*struct_GraphConfig, 0)
	json.Unmarshal(result_setting, &arrSetting)
	for _, setting := range arrSetting {
		objParam := &model_dam_daily.Struct_GetDamGraph_InputParam{}
		objParam.Dam_id = strconv.FormatInt(setting.Dam_id, 10)
		objParam.Dam_type = setting.Dam_data
		objParam.Year = "2"

		result_Graph, err := model_dam_daily.GetDamGraph(objParam)
		if err != nil {
			return errors.Repack(err)
		}

		result_dam, err := model_dam.GetDam(strconv.FormatInt(setting.Dam_id, 10), "")

		objGraph := &Struct_DamGraph{}
		objGraph.Dam = result_dam[0]
		objGraph.GraphLabel = setting.Dam_data_name
		objGraph.GraphData = result_Graph

		arrObjGraph = append(arrObjGraph, objGraph)
	}
	objResult.Graph = result.Result1(arrObjGraph)

	//=== Basin ===//
	resultBasin, err := model_basin.GetAllBasin()
	if err != nil {
		objResult.Basin = result.Result0(err)
	} else {
		objResult.Basin = result.Result1(resultBasin)
	}

	//=== Large Dam ===//
	resultLargeDam, err := model_dam.GetDamGroupByAgency()
	if err != nil {
		objResult.DamLarge = result.Result0(err)
	} else {
		objResult.DamLarge = result.Result1(resultLargeDam)
	}

	//=== Medium Dam ===//
	resultMediumDam, err := model_medium_dam.GetDamGroupByAgency()
	if err != nil {
		objResult.DamMedium = result.Result0(err)
	} else {
		objResult.DamMedium = result.Result1(resultMediumDam)
	}

	//=== Agency ===//
	rsAgency, err := model_agency.GetAgencyByDam(0)
	if err != nil {
		objResult.Agency = result.Result0(err)
	} else {
		objResult.Agency = result.Result1(rsAgency)
	}

	//=== Dam Hour ===//
	p := &Struct_Dam_Inputparam{}
	paramDamHourly := &model_dam_hourly.Struct_DamHourlyLastest_InputParam{}
	paramDamHourly.Basin_id = p.BasinID
	paramDamHourly.Dam_date = p.DamDate
	resultDamHourly, err := model_dam_hourly.GetDamHourlyLastest(paramDamHourly)
	if err != nil {
		objResult.DamHour = result.Result0(err)
	} else {
		objResult.DamHour = result.Result1(resultDamHourly)
	}

	//=== Dam Datatype ===//
	datatype := model_setting.GetSystemSettingJson("Frontend.public.dam_data_type")
	objResult.Datatype = result.Result1(&datatype)

	scale := model_setting.GetSystemSettingJson("Frontend.public.dam_scale_color")
	objResult.Scale = result.Result1(&scale)

	ctx.ReplyJSON(objResult)
	return nil
}

type Struct_getDam struct {
	Result string              `json:"result"` // example:`OK`
	Data   *Struct_getDam_Data `json:"data"`   // ข้อมูลเขื่อน
}
type Struct_getDam_Data struct {
	DamHourly []*model_dam_hourly.Struct_DamHourly       `json:"dam_hourly,omitempty"` // เขื่อนรายขั่วโมง
	DamMedium []*model_medium_dam.Struct_MediumDamLatest `json:"dam_medium,omitempty"` // เขื่อนขนาดกลาง
	DamDaily  []*model_dam_daily.Struct_DamDaily         `json:"dam_daily,omitempty"`  // เขื่อนรายวัน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam
// @Summary			เขื่อน
// @Description		เขื่อนขนาดกลาง, เขื่อนรายขั่วโมง, เขื่อนรายวัน
// @Method			GET
// @Parameter		-	query Struct_Dam_Inputparam
// @Produces		json
// @Response		200	Struct_getDam successful operation
func (srv *HttpService) getDam(ctx service.RequestContext) error {

	objResult := &Struct_getDam_Data{}

	// Map parameters
	p := &Struct_Dam_Inputparam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	// === Large Dam === //
	// เอาเงื่อนไข ishourly ออก เพราะเวลาเรียก เขื่อนขนาดใหญ่ สามารถแสดงผลเขื่อยรายชม.ได้ ด้วย
	if (p.DamType == 0) || (p.DamType == 1) {
		//		if p.IsHourly {
		paramDamHourly := &model_dam_hourly.Struct_DamHourlyLastest_InputParam{}
		paramDamHourly.Basin_id = p.BasinID
		paramDamHourly.Province_id = p.ProvinceID
		paramDamHourly.Region_id = p.RegionID
		paramDamHourly.Dam_date = p.DamDate

		resultDamHourly, err := model_dam_hourly.GetDamHourlyLastest(paramDamHourly)
		if err != nil {
			return errors.Repack(err)
		}
		objResult.DamHourly = resultDamHourly
		//		} else {
		paramDamDaily := &model_dam_daily.Struct_DamDailyLastest_InputParam{}
		paramDamDaily.Basin_id = p.BasinID
		paramDamDaily.Province_id = p.ProvinceID
		paramDamDaily.Region_id = p.RegionID
		paramDamDaily.Dam_date = p.DamDate

		resultDamDaily, err := model_dam_daily.GetDamDailyLastest(paramDamDaily)
		if err != nil {
			return errors.Repack(err)
		}

		objResult.DamDaily = resultDamDaily
		//		}
	}

	// Cancel current date param for latest medium dam
	// Sometime medium dam is delayed and no result if we specific today date
    // if p.DamDate == "" {
	// 	p.DamDate = time.Now().Format("2006-01-02")
	// }

	//=== Medium Dam ===//
	if (p.DamType == 0) || (p.DamType == 2) {
		paramMediumDam := &model_medium_dam.Struct_DamHourlyLastest_InputParam{}
		paramMediumDam.Basin_id = p.BasinID
		paramMediumDam.Province_id = p.ProvinceID
		paramMediumDam.Region_id = p.RegionID
		// paramMediumDam.Dam_date = p.DamDate

		resultMediumDam, err := model_medium_dam.GetDamLatest(paramMediumDam)
		if err != nil {
			return errors.Repack(err)
		}
		objResult.DamMedium = resultMediumDam
	}

	ctx.ReplyJSON(result.Result1(objResult))
	return nil
}

func (srv *HttpService) getDamGraph(ctx service.RequestContext) error {
	/*
		//Map parameters
		p := &agencyParam{}
		if err := ctx.GetRequestParams(p); err != nil {
			return errors.Repack(err)
		}
		ctx.LogRequestParams(p)

		//Get Agency
		result, err := model.GetAgency(ctx.GetServiceParams("id"), p.DepartmentId)
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result)
		}
	*/
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_daily_graph
// @Method			GET
// @Summary			เขื่อนรายวันสำหรับกราฟ
// @Description		Return data dam daily for graph.
// @Parameter		-	query model_dam_daily.GraphAnalystDamDailyInput
// @Produces		json
// @Response		200		DamDailySwagger successful operation

type DamDailySwagger struct {
	Result string                                       `json:"result"` //example:`OK`
	Data   []model_dam_daily.GraphAnalystDamDailyOutput `json:"data"`
}

func (srv *HttpService) getDamGraphDaily(ctx service.RequestContext) error {

	p := &model_dam_daily.GraphAnalystDamDailyInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_dam_daily.GetDamGraphDailyAnalyst(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_yearly_graph
// @Method			GET
// @Summary			เขื่อนรายปีสำหรับกราฟ
// @Description		Return data dam yearly for graph.
// @Parameter		-	query model_dam_yearly.GraphDamYearlyInput
// @Produces		json
// @Response		200		model_dam_yearly.GraphDamOutput successful operation
func (srv *HttpService) getDamGraphYearly(ctx service.RequestContext) error {

	p := &model_dam_yearly.GraphDamYearlyInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_dam_yearly.GetDamGraphYearly(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_medium_graph
// @Method			GET
// @Summary			เขื่อนขนาดกลางสำหรับกราฟ
// @Description		Return data dam medium for graph.
// @Parameter		-	query model_dam_yearly.GraphDamMediumInput
// @Produces		json
// @Response		200		model_dam_yearly.GraphDamOutput successful operation
func (srv *HttpService) getDamMediumGraph(ctx service.RequestContext) error {

	p := &model_dam_yearly.GraphDamMediumInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_dam_yearly.GetDamMediumGraph(p)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_monitoring
// @Method			GET
// @Summary			monitoring 4 เขื่อนหลัก
// @Description		Return data 4 dam.
// @Produces		json
// @Response		200		MonitoringDamSwagger successful operation
type MonitoringDamSwagger struct {
	Result string                `json:"result"` //example:`OK`
	Data   []MonitoringDamOutput `json:"data"`
}

func (srv *HttpService) getDam4Main(ctx service.RequestContext) error {

	rs, err := model_dam_daily.GetDamFourMain()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type MonitoringDamOutput struct {
	DamId     				string					`json:"dam_id"`     	// id เขื่อน
	DamName     			json.RawMessage `json:"dam_name"`     // example:`{"th":"เขื่อนภูมิพล","en":" ","jp":" "}` ชื่อเขื่อน
	DamDate     			string          `json:"dam_date"`     // example:`2006-01-02` วันที่เก็บข้อมูล
	DamStorage  			interface{}     `json:"dam_storage"`  // example:`140` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	DamInflow   			interface{}     `json:"dam_inflow"`   // example:`10` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	DamReleased 			interface{}     `json:"dam_released"` // example:`11` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	DamStoragePercent	interface{}     `json:"dam_storage_percent"`  // example:`140` เปอร์เซ็ฯปริมาณน้ำกักเก็บปัจจุบัน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_daily_sum_by_region_rid
// @Method			GET
// @Summary			คำนวณผลรวมเขื่อนรายภาค รายวัน
// @Description		Return dam sum data for graph
// @Parameter		-	query model_dam_daily_sum_by_region.DamSumByRegionInput
// @Produces		json
// @Response		200		model_dam_daily_sum_by_region.DamSumByRegionByYear successful operation
func (srv *HttpService) getDamSumByRegionRid(ctx service.RequestContext) error {
	// parameters
	p := &model_dam_daily_sum_by_region.DamSumByRegionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	
	// get data
	rs, err := model_dam_daily_sum_by_region.GetSumByRegionRid(p)
	
	// check error and return json
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/dam_daily_compare_sum_region_rid
// @Method			GET
// @Summary			คำนวณผลรวมเขื่อนรายภาค รายวัน
// @Description		Return dam sum data for graph
// @Parameter		-	query model_dam_daily_compare_sum_by_region.DamCompareSumByRegionInput
// @Produces		json
// @Response		200		model_dam_daily_compare_sum_by_region.DamCompareSumByRegion successful operation
func (srv *HttpService) getDamCompareSumByRegionRid(ctx service.RequestContext) error {
	// parameters
	p := &model_dam_daily_sum_by_region.DamCompareSumByRegionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	
	// get data
	rs, err := model_dam_daily_sum_by_region.GetCompareSumByRegionRid(p)
	
	// check error and return json
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}