package migrate

import (
	"database/sql"
	"haii.or.th/api/thaiwater30/util/b64"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"

	//	"log"
	"time"

	"haii.or.th/api/server/model/backgroundjob"
	bgjob "haii.or.th/api/thaiwater30/model/backgroundjob"
)

func GetSummaryMasterData() (interface{}, string, bool, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, "", false, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := "SELECT nhc_table, nhc_row, thaiwater30_table, thaiwater30_row,last_update FROM migrate_log.summary_master_data order by nhc_table"
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, "", false, pqx.GetRESTError(err)
	}
	data := make([]*Summary, 0)
	var last_update time.Time
	for rows.Next() {

		var (
			nhc_table         string
			nhc_row           int64
			thaiwater30_table string
			thaiwater30_row   int64
		)
		rows.Scan(&nhc_table, &nhc_row, &thaiwater30_table, &thaiwater30_row, &last_update)
		dataRow := &Summary{}
		dataRow.NHCRow = nhc_row
		dataRow.NHCTable = nhc_table
		dataRow.TW30Row = thaiwater30_row
		dataRow.TW30Table = thaiwater30_table
		data = append(data, dataRow)
	}
	isBgjobRunning, err := bgjob.IsBgJobRunning(Cmd + " " + Cmd_RegenMasterData)
	if err != nil {
		return nil, "", false, pqx.GetRESTError(err)
	}

	return data, last_update.Format("2006-01-02 15:04"), isBgjobRunning, nil
}

func GetSummaryData() (interface{}, string, bool, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, "", false, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := "SELECT nhc_table, nhc_row, thaiwater30_table, thaiwater30_row,last_update FROM migrate_log.summary_data_row order by nhc_table"
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, "", false, pqx.GetRESTError(err)
	}
	data := make([]*Summary, 0)
	var last_update time.Time
	for rows.Next() {

		var (
			nhc_table         string
			nhc_row           int64
			thaiwater30_table string
			thaiwater30_row   int64
		)
		rows.Scan(&nhc_table, &nhc_row, &thaiwater30_table, &thaiwater30_row, &last_update)
		dataRow := &Summary{}
		dataRow.NHCRow = nhc_row
		dataRow.NHCTable = nhc_table
		dataRow.TW30Row = thaiwater30_row
		dataRow.TW30Table = thaiwater30_table
		data = append(data, dataRow)
	}

	isBgjobRunning, err := bgjob.IsBgJobRunning(Cmd + " " + Cmd_RegenData)
	if err != nil {
		return nil, "", false, pqx.GetRESTError(err)
	}

	return data, last_update.Format("2006-01-02 15:04"), isBgjobRunning, nil
}

type Summary struct {
	NHCTable  string `json:"nhc_table"`
	NHCRow    int64  `json:"nhc_count"`
	TW30Table string `json:"tw30_table"`
	TW30Row   int64  `json:"tw30_count"`
}

type DataOutput struct {
	NHC  interface{} `json:"nhc"`
	TW30 interface{} `json:"tw30"`
}

func GetDataByTable(table, strtDate, endDate string) (*DataOutput, error) {
	q1, h1, q2, h2, flag := getInfoData(table)
	data := &DataOutput{}
	p := []interface{}{}
	if flag {
		p = append(p, strtDate)
		p = append(p, endDate)
	}
	data.TW30, _ = getDataTW30Table(q2, h2, p)
	data.NHC, _ = getDataNHCTable(q1, h1, table, p)
	return data, nil
}

func getInfoData(table string) (string, []string, string, []string, bool) {

	switch table {
	case "wl_canal":
		return NhcCanal, []string{"canal_station_oldid", "wl_canal_date", "wl_canal", "comm_status"}, TW30Canal, []string{"canal_station_oldcode", "cw.id", "cw.canal_station_id", "canal_waterlevel_datetime", "canal_waterlevel_value", "comm_status"}, true
	case "dam_daily":
		return NHCDamDaily, []string{"dam_tname", "dam_date", "dam_level", "dam_storage", "dam_inflow", "dam_released", "dam_spilled", "dam_losses", "dam_evalp", "dam_uses_water", "dam_storage_percent", "dam_uses_water_percent", "dam_inflow_acc_percent"}, TW30DamDaily, []string{"dam_name", "dam_date", "dam_level", "dam_storage", "dam_inflow", "dam_released", "dam_spilled", "dam_losses", "dam_evap", "dam_uses_water", "dam_storage_percent", "dam_uses_water_percent", "dam_inflow_avg", "dam_released_acc", "dam_inflow_acc", "dam_inflow_acc_percent"}, true
	case "dam_hourly":
		return NHCDamHourly, []string{"dam_tname", "dam_datetime", "dam_level", "dam_storage", "dam_inflow", "dam_released", "dam_spilled", "dam_losses", "dam_evalp"}, TW30DamHourly, []string{"dam_name", "dam_datetime", "dam_level", "dam_storage", "dam_inflow", "dam_released", "dam_spilled", "dam_losses", "dam_evap"}, true
	case "humid":
		return NHCHumid, []string{"tele_station_oldcode", "humid_datetime", "humid_value"}, TW30Humid, []string{"tele_station_oldcode", "humid_datetime", "humid_value"}, true
	case "medium_dam":
		return NHCMediumDam, []string{"dam_tname", "dam_date", "dam_level", "dam_storage", "dam_inflow", "dam_released", "dam_spilled", "dam_losses", "dam_evalp", "dam_uses_water", "dam_storage_percent", "dam_uses_water_percent", "dam_inflow_acc_percent"}, TW30MediumDam, []string{"mediumdam_name", "mediumdam_date", "mediumdam_storage", "mediumdam_inflow", "mediumdam_released", "mediumdam_uses_water", "mediumdam_storage_percent"}, true
	case "pressure":
		return NHCPressure, []string{"tele_station_oldcode", "pressure_datetime", "pressure_value"}, TW30Pressure, []string{"tele_station_oldcode", "pressure_datetime", "pressure_value"}, true
	case "rainfall":
		return NHCRainfall, []string{"tele_station_oldid", "rainfall_date", "rainfall_date_calc", "rainfall5m", "rainfall10m", "rainfall15m", "rainfall30m", "rainfall1h", "rainfall3h", "rainfall6h", "rainfall12h", "rainfall24h", "rainfall_acc"}, TW30Rainfall, []string{"tele_station_oldcode", " rainfall_datetime", " rainfall5m", " rainfall_date_calc", " rainfall10m", " rainfall15m", " rainfall30m", " rainfall1h", " rainfall3h", " rainfall6h", " rainfall12h", " rainfall24h", " rainfall_acc", " rainfall_today"}, true
	case "rainfall1h":
		return NHCRainfall1H, []string{"tele_station_oldcode", "rainfall_datetime", "rainfall_datetime_calc", "rainfall1h"}, TW30Rainfall1H, []string{"tele_station_oldcode", "rainfall_datetime", "rainfall_datetime_calc", "rainfall1h"}, true
	case "rainfall24h":
		return NHCRainfall24H, []string{"tele_station_oldcode", "rainfall_datetime", "rainfall_datetime_calc", "rainfall24h"}, TW30Rainfall24H, []string{"tele_station_oldcode", "rainfall_datetime", "rainfall_datetime_calc", "rainfall24h"}, true
	case "rainfall_daily":
		return NHCRainfallDaily, []string{"tele_station_oldcode", "rainfall_datetime", "rainfall_value"}, TW30RainfallDaily, []string{"tele_station_oldcode", "rainfall_datetime", "rainfall_value"}, true
	case "soil_moisture":
		return NHCSoilMoisutre, []string{"tele_station_oldcode", "soil_datetime", "soil_value"}, TW30SoilMoisutre, []string{"tele_station_oldcode", "soil_datetime", "soil_value"}, true
	case "solar":
		return NHCSolar, []string{"tele_station_oldcode", "solar_datetime", "solar_value"}, TW30Solar, []string{"tele_station_oldcode", "solar_datetime", "solar_value"}, true
	case "tele_watergate":
		return NHCTeleWaterGate, []string{"tele_station_oldcode", "watergate_datetime", "watergate_in", "watergate_out", "watergate_out2"}, TW30TeleWaterGate, []string{"tele_station_oldcode", "watergate_datetime", "watergate_in", "watergate_out", "watergate_out2"}, true
	case "tele_waterlevel":
		return NHCTeleWaterlevel, []string{"tele_station_oldcode", "waterlevel_datetime", "waterlevel_m", "waterlevel_msl", "flow_rate", "discharge"}, TW30TeleWaterlevel, []string{"tele_station_oldcode", "waterlevel_datetime", "waterlevel_m", "waterlevel_msl", "flow_rate", "discharge"}, true
	case "temp":
		return NHCTemp, []string{"tele_station_oldcode", "temp_datetime", "temp_value"}, TW30Temp, []string{"tele_station_oldcode", "temp_datetime", "temp_value"}, true
	case "wind":
		return NHCWind, []string{"tele_station_oldcode", "wind_datetime", "wind_speed", "wind_dir", "wind_dir_value"}, TW30Wind, []string{"tele_station_oldcode", "wind_datetime", "wind_speed", "wind_dir", "wind_dir_value"}, true
	case "wl_ford":
		return NHCFord, []string{"ford_station_oldcode", "ford_waterlevel_datetime", "ford_waterlevel_value", "comm_status"}, TW30Ford, []string{"ford_station_oldcode", "ford_waterlevel_datetime", "ford_waterlevel_value", "comm_status"}, true
	case "canal_station":
		return NHCCanal, []string{}, TW30MCanal, []string{}, false
	case "dam":
		return NHCMDam, []string{}, TW30Dam, []string{}, false
	case "ford_station":
		return NHCMFord, []string{}, TW30MFord, []string{}, false
	case "tele_station":
		return NHCTele, []string{}, TW30Tele, []string{}, false
	case "m_medium_dam":
		return NHCMedium, []string{}, TW30Medium, []string{}, false
	case "agency_document":
		return NHCAgency, []string{}, TW30Agency, []string{}, false
	case "basin":
		return NHCBasin, []string{}, TW30Basin, []string{}, false
	case "subbasin":
		return NHCSubbasin, []string{}, TW30Subbasin, []string{}, false
	case "ref_geocode":
		return NHCGeocode, []string{}, TW30Geocode, []string{}, false
	case "media_type":
		return NHCMediaType, []string{}, TW30MediaType, []string{}, false
	case "power_plant":
		return NHCPowerPlant, []string{}, TW30PowerPlant, []string{}, false
	case "egat_new_metada":
		return NHCEgatMetadata, []string{}, TW30EgatMetadata, []string{}, false
	case "safety_zone":
		return NHCMSzone, []string{}, TW30Szone, []string{}, false
	case "xsection_station":
		return NHCXsection, []string{}, TW30Xsection, []string{}, false
	case "ground_wr":
		return NHCMGround, []string{}, TW30MGround, []string{}, false
	case "rule_curve":
		return NHCRuleCurve, []string{}, TW30RuleCurve, []string{}, false
	case "xsection_data":
		return NHCXSectionData, []string{}, TW30XSectionData, []string{}, false
	case "wl_groundwater":
		return NHCGroundWater, []string{}, TW30GroundWater, []string{}, false
	case "lt_region_tmd":
		return NHCTMDRegion, []string{}, TW30TMDRegion, []string{}, false
	case "water_resource":
		return NHCWaterResource, []string{}, TW30WaterResource, []string{}, false
	}

	return "", nil, "", nil, false
}

func getDataTW30Table(q string, header []string, p []interface{}) (interface{}, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, err
	}
	data := make([]interface{}, 0)
	data1 := make([]interface{}, 0)
	if len(header) == 0 {
		header, _ = rows.Columns()
	}

	data = append(data, header)
	for rows.Next() {
		v := make([]interface{}, len(header))
		for i := range v {
			v[i] = new(interface{})
		}
		rows.Scan(v...)
		data1 = append(data1, v)
	}
	data = append(data, data1)
	return data, nil
}

func getDataNHCTable(q string, header []string, table string, p []interface{}) (interface{}, error) {
	var db *pqx.DB
	var err error
	if table == "ref_geocode" {
		//		db, err = sql.Open("postgres", "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/staging?sslmode=disable")
		db, err = OpenStaging()
	} else {
		//		db, err = sql.Open("postgres", "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/nhc?sslmode=disable")
		db, err = OpenNhc()
	}

	if err != nil {
		return nil, err
	}
	//	defer db.Close()
	rows, err := db.Query(q, p...)
	data := make([]interface{}, 0)
	data1 := make([]interface{}, 0)
	if len(header) == 0 {
		header, _ = rows.Columns()
	}
	if err != nil {
		return nil, err
	}
	data = append(data, header)
	for rows.Next() {
		v := make([]interface{}, len(header))
		for i := range v {
			v[i] = new(interface{})
		}
		rows.Scan(v...)
		data1 = append(data1, v)
	}
	defer db.Close()
	data = append(data, data1)
	return data, nil
}

//	run background job - count data
func RegenData() (interface{}, error) {

	//	log.WrapPanicGO(updateData)
	_, err := backgroundjob.RunBackgroundJob(Cmd, Cmd_RegenData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//	run background job - count master data
func RegenMasterData() (interface{}, error) {

	//	log.WrapPanicGO(updateMasterData)
	_, err := backgroundjob.RunBackgroundJob(Cmd, Cmd_RegenMasterData)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//	count master data
func UpdateMasterData() error {
	a, _ := getLastMasterDataNHC()
	b, _ := getLastMasterDataTW30()
	q := "UPDATE migrate_log.summary_master_data SET nhc_row=$2, thaiwater30_row=$3, last_update=NOW() WHERE nhc_table=$1"
	for k := range a {
		v1 := a[k]
		v2 := b[k]
		p := []interface{}{k, v1, v2}
		err := uData(q, p)
		if err != nil {
			return err
		}
	}
	return nil
}

//	count data
func UpdateData() error {
	a, err := GetLastDataNHC()
	if err != nil {
		return err
	}
	b, err := getLastDataTW30()
	if err != nil {
		return err
	}
	q := "UPDATE migrate_log.summary_data_row SET nhc_row=$2, thaiwater30_row=$3, last_update=NOW() WHERE nhc_table=$1"
	for k := range a {
		v1 := a[k]
		v2 := b[k]
		p := []interface{}{k, v1, v2}
		err := uData(q, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func uData(q string, p []interface{}) error {

	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func getLastMasterDataNHC() (map[string]interface{}, error) {

	//	db, err := sql.Open("postgres", "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/nhc?sslmode=disable")
	db, err := OpenNhc()
	if err != nil {
		return nil, err
	}
	//	defer db.Close()
	q := SumMasterNHC
	rows, err := db.Query(q)
	data := map[string]interface{}{}
	for rows.Next() {
		var (
			table string
			count int64
		)
		rows.Scan(&table, &count)
		data[table] = count
	}

	return data, nil
}

func getLastMasterDataTW30() (map[string]interface{}, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := SumMasterTW30
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := map[string]interface{}{}
	for rows.Next() {
		var (
			table string
			count int64
		)
		rows.Scan(&table, &count)
		data[table] = count
	}

	return data, nil
}

func Query3Time(db *pqx.DB, SumQuery []string) (map[string]interface{}, error) {
	var err error
	data := map[string]interface{}{}
	mapIsError := map[int]error{}
	var hasErr bool = true
	for count := 1; count <= 3 && hasErr; count++ { // วนลูป 3ครั้ง เพื่อคิวรี่ข้อมูลให้ได้
		log.Log("try ", count)
		hasErr = false
		for index, q := range SumQuery {
			log.Log("query ", index)
			if b, ok := mapIsError[index]; ok { // เคยคิวรี่ผ่านแล้วให้ข้าม เพื่อที่จะให้คิวรี่เฉพาะตัวที่เคย err
				if b == nil {
					log.Log("found and no error")
					continue
				} else {
					hasErr = true
					log.Log("found with error", b)
				}

			}

			//		log.Println(q)

			rows, err := db.Query(q)
			if err != nil {
				log.Log("error", err)
				hasErr = true
				mapIsError[index] = err
				continue
				//			return nil, err
			}
			for rows.Next() {
				var (
					table string
					count int64
				)
				rows.Scan(&table, &count)
				data[table] = count
				mapIsError[index] = nil
			}
		}
	}
	return data, err
}

func GetLastDataNHC() (map[string]interface{}, error) {

	//	db, err := sql.Open("postgres", "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/nhc?sslmode=disable")
	db, err := OpenNhc()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	//	q := SumNHC
	//	rows, err := db.Query(q)
	//	if err != nil {
	//		return nil, err
	//	}
	//	data := map[string]interface{}{}
	//	for rows.Next() {
	//		var (
	//			table string
	//			count int64
	//		)
	//		rows.Scan(&table, &count)
	//		data[table] = count
	//	}

	return Query3Time(db, SumNHC)
}

func getLastDataTW30() (map[string]interface{}, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	//	q := SumTW30
	//	p := []interface{}{}
	//	rows, err := db.Query(q, p...)
	//	if err != nil {
	//		return nil, pqx.GetRESTError(err)
	//	}
	//	data := map[string]interface{}{}
	//	for rows.Next() {
	//		var (
	//			table string
	//			count int64
	//		)
	//		rows.Scan(&table, &count)
	//		data[table] = count
	//	}

	data := map[string]interface{}{}
	for _, q := range SumTW30 {
		rows, err := db.Query(q)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var (
				table string
				count int64
			)
			rows.Scan(&table, &count)
			data[table] = count
		}
	}

	return data, nil
}

func GetSummaryImage() (interface{}, string, bool, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, "", false, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := "SELECT media_type_id, nhc_row, nhc_file_count, tw_cv_ss, tw_cv_err, rec_diff, tw30_file_count,last_import_date,media_type_name,media_subtype_name,last_update_media_type FROM migrate_log.summary_img order by media_type_id"
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, "", false, pqx.GetRESTError(err)
	}
	var LupDate time.Time
	data := make([]*SummaryImg, 0)
	for rows.Next() {

		var (
			media_type_id          int64
			nhc_row                int64
			nhc_file_count         int64
			tw_cv_ss               int64
			tw_cv_err              int64
			rec_diff               int64
			tw30_file_count        int64
			date                   time.Time
			media_type_name        sql.NullString
			media_subtype_name     sql.NullString
			last_update_media_type time.Time
		)
		rows.Scan(&media_type_id, &nhc_row, &nhc_file_count, &tw_cv_ss,
			&tw_cv_err, &rec_diff, &tw30_file_count, &date, &media_type_name, &media_subtype_name, &last_update_media_type)
		dataRow := &SummaryImg{}
		dataRow.MediaTypeID = media_type_id
		dataRow.NHCRow = nhc_row
		dataRow.NHCFileCount = nhc_file_count
		dataRow.TW30CVSS = tw_cv_ss
		dataRow.TW30CVErr = tw_cv_err
		dataRow.RecDiff = rec_diff
		dataRow.TW30FileCount = tw30_file_count
		dataRow.LastMigrateDate = date.Format(time.RFC3339)
		dataRow.MediaTypeName = media_type_name.String
		dataRow.MediaSubTypeName = media_subtype_name.String
		dataRow.LastUpdate = last_update_media_type.Format(time.RFC3339)
		data = append(data, dataRow)
		if LupDate.Before(last_update_media_type) {
			LupDate = last_update_media_type
		}
	}

	isBgjobRunning, err := bgjob.IsBgJobRunning(Cmd + " " + Cmd_RegenDataImg)
	if err != nil {
		return nil, "", false, pqx.GetRESTError(err)
	}

	return data, LupDate.Format("2006-01-02 15:04"), isBgjobRunning, nil
}

type SummaryImg struct {
	MediaTypeID      int64  `json:"media_type_id"`
	NHCRow           int64  `json:"nhc_row_count"`
	NHCFileCount     int64  `json:"nhc_file_count"`
	TW30CVSS         int64  `json:"thaiwater30_convert_success"`
	TW30CVErr        int64  `json:"thaiwater30_convert_error"`
	RecDiff          int64  `json:"row_diff"`
	TW30FileCount    int64  `json:"thaiwater30_file_count"`
	LastMigrateDate  string `json:"last_migrate_date"`
	MediaTypeName    string `json:"media_type_name"`
	MediaSubTypeName string `json:"media_subtype_name"`
	LastUpdate       string `json:"last_update"`
}

type ImageOutput struct {
	NHCImageUrl  string `json:"nhc_image_url"`
	TW30ImageUrl string `json:"thaiwater30_image_url"`
}

func GetImgByMedia(media_type_id int64) (interface{}, error) {
	data := &ImageOutput{}
	switch media_type_id {
	case 1:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/ContourImg/2013/03/30/hahumidY2013M03D30T21.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/countour_image/humidity/haii/media/2013/03/30/hahumidY2013M03D30T21.png")
	case 2:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/ContourImg/2013/03/11/hatempY2013M03D11T04.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/countour_image/temperature/haii/media/2013/03/11/hatempY2013M03D11T04.png")
	case 3:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/ContourImg/2014/05/11/hapressY2014M05D11T09.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/countour_image/pressure/haii/media/2014/05/11/hapressY2014M05D11T09.png")
	case 4:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrfroms/v2/domain03/2016/08/04/19/rain_init_201608041201.d03.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/rain_accumulate/thailand/haii/media/2016/08/04/19/rain_init_201608041201.d03.jpg")
	case 5:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/wrf_image/rain_accumulation/domain01/7days/2014/04/05/19/wrfout_racc_d01_init201404051900_20140407.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/rain_accumulate/asia/haii/media/2014/04/05/19/wrfout_racc_d01_init201404051900_20140407.png")
	case 6:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/rain_accumulation/domain02/7days/2015/07/09/07/wrfout_racc_d02_init201507090700_20150712.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/rain_accumulate/southeast_asia/haii/media/2015/07/09/07/wrfout_racc_d02_init201507090700_20150712.png")
	case 7:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/wind/domain01/7days/2014/07/04/19/wrfout_wind_d01_init201407041900_20140707_070000.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/wind/asia/haii/media/2014/07/04/19/wrfout_wind_d01_init201407041900_20140707_070000.png")
	case 8:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/wind/domain02/7days/2013/05/10/07/wrfout_wind_d02_init201305100700_20130512_060000.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/wind/southeast_asia/haii/media/2013/05/10/07/wrfout_wind_d02_init201305100700_20130512_060000.png")
	case 9:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/wind/domain03/3days/2014/06/09/19/wrfout_wind_d03_init201406091900_20140610_200000.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/wind/thailand/haii/media/2014/06/09/19/wrfout_wind_d03_init201406091900_20140610_200000.png")
	case 10:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/upper_wind_600m/domain01/nowcast/2015/03/upper_wind_0.6km_ini_2015-03-10_12_plot_2015-03-10_12.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/upper_wind_600m/asia/haii/media/2015/03/10/07/upper_wind_0.6km_ini_2015-03-10_12_plot_2015-03-10_12.png")
	case 11:
		data.NHCImageUrl = ""
		data.TW30ImageUrl = ""
	case 14:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/SST_W/2016/diff_20160131_20160227.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/sst_w/global/haii/media/2016/diff_20160131_20160227.png")
	case 15:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/pressure_600m/domain01/nowcast/2015/09/pressure_0.6km_ini_2015-09-05_00_plot_2015-09-05_00.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/pressure_600m/asia/haii/media/2015/09/05/19/pressure_0.6km_ini_2015-09-05_00_plot_2015-09-05_00.png")
	case 16:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/pressure_1500m/domain01/nowcast/2014/08/pressure_1.5km_ini_2014-08-07_12_plot_2014-08-07_12.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/pressure_1500m/asia/haii/media/2014/08/08/07/pressure_1.5km_ini_2014-08-08_12_plot_2014-08-08_12.png")
	case 17:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/precipitation/domain03/3days/2015/08/11/19/wrfout_prec_d03_init201508111900_20150812_040000.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/precipitation/thailand/haii/media/2015/08/11/19/wrfout_prec_d03_init201508111900_20150812_040000.png")
	case 18:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/precipitation/domain01/7days/2015/03/03/19/wrfout_prec_d01_init201503031900_20150305_170000.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/precipitation/asia/haii/media/2015/03/03/19/wrfout_prec_d01_init201503031900_20150305_170000.png")
	case 19:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/wrf_image/precipitation/domain02/7days/2014/06/10/19/wrfout_prec_d02_init201406101900_20140612_160000.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/precipitation/southeast_asia/haii/media/2014/06/10/19/wrfout_prec_d02_init201406101900_20140612_160000.png")
	case 23:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/wave_height/sea/navy/2015/05/20/19/sea006.gif"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/wave_height/south_east_asia/hd/media/2015/05/20/19/sea006.gif")
	case 24:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/wave_height/thai/navy/2014/09/06/19/thai013.gif"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/wave_height/south_china_sea/hd/media/2014/09/06/19/thai013.gif")
	case 25:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/wave_height/ind/navy/2013/03/04/19/ind010.gif"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/wave_height/indian_ocean_northern/hd/media/2013/03/04/19/ind010.gif")
	case 26:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/pom/navy/2014/09/05/19/pom012.gif"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/sea_surface_elevation/global/hd/media/2014/09/05/19/pom012.gif")
	case 27:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/map/weather_map/tmd/2014/03/05/13/2014-03-05_TopChart_13.JPG"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/weather_map/thailand/tmd/media/2014/03/05/13/2014-03-05_TopChart_13.JPG")
	case 28:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/map/upper_wind_850hPa/tmd/2016/08/03/13/2016-08-03_13_UpperWind850.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/upper_wind_850hpa/thailand/tmd/media/2016/08/03/13/2016-08-03_13_UpperWind850.jpg")
	case 29:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/map/upper_wind_925hPa/tmd/2014/06/03/19/2014-06-03_19_UpperWind925.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/upper_wind_925hpa/thailand/tmd/media/2014/06/03/19/2014-06-03_19_UpperWind925.jpg")
	case 30:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/radar/phs/2015/07/10/phs240_201507100100.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/radar/phs/tmd/media/2015/07/10/phs240_201507100100.jpg")
	case 32:
		data.NHCImageUrl = ""
		data.TW30ImageUrl, _ = b64.EncryptText("")
	case 33:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/map/coastalradar/2016/10/Current_GOT_20161011_Av25hr_0000_GULF1_resize.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/radar/coastal/gistda/media/2016/10/Current_GOT_20161011_Av25hr_0000_GULF1_resize.jpg")
	case 140:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/sea/SSH_W/2013/sshlow_diff_20130224_20130303.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/ssh_w/global/haii/media/2013/sshlow_diff_20130224_20130303.png")
	case 141:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sat/mtsat/nexsat/2015/03/02/se.15030201.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/himawari-8/thailand/us_naval/media/2015/03/se.15030201.jpg")
	case 142:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/SSH_EVENT/2012/MARIA-12/ssh_diff_20120923_20120930.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/ssh_event/global/haii/media/2012/ssh_diff_20120923_20120930.png")
	case 149:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/map/precip_usda/AFWA/2014/01/precip.20140110.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/modis/precipitaion/usda/media/2014/precip.20140110.jpg")
	case 150:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/map/precip_usda/AFWA/2014/01/precip_n.20140110.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/modis/precipitaion/usda/media/2014/precip_n.20140110.jpg")
	case 151:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/map/precip_usda/WMO/2014/01/precip_n_s.20140110.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/modis/precipitaion/usda/media/2014/precip_n_s.20140110.jpg")
	case 152:
		data.NHCImageUrl = "http://www.nhc.in.th/product/report/map/precip_usda/WMO/2014/01/precip_s.20140110.jpg"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/modis/precipitaion/usda/media/2014/precip_s.20140110.jpg")
	case 157:
		data.NHCImageUrl = "http://www.nhc.in.th/product/history/sea/SST_M/2015/diff_20150101_20150228.png"
		data.TW30ImageUrl, _ = b64.EncryptText("product/image/satellite_rain/10km/haii/media/2015/diff_20150101_20150228.png")

	}

	return data, nil
}
