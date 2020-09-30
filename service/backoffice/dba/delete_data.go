package dba

import (
	"encoding/json"
	model_setting "haii.or.th/api/server/model/setting"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_dam_hourly "haii.or.th/api/thaiwater30/model/dam_hourly"
	model_rainfall "haii.or.th/api/thaiwater30/model/rainfall"
	model_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	uSeeting "haii.or.th/api/thaiwater30/util/setting"
)

type deleteDataInputParam struct {
	ID        string `json:"id"`
	DataType  string `json:"data_type"` // enum:[dam_daily,dam_hourly,rainfall,tele_waterlevel]
	StationID string `json:"station_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Struct_onLoadDeleteData struct {
	Data   *Struct_onLoadDeleteData_Data `json:"data"`   // ประเภทข้อมูล
	Result string                        `json:"result"` // example:`OK`
}
type Struct_onLoadDeleteData_Data struct {
	DataType *[]uSeeting.Struct_DataTypeOption `json:"data_type"` // ประเภทข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/delete_data_load
// @Summary			เริ่มต้นหน้าลบข้อมูล
// @Description		ประเภทข้อมูล
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadDeleteData successful operation
func (srv *HttpService) onLoadDeleteData(ctx service.RequestContext) error {
	type rs struct {
		DataType json.RawMessage `json:"data_type"`
	}
	//Get List of DataTypeOption Data
	dataResult := &rs{}
	dataResult.DataType = model_setting.GetSystemSettingJson("bof.Shared.DataTypeOption")

	ctx.ReplyJSON(result.Result1(dataResult))

	return nil
}

type Param_getDeleteData struct {
	//	DataType  string `json:"data_type"` // enum:[dam_daily,dam_hourly,rainfall,tele_waterlevel]
	StationID string `json:"station_id"` // example: 17 รหัสสถานี
	StartDate string `json:"start_date"` // example: 2017-06-13 วันที่เริ่ม้ตน
	EndDate   string `json:"end_date"`   // example: 2017-06-14 วันที่สิ้นสุด
}
type Struct_getDeleteData_dam_daily struct {
	Data   *model_dam_daily.Struct_DamDailyLastest_OutputParam `json:"data"`   // ข้อมูลเขื่อนรายวัน
	Result string                                              `json:"result"` // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/delete_data?data_type=dam_daily
// @Summary			ดูข้อมูลเขื่อนรายวัน
// @Method			GET
// @Parameter		- query Param_getDeleteData
// @Produces		json
// @Response		200	Struct_getDeleteData_dam_daily successful operation

type Struct_getDeleteData_dam_hourly struct {
	Data   *model_dam_hourly.DamHourlyLastest_OutputParam `json:"data"`   // ข้อมูลเขื่อนรายชั่วโมง
	Result string                                         `json:"result"` // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/delete_data?data_type=dam_hourly
// @Summary			ดูข้อมูลเขื่อนรายชั่วโมง
// @Method			GET
// @Parameter		- query Param_getDeleteData
// @Produces		json
// @Response		200	Struct_getDeleteData_dam_hourly successful operation

type Struct_getDeleteData_rainfall struct {
	Data   *model_rainfall.GetRainfallLastest_OutputParam `json:"data"`   // ข้อมูลฝน
	Result string                                         `json:"result"` // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/delete_data?data_type=rainfall
// @Summary			ดูข้อมูลฝน
// @Method			GET
// @Parameter		- query Param_getDeleteData
// @Produces		json
// @Response		200	Struct_getDeleteData_rainfall successful operation

type Struct_getDeleteData_tele_waterlevel struct {
	Data   *model_waterlevel.GetWaterlevelLastest_OutputParam `json:"data"`   // ข้อมูลระดับน้ำ
	Result string                                             `json:"result"` // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/delete_data?data_type=tele_waterlevel
// @Summary			ดูข้อมูลระดับน้ำ
// @Method			GET
// @Parameter		- query Param_getDeleteData
// @Produces		json
// @Response		200	Struct_getDeleteData_tele_waterlevel successful operation
func (srv *HttpService) getDeleteData(ctx service.RequestContext) error {
	//Map parameters
	p := &deleteDataInputParam{}
	err := ctx.GetRequestParams(p)

	if err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	var result *result.Result

	//Get Data
	switch p.DataType {
	case "dam_daily":
		result, err = getDamDaily(p)
	case "dam_hourly":
		result, err = getDamHourly(p)
	case "rainfall":
		result, err = getRainfall(p)
	case "tele_waterlevel":
		result, err = getWaterlevel(p)
	default:
		return rest.NewError(404, "mismatch param 'data_type' ", nil)
		//err = errors.New("mismatch query string 'data_type' ")
	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

type Param_deleteDeleteData struct {
	ID       string `json:"id"`        // example: 101 รหัสข้อมูล
	DataType string `json:"data_type"` // enum:[dam_daily,dam_hourly,rainfall,tele_waterlevel] example: dam_daily ประเภทข้อมูล
}

type Struct_deleteDeleteData struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/delete_data
// @Summary			ลบข้อมูล
// @Method			DELETE
// @Parameter		- query Param_deleteDeleteData
// @Produces		json
// @Response		200	Struct_deleteDeleteData successful operation
func (srv *HttpService) deleteDeleteData(ctx service.RequestContext) error {
	//Map parameters
	p := &deleteDataInputParam{}
	err := ctx.GetRequestParams(p)

	if err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	var result *result.Result

	//Get Data
	switch p.DataType {
	case "dam_daily":
		result, err = deleteDamDaily(p, ctx.GetUserID())
	case "dam_hourly":
		result, err = deleteDamHourly(p, ctx.GetUserID())
	case "rainfall":
		result, err = deleteRainfall(p, ctx.GetUserID())
	case "tele_waterlevel":
		result, err = deleteWaterlevel(p, ctx.GetUserID())
	default:
		return rest.NewError(404, "mismatch param 'data_type' ", nil)
		//err = errors.New("mismatch query string 'data_type' ")
	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func getDamDaily(p *deleteDataInputParam) (*result.Result, error) {

	objParam := &model_dam_daily.Struct_DamDaily_InputParam{}
	objParam.Dam_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate

	rs, err := model_dam_daily.GetDamDaily(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func getDamHourly(p *deleteDataInputParam) (*result.Result, error) {

	objParam := &model_dam_hourly.Struct_DamHourly_InputParam{}
	objParam.Dam_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate

	rs, err := model_dam_hourly.GetDamHourly(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func getRainfall(p *deleteDataInputParam) (*result.Result, error) {

	objParam := &model_rainfall.Rainfall_InputParam{}
	objParam.Station_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate

	rs, err := model_rainfall.GetRainfallByStationAndDate(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func getWaterlevel(p *deleteDataInputParam) (*result.Result, error) {

	objParam := &model_waterlevel.Waterlevel_InputParam{}
	objParam.Station_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate

	rs, err := model_waterlevel.GetWaterlevelByStationAndDate(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteDamDaily(p *deleteDataInputParam, userId int64) (*result.Result, error) {

	objParam := &model_dam_daily.Struct_DamDaily_InputParam{}
	objParam.Id = p.ID

	rs, err := model_dam_daily.UpdatetoDeleteDamDaily(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteDamHourly(p *deleteDataInputParam, userId int64) (*result.Result, error) {

	objParam := &model_dam_hourly.Struct_DamHourly_InputParam{}
	objParam.Id = p.ID

	rs, err := model_dam_hourly.UpdateToDeleteDamHourly(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteRainfall(p *deleteDataInputParam, userId int64) (*result.Result, error) {

	objParam := &model_rainfall.Rainfall_InputParam{}
	objParam.Id = p.ID

	rs, err := model_rainfall.UpdateToDeleteRainfall(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteWaterlevel(p *deleteDataInputParam, userId int64) (*result.Result, error) {

	objParam := &model_waterlevel.Waterlevel_InputParam{}
	objParam.Id = p.ID

	rs, err := model_waterlevel.UpdateToDeleteWaterlevel(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}
