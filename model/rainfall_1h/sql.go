// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainfall_1h is a model for public.rainfall_1h table. This table store rainfall_1h.
package rainfall_1h

import (
	//	"fmt"
	"time"
)

//	genarate sql string
func Gen_SQL_GetRainfallGraph(p *Param_Rainfall1h_Graph) (string, []interface{}) {
	var param []interface{}
	param = append(param, p.StationId)
	// strSQL := " SELECT gs.datetime, d.rainfall1h FROM rainfall_1h d "
	// var (
	// 	gsFrom      string
	// 	gsTo        string
	// 	gsInterval  string = " '1 hour' "
	// 	gsFormTable string = ""
	// )
	// if p.Is24 {
	// 	// ย้อนหลัง 24 ชั่วโมง
	// 	gsFrom = " date_trunc('hour', max(rainfall_datetime) - interval '24 hours') "
	// 	gsTo = " date_trunc('hour', max(rainfall_datetime)) "
	// 	gsFormTable = " FROM rainfall_1h WHERE tele_station_id = $1 "
	// } else if p.IsToday {
	// 	//	7โมง - ปัจจุบัน
	// 	gsFrom = " CURRENT_DATE + interval '7 hours' "
	// 	gsTo = " NOW() "
	// } else if p.IsDaily {
	// 	if p.DateStart == "" || p.DateEnd == "" {
	// 		//	7โมง 2วันที่แล้ว - 7โมง เมื่อวาน
	// 		gsFrom = " CURRENT_DATE - interval '2 day' + interval '7 hours' "
	// 		gsTo = " CURRENT_DATE - interval '1 day' + interval '7 hours' "
	// 	} else {
	// 		param = append(param, p.DateStart)
	// 		gsFrom = "$" + strconv.Itoa(len(param))

	// 		param = append(param, p.DateEnd)
	// 		gsTo = "$" + strconv.Itoa(len(param))
	// 	}

	// }
	// strSQL += `RIGHT JOIN (
	// 		SELECT generate_series(` + gsFrom + `, ` + gsTo + `,  ` + gsInterval + ` ) as datetime `
	// strSQL += gsFormTable + ") gs"
	// strSQL += " ON d.tele_station_id = $1 AND d.deleted_at = '1970-01-01 07:00:00+07' AND gs.datetime = d.rainfall_datetime "
	// strSQL += " WHERE ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') "
	// strSQL += " ORDER BY gs.datetime ASC"

	now := time.Now()
	//	var (
	//		gsInterval = " '1 hour' "
	//	)
	if p.Is24 {
		// ย้อนหลัง 24 ชั่วโมง
		p.DateStart = now.AddDate(0, 0, -1).Format("2006-01-02 15:00") // date_trunc('hour', max(rainfall_datetime) - interval '24 hours') "
		p.DateEnd = now.Format("2006-01-02 15:00")                     // date_trunc('hour', max(rainfall_datetime)) "
	} else if p.IsToday {
		//	7โมง - ปัจจุบัน
		p.DateStart = now.Format("2006-01-02") + " 07:00" // CURRENT_DATE + interval '7 hours' "
		p.DateEnd = now.Format(time.RFC3339Nano)          // NOW()
	} else if p.IsDaily {
		if p.DateStart == "" || p.DateEnd == "" {
			//	7โมง 2วันที่แล้ว - 7โมง เมื่อวาน
			p.DateStart = now.AddDate(0, 0, -2).Format("2006-01-02") + " 07:00" // CURRENT_DATE - interval '2 day' + interval '7 hours'
			p.DateEnd = now.AddDate(0, 0, -1).Format("2006-01-02") + " 07:00"   // CURRENT_DATE - interval '1 day' + interval '7 hours'
		}
	}
	param = append(param, p.DateStart)
	param = append(param, p.DateEnd)

	strSQL := `SELECT gs.datetime, rainfall1h
	FROM public.rainfall_1h data
	INNER JOIN m_tele_station m ON m.id = data.tele_station_id  AND data.tele_station_id = $1
	RIGHT JOIN ( SELECT generate_series ($2::date, $3, '1 hour' ) AS datetime ) gs 
		ON data.rainfall_datetime between $2 AND $3
		AND data.rainfall_datetime = gs.datetime  AND data.deleted_at = to_timestamp (0)  
		AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
	ORDER BY gs.datetime ASC `

	return strSQL, param
}
