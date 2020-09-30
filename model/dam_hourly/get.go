package dam_hourly

import (
	"database/sql"
	"encoding/json"
	//	"fmt"
	"strconv"
	"strings"
	"time"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
	float "haii.or.th/api/thaiwater30/util/float"
	"haii.or.th/api/util/errors"
	logx "haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
	//"log"
)

//	get dam hourly
//	Parameters:
//		param
//			ใช้ในส่วน Id เพื่อหาแค่ เขื่อนนั้น
//			ใช้ในส่วน Start_date, End_date เพื่อกำหนดช่วงเวลาของข้อมูล
//	Return:
//		 DamHourlyLastest_OutputParam
func GetDamHourly(param *Struct_DamHourly_InputParam) (*DamHourlyLastest_OutputParam, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data         []*Struct_Dam_Houly
		objDamHourly *Struct_Dam_Houly

		_id sql.NullInt64
		//_dam_id					sql.NullInt64
		_dam_date               time.Time
		_dam_level              sql.NullFloat64
		_dam_storage            sql.NullFloat64
		_dam_storage_percent    sql.NullFloat64
		_dam_inflow             sql.NullFloat64
		_dam_inflow_acc_percent sql.NullFloat64
		_dam_uses_water         sql.NullFloat64
		_dam_uses_water_percent sql.NullFloat64
		_dam_released           sql.NullFloat64
		_dam_spilled            sql.NullFloat64
		_dam_losses             sql.NullFloat64
		_dam_evap               sql.NullFloat64

		_result *sql.Rows
	)
	//Query
	//log.Printf(sqlGetHourlyByStationAndDate + sqlGetHourlyByStationAndDateOrderBy, param.Station_id, param.Start_date, param.End_date)
	_result, err = db.Query(sqlGetHourlyByStationAndDate+sqlGetHourlyByStationAndDateOrderBy,
		param.Dam_id, param.Start_date, param.End_date)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_Dam_Houly, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_dam_date,
			&_dam_storage, &_dam_storage_percent,
			&_dam_inflow, &_dam_inflow_acc_percent,
			&_dam_uses_water, &_dam_uses_water_percent,
			&_dam_level, &_dam_released, &_dam_spilled, &_dam_losses, &_dam_evap)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily object
		objDamHourly = &Struct_Dam_Houly{}
		objDamHourly.Dam_id, _ = strconv.ParseInt(param.Dam_id, 10, 64)
		objDamHourly.Id = _id.Int64
		objDamHourly.Dam_date = _dam_date.Format(strDatetimeFormat)
		objDamHourly.Dam_storage = ValidData(_dam_storage.Valid, _dam_storage.Float64)
		objDamHourly.Dam_storage_percent = ValidData(_dam_storage_percent.Valid, _dam_storage_percent.Float64)
		objDamHourly.Dam_inflow = ValidData(_dam_inflow.Valid, _dam_inflow.Float64)
		objDamHourly.Dam_inflow_acc_percent = ValidData(_dam_inflow_acc_percent.Valid, _dam_inflow_acc_percent.Float64)
		objDamHourly.Dam_uses_water = ValidData(_dam_uses_water.Valid, _dam_uses_water.Float64)
		objDamHourly.Dam_uses_water_percent = ValidData(_dam_uses_water_percent.Valid, _dam_uses_water_percent.Float64)
		objDamHourly.Dam_level = ValidData(_dam_level.Valid, _dam_level.Float64)
		objDamHourly.Dam_released = ValidData(_dam_released.Valid, _dam_released.Float64)
		objDamHourly.Dam_spilled = ValidData(_dam_spilled.Valid, _dam_spilled.Float64)
		objDamHourly.Dam_losses = ValidData(_dam_losses.Valid, _dam_losses.Float64)
		objDamHourly.Dam_evap = ValidData(_dam_evap.Valid, _dam_evap.Float64)
		objDamHourly.Station_type = "dam_hourly"

		data = append(data, objDamHourly)
	}

	resultData := &DamHourlyLastest_OutputParam{}
	resultData.Data = data
	resultData.Header = arrDamHourlyByStationAndDateColumn

	return resultData, nil
}

//	get dam hourly latest
//	Parameters:
//		param
//			ใช้ในส่วน Dam_date ถ้าอยากเลือกเฉพาะวันที่
//			ใช้ในส่วน Basin_id ถ้าอยากเลือกเฉพาะลุ่มน้ำ
//	Return:
//		[]Struct_DamHourly
func GetDamHourlyLastest(param *Struct_DamHourlyLastest_InputParam) ([]*Struct_DamHourly, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data         []*Struct_DamHourly
		objDamHourly *Struct_DamHourly

		_id                     sql.NullInt64
		_dam_date               time.Time
		_dam_storage            sql.NullFloat64
		_dam_storage_percent    sql.NullFloat64
		_dam_inflow             sql.NullFloat64
		_dam_inflow_acc_percent sql.NullFloat64
		_dam_uses_water         sql.NullFloat64
		_dam_uses_water_percent sql.NullFloat64
		_dam_level              sql.NullFloat64
		_dam_released           sql.NullFloat64
		_dam_spilled            sql.NullFloat64
		_dam_losses             sql.NullFloat64
		_dam_evap               sql.NullFloat64

		_dam_id   sql.NullInt64
		_dam_name sql.NullString
		_dam_lat  sql.NullFloat64
		_dam_long sql.NullFloat64
		//_dam_min_storage		sql.NullFloat64
		//_dam_max_storage		sql.NullFloat64
		//_dam_min_waterlevel	sql.NullFloat64
		//_dam_max_waterlevel	sql.NullFloat64

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_basin_id   sql.NullInt64
		_basin_code sql.NullInt64
		_basin_name sql.NullString

		_geocode_id sql.NullInt64
		_geocode    sql.NullString
		_area_code  sql.NullString

		_area_name     sql.NullString
		_province_name sql.NullString
		_amphoe_name   sql.NullString
		_tumbon_name   sql.NullString
		_province_code sql.NullString

		_dam_oldcode sql.NullString

		_sub_basin_id sql.NullInt64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""
	arrBasinID := []string{}
	if param.Basin_id != "" {
		arrBasinID = strings.Split(param.Basin_id, ",")
	}
	//Check Filter basin_id
	if len(arrBasinID) > 0 {
		if len(arrBasinID) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Basin_id, " "))
			sqlCmdWhere += " AND b.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrBasinID {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND b.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter province_id
	arrProvinceId := []string{}
	if param.Province_id != "" {
		arrProvinceId = strings.Split(param.Province_id, ",")
	}
	if len(arrProvinceId) > 0 {
		if len(arrProvinceId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Province_id, " "))
			sqlCmdWhere += " AND province_code = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND province_code IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	if param.Region_id != "" {
		sqlCmdWhere += " AND area_code like '" + param.Region_id + "' "
	}

	sqlCmdWhereMax := ""
	if param.Dam_date != "" {
		arrParam = append(arrParam, param.Dam_date+" 00:00")
		// sqlCmdWhereMax = " AND dam_datetime >= $" + strconv.Itoa(len(arrParam)) + " AND dam_datetime < $" + strconv.Itoa(len(arrParam)) + " + interval' 1 day'"
		sqlCmdWhereMax += " AND dam_datetime >= $" + strconv.Itoa(len(arrParam))
		arrParam = append(arrParam, param.Dam_date+" 23:59")
		sqlCmdWhereMax += " AND dam_datetime <= $" + strconv.Itoa(len(arrParam))

		//		sqlCmdWhereMax = " AND dam_datetime = '" + param.Dam_date + "' "
	}

	//Query
	//	fmt.Println(SQL_GetDamHourlyLastest + sqlCmdWhereMax + SQL_GetDamHourlyLastest2 + sqlCmdWhere + SQL_GetDamHourlyLastestOrderBy)
	_result, err = db.Query(SQL_GetDamHourlyLastest+sqlCmdWhereMax+SQL_GetDamHourlyLastest2+sqlCmdWhere+SQL_GetDamHourlyLastestOrderBy, arrParam...)
	if err != nil {
		logx.Log(SQL_GetDamHourlyLastest + sqlCmdWhereMax + SQL_GetDamHourlyLastest2 + sqlCmdWhere + SQL_GetDamHourlyLastestOrderBy)
		logx.Log(arrParam)
		return nil, errors.Repack(err)
	}

	defer _result.Close()

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_dam_date,
			&_dam_id, &_dam_name, &_dam_lat, &_dam_long,
			&_dam_inflow, &_dam_inflow_acc_percent,
			&_dam_storage, &_dam_storage_percent,
			&_dam_uses_water, &_dam_uses_water_percent,
			&_dam_released, &_dam_level, &_dam_spilled, &_dam_losses, &_dam_evap,
			&_agency_id, &_agency_shortname, &_agency_name,
			&_basin_id, &_basin_code, &_basin_name,
			&_geocode_id, &_geocode, &_area_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_province_code, &_dam_oldcode,
			&_sub_basin_id)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily Object
		/*----- DamDaily -----*/
		objDamHourly = &Struct_DamHourly{}
		objDamHourly.Station_type = "dam_hourly"

		objDamHourly.Id = _id.Int64
		objDamHourly.Dam_date = _dam_date.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		objDamHourly.Dam_storage = float.TwoDigit(_dam_storage.Float64)
		objDamHourly.Dam_storage_percent = float.TwoDigit(_dam_storage_percent.Float64)
		objDamHourly.Dam_inflow = float.TwoDigit(_dam_inflow.Float64)
		objDamHourly.Dam_inflow_acc_percent = float.TwoDigit(_dam_inflow_acc_percent.Float64)
		objDamHourly.Dam_uses_water = float.TwoDigit(_dam_uses_water.Float64)
		objDamHourly.Dam_uses_water_percent = float.TwoDigit(_dam_uses_water_percent.Float64)
		objDamHourly.Dam_level = float.TwoDigit(_dam_level.Float64)
		objDamHourly.Dam_released = float.TwoDigit(_dam_released.Float64)
		objDamHourly.Dam_spilled = float.TwoDigit(_dam_spilled.Float64)
		objDamHourly.Dam_losses = float.TwoDigit(_dam_losses.Float64)
		objDamHourly.Dam_evap = float.TwoDigit(_dam_evap.Float64)

		/*----- Dam -----*/
		objDamHourly.Dam = &model_dam.Struct_Dam{}
		objDamHourly.Dam.Id = _dam_id.Int64
		objDamHourly.Dam.Dam_lat, _ = _dam_lat.Value()
		objDamHourly.Dam.Dam_long, _ = _dam_long.Value()
		objDamHourly.Dam.Dam_oldcode = _dam_oldcode.String
		objDamHourly.Dam.SubBasin_id = _sub_basin_id.Int64

		if _dam_name.String == "" {
			_dam_name.String = "{}"
		}
		objDamHourly.Dam.Dam_name = json.RawMessage(_dam_name.String)

		/*----- Agency -----*/
		objDamHourly.Agency = &model_agency.Struct_Agency{}
		objDamHourly.Agency.Id = _agency_id.Int64

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		objDamHourly.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)

		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		objDamHourly.Agency.Agency_name = json.RawMessage(_agency_name.String)

		/*----- Basin -----*/
		objDamHourly.Basin = &model_basin.Struct_Basin{}
		objDamHourly.Basin.Id = _basin_id.Int64
		objDamHourly.Basin.Basin_code = _basin_code.Int64

		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		objDamHourly.Basin.Basin_name = json.RawMessage(_basin_name.String)

		/*----- Geocode -----*/
		objDamHourly.Geocode = &model_geocode.Struct_Geocode{}
		objDamHourly.Geocode.Id = _geocode_id.Int64
		objDamHourly.Geocode.Geocode = _geocode.String
		objDamHourly.Geocode.Area_code = _area_code.String

		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		objDamHourly.Geocode.Area_name = json.RawMessage(_area_name.String)

		objDamHourly.Geocode.Province_code = _province_code.String

		if _province_name.String == "" {
			_province_name.String = "{}"
		}
		objDamHourly.Geocode.Province_name = json.RawMessage(_province_name.String)

		if _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		objDamHourly.Geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)

		if _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		objDamHourly.Geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)

		//objDamHourly.ProvinceName = json.RawMessage(_province_name.String)
		//objDamHourly.Name = json.RawMessage(_dam_name.String)
		//objDamHourly.Datetime = _dam_date.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		//objDamHourly.Oldcode = _dam_oldcode.String
		//objDamHourly.Value = ValidData(_dam_storage.Valid, _dam_storage.Float64)
		//objDamHourly.DataID = _id.Int64
		//objDamHourly.StationID = _dam_id.Int64

		data = append(data, objDamHourly)
	}

	//Return Data
	return data, nil
}

//	get dam hourly ที่มีค่า error
//	Parameters:
//		param
//			ใช้ในส่วน Agency_id เพื่อหาเขื่อนเฉพาะหน่วยงาน
//	Return:
//		[]Struct_DamHourly_ErrorData
func GetErrorDamHourly(param *Struct_DamHourlyLastest_InputParam) ([]*Struct_DamHourly_ErrorData, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data         []*Struct_DamHourly_ErrorData
		objDamHourly *Struct_DamHourly_ErrorData

		_id                    sql.NullInt64
		_station_id            sql.NullInt64
		_oldcode               sql.NullString
		_date                  time.Time
		_station_name          sql.NullString
		_station_province_name sql.NullString
		_agency_name           sql.NullString
		_agency_shortname      sql.NullString

		_dam_level              sql.NullFloat64
		_dam_storage            sql.NullFloat64
		_dam_storage_percent    sql.NullFloat64
		_dam_inflow             sql.NullFloat64
		_dam_inflow_acc_percent sql.NullFloat64
		_dam_uses_water         sql.NullFloat64
		_dam_uses_water_percent sql.NullFloat64
		_dam_released           sql.NullFloat64
		_dam_spilled            sql.NullFloat64
		_dam_losses             sql.NullFloat64
		_dam_evap               sql.NullFloat64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""
	var arrAgencyID = []string{}

	if param.Agency_id != "" {
		arrAgencyID = strings.Split(param.Agency_id, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyID) > 0 {
		if len(arrAgencyID) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Agency_id, " "))
			sqlCmdWhere = " AND d.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strID := range arrAgencyID {
				arrParam = append(arrParam, strings.Trim(strID, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere = " AND d.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	if param.Start_date != "" && param.End_date != "" {
		arrParam = append(arrParam, param.Start_date)
		sqlCmdWhere += " AND dam_datetime >= $" + strconv.Itoa(len(arrParam))

		arrParam = append(arrParam, param.End_date+" 23:59")
		sqlCmdWhere += " AND dam_datetime <= $" + strconv.Itoa(len(arrParam))
	}

	//Query
	_result, err = db.Query(sqlGetErrorDamHourly+sqlCmdWhere, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Loop data result
	data = make([]*Struct_DamHourly_ErrorData, 0)
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_oldcode, &_date, &_station_name, &_station_province_name, &_agency_name, &_agency_shortname,
			&_dam_storage, &_dam_storage_percent,
			&_dam_inflow, &_dam_inflow_acc_percent,
			&_dam_uses_water, &_dam_uses_water_percent,
			&_dam_level, &_dam_released, &_dam_spilled, &_dam_losses, &_dam_evap, &_station_id)
		if err != nil {
			return nil, err
		}

		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}
		if !_station_province_name.Valid || _station_province_name.String == "" {
			_station_province_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}

		//Generate DamDaily object
		objDamHourly = &Struct_DamHourly_ErrorData{}
		objDamHourly.ID = _id.Int64
		objDamHourly.StationID = _station_id.Int64
		objDamHourly.StationOldCode = _oldcode.String
		objDamHourly.Datetime = _date.Format("2006-01-02 15:04")
		objDamHourly.StationName = json.RawMessage(_station_name.String)
		objDamHourly.ProvinceName = json.RawMessage(_station_province_name.String)
		objDamHourly.AgencyName = json.RawMessage(_agency_name.String)
		objDamHourly.AgencyShortName = json.RawMessage(_agency_shortname.String)

		objDamHourly.DamStorage = ValidData(_dam_storage.Valid, _dam_storage.Float64)
		objDamHourly.DamStoragePercent = ValidData(_dam_storage_percent.Valid, _dam_storage_percent.Float64)
		objDamHourly.DamInflow = ValidData(_dam_inflow.Valid, _dam_inflow.Float64)
		objDamHourly.DamInflowAccPercent = ValidData(_dam_inflow_acc_percent.Valid, _dam_inflow_acc_percent.Float64)
		objDamHourly.DamUsesWater = ValidData(_dam_uses_water.Valid, _dam_uses_water.Float64)
		objDamHourly.DamUsesWaterPercent = ValidData(_dam_uses_water_percent.Valid, _dam_uses_water_percent.Float64)
		objDamHourly.DamLevel = ValidData(_dam_level.Valid, _dam_level.Float64)
		objDamHourly.DamReleased = ValidData(_dam_released.Valid, _dam_released.Float64)
		objDamHourly.DamSpilled = ValidData(_dam_spilled.Valid, _dam_spilled.Float64)
		objDamHourly.DamLosses = ValidData(_dam_losses.Valid, _dam_losses.Float64)
		objDamHourly.DamEvap = ValidData(_dam_evap.Valid, _dam_evap.Float64)

		data = append(data, objDamHourly)
	}

	return data, nil
}

//	เช็คว่าค่าจาก sql.null?? ว่า valid หรือไม่
//	Parameter:
//		valid
//			.valid จากตัว sq.null??
//		value
//			ค่าจาก sql.null??
//	Return:
//		value ถ้า valid เป็น true
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
