package tool

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_dam_hourly "haii.or.th/api/thaiwater30/model/dam_hourly"
	model "haii.or.th/api/thaiwater30/model/ignore"
	model_history "haii.or.th/api/thaiwater30/model/ignore_history"
	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
	model_rainfall "haii.or.th/api/thaiwater30/model/rainfall"

	"haii.or.th/api/thaiwater30/service/frontend/public"

	result "haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	//"encoding/json"
	//"sort"
)

/*
type Struct_IgnoreStation_InputParam struct {
	ID				string     	`json:"id"`
	TableName 		string     	`json:"table_name"`
	IsIgnore		string 		`json:"is_ignore"`
}
*/

type Struct_getIgnoreTable struct {
	Result string                                                  `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_Inore_TableList_LatestIgnoreDatetime `json:"data"`   // ชนิดข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/ignore_table
// @Summary			ชนิดข้อมูลทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getIgnoreTable successful operation
func (srv *HttpService) getIgnoreTable(ctx service.RequestContext) error {
	//Get List of DataTypeOption Data
	ignoreStationList := model_setting.GetSystemSettingJson("bof.Tool.Ignore.TableList")
	dataResult, err := model.GetLastestIgnoreStation(ignoreStationList)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Struct_getIgnoreHistory struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model_history.Struct_IgnoreHistory `json:"data"`   // ประวัติการ ignore
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/ignore_history
// @Summary			ประวัติการ ignore
// @Method			GET
// @Parameter		- query model_history.Struct_IgnoreHistory_InputParam{table_name}
// @Produces		json
// @Response		200	Struct_getIgnoreHistory successful operation
func (srv *HttpService) getIgnoreHistory(ctx service.RequestContext) error {
	//Map parameters
	param := &model_history.Struct_IgnoreHistory_InputParam{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Get List of History
	resultData, err := model_history.GetIgnoreHistory(param)
	if err != nil {
		return errors.Repack(err)
	}

	ctx.ReplyJSON(result.Result1(resultData))
	return nil
}

type Struct_getIgnoreRainfallDetail struct {
	Result string                                                  `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_Inore_TableList_LatestIgnoreDatetime `json:"data"`   // ชนิดข้อมูล
}

// by 				Peerapong (peerapong@haii.or.th)
// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/ignore_rainfall_detail
// @Summary			รายละเอียดข้อมูลฝนที่ถูก ignore รายสถานี
// @Method			GET
// @Parameter		station_id	query	string	example:`13` รหัสสถานี
// @Parameter		start_date	query	string	example:`2018-01-01` วันที่เริ่มต้น
// @Parameter		end_date	query	string	example:`2018-01-02` วันที่สิ้นสุด
// @Produces		json
// @Response		200	Struct_getIgnoreRainfallDetail successful operation
func (srv *HttpService) getIgnoreRainfallDetail(ctx service.RequestContext) error {
	//Get List of DataTypeOption Data
	var err error
	//Map parameters
	param := &model_rainfall.Rainfall_InputParam{}
	err = ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}
	
	paramRain := &model_rainfall.Rainfall_InputParam {
		Station_id : param.Station_id, // รหัสสถานี
		Start_date : param.Start_date, // วันที่เริ่มต้น
		End_date   : param.End_date,   // วันที่สิ้นสุด
	}
	dataResult, err := model_rainfall.GetRainfallByStationAndDate(paramRain)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Struct_IgnoreStation struct {
	Data       *result.Result `json:"data"`
	IgnoreData *result.Result `json:"ignore_data"`
}

type Struct_getIgnoreStationLoad struct {
	Data       *Struct_getIgnoreStationLoad_Data       `json:"data"`        // ข้อมูลที่แสดงในหน้า main
	IgnoreData *Struct_getIgnoreStationLoad_IgnoreData `json:"ignore_data"` // สถานีที่ถูก ignore
}
type Struct_getIgnoreStationLoad_Data struct {
	Result string                         `json:"result"` // example:`OK`
	Data   *public.Struct_WaterLevel_Data `json:"data"`   // ข้อมูลที่แสดงในหน้า main
}
type Struct_getIgnoreStationLoad_IgnoreData struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model_history.Struct_IgnoreHistory `json:"data"`   // สถานีที่ถูก ignore
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/ignore
// @Summary			รายชื่อสถานี่ที่ถูก ignore / ไม่ถูก ignore
// @Description		ข้อมูลที่ได้จะแตกต่างกันไปตามชนิดข้อมูล
// @Method			GET
// @Parameter		- query model_history.Struct_IgnoreHistory_InputParam{table_name}
// @Produces		json
// @Response		200	Struct_getIgnoreStationLoad successful operation
func (srv *HttpService) getIgnoreStationLoad(ctx service.RequestContext) error {

	dataResult := &Struct_IgnoreStation{}

	var objData interface{}
	var objIgnoreData interface{}
	var err error

	//Map parameters
	param := &model.Struct_IgnoreStation_InputParam{}
	err = ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}

	switch param.TableName {
	case "rainfall_24h":
		paramRain := &model_rainfall24hr.Param_Rainfall24{}
		objData, err = model_rainfall24hr.GetRainfallThailandDataCache(paramRain)
	case "dam_daily":
		paramDamDaily := &model_dam_daily.Struct_DamDailyLastest_InputParam{}
		objData, err = model_dam_daily.GetDamDailyLastest(paramDamDaily)
	case "dam_hourly":
		paramDamHourly := &model_dam_hourly.Struct_DamHourlyLastest_InputParam{}
		objData, err = model_dam_hourly.GetDamHourlyLastest(paramDamHourly)
	case "tele_waterlevel":
		paramWaterlevel := &model_waterlevel.Waterlevel_InputParam{}
		objData, err = model_waterlevel.GetWaterLevelThailandDataCache(paramWaterlevel)
	case "waterquality":
		objData, err = model_waterquality.GetWaterQualityThailandDataCache(&model_waterquality.Param_WaterQualityCache{})
	default:
		return rest.NewError(404, "Unknown table_name", nil)
	}

	if err != nil {
		dataResult.Data = result.Result0(err)
	} else {
		dataResult.Data = result.Result1(objData)
	}

	paramIgnoreData := &model_history.Struct_IgnoreHistory_InputParam{}
	paramIgnoreData.TableName = param.TableName
	objIgnoreData, err = model_history.GetIgnoreData(paramIgnoreData)
	if err != nil {
		dataResult.IgnoreData = result.Result0(err)
	} else {
		dataResult.IgnoreData = result.Result1(objIgnoreData)
	}

	ctx.ReplyJSON(dataResult)
	return nil
}

type Struct_patchIgnoreStation struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Ignore Data Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/ignore
// @Summary			เปลี่ยนสถานะ ignore
// @Method			PATCH
// @Consumes		json
// @Parameter		-	body	model.Struct_IgnoreStation_InputParam
// @Produces		json
// @Response		200	Struct_patchIgnoreStation successful operation
func (srv *HttpService) patchIgnoreStation(ctx service.RequestContext) error {
	//Map parameters
	param := &model.Struct_IgnoreStation_InputParam{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Get List of History
	resultData, err := model.PatchIgnoreStation(ctx.GetUserID(), param)
	if err != nil {
		return errors.Repack(err)
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}

	return nil
}
