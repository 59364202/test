// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package rainfall_1d 

import (
	"strings"
	"strconv"
	"time"
)

//	genarate sql string
func Gen_SQL_GetRainfall1d(p *Param_Rainfall1d) (string, []interface{}){
	var param []interface{}
	var (
		strSql        string
		strWhere      string 
	)
	
	// Province Fillter
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}

	if len(arrProvinceId) > 0 {
		strWhere += " AND "
		if len(arrProvinceId) == 1 {
			param = append(param, strings.Trim(p.Province_Code, " "))
			strWhere += " province_code = $" + strconv.Itoa(len(param))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				param = append(param, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(param)))
			}
			strWhere += " province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}
	
	// Date Start Format '2006-01-02'
	strStartDate := time.Now().AddDate(0, 0, 0).Format("2006-01-02")
	 
	strSql = `
	WITH sum_query AS (
		SELECT 
			tele_station_id,sum( rainfall_value) AS rainfall1d,
			array_to_string(array_agg(rd.rainfall_date_calc), '')  as rainfall_date_calc,
			ROW_NUMBER () OVER (PARTITION BY province_code ORDER BY sum( rainfall_value) DESC) AS ROW
		FROM rainfall_daily rd
		LEFT JOIN m_tele_station m  ON rd.tele_station_id = m.id 
		LEFT JOIN lt_geocode lg  ON m.geocode_id = lg.id  
			AND m.is_ignore = 'false' 
			AND lg.geocode <> '999999' 
		WHERE rainfall_datetime >= '` + strStartDate + `' ` + strWhere + `
		GROUP BY tele_station_id,province_code HAVING sum(rainfall_value) < 2000::double precision  
		ORDER BY province_code,rainfall1d DESC
	)
	SELECT 
		tele_station_id, date_trunc('day'::text, now()) AS start_date, 
		date_trunc('day'::text, now()) AS end_date,rainfall_date_calc::TIMESTAMP,rainfall1d,
		tele_station_oldcode,tele_station_lat,tele_station_long,
		area_code,province_code,area_name,province_name,amphoe_name,tumbon_name,agency_id,agency_name,tele_station_name 
	FROM sum_query sq
	LEFT JOIN m_tele_station m  ON sq.tele_station_id = m.id 
	LEFT JOIN lt_geocode lg  ON m.geocode_id = lg.id  
		AND m.is_ignore = 'false' 
		AND lg.geocode <> '999999' 
	LEFT JOIN agency a ON m.agency_id = a.id 
	WHERE sq.row = 1 
	`
		
	return strSql, param
}
