package data_management

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
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

type Struct_CheckErrordata_InputParam struct {
	StationName       string `json:"data_type"`
	StationTableName  string `json:"station_table"`
	AgencyID          string `json:"agency_id"`
	StationColumnName string `json:"station_column_name"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}

type Struct_CheckErrordata struct {
	ColumnName []string    `json:"column_name"` // example:`["id", "station_oldcode", "datetime", "station_name", "station_province_name", "agency_name"]` หัวตาราง
	Data       interface{} `json:"data"`        // ข้อมูลในตาราง
}
type Struct_onLoadCheckErrorData struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_DataTypeOption `json:"data"`   // ประเภทข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/check_error_data_load
// @Summary			เริ่มต้นหน้าค้นหาข้อมูลทีมีค่าผิดพลาด
// @Description		รายละเอียด ประเภทข้อมูล
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadCheckErrorData successful operation
func (srv *HttpService) onLoadCheckErrorData(ctx service.RequestContext) error {
	//Get List of DataTypeOption Data
	dataResult := model_setting.GetSystemSettingJson("bof.Shared.DataTypeOption")
	ctx.ReplyJSON(result.Result1(&dataResult))
	return nil
}

type Struct_getCheckErrorData struct {
	Result string                         `json:"result"` // example:`OK`
	Data   *Struct_getCheckErrorData_Data `json:"data"`   //
}
type Struct_getCheckErrorData_Data struct {
	ColumnName []string                                    `json:"column_name"` // example:`["id","station_oldcode","datetime","station_name","station_province_name","agency_name","agency_shortname","rainfall5m","rainfall10m","rainfall15m","rainfall30m","rainfall1h","rainfall3h","rainfall6h","rainfall12h","rainfall24h","rainfall_acc"]` หัวตาราง
	Data       []*model_rainfall.Struct_Rainfall_ErrorData `json:"data"`        // ข้อมูลที่ผิดพลาด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/check_error_data
// @Summary			ค้นหาข้อมูลทีมีค่าผิดพลาด
// @Method			GET
// @Parameter		agency_id query string example:9  รหัสหน่วยงาน
// @Parameter		data_type query	string enum:[dam_daily,dam_hourly,rainfall,waterlevel] ประเภทข้อมูล
// @Produces		json
// @Response		200	Struct_getCheckErrorData successful operation
func (srv *HttpService) getCheckErrorData(ctx service.RequestContext) error {

	//Map parameters
	param := &Struct_CheckErrordata_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	resultData := &Struct_CheckErrordata{}

	var err error
	var resultErrorData interface{}

	//Get Data
	switch param.StationName {
	case "dam_daily":
		resultData.ColumnName = model_dam_daily.DamDailyColumnForCheckErrordata
		paramDamDaily := &model_dam_daily.Struct_DamDailyLastest_InputParam{}
		paramDamDaily.Agency_id = param.AgencyID
		paramDamDaily.Start_date = param.StartDate
		paramDamDaily.End_date = param.EndDate
		resultErrorData, err = model_dam_daily.GetErrorDamDaily(paramDamDaily)
	case "dam_hourly":
		resultData.ColumnName = model_dam_hourly.DamHourlyColumnForCheckErrordata
		paramDamHourly := &model_dam_hourly.Struct_DamHourlyLastest_InputParam{}
		paramDamHourly.Agency_id = param.AgencyID
		paramDamHourly.Start_date = param.StartDate
		paramDamHourly.End_date = param.EndDate
		resultErrorData, err = model_dam_hourly.GetErrorDamHourly(paramDamHourly)
	case "rainfall":
		resultData.ColumnName = model_rainfall.RainfallColumnForCheckErrordata
		paramRainfall := &model_rainfall.Rainfall_InputParam{}
		paramRainfall.Agency_id = param.AgencyID
		paramRainfall.Start_date = param.StartDate
		paramRainfall.End_date = param.EndDate
		resultErrorData, err = model_rainfall.GetErrorRainfall(paramRainfall)
	case "waterlevel", "tele_waterlevel":
		resultData.ColumnName = model_waterlevel.WaterlevelColumnForCheckErrordata
		paramWaterlevel := &model_waterlevel.Waterlevel_InputParam{}
		paramWaterlevel.Agency_id = param.AgencyID
		paramWaterlevel.Start_date = param.StartDate
		paramWaterlevel.End_date = param.EndDate
		resultErrorData, err = model_waterlevel.GetErrorWaterlevel(paramWaterlevel)
	default:
		return rest.NewError(422, "Unknown table_name id", nil)
	}

	if err != nil {
		//return errors.Repack(err)
		ctx.ReplyJSON(result.Result0(err.Error()))
		return nil
	}
	resultData.Data = resultErrorData

	//Return Data
	ctx.ReplyJSON(result.Result1(resultData))
	return nil
}

type Struct_getAgencyByStationTable struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/check_error_data_agency
// @Summary			หน่วยงาน ตามประเภทข้อมูล
// @Method			GET
// @Parameter		data_type query string example:dam_hourly ชื่อตาราง
// @Parameter		station_table query string example: m_dam ชื่อ master table ของสถานี
// @Parameter		station_column_name query string example:dam_id ชื่อ field id ของตาราง
// @Produces		json
// @Response		200	Struct_getAgencyByStationTable successful operation
func (srv *HttpService) getAgencyByStationTable(ctx service.RequestContext) error {
	//Map parameters
	p := &Struct_CheckErrordata_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	dataResult, err := model_agency.GetAgencyInStationTable(p.StationName, p.StationTableName, p.StationColumnName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
