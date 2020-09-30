package data_management

import (
	"encoding/json"
	model_setting "haii.or.th/api/server/model/setting"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_dam_hourly "haii.or.th/api/thaiwater30/model/dam_hourly"
	model_rainfall "haii.or.th/api/thaiwater30/model/rainfall"
	model_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	result "haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

type Struct_OnloadLastestData struct {
	DataType  json.RawMessage `json:"data_type"`  // ประเภทข้อมูล
	DateRange int64           `json:"date_range"` // example:`31` ช่วงวันที่ ที่เลือกได้
}

type Struct_onLoadLastestData struct {
	Result string                         `json:"result"` // example:`OK`
	Data   *Struct_onLoadLastestData_Data `json:"data"`   // ประเภทข้อมูล, ช่วงวันที่
}
type Struct_onLoadLastestData_Data struct {
	DataType  []*uSetting.Struct_DataTypeOption `json:"data_type"`  // ประเภทข้อมูล
	DateRange int64                             `json:"date_range"` // example:`31` ช่วงวันที่ ที่เลือกได้
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/lastest_data_load
// @Summary			เริ่มต้นหน้า ค้นหาข้อมูลล่าสุด
// @Description		ประเภทข้อมูล, ช่วงวันที่
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadLastestData successful operation
func (srv *HttpService) onLoadLastestData(ctx service.RequestContext) error {

	//Get List of DataTypeOption Data
	dataResult := &Struct_OnloadLastestData{}
	dataResult.DataType = model_setting.GetSystemSettingJson("bof.Shared.DataTypeOption")

	//Get DateRange Data
	dataResult.DateRange = model_setting.GetSystemSettingInt("bof.DataMgt.LastestData.DateRange")
	if (dataResult.DateRange) == 0 {
		dataResult.DateRange = model_setting.GetSystemSettingInt("setting.Default.DateRange")
	}

	ctx.ReplyJSON(result.Result1(dataResult))

	return nil
}

type Param_getLastestData struct {
	DataType  string `json:"data_type"`  // enum:[dam_daily,dam_hourly,rainfall,tele_waterlevel] ชื่อตาราง
	StationID string `json:"station_id"` // example: 17 สถานี
	StartDate string `json:"start_date"` // example: 2006-01-02 วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // example: 2006-01-02 วันที่สิ้นสุด
}
type Sturct_getLastestData struct {
	Result string                                              `json:"result"` // example:`OK`
	Data   *model_dam_daily.Struct_DamDailyLastest_OutputParam `json:"data"`   // ข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/lastest_data
// @Summary			ค้นหาข้อมูลล่าสุด
// @Method			GET
// @Parameter		-	query	Param_getLastestData
// @Produces		json
// @Response		200		Sturct_getLastestData successful operation
func (srv *HttpService) getLastestData(ctx service.RequestContext) error {
	//Map parameters
	p := &Param_getLastestData{}
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
	case "waterlevel", "tele_waterlevel":
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

type Param_deleteLastestData struct {
	ID       string `json:"id"`        // example:3473 รหัสของข้อมูล
	DataType string `json:"data_type"` // enum:[dam_daily,dam_hourly,rainfall,tele_waterlevel] example:dam_daily ชื่อตาราง
}

type Sturct_deleteLastestData struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/lastest_data
// @Summary			ลบข้อมูล
// @Description 	เปลี่ยนค่า deleted_at เป็นวันเวลาปัจจุบัน
// @Method			DELETE
// @Parameter		-	query	Param_deleteLastestData
// @Produces		json
// @Response		200 Sturct_deleteLastestData successful operation
func (srv *HttpService) deleteLastestData(ctx service.RequestContext) error {
	//Map parameters
	p := &Param_deleteLastestData{}
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

func getDamDaily(p *Param_getLastestData) (*result.Result, error) {

	objParam := &model_dam_daily.Struct_DamDaily_InputParam{}
	objParam.Dam_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate + " 23:59"

	rs, err := model_dam_daily.GetDamDaily(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func getDamHourly(p *Param_getLastestData) (*result.Result, error) {

	objParam := &model_dam_hourly.Struct_DamHourly_InputParam{}
	objParam.Dam_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate + " 23:59"

	rs, err := model_dam_hourly.GetDamHourly(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func getRainfall(p *Param_getLastestData) (*result.Result, error) {

	objParam := &model_rainfall.Rainfall_InputParam{}
	objParam.Station_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate + " 23:59"

	rs, err := model_rainfall.GetRainfallByStationAndDate(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func getWaterlevel(p *Param_getLastestData) (*result.Result, error) {

	objParam := &model_waterlevel.Waterlevel_InputParam{}
	objParam.Station_id = p.StationID
	objParam.Start_date = p.StartDate
	objParam.End_date = p.EndDate + " 23:59"

	rs, err := model_waterlevel.GetWaterlevelByStationAndDate(objParam)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteDamDaily(p *Param_deleteLastestData, userId int64) (*result.Result, error) {

	objParam := &model_dam_daily.Struct_DamDaily_InputParam{}
	objParam.Id = p.ID

	rs, err := model_dam_daily.UpdatetoDeleteDamDaily(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteDamHourly(p *Param_deleteLastestData, userId int64) (*result.Result, error) {

	objParam := &model_dam_hourly.Struct_DamHourly_InputParam{}
	objParam.Id = p.ID

	rs, err := model_dam_hourly.UpdateToDeleteDamHourly(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteRainfall(p *Param_deleteLastestData, userId int64) (*result.Result, error) {

	objParam := &model_rainfall.Rainfall_InputParam{}
	objParam.Id = p.ID

	rs, err := model_rainfall.UpdateToDeleteRainfall(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}

func deleteWaterlevel(p *Param_deleteLastestData, userId int64) (*result.Result, error) {

	objParam := &model_waterlevel.Waterlevel_InputParam{}
	objParam.Id = p.ID

	rs, err := model_waterlevel.UpdateToDeleteWaterlevel(objParam, userId)
	if err != nil {
		return nil, err
	}

	return result.Result1(rs), nil
}
