package dam_daily

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"

	"haii.or.th/api/server/model/setting"
	float "haii.or.th/api/thaiwater30/util/float"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	tw30_sort "haii.or.th/api/thaiwater30/util/sort"
	"haii.or.th/api/util/datatype"

	"database/sql"
	"encoding/json"
	_ "log"
	"sort"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	get dam daily ของ dam id ในช่วง start date, end date
//	Parameters:
//		param
//			ใช้ในส่วนของ Dam_id, Start_date, End_date
//	Return:
//		Struct_DamDailyLastest_OutputParam
func GetDamDaily(param *Struct_DamDaily_InputParam) (*Struct_DamDailyLastest_OutputParam, error) {

	//Find datetime default format
	strDateFormat := model_setting.GetSystemSetting("bof.Default.DateFormat")
	if strDateFormat == "" {
		strDateFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_Dam_Daily
		objDamDaily *Struct_Dam_Daily

		_id sql.NullInt64
		//_dam_id				sql.NullInt64
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
	//log.Printf(sqlGetDamDailyByStationAndDate + sqlGetDamDailyByStationAndDateOrderBy , param.Station_id, param.Start_date, param.End_date)
	_result, err = db.Query(sqlGetDamDaily+sqlGetDamDailyOrderBy, param.Dam_id, param.Start_date, param.End_date)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	//defer _result.Close()

	// Loop data result
	data = make([]*Struct_Dam_Daily, 0)

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
		objDamDaily = &Struct_Dam_Daily{}
		objDamDaily.Id = _id.Int64
		objDamDaily.Dam_date = _dam_date.Format(strDateFormat)
		objDamDaily.Dam_storage = _dam_storage.Float64
		objDamDaily.Dam_storage_percent = _dam_storage_percent.Float64
		objDamDaily.Dam_inflow = ValidData(_dam_inflow.Valid, _dam_inflow.Float64)
		objDamDaily.Dam_inflow_acc_percent = ValidData(_dam_inflow_acc_percent.Valid, _dam_inflow_acc_percent.Float64)
		objDamDaily.Dam_uses_water = ValidData(_dam_uses_water.Valid, _dam_uses_water.Float64)
		objDamDaily.Dam_uses_water_percent = ValidData(_dam_uses_water_percent.Valid, _dam_uses_water_percent.Float64)
		objDamDaily.Dam_level = ValidData(_dam_level.Valid, _dam_level.Float64)
		objDamDaily.Dam_released = ValidData(_dam_released.Valid, _dam_released.Float64)
		objDamDaily.Dam_spilled = ValidData(_dam_spilled.Valid, _dam_spilled.Float64)
		objDamDaily.Dam_losses = ValidData(_dam_losses.Valid, _dam_losses.Float64)
		objDamDaily.Dam_evap = ValidData(_dam_evap.Valid, _dam_evap.Float64)

		data = append(data, objDamDaily)
	}

	resultData := &Struct_DamDailyLastest_OutputParam{}
	resultData.Data = data
	resultData.Header = arrDamDailyLastestColumn

	return resultData, err
}

//	get dam daily ที่มีค่า error(-9999, 999999)
//	Parameters:
//		param
//			ใช้ในส่วน Agency_id
//	Return:
//		[]Struct_DamDaily_ErrorData
func GetErrorDamDaily(param *Struct_DamDailyLastest_InputParam) ([]*Struct_DamDaily_ErrorData, error) {

	//Find datetime default format
	strDateFormat := model_setting.GetSystemSetting("bof.Default.DateFormat")
	if strDateFormat == "" {
		strDateFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_DamDaily_ErrorData
		objDamDaily *Struct_DamDaily_ErrorData

		_station_id            sql.NullInt64
		_oldcode               sql.NullString
		_date                  time.Time
		_station_name          sql.NullString
		_station_province_name sql.NullString
		_agency_name           sql.NullString
		_agency_shortname      sql.NullString

		_id                     sql.NullInt64
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
		sqlCmdWhere += " AND dam_date >= $" + strconv.Itoa(len(arrParam))

		arrParam = append(arrParam, param.End_date)
		sqlCmdWhere += " AND dam_date <= $" + strconv.Itoa(len(arrParam))
	}

	//Query
	_result, err = db.Query(sqlGetErrorDamDaily+sqlCmdWhere, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Loop data result
	data = make([]*Struct_DamDaily_ErrorData, 0)
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
		objDamDaily = &Struct_DamDaily_ErrorData{}
		objDamDaily.ID = _id.Int64
		objDamDaily.StationID = _station_id.Int64
		objDamDaily.StationOldCode = _oldcode.String
		objDamDaily.Datetime = _date.Format(strDateFormat)
		objDamDaily.StationName = json.RawMessage(_station_name.String)
		objDamDaily.ProvinceName = json.RawMessage(_station_province_name.String)
		objDamDaily.AgencyName = json.RawMessage(_agency_name.String)
		objDamDaily.AgencyShortName = json.RawMessage(_agency_shortname.String)

		objDamDaily.DamStorage = ValidData(_dam_storage.Valid, _dam_storage.Float64)
		objDamDaily.DamStoragePercent = ValidData(_dam_storage_percent.Valid, _dam_storage_percent.Float64)
		objDamDaily.DamInflow = ValidData(_dam_inflow.Valid, _dam_inflow.Float64)
		objDamDaily.DamInflowAccPercent = ValidData(_dam_inflow_acc_percent.Valid, _dam_inflow_acc_percent.Float64)
		objDamDaily.DamUsesWater = ValidData(_dam_uses_water.Valid, _dam_uses_water.Float64)
		objDamDaily.DamUsesWaterPercent = ValidData(_dam_uses_water_percent.Valid, _dam_uses_water_percent.Float64)
		objDamDaily.DamLevel = ValidData(_dam_level.Valid, _dam_level.Float64)
		objDamDaily.DamReleased = ValidData(_dam_released.Valid, _dam_released.Float64)
		objDamDaily.DamSpilled = ValidData(_dam_spilled.Valid, _dam_spilled.Float64)
		objDamDaily.DamLosses = ValidData(_dam_losses.Valid, _dam_losses.Float64)
		objDamDaily.DamEvap = ValidData(_dam_evap.Valid, _dam_evap.Float64)

		data = append(data, objDamDaily)
	}

	return data, nil
}

//	get ข้อมูลเขื่อนล่าสุด
//	Parameters:
//		param
//			ใช้ในส่วน Agency_id, Basin_id, Dam_date
//	Return:
//		[]Struct_DamDaily
func GetDamDailyLastest(param *Struct_DamDailyLastest_InputParam) ([]*Struct_DamDaily, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02"
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_DamDaily
		objDamDaily *Struct_DamDaily

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

		_dam_id      sql.NullInt64
		_dam_name    sql.NullString
		_dam_lat     sql.NullFloat64
		_dam_long    sql.NullFloat64
		_max_storage sql.NullFloat64

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_basin_id   sql.NullInt64
		_basin_code sql.NullInt64
		_basin_name sql.NullString

		_geocode_id    sql.NullInt64
		_geocode       sql.NullString
		_rid_area_code sql.NullString
		_rid_area_name sql.NullString
		_province_name sql.NullString
		_amphoe_name   sql.NullString
		_tumbon_name   sql.NullString
		_province_code sql.NullString

		_dam_oldcode 	sql.NullString

		_sub_basin_id 	sql.NullInt64

		_cctv_id		sql.NullInt64
		_cctv_url		sql.NullString
		_cctv_filename	sql.NullString

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""
	arrAgencyId := []string{}
	arrBasinId := []string{}

	if param.Agency_id != "" {
		arrAgencyId = strings.Split(param.Agency_id, ",")
	}
	if param.Basin_id != "" {
		arrBasinId = strings.Split(param.Basin_id, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyId) > 0 {
		if len(arrAgencyId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Agency_id, " "))
			sqlCmdWhere += " AND st.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrAgencyId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND st.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter basin_id
	if len(arrBasinId) > 0 {
		if len(arrBasinId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Basin_id, " "))
			sqlCmdWhere += " AND b.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrBasinId {
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
		sqlCmdWhere += " AND rid_area_code like '" + param.Region_id + "' "
	}

	// sqlCmdWhereMax := ""
	// if param.Dam_date != "" {
	// 	sqlCmdWhereMax = " AND dam_date = '" + param.Dam_date + "' "
	// }
	qry := SQL_GetDamDaily_
	if param.Dam_date != "" {
		qry = strings.Replace(qry, "--replace", "dd.id", 1)
		qry += " FROM dam_daily dd "
		sqlCmdWhere += " AND dam_date = '" + param.Dam_date + "' "
	} else {
		qry = strings.Replace(qry, "--replace", "dd.master_row_id", 1)
		qry += " FROM latest.dam_daily dd "
	}
	qry += SQL_GetDamDaily_2

	//Query
	//	log.Println(SQL_GetDamDailyLastest + sqlCmdWhereMax + SQL_GetDamDailyLastest2 + sqlCmdWhere + SQL_GetDamDailyLastestOrderBy)
	// _result, err = db.Query(SQL_GetDamDailyLastest+sqlCmdWhereMax+SQL_GetDamDailyLastest2+sqlCmdWhere+SQL_GetDamDailyLastestOrderBy, arrParam...)
	_result, err = db.Query(qry+sqlCmdWhere+SQL_GetDamDailyLastestOrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	//	arrResult := map[int64]*Struct_DamDailyLastest_Output{}

	data = make([]*Struct_DamDaily, 0)
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_dam_date,
			&_dam_id, &_dam_name, &_dam_lat, &_dam_long, &_max_storage,
			&_dam_inflow, &_dam_inflow_acc_percent,
			&_dam_storage, &_dam_storage_percent,
			&_dam_uses_water, &_dam_uses_water_percent,
			&_dam_released, &_dam_level, &_dam_spilled, &_dam_losses, &_dam_evap,
			&_agency_id, &_agency_shortname, &_agency_name,
			&_basin_id, &_basin_code, &_basin_name,
			&_geocode_id, &_geocode, &_rid_area_code, &_rid_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_province_code, &_dam_oldcode,
			&_sub_basin_id,
			&_cctv_id, &_cctv_url, &_cctv_filename,
		)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily Object
		/*----- DamDaily -----*/
		objDamDaily = &Struct_DamDaily{}
		objDamDaily.Station_type = "dam_daily"

		objDamDaily.Id = _id.Int64
		objDamDaily.Dam_date = _dam_date.Format(strDatetimeFormat)
		objDamDaily.Dam_storage = float.TwoDigit(_dam_storage.Float64)
		objDamDaily.Dam_storage_percent = float.TwoDigit(_dam_storage_percent.Float64)
		objDamDaily.Dam_inflow = float.TwoDigit(_dam_inflow.Float64)
		objDamDaily.Dam_inflow_acc_percent = float.TwoDigit(_dam_inflow_acc_percent.Float64)
		objDamDaily.Dam_uses_water = float.TwoDigit(_dam_uses_water.Float64)
		objDamDaily.Dam_uses_water_percent = float.TwoDigit(_dam_uses_water_percent.Float64)
		objDamDaily.Dam_level = float.TwoDigit(_dam_level.Float64)
		objDamDaily.Dam_released = float.TwoDigit(_dam_released.Float64)
		objDamDaily.Dam_spilled = float.TwoDigit(_dam_spilled.Float64)
		objDamDaily.Dam_losses = float.TwoDigit(_dam_losses.Float64)
		objDamDaily.Dam_evap = float.TwoDigit(_dam_evap.Float64)

		/*----- Dam -----*/
		objDamDaily.Dam = &model_dam.Struct_D{}
		objDamDaily.Dam.Id = _dam_id.Int64
		objDamDaily.Dam.Dam_lat, _ = _dam_lat.Value()
		objDamDaily.Dam.Dam_long, _ = _dam_long.Value()
		objDamDaily.Dam.Max_storage, _ = _max_storage.Value()
		objDamDaily.Dam.Dam_oldcode = _dam_oldcode.String
		objDamDaily.Dam.SubBasin_id = _sub_basin_id.Int64

		/* -- CCTV --*/
		objDamDaily.Cctv.Cctv_id = _cctv_id.Int64
		objDamDaily.Cctv.Cctv_url = _cctv_url.String + _cctv_filename.String

		if _dam_name.String == "" {
			_dam_name.String = "{}"
		}
		objDamDaily.Dam.Dam_name = json.RawMessage(_dam_name.String)

		/*----- Agency -----*/
		objDamDaily.Agency = &model_agency.Struct_Agency{}
		objDamDaily.Agency.Id = _agency_id.Int64

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		objDamDaily.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)

		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		objDamDaily.Agency.Agency_name = json.RawMessage(_agency_name.String)

		/*----- Basin -----*/
		objDamDaily.Basin = &model_basin.Struct_Basin{}
		objDamDaily.Basin.Id = _basin_id.Int64
		objDamDaily.Basin.Basin_code = _basin_code.Int64

		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		objDamDaily.Basin.Basin_name = json.RawMessage(_basin_name.String)

		/*----- Geocode -----*/
		objDamDaily.Geocode = &model_geocode.Struct_Geocode{}
		objDamDaily.Geocode.Id = _geocode_id.Int64
		objDamDaily.Geocode.Geocode = _geocode.String
		objDamDaily.Geocode.Area_code = _rid_area_code.String

		if _rid_area_name.String == "" {
			_rid_area_name.String = "{}"
		}
		objDamDaily.Geocode.Area_name = json.RawMessage(_rid_area_name.String)

		objDamDaily.Geocode.Province_code = _province_code.String

		if _province_name.String == "" {
			_province_name.String = "{}"
		}
		objDamDaily.Geocode.Province_name = json.RawMessage(_province_name.String)

		if _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		objDamDaily.Geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)

		if _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		objDamDaily.Geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)

		data = append(data, objDamDaily)
	}
	return data, nil
}

//	get dam daily graph
//	Parameters:
//		param
//			ใช้ Dam_type เพื่อบอกว่าต้องการค่าจาก คอลั่มไหน
//			ใช้ Year เพื่อบอกว่าต้องการข้อมูลย้อนหลังกี่ปี
//	Return:
//		Struct_DamGraph
//			min_storage, max_storage, ข้อมูลกราฟย้อนหลังตามจำนวนปี
func GetDamGraph(param *Struct_GetDamGraph_InputParam) (*Struct_DamGraph, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		_time  time.Time
		_value sql.NullFloat64
	)
	// query
	_row, err := db.Query(SQL_GetDamGraph(param.Dam_type, param.Year), param.Dam_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	rs := &Struct_DamGraph{}

	for _row.Next() {
		_row.Scan(&_time, &_value)

		rs_data := &Struct_GraphData{}
		rs_data.Date = _time.Format("2006-01-02")
		rs_data.Value, _ = _value.Value()

		rs.Data = append(rs.Data, rs_data)
	}
	// ดึงข้อมูลเขื่อน
	master_dams, err := model_dam.GetDamFromId(param.Dam_id)
	if err != nil {
		return nil, err
	}
	master_dam := master_dams[0]

	rs.Max_storage = master_dam.Max_storage
	rs.Min_storage = master_dam.Min_storage

	return rs, nil
}

// Get infomation dam_daily for analyst
//  Parameters:
//		inputData
//				GraphAnalystDamDailyInput
//  Return:
//		Array damdaily data
func GetDamGraphDailyAnalyst(inputData *GraphAnalystDamDailyInput) ([]*GraphAnalystDamDailyOutput, error) {

	// get sql
	q := sqlSelectDamDailyGraphAnalyst
	inputDamID := ""
	p := []interface{}{}
	// add condition
	for i, v := range inputData.DamID {
		if i > 0 {
			inputDamID += " OR dd.dam_id=$" + strconv.Itoa(i+1)
			p = append(p, v)
		} else {
			inputDamID = "dd.dam_id=$" + strconv.Itoa(i+1)
			p = append(p, v)
		}
	}

	if inputDamID != "" {
		q += " AND (" + inputDamID + ")"
	}

	inputDamDate := ""
	count := len(inputData.DamID)
	dd := make(tw30_sort.DataRange, 0)
	dd = append(dd, inputData.Year...)
	sort.Sort(dd)
	inputData.Year = dd
	// conbine sql
	for i, v := range inputData.Year {
		t, _ := time.Parse("2006-1-2", strconv.FormatInt(v, 10)+"-"+strconv.FormatInt(inputData.Month, 10)+"-"+strconv.FormatInt(inputData.Day, 10))
		damDate := t.Format("2006-01-02")
		if i > 0 {
			inputDamDate += " OR dd.dam_date=$" + strconv.Itoa(count+1)
		} else {
			inputDamDate = "dd.dam_date=$" + strconv.Itoa(count+1)
		}
		p = append(p, damDate)
		count++
	}

	if inputDamDate != "" {
		q += " AND (" + inputDamDate + ")"
	}

	q += " GROUP BY dd.dam_date ORDER BY dd.dam_date DESC "

	// open db
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	dataRow := make(map[int]interface{})
	defer rows.Close()
	for rows.Next() {
		var (
			dam_date         time.Time
			dam_storage      sql.NullFloat64
			dam_inflow       sql.NullFloat64
			dam_released     sql.NullFloat64
			dam_spilled      sql.NullFloat64
			dam_losses       sql.NullFloat64
			dam_evap         sql.NullFloat64
			dam_uses_water   sql.NullFloat64
			dam_inflow_avg   sql.NullFloat64
			dam_released_acc sql.NullFloat64
			dam_inflow_acc   sql.NullFloat64
		)
		rows.Scan(&dam_date, &dam_storage, &dam_inflow, &dam_released, &dam_spilled, &dam_losses, &dam_evap, &dam_uses_water, &dam_inflow_avg, &dam_released_acc, &dam_inflow_acc)

		// select value for return
		dd := dam_date.Year()
		if inputData.DataType == "dam_storage" {
			dataRow[dd] = ValidData(dam_storage.Valid, dam_storage.Float64)
		} else if inputData.DataType == "dam_inflow" {
			dataRow[dd] = ValidData(dam_inflow.Valid, dam_inflow.Float64)
		} else if inputData.DataType == "dam_released" {
			dataRow[dd] = ValidData(dam_released.Valid, dam_released.Float64)
		} else if inputData.DataType == "dam_spilled" {
			dataRow[dd] = ValidData(dam_spilled.Valid, dam_spilled.Float64)
		} else if inputData.DataType == "dam_losses" {
			dataRow[dd] = ValidData(dam_losses.Valid, dam_losses.Float64)
		} else if inputData.DataType == "dam_evap" {
			dataRow[dd] = ValidData(dam_evap.Valid, dam_evap.Float64)
		} else if inputData.DataType == "dam_uses_water" {
			dataRow[dd] = ValidData(dam_uses_water.Valid, dam_uses_water.Float64)
		} else if inputData.DataType == "dam_inflow_avg" {
			dataRow[dd] = ValidData(dam_inflow_avg.Valid, dam_inflow_avg.Float64)
		} else if inputData.DataType == "dam_released_acc" {
			dataRow[dd] = ValidData(dam_released_acc.Valid, dam_released_acc.Float64)
		} else if inputData.DataType == "dam_inflow_acc" {
			dataRow[dd] = ValidData(dam_inflow_acc.Valid, dam_inflow_acc.Float64)
		} else {
			return nil, rest.NewError(422, "No Data Type", nil)
		}
	}

	data := make([]*GraphAnalystDamDailyOutput, 0)

	// loop add by year
	for _, v := range inputData.Year {
		dd := &GraphAnalystDamDailyOutput{}
		dd.Year = v
		dd.Date = datatype.MakeString(v) + "-" + strconv.FormatInt(inputData.Month, 10) + "-" + strconv.FormatInt(inputData.Day, 10)
		dd.Data = dataRow[int(v)]

		data = append(data, dd)
	}

	return data, nil
}

// Get infomation four dam main
//  Parameters:
//		None
//  Return:
//		Array information dam
func GetDamFourMain() ([]*MonitoringDamOutput, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql get latest dam
	q := sqlSelectFourDamLatest
	p := []interface{}{1, 12, 36, 11}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	defer rows.Close()

	dataRow := make([]*MonitoringDamOutput, 0)
	for rows.Next() {
		var (
			dam_id              string
			dam_date            time.Time
			dam_name            pqx.JSONRaw
			dam_storage         sql.NullFloat64
			dam_inflow          sql.NullFloat64
			dam_released        sql.NullFloat64
			dam_storage_percent sql.NullFloat64
			dam_uses_water      sql.NullFloat64
		)
		rows.Scan(&dam_date, &dam_id, &dam_name, &dam_storage, &dam_inflow, &dam_released, &dam_storage_percent, &dam_uses_water)
		data := &MonitoringDamOutput{}
		data.DamId = dam_id
		data.DamDate = dam_date.Format("2006-01-02")
		data.DamName = dam_name.JSON()
		data.DamStorage = ValidData(dam_storage.Valid, dam_storage.Float64)
		data.DamInflow = ValidData(dam_inflow.Valid, dam_inflow.Float64)
		data.DamReleased = ValidData(dam_released.Valid, dam_released.Float64)
		data.DamStoragePercent = ValidData(dam_storage_percent.Valid, dam_storage_percent.Float64)
		data.DamUsesWater = ValidData(dam_uses_water.Valid, dam_uses_water.Float64)
		// add value to array
		dataRow = append(dataRow, data)
	}

	return dataRow, nil
}

// check valid data return valid return data, invalid return null
//  Parameters:
//		valid
//			boolean check valid
//		value
//			value scan from db
//  Return:
//		if true return valid return data if not invalid return null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

//	ข้อมูลน้ำในเขื่อนหลัก
//	Returns:
//		array Struct_WaterInformation
func GetWaterInformationMainDam() ([]*Struct_WaterInformation, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// คิวรี่ เขื่อนไอดี  1, 12, 36, 11
	// remove criteria " AND ( dd.qc_status IS NULL OR dd.qc_status->>'is_pass' = 'true' )"
	// sometimes qc_status not pass so no data show on mobile
	var q string = `
	SELECT m.id
	       , m.dam_name ->> 'th' AS dam_name
	       , m.dam_lat
	       , m.dam_long
	       , m.normal_storage
	       , dd.dam_uses_water
	       , dd.dam_uses_water_percent
	FROM   latest.dam_daily dd
	       INNER JOIN m_dam m
	               ON dd.dam_id = m.id
	WHERE  m.id IN ( 1, 12, 36, 11 )
	ORDER  BY CASE
	            WHEN m.id = 1 THEN 1
	            WHEN m.id = 12 THEN 2
	            WHEN m.id = 36 THEN 3
	            WHEN m.id = 11 THEN 4
	          END ASC
	`

	// แปลง Frontend.public.dam_scale_color จาก setting ให้เป็น uSetting.Struct_DamScaleColor
	dam_scale_color := &uSetting.Struct_DamScaleColor{}
	err = setting.GetSystemSettingPtr("Frontend.public.dam_scale_color", &dam_scale_color)
	if err != nil {
		return nil, err
	}

	row, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	rs := make([]*Struct_WaterInformation, 0)
	// dam 4 เขื่อนหลักลุ่มน้ำเจ้าพระยา fix ค่าไว้
	dam0 := &Struct_WaterInformation{
		Dam_id:   "0",
		Dam_name: "4 เขื่อนหลักลุ่มน้ำเจ้าพระยา",
		Dam_lat:  "13.7563",
		Dam_long: "100.502",
	}
	rs = append(rs, dam0)

	var sumNomarlStorage float64 = 0 // sum normal_storage ใช้คำนวณตอนสุดท้าย
	var sumUsesWater float64 = 0     // sum dam_uses_water ใช้คำนวณตอนสุดท้าย

	for row.Next() {
		var (
			_id                     sql.NullInt64
			_dam_name               sql.NullString
			_dam_lat                sql.NullFloat64
			_dam_long               sql.NullFloat64
			_normal_storage         sql.NullFloat64
			_dam_uses_water         sql.NullFloat64
			_dam_uses_water_percent sql.NullFloat64
		)
		err = row.Scan(&_id, &_dam_name, &_dam_lat, &_dam_long, &_normal_storage, &_dam_uses_water, &_dam_uses_water_percent)
		if err != nil {
			return nil, err
		}
		dam_uses_water := float.TwoDigit(_dam_uses_water.Float64)                 // ปรับทศนิยมเป็น 2 ตำแหน่ง
		dam_uses_water_percent := float.TwoDigit(_dam_uses_water_percent.Float64) // ปรับทศนิยมเป็น 2 ตำแหน่ง
		d := &Struct_WaterInformation{
			Dam_id:                 datatype.MakeString(_id.Int64),
			Dam_name:               "เขื่อน" + _dam_name.String,
			Dam_lat:                datatype.MakeString(_dam_lat.Float64),
			Dam_long:               datatype.MakeString(_dam_long.Float64),
			Dam_uses_water:         datatype.MakeString(dam_uses_water),
			Dam_uses_water_percent: datatype.MakeString(dam_uses_water_percent),
		}
		sumUsesWater += _dam_uses_water.Float64
		sumNomarlStorage += _normal_storage.Float64

		st := dam_scale_color.CompareScale(dam_uses_water_percent) // หา setting ตามค่า dam_uses_water_percent
		if st != nil {
			d.Status_color = st.Color
		}

		rs = append(rs, d)
	}
	if sumUsesWater != 0 && sumNomarlStorage != 0 {
		dam0_uses_water_percent := (sumUsesWater / sumNomarlStorage) * 100 // คำนวณ uses_water_percent ของ dam0

		dam0.Dam_uses_water = datatype.MakeString(float.OneDigit(sumUsesWater))
		dam0.Dam_uses_water_percent = datatype.MakeString(float.NoDigit(dam0_uses_water_percent))

		st := dam_scale_color.CompareScale(dam0_uses_water_percent) // หา setting ตามค่า dam0_uses_water_percent
		if st != nil {
			dam0.Status_color = st.Color
		}
	}

	return rs, nil
}

//	ข้อมูลน้ำในเขื่อนหลัก
//	Returns:
//		array Struct_WaterInformation_DamList
func GetWaterInformationDamList() ([]*Struct_WaterInformation_DamList, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// query
	q := `
SELECT     m.id
           , m.dam_name->> 'th' AS dam_name
           , g.province_code
           , g.province_name ->> 'th' AS province_name
           , dd.dam_date
           , dd.dam_storage
           , dd.dam_storage_percent
           , dd.dam_uses_water
           , dd.dam_uses_water_percent
           , dd.dam_inflow
           , dd.dam_inflow_acc_percent
           , dd.dam_released
           , m.dam_lat
           , m.dam_long
           , m.normal_storage
           , dd.dam_uses_water_percent AS uses_water_percent
           , g.amphoe_name->> 'th' AS amphoe_name
           , g.tumbon_name->> 'th' AS district_name
FROM       (
                    SELECT   dam_id
                             , MAX ( dam_date ) AS dam_date
                    FROM     dam_daily
                    WHERE    dam_date > $1
                    GROUP BY dam_id
           ) d
INNER JOIN dam_daily dd ON dd.dam_id    = d.dam_id AND dd.dam_date = d.dam_date
INNER JOIN m_dam m      ON d.dam_id     = m.id AND dd.dam_id = m.id AND m.agency_id=12
INNER JOIN lt_geocode g ON m.geocode_id = g.id
WHERE      dd.dam_date                  > $1
-- AND ( dd.qc_status IS NULL OR dd.qc_status->>'is_pass' = 'true' )
ORDER BY
           CASE
                      WHEN m.id = 1 THEN 1
                      WHEN m.id = 12 THEN 2
                      WHEN m.id = 36 THEN 3
                      WHEN m.id = 11 THEN 4
           END
           , dam_storage_percent DESC
	`
	// แปลง Frontend.public.dam_scale_color จาก setting ให้เป็น uSetting.Struct_DamScaleColor
	dam_scale_color := &uSetting.Struct_DamScaleColor{}
	err = setting.GetSystemSettingPtr("Frontend.public.dam_scale_color", &dam_scale_color)
	if err != nil {
		return nil, err
	}
	currentDate := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	// execute query
	row, err := db.Query(q, currentDate)
	if err != nil {
		return nil, err
	}
	// วนลูปใส่ struct
	rs := make([]*Struct_WaterInformation_DamList, 0)
	for row.Next() {
		var (
			_dam_id                 int64
			_dam_name               sql.NullString
			_province_id            int64
			_province_name          sql.NullString
			_dam_date               sql.NullString
			_dam_storage            sql.NullFloat64
			_dam_storage_percent    sql.NullFloat64
			_dam_uses_water         sql.NullFloat64
			_dam_uses_water_percent sql.NullFloat64
			_dam_inflow             sql.NullFloat64
			_dam_inflow_acc_percent sql.NullFloat64
			_dam_released           sql.NullFloat64
			_dam_lat                sql.NullFloat64
			_dam_long               sql.NullFloat64
			_normal_storage         sql.NullFloat64
			_uses_water_percent     sql.NullFloat64
			_amphoe_name            sql.NullString
			_district_name          sql.NullString
		)
		err = row.Scan(&_dam_id, &_dam_name, &_province_id, &_province_name, &_dam_date, &_dam_storage, &_dam_storage_percent, &_dam_uses_water, &_dam_uses_water_percent,
			&_dam_inflow, &_dam_inflow_acc_percent, &_dam_released, &_dam_lat, &_dam_long, &_normal_storage, &_uses_water_percent, &_amphoe_name, &_district_name)
		if err != nil {
			return nil, err
		}
		d := &Struct_WaterInformation_DamList{
			Dam_id:                 datatype.MakeString(_dam_id),
			Dam_name:               "เขื่อน" + _dam_name.String,
			Province_id:            _province_id,
			Province_name:          _province_name.String,
			Dam_date:               pqx.NullStringToTime(_dam_date).Format(setting.GetSystemSetting("setting.Default.DateFormat")),
			Dam_storage:            float.Comma(_dam_storage.Float64, 0),
			Dam_storage_percent:    float.Comma(_dam_storage_percent.Float64, 2) + "%",
			Dam_uses_water:         float.Comma(_dam_uses_water.Float64, 0),
			Dam_uses_water_percent: float.Comma(_dam_uses_water_percent.Float64, 2) + "%",
			Dam_inflow:             float.Comma(_dam_inflow.Float64, 0),
			Dam_inflow_acc_percent: float.Comma(_dam_inflow_acc_percent.Float64, 2) + "%",
			Dam_released:           float.Comma(_dam_released.Float64, 0),
			Dam_lat:                datatype.MakeString(_dam_lat.Float64),
			Dam_long:               datatype.MakeString(_dam_long.Float64),
			Normal_storage:         float.Comma(_normal_storage.Float64, 0),
			Uses_water_percent:     datatype.MakeString(float.TwoDigit(_uses_water_percent.Float64)),
			Amphoe_name:            _amphoe_name.String,
			District_name:          _district_name.String,
		}

		// หา config จาก dam_storage
		dam_storage := dam_scale_color.CompareScale(_dam_storage_percent.Float64)
		if dam_storage != nil {
			d.Dam_storage_status_color = dam_storage.Color
			d.Dam_storage_status_level = datatype.MakeInt(dam_storage.Level)
		}
		// หา config จาก dam_uses
		dam_use := dam_scale_color.CompareScale(_dam_uses_water_percent.Float64)
		if dam_use != nil {
			d.Dam_uses_status_color = dam_use.Color
		}

		rs = append(rs, d)
	}

	return rs, nil
}

//	ข้อมูลกราฟเขื่อน ตามปี
//	Parameters:
//		dam_id
//			รหัสเขื่อน
//		startYear
//			ปีเริ่มต้น
//		endYear
//			ปีสิ้นสุด
//	Return:
//		array Struct_DamGraphHistory
func GetDamGraphByYear(dam_id, startYear, endYear string) ([]*Struct_DamGraphHistory, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// validate dam_id
	if dam_id == "" {
		return nil, errors.New("no dam_id")
	}

	// query
	q := `
SELECT gen.date
       , dam_storage
FROM   (SELECT Generate_series('` + startYear + `-01-01' :: date, '` + endYear + `-12-31', '1 day') :: date AS date) gen
       LEFT JOIN (SELECT *
                  FROM   dam_daily
                  WHERE  dam_date >= '` + startYear + `-01-01'
		                 AND dam_date <= '` + endYear + `-12-31'
                         AND dam_id = $1) d
			  ON gen.date = d.dam_date

ORDER  BY date ASC
	`
	//WHERE dam_storage != -9999
	// execute query
	row, err := db.Query(q, dam_id)
	if err != nil {
		return nil, err
	}

	// วนลูปใส่ struct
	rs := make([]*Struct_DamGraphHistory, 0)
	for row.Next() {
		var (
			_date        time.Time
			_dam_storage sql.NullFloat64
		)
		err = row.Scan(&_date, &_dam_storage)
		if err != nil {
			return nil, err
		}

		//  dam_storage เป็น null ให้แปลงเป็น -9999
		if !_dam_storage.Valid {
			_dam_storage.Float64 = -9999
		}

		d := &Struct_DamGraphHistory{
			Dt_date:     _date.Format(setting.GetSystemSetting("setting.Default.DateFormat")),
			Dam_storage: float.String(_dam_storage.Float64, 2),
		}
		rs = append(rs, d)
	}

	return rs, nil
}

//	get ข้อมูลเขื่อนล่าสุด รายจังหวัด
//	Parameters:
//		param
//			ใช้ในส่วน Agency_id, Basin_id, Dam_date
//	Return:
//		[]Struct_DamDaily
func GetDamDailyLastestProvince(param *Struct_DamDailyLastest_InputParam) ([]*Struct_DamDaily, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02"
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_DamDaily
		objDamDaily *Struct_DamDaily

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

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_basin_id   sql.NullInt64
		_basin_code sql.NullInt64
		_basin_name sql.NullString

		_geocode_id    sql.NullInt64
		_geocode       sql.NullString
		_rid_area_code sql.NullString
		_rid_area_name sql.NullString
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
	arrAgencyId := []string{}
	arrBasinId := []string{}

	if param.Agency_id != "" {
		arrAgencyId = strings.Split(param.Agency_id, ",")
	}
	if param.Basin_id != "" {
		arrBasinId = strings.Split(param.Basin_id, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyId) > 0 {
		if len(arrAgencyId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Agency_id, " "))
			sqlCmdWhere += " AND st.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrAgencyId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND st.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter basin_id
	if len(arrBasinId) > 0 {
		if len(arrBasinId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Basin_id, " "))
			sqlCmdWhere += " AND b.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrBasinId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND b.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	// check province dam near
	if param.Province_id != "" {
		db, err := pqx.Open()
		if err != nil {
			return nil, err
		}
		var q string = sqlSelectMappingNearDam
		rows, err := db.Query(q, param.Province_id)
		if err != nil {
			return nil, err
		}
		var rs string
		for rows.Next() {
			var (
				_id_dam sql.NullString
			)
			err = rows.Scan(&_id_dam)
			if err != nil {
				return nil, err
			}
			rs = _id_dam.String
		}
		sqlCmdWhere += " AND st.id IN (" + rs + ")"
		sqlCmdWhere += " AND province_code <> '" + param.Province_id + "'"
	}

	if param.Region_id != "" {
		sqlCmdWhere += " AND rid_area_code like '" + param.Region_id + "' "
	}
	sqlCmdWhereMax := ""
	if param.Dam_date != "" {
		sqlCmdWhereMax = " AND dam_date = '" + param.Dam_date + "' "
	}

	//Query
	//	if(SQL_GetDamDailyLastest != ""){
	//		return nil,errors.New(SQL_GetDamDailyLastest+sqlCmdWhereMax+SQL_GetDamDailyLastest2+sqlCmdWhere+SQL_GetDamDailyLastestProvinceOrderBy)
	//	}
	_result, err = db.Query(SQL_GetDamDailyLastestProvince+sqlCmdWhereMax+SQL_GetDamDailyLastestProvince_2+sqlCmdWhere+SQL_GetDamDailyLastestProvinceOrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	//	arrResult := map[int64]*Struct_DamDailyLastest_Output{}

	data = make([]*Struct_DamDaily, 0)
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
			&_geocode_id, &_geocode, &_rid_area_code, &_rid_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_province_code, &_dam_oldcode,
			&_sub_basin_id)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily Object
		/*----- DamDaily -----*/
		objDamDaily = &Struct_DamDaily{}
		objDamDaily.Station_type = "dam_daily"

		objDamDaily.Id = _id.Int64
		objDamDaily.Dam_date = _dam_date.Format(strDatetimeFormat)
		objDamDaily.Dam_storage = float.TwoDigit(_dam_storage.Float64)
		objDamDaily.Dam_storage_percent = float.TwoDigit(_dam_storage_percent.Float64)
		objDamDaily.Dam_inflow = float.TwoDigit(_dam_inflow.Float64)
		objDamDaily.Dam_inflow_acc_percent = float.TwoDigit(_dam_inflow_acc_percent.Float64)
		objDamDaily.Dam_uses_water = float.TwoDigit(_dam_uses_water.Float64)
		objDamDaily.Dam_uses_water_percent = float.TwoDigit(_dam_uses_water_percent.Float64)
		objDamDaily.Dam_level = float.TwoDigit(_dam_level.Float64)
		objDamDaily.Dam_released = float.TwoDigit(_dam_released.Float64)
		objDamDaily.Dam_spilled = float.TwoDigit(_dam_spilled.Float64)
		objDamDaily.Dam_losses = float.TwoDigit(_dam_losses.Float64)
		objDamDaily.Dam_evap = float.TwoDigit(_dam_evap.Float64)

		/*----- Dam -----*/
		objDamDaily.Dam = &model_dam.Struct_D{}
		objDamDaily.Dam.Id = _dam_id.Int64
		objDamDaily.Dam.Dam_lat, _ = _dam_lat.Value()
		objDamDaily.Dam.Dam_long, _ = _dam_long.Value()
		objDamDaily.Dam.Dam_oldcode = _dam_oldcode.String
		objDamDaily.Dam.SubBasin_id = _sub_basin_id.Int64

		if _dam_name.String == "" {
			_dam_name.String = "{}"
		}
		objDamDaily.Dam.Dam_name = json.RawMessage(_dam_name.String)

		/*----- Agency -----*/
		objDamDaily.Agency = &model_agency.Struct_Agency{}
		objDamDaily.Agency.Id = _agency_id.Int64

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		objDamDaily.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)

		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		objDamDaily.Agency.Agency_name = json.RawMessage(_agency_name.String)

		/*----- Basin -----*/
		objDamDaily.Basin = &model_basin.Struct_Basin{}
		objDamDaily.Basin.Id = _basin_id.Int64
		objDamDaily.Basin.Basin_code = _basin_code.Int64

		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		objDamDaily.Basin.Basin_name = json.RawMessage(_basin_name.String)

		/*----- Geocode -----*/
		objDamDaily.Geocode = &model_geocode.Struct_Geocode{}
		objDamDaily.Geocode.Id = _geocode_id.Int64
		objDamDaily.Geocode.Geocode = _geocode.String
		objDamDaily.Geocode.Area_code = _rid_area_code.String

		if _rid_area_name.String == "" {
			_rid_area_name.String = "{}"
		}
		objDamDaily.Geocode.Area_name = json.RawMessage(_rid_area_name.String)

		objDamDaily.Geocode.Province_code = _province_code.String

		if _province_name.String == "" {
			_province_name.String = "{}"
		}
		objDamDaily.Geocode.Province_name = json.RawMessage(_province_name.String)

		if _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		objDamDaily.Geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)

		if _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		objDamDaily.Geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)

		data = append(data, objDamDaily)
	}
	return data, nil
}

//  get ข้อมูลเขื่อนล่าสุด รวม 6 ภาค ("ภาคเหนือ","ภาคตะวันออกเฉียงเหนือ","ภาคกลาง","ภาคตะวันตก","ภาคตะวันออก","ภาคใต้")
//  Parameters:
//    param
//      ใช้ในส่วน Agency_id, Basin_id, Dam_date
//  Return:
//    []Struct_DamDailySummary
func GetDamDailyLastestSummary() ([]*Struct_DamDailySummary, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02"
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_DamDailySummary
		objDamDaily *Struct_DamDailySummary

		_dam_date               time.Time
		_region                 sql.NullString
		_dam_storage            sql.NullFloat64
		_dam_uses_water         sql.NullFloat64
		_dam_storage_percent    sql.NullFloat64
		_dam_uses_water_percent sql.NullFloat64
		_row                    sql.NullString

		_geocode       sql.NullString
		_area_code     sql.NullString
		_area_name     sql.NullString
		_province_name sql.NullString
		_amphoe_code   sql.NullString
		_amphoe_name   sql.NullString
		_tumbon_code   sql.NullString
		_tumbon_name   sql.NullString
		_province_code sql.NullString
		_rid_area_code sql.NullString

		_result *sql.Rows
	)
	_result, err = db.Query(SQL_GetDamDailySummary)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_DamDailySummary, 0)
	for _result.Next() {

		//			&_dam_date, &_region, &_dam_storage, &_dam_uses_water, &_geocode, &_area_code, &_province_code, &_amphoe_code, &_tumbon_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_region_order, &_row, &_region_order, &_dam_storage_percent, &_dam_uses_water_percent
		err := _result.Scan(&_dam_date, &_region, &_dam_storage, &_dam_uses_water, &_dam_storage_percent, &_dam_uses_water_percent, &_geocode, &_area_code, &_province_code, &_amphoe_code, &_tumbon_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_rid_area_code, &_row)

		if err != nil {
			return nil, err
		}

		objDamDaily = &Struct_DamDailySummary{}
		objDamDaily.Dam_date = _dam_date.Format(strDatetimeFormat)
		objDamDaily.Region_name = _region.String
		objDamDaily.Dam_storage = float.TwoDigit(_dam_storage.Float64)
		objDamDaily.Dam_storage_percent = float.TwoDigit(_dam_storage_percent.Float64)
		objDamDaily.Dam_uses_water = float.TwoDigit(_dam_uses_water.Float64)
		objDamDaily.Dam_uses_water_percent = float.TwoDigit(_dam_uses_water_percent.Float64)

		objDamDaily.Geocode = &model_geocode.Struct_Geocode{}
		objDamDaily.Geocode.Geocode = _geocode.String
		objDamDaily.Geocode.Area_code = _area_code.String

		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		objDamDaily.Geocode.Area_name = json.RawMessage(_area_name.String)

		objDamDaily.Geocode.Province_code = _province_code.String

		if _province_name.String == "" {
			_province_name.String = "{}"
		}
		objDamDaily.Geocode.Province_name = json.RawMessage(_province_name.String)

		if _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		objDamDaily.Geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)

		if _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		objDamDaily.Geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)

		data = append(data, objDamDaily)
	}
	return data, nil
}

//  get ข้อมูลน้ำใช้การของ 4 เขื่อนหลัก รายวัน
//  Return:
//    []Struct_DamDailySummary4Dam
func GetDamDailySummary4Dam() ([]*Struct_DamDailySummary4Dam, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02"
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_DamDailySummary4Dam
		objDamDaily *Struct_DamDailySummary4Dam

		_dam_date       time.Time
		_dam_storage    sql.NullFloat64
		_dam_uses_water sql.NullFloat64
		_dam_inflow     sql.NullFloat64
		_dam_released   sql.NullFloat64
		_row            sql.NullString

		_result *sql.Rows
	)
	_result, err = db.Query(SQL_GetDamDailySummary4Dam)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_DamDailySummary4Dam, 0)
	for _result.Next() {

		err := _result.Scan(&_dam_date, &_dam_storage, &_dam_uses_water, &_dam_inflow, &_dam_released, &_row)

		if err != nil {
			return nil, err
		}

		objDamDaily = &Struct_DamDailySummary4Dam{}
		objDamDaily.Dam_date = _dam_date.Format(strDatetimeFormat)
		objDamDaily.Dam_storage = float.TwoDigit(_dam_storage.Float64)
		objDamDaily.Dam_inflow = float.TwoDigit(_dam_inflow.Float64)
		objDamDaily.Dam_uses_water = float.TwoDigit(_dam_uses_water.Float64)
		objDamDaily.Dam_released = float.TwoDigit(_dam_released.Float64)

		data = append(data, objDamDaily)
	}
	return data, nil
}
