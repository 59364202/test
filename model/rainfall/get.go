// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainfall is a model for public.rainfall table. This table store rainfall information.
package rainfall

import (
	"database/sql"
	"encoding/json"

	"haii.or.th/api/server/model/setting"
	model_setting "haii.or.th/api/server/model/setting"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	//	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
	"strconv"
	"strings"
	"time"
)

//func GetRainfallGraph(p *GetRainGraphParam) (*highchart.HighChart, error) {
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, err
//	}
//
//	var (
//		strSql string
//
//		_result *sql.Rows
//
//		_time time.Time
//		_v    sql.NullFloat64
//		_sum  float64
//
//		series2 *highchart.Series
//	)
//	strSql = SQL_GetRainfallGraph(p.Datatype)
//	log.Println(strSql, p.Id)
//	_result, err = db.Query(strSql, p.Id)
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//
//	hc := highchart.NewChart()
//	series := hc.NewColumnSerie("")
//	if p.Datatype == "1" {
//		series2 = hc.NewLineSerie("")
//		series2.YRight()
//		_sum = 0
//	}
//
//	for _result.Next() {
//		err = _result.Scan(&_time, &_v)
//		if err != nil {
//			return nil, err
//		}
//		if _v.Valid {
//			series.AddDateData(_time, _v.Float64)
//			_sum += _v.Float64
//		} else {
//			series.AddDateData(nil, nil)
//		}
//
//		if p.Datatype == "1" {
//			series2.AddDateData(_time, _sum)
//		}
//	}
//
//	return hc, nil
//}

//	get rainfall by station and date
//	Parameters:
//		param
//			Rainfall_InputParam
//	Return:
//		GetRainfallLastest_OutputParam
func GetRainfallByStationAndDate(param *Rainfall_InputParam) (*GetRainfallLastest_OutputParam, error) {

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
		data        []*RainfallStruct
		objRainfall *RainfallStruct

		_id sql.NullInt64
		//_tele_station_id	sql.NullInt64
		_rainfall_datetime time.Time
		_rainfall5m        sql.NullFloat64
		_rainfall10m       sql.NullFloat64
		_rainfall15m       sql.NullFloat64
		_rainfall30m       sql.NullFloat64
		_rainfall1h        sql.NullFloat64
		_rainfall3h        sql.NullFloat64
		_rainfall6h        sql.NullFloat64
		_rainfall12h       sql.NullFloat64
		_rainfall24h       sql.NullFloat64
		_rainfall_acc      sql.NullFloat64

		_result *sql.Rows
	)
	//Query
	//	log.Printf(sqlGetRainfallByStationAndDate+sqlGetRainfallByStationAndDateOrderBy, param.Station_id, param.Start_date, param.End_date)
	_result, err = db.Query(sqlGetRainfallByStationAndDate+sqlGetRainfallByStationAndDateOrderBy,
		param.Station_id, param.Start_date, param.End_date)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*RainfallStruct, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_rainfall_datetime,
			&_rainfall5m, &_rainfall10m, &_rainfall15m, &_rainfall30m,
			&_rainfall1h, &_rainfall3h, &_rainfall6h, &_rainfall12h, &_rainfall24h,
			&_rainfall_acc)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily object
		objRainfall = &RainfallStruct{}
		objRainfall.Id = _id.Int64
		objRainfall.Rainfall_date = _rainfall_datetime.Format(strDatetimeFormat)
		objRainfall.Tele_station_id, _ = strconv.ParseInt(param.Station_id, 10, 64)

		objRainfall.Rainfall5m = ValidData(_rainfall5m.Valid, _rainfall5m.Float64)
		objRainfall.Rainfall10m = ValidData(_rainfall10m.Valid, _rainfall10m.Float64)
		objRainfall.Rainfall15m = ValidData(_rainfall15m.Valid, _rainfall15m.Float64)
		objRainfall.Rainfall30m = ValidData(_rainfall30m.Valid, _rainfall30m.Float64)
		objRainfall.Rainfall1h = ValidData(_rainfall1h.Valid, _rainfall1h.Float64)
		objRainfall.Rainfall3h = ValidData(_rainfall3h.Valid, _rainfall3h.Float64)
		objRainfall.Rainfall6h = ValidData(_rainfall6h.Valid, _rainfall6h.Float64)
		objRainfall.Rainfall12h = ValidData(_rainfall12h.Valid, _rainfall12h.Float64)
		objRainfall.Rainfall24h = ValidData(_rainfall24h.Valid, _rainfall24h.Float64)
		objRainfall.Rainfall_acc = ValidData(_rainfall_acc.Valid, _rainfall_acc.Float64)

		data = append(data, objRainfall)
	}

	resultData := &GetRainfallLastest_OutputParam{}
	resultData.Data = data
	resultData.Header = arrRainfallHeaderByStationAndDate

	return resultData, nil
}

//	get rainfall error data
//	Parameters:
//		param
//			Rainfall_InputParam
//	Return:
//		Array Struct_Rainfall_ErrorData
func GetErrorRainfall(param *Rainfall_InputParam) ([]*Struct_Rainfall_ErrorData, error) {

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
		data        []*Struct_Rainfall_ErrorData
		objRainfall *Struct_Rainfall_ErrorData

		_id                    sql.NullInt64
		_station_id            sql.NullInt64
		_oldcode               sql.NullString
		_date                  time.Time
		_station_name          sql.NullString
		_station_province_name sql.NullString
		_agency_name           sql.NullString
		_agency_shortname      sql.NullString

		_rainfall5m   sql.NullFloat64
		_rainfall10m  sql.NullFloat64
		_rainfall15m  sql.NullFloat64
		_rainfall30m  sql.NullFloat64
		_rainfall1h   sql.NullFloat64
		_rainfall3h   sql.NullFloat64
		_rainfall6h   sql.NullFloat64
		_rainfall12h  sql.NullFloat64
		_rainfall24h  sql.NullFloat64
		_rainfall_acc sql.NullFloat64

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
		sqlCmdWhere += " AND rainfall_datetime >= $" + strconv.Itoa(len(arrParam))

		arrParam = append(arrParam, param.End_date+" 23:59")
		sqlCmdWhere += " AND rainfall_datetime <= $" + strconv.Itoa(len(arrParam))
	}

	//Query
	_result, err = db.Query(sqlGetErrorRainfall+sqlCmdWhere, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Loop data result
	data = make([]*Struct_Rainfall_ErrorData, 0)
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_oldcode, &_date, &_station_name, &_station_province_name, &_agency_name, &_agency_shortname,
			&_rainfall5m, &_rainfall10m, &_rainfall15m, &_rainfall30m,
			&_rainfall1h, &_rainfall3h, &_rainfall6h, &_rainfall12h, &_rainfall24h,
			&_rainfall_acc, &_station_id)
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
		objRainfall = &Struct_Rainfall_ErrorData{}
		objRainfall.ID = _id.Int64
		objRainfall.StationID = _station_id.Int64
		objRainfall.StationOldCode = _oldcode.String
		objRainfall.Datetime = _date.Format(strDatetimeFormat)
		objRainfall.StationName = json.RawMessage(_station_name.String)
		objRainfall.ProvinceName = json.RawMessage(_station_province_name.String)
		objRainfall.AgencyName = json.RawMessage(_agency_name.String)
		objRainfall.AgencyShortName = json.RawMessage(_agency_shortname.String)

		objRainfall.Rainfall5m = ValidData(_rainfall5m.Valid, _rainfall5m.Float64)
		objRainfall.Rainfall10m = ValidData(_rainfall10m.Valid, _rainfall10m.Float64)
		objRainfall.Rainfall15m = ValidData(_rainfall15m.Valid, _rainfall15m.Float64)
		objRainfall.Rainfall30m = ValidData(_rainfall30m.Valid, _rainfall30m.Float64)
		objRainfall.Rainfall1h = ValidData(_rainfall1h.Valid, _rainfall1h.Float64)
		objRainfall.Rainfall3h = ValidData(_rainfall3h.Valid, _rainfall3h.Float64)
		objRainfall.Rainfall6h = ValidData(_rainfall6h.Valid, _rainfall6h.Float64)
		objRainfall.Rainfall12h = ValidData(_rainfall12h.Valid, _rainfall12h.Float64)
		objRainfall.Rainfall24h = ValidData(_rainfall24h.Valid, _rainfall24h.Float64)
		objRainfall.RainfallAcc = ValidData(_rainfall_acc.Valid, _rainfall_acc.Float64)

		data = append(data, objRainfall)
	}

	return data, nil
}

//
//func GetRainfallCurrentTime() (string, error) {
//	return "", nil
//}

//	get rain monthly by station for graph
//	Parameters:
//		inputData
//			GetAdvRainMonthStationGraphInput
//	Return:
//		GraphData
func GetAdvRainMonthlyStationGraph(inputData *GetAdvRainMonthStationGraphInput) (*GraphData, error) {
	if inputData.StationID == 0 {
		return nil, rest.NewError(422, "No Input Station", nil)
	}

	if len(inputData.Year) == 0 {
		return nil, rest.NewError(422, "No Input Year", nil)
	}
	data := make([]*GetAdvRainMonthStationGraphOutput, 0)
	// loop get data by year
	for _, v := range inputData.Year {

		dY := &GetAdvRainMonthStationGraphOutput{}
		dM, err := getAdvRainMonthlyStationGraphByYear(inputData.StationID, strconv.FormatInt(v, 10), "")
		if err != nil {
			return nil, err
		}
		dY.Year = v
		dY.GSeries = dM
		// add data by year
		data = append(data, dY)
	}
	dd := &GraphData{}
	// return data
	dd.GData = data
	return dd, nil
}

//	get rain monthly station for graph by year
//	Parameters:
//		stationID
//			รหัสสถานี
//		year
//			ปี
//		code
//			"area" ถ้าต้องการหาโดยใช้ area
//	Return:
//		Array GetAdvRainMonthStationMonthOutput
func getAdvRainMonthlyStationGraphByYear(stationID interface{}, year, code string) ([]*GetAdvRainMonthStationMonthOutput, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	var q string
	p := []interface{}{}
	// select sql by area code
	if code == "area" {
		q = sqlAdvRainMonthlyGraphByArea
		if stationID.(string) != "" {
			q += " AND lg.tmd_area_code=$3"
			p = []interface{}{year + "-01-01", year + "-12-31", stationID}
		} else {
			p = []interface{}{year + "-01-01", year + "-12-31"}
		}
		q += " GROUP BY rainfall_datetime ORDER BY rainfall_datetime ASC"
	} else {
		q = sqlAdvRainMonthlyStationGraph
		p = []interface{}{stationID, year + "-01-01", year + "-12-31"}
	}

	// query
	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dm := make([]*GetAdvRainMonthStationMonthOutput, 12)
	// loop add data
	for i := 0; i < 12; i++ {
		m := &GetAdvRainMonthStationMonthOutput{}
		m.Month = int64(i + 1)
		dm[i] = m
	}

	for rows.Next() {
		var (
			datetime time.Time
			value    sql.NullFloat64
		)
		// scan data from database
		rows.Scan(&datetime, &value)
		m := int(datetime.Month())
		dataRow := dm[m-1]
		dataRow.Value = ValidData(value.Valid, value.Float64)
		// add data to array
		dm[m-1] = dataRow
	}

	return dm, nil
}

//	get rain monthly for graph
//	Parameters:
//		inputData
//			GetAdvRainMonthGraphInput
//	Return:
//		GraphData
func GetAdvRainMonthlyGraph(inputData *GetAdvRainMonthGraphInput) (*GraphData, error) {
	// validate input
	if inputData.StationID == 0 {
		return nil, rest.NewError(422, "No Input Station", nil)
	}

	if len(inputData.Month) == 0 {
		return nil, rest.NewError(422, "No Input Month", nil)
	}

	if inputData.StartYear == 0 || inputData.EndYear == 0 {
		return nil, rest.NewError(422, "No Input Year", nil)
	}

	y := inputData.StartYear
	data := make([]*GetAdvRainMonthStationGraphOutput, 0)
	// loop get data by year
	for ; y <= inputData.EndYear; y++ {

		dY := &GetAdvRainMonthStationGraphOutput{}
		// get data by year
		dM, err := getAdvRainMonthlyGraphByYear(inputData.StationID, strconv.FormatInt(y, 10), inputData.Month)
		if err != nil {
			return nil, err
		}
		dY.Year = y
		dY.GSeries = dM
		data = append(data, dY)
	}
	dd := &GraphData{}
	dd.GData = data
	return dd, nil
}

//	get rainmothly by station id and year
//	Parameters:
//		stationID
//			รหัสสถานี
//		year
//			ปี
//		month
//			เดือน
//	Return:
//		Array GetAdvRainMonthStationMonthOutput
func getAdvRainMonthlyGraphByYear(stationID int64, year string, month []int64) ([]*GetAdvRainMonthStationMonthOutput, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get rain monthly
	q := sqlAdvRainMonthlyStationGraph
	p := []interface{}{stationID, year + "-01-01", year + "-12-31"}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dm := make([]*GetAdvRainMonthStationMonthOutput, 12)

	// loop define data 12 month
	for i := 0; i < 12; i++ {
		m := &GetAdvRainMonthStationMonthOutput{}
		m.Month = int64(i + 1)
		dm[i] = m
	}

	for rows.Next() {
		var (
			datetime time.Time
			value    sql.NullFloat64
		)
		// scan data
		rows.Scan(&datetime, &value)
		// add data by month to array
		m := int(datetime.Month())
		dataRow := dm[m-1]
		dataRow.Value = ValidData(value.Valid, value.Float64)
		dm[m-1] = dataRow
	}
	data := make([]*GetAdvRainMonthStationMonthOutput, 0)
	// loop get month add to array of year
	for _, v := range month {
		m := &GetAdvRainMonthStationMonthOutput{}
		m.Month = v
		m.Value = dm[int(v-1)].Value
		data = append(data, m)
	}
	return data, nil
}

//	get rain yearly for graph
//	Parameters:
//		inputData
//			GetAdvRainYearlyGraphInput
//	Return:
//		GraphDataYearly
func GetAdvRainYearlyGraph(inputData *GetAdvRainYearlyGraphInput) (*GraphDataYearly, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// get sql rain yearly
	q := sqlAdvRainYearlyGraph
	p := []interface{}{inputData.StationID, strconv.FormatInt(inputData.StartYear, 10) + "-01-01", strconv.FormatInt(inputData.EndYear, 10) + "-12-31"}

	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	// define data structure
	dm := make(map[int64]*GetAdvRainYearlyOutput)
	y := inputData.StartYear
	for ; y <= inputData.EndYear; y++ {
		m := &GetAdvRainYearlyOutput{}
		m.Year = y
		dm[y] = m
	}

	for rows.Next() {
		var (
			datetime time.Time
			value    sql.NullFloat64
		)
		// scan data
		rows.Scan(&datetime, &value)

		m := int64(datetime.Year())

		// add data to array month
		dataRow := dm[m]
		dataRow.Value = ValidData(value.Valid, value.Float64)
		dm[m] = dataRow
	}
	data := make([]*GetAdvRainYearlyOutput, 0)
	y = inputData.StartYear
	// loop add data by tear
	for ; y <= inputData.EndYear; y++ {
		data = append(data, dm[y])
	}
	dd := &GraphDataYearly{}

	dd.GData = data

	// sql raini normal
	q = sqlAdvRainNormal
	p = []interface{}{inputData.StationID}

	// query
	rows, err = db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var (
		y30 sql.NullFloat64
		y48 sql.NullFloat64
	)

	for rows.Next() {
		// scan y30,y48
		rows.Scan(&y30, &y48)
	}

	// check valid data
	dd.Normal30 = ValidData(y30.Valid, y30.Float64)
	dd.Normal48 = ValidData(y48.Valid, y48.Float64)

	return dd, nil
}

//	get rain monthly by area for graph
//	Parameters:
//		inputData
//			GetAdvRainMonthAreaGraphInput
//	Return:
//		GetAdvRainAreaOutput
func GetAdvRainMonthlyAreaGraph(inputData *GetAdvRainMonthAreaGraphInput) (*GetAdvRainAreaOutput, error) {
	if len(inputData.Year) == 0 {
		return nil, rest.NewError(422, "No Input Year", nil)
	}

	if inputData.BaseLine == "" {
		return nil, rest.NewError(422, "No Input Baseline", nil)
	}
	output := &GetAdvRainAreaOutput{}
	data := make([]*GetAdvRainMonthStationGraphOutput, 0)
	dataBaseLine := make([]*GetAdvRainMonthStationGraphOutput, 0)
	// loop get data by tear
	for _, v := range inputData.Year {
		dY := &GetAdvRainMonthStationGraphOutput{}
		dM, err := getAdvRainMonthlyStationGraphByYear(inputData.AreaCode, strconv.FormatInt(v, 10), "area")
		if err != nil {
			return nil, err
		}
		dY.Year = v
		dY.GSeries = dM
		data = append(data, dY)
	}

	// get data base line
	bY := &GetAdvRainMonthStationGraphOutput{}
	bM, err := getAdvBaseline(inputData.AreaCode, inputData.BaseLine)
	if err != nil {
		return nil, err
	}

	bY.Year, _ = strconv.ParseInt(inputData.BaseLine, 10, 64)
	bY.GSeries = bM
	// add base line
	dataBaseLine = append(dataBaseLine, bY)

	output.GraphData = data
	output.Baseline = dataBaseLine

	return output, nil
}

//	get base line by station
//	Parameters:
//		stationID
//			รหัสสถานี
//		baseline
//			30,48
//	Return:
//		Array GetAdvRainMonthStationMonthOutput
func getAdvBaseline(stationID string, baseline string) ([]*GetAdvRainMonthStationMonthOutput, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	var q string
	if baseline == "48" {
		q = sqlAdvRainBaselineAreaY48
	} else {
		q = sqlAdvRainBaselineAreaY30
	}

	p := []interface{}{}

	// get base line by region id
	if stationID != "" {
		q += " WHERE reg_id=$1 GROUP BY month_id ORDER BY month_id"
		p = append(p, stationID)
	} else {
		q += " WHERE reg_id='0' GROUP BY month_id ORDER BY month_id"
	}

	// query
	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dm := make([]*GetAdvRainMonthStationMonthOutput, 12)

	// loop define data by month
	for i := 0; i < 12; i++ {
		m := &GetAdvRainMonthStationMonthOutput{}
		m.Month = int64(i + 1)
		dm[i] = m
	}

	for rows.Next() {
		var (
			month sql.NullString
			value sql.NullFloat64
		)
		// scan data
		rows.Scan(&month, &value)
		m, _ := strconv.Atoi(strings.Trim(month.String, " "))
		dataRow := dm[m-1]
		dataRow.Value = ValidData(value.Valid, value.Float64)
		// add data by month
		dm[m-1] = dataRow
	}
	return dm, nil
}

//	check validdata
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

//Api Mobile Rainfall24h_latest
type Struct_rainfall24h_latest struct {
	Flag     bool   `json:"rain24_flag"`
	Province string `json:"rain24_text"`
}

func Get_Rainfall24h_latest() ([]*Struct_rainfall24h_latest, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//	change lr.rainfall24h <= 500
	//	<= 600 ตามที่ผู้ใช้แจ้งกรณีพายุเข้าภาคใต้
	var q string = ` 
	select province_name, max(t2.rainfall24h) as rainfall24h
    from (
	select  lr.tele_station_id, lr.rainfall_datetime, lr.rainfall24h,pn.province_name ->> 'th' as province_name
	from cache.latest_rainfall24h lr
	left join public.m_tele_station st
	on lr.tele_station_id = st.id
	left join public.lt_geocode pn
	on st.geocode_id = pn.id
	LEFT  JOIN ignore ig ON st.id = ig.station_id ::int
	where lr.rainfall_datetime  >= $1
	and lr.rainfall_datetime <= $2
	and lr.rainfall24h <> '999999'
    and lr.rainfall24h IS NOT NULL
	and lr.rainfall24h is not null
	and st.geocode_id is not null
    and lr.qc_status ->> 'is_pass' = 'true'
    -- AND ig.data_category = 'rainfall_24h'
	AND (ig.is_ignore IS FALSE OR ig.is_ignore IS NULL)
	order by lr.rainfall24h desc
    ) t2
    group by province_name
    order by max(t2.rainfall24h) desc
    limit 3`

	var rs []*Struct_rainfall24h_latest = make([]*Struct_rainfall24h_latest, 0)

	ds := time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04")
	de := time.Now().Format("2006-01-02 15:04")

	rows, err := db.Query(q, ds, de)
	if rows != nil {
		var province_list string

		for i := 0; rows.Next(); i++ {
			var (
				_Province sql.NullString
				_Rainfall sql.NullString
			)

			err = rows.Scan(&_Province, &_Rainfall)
			if err != nil {
				return nil, err
			}

			if i > 0 {
				province_list = province_list + " " + _Province.String
			} else {
				province_list = province_list + _Province.String
			}
		}

		s := &Struct_rainfall24h_latest{
			Flag:     true,
			Province: province_list,
		}
		rs = append(rs, s)
	} else {
		s := &Struct_rainfall24h_latest{
			Flag:     false,
			Province: "ไม่มีฝน",
		}
		rs = append(rs, s)
	}

	return rs, err
}

type Struct_rainfall_latest_list struct {
	Rainfall_datas []*Struct_rainfall_datas `json:"rainfall_datas"`
	Map            []*Struct_media_link     `json:"map"`
}

type Struct_rainfall_datas struct {
	Station_name    string `json:"station_name"`    //example:"สถานีบ้านโคกแมว จ.พัทลุง"
	Current_time    string `json:"current_time"`    //example:"2017-11-07 13:30ง"
	Station_id      string `json:"station_id"`      //example:"teledwr1022"
	Station_lat     string `json:"station_lat"`     //example:"7.393765"
	Station_long    string `json:"station_long"`    //example:"100.083260"
	Rainfall        string `json:"rainfall"`        //example:"71.00"
	Criterion_color string `json:"criterion_color"` //example:"#ca6504"
	Province_id     string `json:"province_id"`     //example:"93"
	Province_name   string `json:"province_name"`   //example:"พัทลุง"
	Amphoe_name     string `json:"amphoe_name"`     //example:"เขาชัยสน"
	District_name   string `json:"district_name"`   //example:"เขาชัยสน"
	Criterion_text  string `json:"criterion_text"`  //example:"ไม่มีฝน"
	Geocode         string `json:"geocode"`         //example:"830203"
	Agency_id       string `json:"agency_id"`       //example:"10"
}

type Struct_media_link struct {
	Image_url   string `json:"image_url"`   //example:"http://www.nhc.in.th/product/latest/img/rain24.jpg"
	Image_time  string `json:"image_time"`  //example:"15:00 น."
	Detail_link string `json:"detail_link"` //example:"http://www.nhc.in.th/web/index.php?model=telemetering&view=weather"
}

func Get_Rainfall_latest_list() (*Struct_rainfall_latest_list, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var q string = `
	SELECT Concat('สถานี', st.tele_station_name ->> 
				'th', ' จ.', pn.province_name ->> 'th')      AS station_name 
		, To_char(lr.rainfall_datetime, 'YYYY-MM-DD HH24:MI') AS 
			rainfall_datetime 
		, lr.tele_station_id 
		, st.tele_station_lat 
		, st.tele_station_long 
		, lr.rainfall24h 
		, pn.province_code 
		, pn.province_name ->> 'th'                           AS province_name 
		, pn.amphoe_name ->> 'th'                             AS amphoe_name 
		, pn.tumbon_name ->> 'th'                             AS tumbon_name 
	FROM   cache.latest_rainfall24h lr 
		LEFT JOIN public.m_tele_station st 
				ON lr.tele_station_id = st.id 
		LEFT JOIN public.lt_geocode pn 
				ON st.geocode_id = pn.id 
		LEFT  JOIN public.ignore ig 
				ON st.id = ig.station_id ::int
	WHERE  lr.rainfall24h <= 600 
		AND lr.rainfall24h IS NOT NULL 
		AND st.geocode_id IS NOT NULL 
		AND lr.rainfall_datetime :: DATE = CURRENT_DATE 
		AND ( lr.qc_status IS NULL OR lr.qc_status->>'is_pass' = 'true' )
		-- AND ig.data_category = 'rainfall_24h'
		AND (ig.is_ignore IS FALSE OR ig.is_ignore IS NULL)
	ORDER  BY lr.rainfall24h DESC `

	// แปลง Frontend.public.rain_setting จาก setting ให้เป็น uSetting.Struct_RainSetting
	rain_scale_color := &uSetting.Struct_RainSetting{}
	err = setting.GetSystemSettingPtr("Frontend.public.rain_setting", &rain_scale_color)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	var rs *Struct_rainfall_latest_list = &Struct_rainfall_latest_list{}
	var rf []*Struct_rainfall_datas = make([]*Struct_rainfall_datas, 0)
	var lk []*Struct_media_link = make([]*Struct_media_link, 0)

	for rows.Next() {
		var (
			_Station_name      sql.NullString
			_Rainfall_datetime sql.NullString
			_Tele_station_id   sql.NullString
			_Tele_station_lat  sql.NullString
			_Tele_station_long sql.NullString
			_Rainfall24h       sql.NullFloat64
			_Province_code     sql.NullString
			_Province_name     sql.NullString
			_Amphoe_name       sql.NullString
			_Tumbon_name       sql.NullString
		)

		err := rows.Scan(&_Station_name, &_Rainfall_datetime, &_Tele_station_id, &_Tele_station_lat, &_Tele_station_long, &_Rainfall24h, &_Province_code, &_Province_name, &_Amphoe_name, &_Tumbon_name)
		if err != nil {
			return nil, err
		}

		//rf_value := float64(float64 (_Rainfall24h.Float64*100))/100 //Convert rainfall is int64 to compare for get color code.

		s := &Struct_rainfall_datas{
			Station_name:  _Station_name.String,
			Current_time:  _Rainfall_datetime.String,
			Station_id:    _Tele_station_id.String,
			Station_lat:   _Tele_station_lat.String,
			Station_long:  _Tele_station_long.String,
			Rainfall:      strconv.FormatFloat(_Rainfall24h.Float64, 'f', 2, 64),
			Province_id:   _Province_code.String,
			Province_name: _Province_name.String,
			Amphoe_name:   _Amphoe_name.String,
			District_name: _Tumbon_name.String,
		}

		// หา config จาก rain_storage
		rain_storage := rain_scale_color.CompareRain(_Rainfall24h.Float64)
		if rain_storage != nil && _Rainfall24h.Float64 > 0 {
			s.Criterion_color = strings.ToLower(rain_storage.Color)
			s.Criterion_text = rain_storage.Criterion_text
		} else {
			s.Criterion_color = "#e0e0e0"
			s.Criterion_text = "ไม่มีฝน"
		}

		rf = append(rf, s)
	}

	current_date := time.Now()
	cur_time := current_date.Format("15.04")

	media := &Struct_media_link{
		Image_time:  cur_time + " น.",
		Image_url:   "http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/image?image=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJAMg9OoFxEIwU2CpA5uqnpJs=?time=" + cur_time,
		Detail_link: "http://www.nhc.in.th/web/index.php?model=telemetering&view=weather",
	}

	lk = append(lk, media)

	rs = &Struct_rainfall_latest_list{
		Rainfall_datas: rf,
		Map:            lk,
	}

	return rs, err
}

func Get_rainfall_province(prov_id string) (*Struct_rainfall_latest_list, error) {
	if prov_id == "" {
		rs, err := Get_Rainfall_latest_list()
		return rs, err
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var q string = `
	SELECT Concat('สถานี', st.tele_station_name ->> 
				'th', ' จ.', pn.province_name ->> 'th')      AS station_name 
		, To_char(lr.rainfall_datetime, 'YYYY-MM-DD HH24:MI') AS 
			rainfall_datetime 
		, lr.tele_station_id 
		, st.tele_station_lat 
		, st.tele_station_long 
		, lr.rainfall24h 
		, pn.province_code 
		, pn.province_name ->> 'th'	AS province_name 
		, pn.amphoe_name ->> 'th'	AS amphoe_name 
		, pn.tumbon_name ->> 'th'	AS tumbon_name 
		, pn.id	AS geocode_id 
		, st.agency_id 
	FROM   cache.latest_rainfall24h lr 
		LEFT JOIN public.m_tele_station st 
			ON lr.tele_station_id = st.id 
		LEFT JOIN public.lt_geocode pn 
			ON st.geocode_id = pn.id 
		LEFT  JOIN public.ignore ig 
			ON st.id = ig.station_id ::int
	WHERE  lr.rainfall24h <= 600 
		AND pn.province_code = $1 
		AND lr.rainfall24h IS NOT NULL 
		AND st.geocode_id IS NOT NULL 
		AND rainfall_datetime :: DATE = CURRENT_DATE 
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
		AND (ig.is_ignore IS FALSE OR ig.is_ignore IS NULL)
	ORDER  BY lr.rainfall24h DESC `

	// แปลง Frontend.public.rain_setting จาก setting ให้เป็น uSetting.Struct_RainSetting
	rain_scale_color := &uSetting.Struct_RainSetting{}
	err = setting.GetSystemSettingPtr("Frontend.public.rain_setting", &rain_scale_color)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(q, prov_id)
	if err != nil {
		return nil, err
	}

	var rs *Struct_rainfall_latest_list = &Struct_rainfall_latest_list{}
	var rf []*Struct_rainfall_datas = make([]*Struct_rainfall_datas, 0)
	var lk []*Struct_media_link = make([]*Struct_media_link, 0)

	for rows.Next() {
		var (
			_Station_name      sql.NullString
			_Rainfall_datetime sql.NullString
			_Tele_station_id   sql.NullString
			_Tele_station_lat  sql.NullString
			_Tele_station_long sql.NullString
			_Rainfall24h       sql.NullFloat64
			_Province_code     sql.NullString
			_Province_name     sql.NullString
			_Amphoe_name       sql.NullString
			_Tumbon_name       sql.NullString
			_Geocode_id        sql.NullString
			_Agency_id         sql.NullString
		)

		err := rows.Scan(&_Station_name, &_Rainfall_datetime, &_Tele_station_id, &_Tele_station_lat, &_Tele_station_long, &_Rainfall24h, &_Province_code, &_Province_name, &_Amphoe_name, &_Tumbon_name, &_Geocode_id, &_Agency_id)
		if err != nil {
			return nil, err
		}

		strconv.FormatFloat(_Rainfall24h.Float64, 'f', 2, 64)

		s := &Struct_rainfall_datas{
			Station_name:  _Station_name.String,
			Current_time:  _Rainfall_datetime.String,
			Station_id:    _Tele_station_id.String,
			Station_lat:   _Tele_station_lat.String,
			Station_long:  _Tele_station_long.String,
			Rainfall:      strconv.FormatFloat(_Rainfall24h.Float64, 'f', 2, 64),
			Province_id:   _Province_code.String,
			Province_name: _Province_name.String,
			Amphoe_name:   _Amphoe_name.String,
			District_name: _Tumbon_name.String,
			Geocode:       _Geocode_id.String,
			Agency_id:     _Agency_id.String,
		}

		//		rf_value, err := strconv.Atoi(_Rainfall24h.Float64)

		// หา config จาก rain_storage
		rain_storage := rain_scale_color.CompareRain(_Rainfall24h.Float64)
		if rain_storage != nil && _Rainfall24h.Float64 > 0 {
			s.Criterion_color = rain_storage.Color
			s.Criterion_text = strings.ToLower(rain_storage.Criterion_text)
		} else {
			s.Criterion_color = "#e0e0e0"
			s.Criterion_text = "ไม่มีฝน"
		}

		rf = append(rf, s)

	}

	current_date := time.Now()
	cur_time := current_date.Format("15.04")

	media := &Struct_media_link{
		Image_url:   "http://www.nhc.in.th/product/latest/img/rain24.jpg",
		Image_time:  cur_time + " น.",
		Detail_link: "http://www.nhc.in.th/web/index.php?model=telemetering&view=weather",
	}

	lk = append(lk, media)

	rs = &Struct_rainfall_latest_list{
		Rainfall_datas: rf,
		Map:            lk,
	}

	return rs, err
}

type Struct_rainfall_station_graph struct {
	Rainfall1h     []*Struct_rainfall_1h    `json:"rainfall1h"`
	Rainfall_daily []*Struct_rainfall_daily `json:"rainfall_daily"`
}

type Struct_rainfall_1h struct {
	Tele_station_id string `json:"tele_station_id"` //example:"telebma0050"
	Rainfall_date   string `json:"rainfall_date"`   //example:"2017-11-08"
	Rainfall_time   string `json:"rainfall_time"`   //example:"11:00:00"
	Rainfall        string `json:"rainfall"`        //example:"0"
	Unit            string `json:"unit"`            //example:"มม."
}

type Struct_rainfall_daily struct {
	Tele_station_id string `json:"tele_station_id"` //example:"telebma0050"
	Rainfall_date   string `json:"rainfall_date"`   //example:"2017-11-08"
	Rainfall_daily  string `json:"rainfall_daily"`  //example:"0"
	Unit            string `json:"unit"`            //example:"มม."
}

func Get_rainfall_lateest_station_graph(station_id string) (*Struct_rainfall_station_graph, error) {
	if station_id == "" {
		return nil, errors.New("no station_id")
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	currentTimestamp := now.Format("2006-01-02 15:04")
	yesterday := now.AddDate(0, 0, -1).Format("2006-01-02 15:00")
	lastWeek := now.AddDate(0, 0, -7).Format("2006-01-02 15:00")
	var q_rainfall1h string = `
	SELECT tele_station_id, 
	to_char(rainfall_datetime::date, 'YYYY-MM-DD') AS rainfall_date, 
	to_char(rainfall_datetime::time, 'HH24:MI:SS') AS rainfall_time, 
	rainfall1h AS rainfall 
	FROM public.rainfall_1h
	WHERE tele_station_id = $1
	AND rainfall_datetime BETWEEN '` + yesterday + `' AND '` + currentTimestamp + `'` +
		// AND rainfall_datetime BETWEEN DATE_TRUNC('HOUR','` + currentTimestamp + `'::TIMESTAMP - INTERVAL '24 HOUR') AND '` + currentTimestamp + `'
		`AND rainfall_datetime = DATE_TRUNC('hour', rainfall_datetime)
	AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
	ORDER BY rainfall_date, rainfall_time`

	var q_rainfall_daily string = `
	select t1.tele_station_id, to_char(t1.rainfall_datetime,'YYYY-MM-DD') as rainfall_datetime,  t1.rainfall_value
	from rainfall_daily t1
	inner join (
		select
		  max(rainfall_datetime) as rainfall_datetime
		from rainfall_daily
		where rainfall_datetime BETWEEN '` + lastWeek + `' AND '` + currentTimestamp + `'` +
		// where rainfall_datetime BETWEEN DATE_TRUNC('HOUR','` + currentTimestamp + `'::TIMESTAMP - INTERVAL '7 DAYS') AND '` + currentTimestamp + `'
		`and tele_station_id = $1
		and deleted_at = to_timestamp(0)
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
		group by rainfall_datetime::date
		order by rainfall_datetime desc
	) t2
	on t2.rainfall_datetime = t1.rainfall_datetime
	where t1.rainfall_datetime BETWEEN '` + lastWeek + `' AND '` + currentTimestamp + `'` +
		// where t1.rainfall_datetime BETWEEN t1.rainfall_datetime - interval '7 days' AND '` + currentTimestamp + `'` +
		`and t1.tele_station_id = $1
	and t1.deleted_at = to_timestamp(0)
	order by t1.rainfall_datetime`

	rows_rainfall1h, err := db.Query(q_rainfall1h, station_id)
	if err != nil {
		return nil, err
	}

	rows_rainfall_daily, err := db.Query(q_rainfall_daily, station_id)
	if err != nil {
		return nil, err
	}

	var rs *Struct_rainfall_station_graph = &Struct_rainfall_station_graph{}
	var rs_1h []*Struct_rainfall_1h = make([]*Struct_rainfall_1h, 0)
	var rs_daily []*Struct_rainfall_daily = make([]*Struct_rainfall_daily, 0)

	for rows_rainfall1h.Next() {
		var (
			_tele_station_id sql.NullString
			_rainfall_date   sql.NullString
			_rainfall_time   sql.NullString
			_rainfall        sql.NullString
		)

		err = rows_rainfall1h.Scan(&_tele_station_id, &_rainfall_date, &_rainfall_time, &_rainfall)
		if err != nil {
			return nil, err
		}

		s_rf_1h := &Struct_rainfall_1h{
			Tele_station_id: _tele_station_id.String,
			Rainfall_date:   _rainfall_date.String,
			Rainfall_time:   _rainfall_time.String,
			Rainfall:        _rainfall.String,
			Unit:            "มม.",
		}

		rs_1h = append(rs_1h, s_rf_1h)
	}

	for rows_rainfall_daily.Next() {
		var (
			_tele_station_id sql.NullString
			_rainfall_date   sql.NullString
			_rainfall_daily  sql.NullString
		)

		err = rows_rainfall_daily.Scan(&_tele_station_id, &_rainfall_date, &_rainfall_daily)
		if err != nil {
			return nil, err
		}

		s_rf_daily := &Struct_rainfall_daily{
			Tele_station_id: _tele_station_id.String,
			Rainfall_date:   _rainfall_date.String,
			Rainfall_daily:  _rainfall_daily.String,
			Unit:            "มม.",
		}

		rs_daily = append(rs_daily, s_rf_daily)
	}

	rs = &Struct_rainfall_station_graph{
		Rainfall1h:     rs_1h,
		Rainfall_daily: rs_daily,
	}

	return rs, err
}
