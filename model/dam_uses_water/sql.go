// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package dam_uses_water

import (
	"strconv"
	"strings"
	"time"
	//	"os"
	//	"fmt"
)

//	genarate sql string
func Gen_SQL_GetDamUsesWater(p *Struct_DamUsesWater_InputParam) (string, []interface{}) {

	year, _, _ := time.Now().Date()

	var param []interface{}
	var (
		strDateStart  = strconv.Itoa(year-1) + "-01-01"
		strDateEnd    = strconv.Itoa(year-1) + "-12-31"
		strSql        string
		strSqlByDamId string
		strWhere      string
	)

	// Dam Fillter
	arrDamId := []string{}
	if p.Dam_id != "" {
		arrDamId = strings.Split(p.Dam_id, ",")
	}

	if len(arrDamId) > 0 {
		strWhere += " AND "
		if len(arrDamId) == 1 {
			param = append(param, strings.Trim(p.Dam_id, " "))
			strWhere += " dam_id = $" + strconv.Itoa(len(param))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrDamId {
				param = append(param, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(param)))
			}
			strWhere += " dam_id IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	// Province Fillter
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}

	if len(arrProvinceId) > 0 && len(arrDamId) <= 0 {
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

	// Date Fillter
	if p.Start_date != "" {
		strDateStart = p.Start_date
	}

	if p.End_date != "" {
		strDateEnd = p.End_date
	}

	// Default date
	if p.Start_date == "" {
		strDateStart = "(SELECT dam_date FROM dam_daily WHERE dam_uses_water IS NOT NULL ORDER BY dam_date DESC LIMIT 1)"
	}

	if p.End_date == "" {
		strDateEnd = "(SELECT dam_date FROM dam_daily WHERE dam_uses_water IS NOT NULL  ORDER BY dam_date DESC LIMIT 1)"
	}

	if len(arrProvinceId) > 0 && len(arrDamId) <= 0 {

	}

	strSql = `SELECT `

	if p.Start_date == "" && p.End_date == "" {
		strSql += `` + strDateStart + ` AS start_date, ` + strDateEnd + ` AS end_date,`
	} else {
		strSql += `'` + strDateStart + `' AS start_date, '` + strDateEnd + `' AS end_date,`
	}
	strSql += ` 
		SUM (dam_uses_water) AS uses,
		AVG (dam_uses_water_percent) AS percent,
		SUM (dam_released) AS released,
		geo.province_code AS province_code,geo.province_name::TEXT AS province_name,
		geo.area_code AS area_code,geo.area_name::TEXT AS area_name
	FROM
		dam_daily dd
	LEFT JOIN m_dam st ON dd.dam_id = st.id
	LEFT JOIN agency agt ON agt.id = st.agency_id
	LEFT JOIN subbasin sb ON sb.id = st.subbasin_id
	LEFT JOIN basin b ON b.id = sb.basin_id
	LEFT JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE
		dam_id IN (
			SELECT
				d. ID
			FROM
				"public"."m_dam" d
			LEFT JOIN "public"."lt_geocode" ge ON d.geocode_id = ge. ID
		)`

	if p.Start_date == "" {
		strSql += `AND dam_date >= ` + strDateStart + ` `
	} else {
		strSql += `AND dam_date >= '` + strDateStart + ` `
	}

	if p.End_date == "" {
		strSql += `AND dam_date <= ` + strDateEnd + ` `
	} else {
		strSql += `AND dam_date <= '` + strDateEnd + `' `
	}

	strSql += strWhere

	strSql += `
	GROUP BY province_code,province_name::TEXT,area_code,area_name::TEXT
	ORDER BY province_code `

	strSqlByDamId = `SELECT `

	if p.Start_date == "" {
		strSqlByDamId += `` + strDateStart + ` AS start_date, `
	} else {
		strSqlByDamId += `'` + strDateStart + `' AS start_date, `
	}

	if p.End_date == "" {
		strSqlByDamId += `` + strDateEnd + ` AS end_date, `
	} else {
		strSqlByDamId += `'` + strDateEnd + `' AS end_date, `
	}
	strSqlByDamId += ` 
		SUM (dam_uses_water) AS uses,
		AVG (dam_uses_water_percent) AS percent,
		SUM (dam_released) AS released,
		'' AS province_code,'' AS province_name,
		'' AS area_code,'' AS area_name
	FROM
		dam_daily dd
	LEFT JOIN m_dam st ON dd.dam_id = st.id
	LEFT JOIN agency agt ON agt.id = st.agency_id
	LEFT JOIN subbasin sb ON sb.id = st.subbasin_id
	LEFT JOIN basin b ON b.id = sb.basin_id
	LEFT JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE `

	if p.Start_date == "" {
		strSqlByDamId += ` dam_date >= ` + strDateStart + ` `
	} else {
		strSqlByDamId += ` dam_date >= '` + strDateStart + `' `
	}

	if p.End_date == "" {
		strSqlByDamId += `AND dam_date <= ` + strDateEnd + ` `
	} else {
		strSqlByDamId += `AND dam_date <= '` + strDateEnd + `' `
	}

	strSqlByDamId += strWhere

	if len(arrDamId) > 0 {
		strSql = strSqlByDamId
	}

	//	fmt.Println(strSql)
	//	os.Exit(3)
	return strSql, param
}
