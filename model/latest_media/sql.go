// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package latest_media is a model for cache.latest_media This table store latest_media information.
package latest_media

import (
	"encoding/json"
	//	"strconv"
	model_setting "haii.or.th/api/server/model/setting"
	//	"fmt"
)

const cacheTableName = "cache.latest_media"

var SQL_Select = `
SELECT media_path, filename, media_datetime FROM ` + cacheTableName + `
WHERE agency_id = $1 AND media_type_id = $2
ORDER BY media_datetime
`

var SQL_SelectPreRain = `
SELECT * FROM (
  SELECT media_path, filename, media_datetime FROM ` + cacheTableName + `
  WHERE agency_id = 9 AND media_type_id = '17'
  UNION ALL
  SELECT media_path, filename, media_datetime FROM ` + cacheTableName + `
  WHERE agency_id = 9 AND media_type_id = '19'
  AND index IN ('d02_day04.jpg', 'd02_day05.jpg', 'd02_day06.jpg', 'd02_day07.jpg')
) a ORDER BY a.media_datetime
`
  
var SQL_Select_Rainforcase_Province = `
SELECT media_path, split_part(filename, '_', 2) AS filename, media_datetime FROM ` + cacheTableName + `
WHERE agency_id = $1 AND media_type_id = $2
AND media_path LIKE 'product/image/rain_forecast_7day/latest/haii/D%'
ORDER BY media_datetime
`

//var SQL_SelectPreRainAnimation = `
//SELECT md.media_path,
//       md.filename,
//       md.media_datetime
//FROM   (SELECT filename,
//               Max(media_datetime) AS media_datetime
//        FROM   media_animation
//        WHERE  filename IN ( 'ani_d03_large.mp4', 'ani_d02_large.mp4',
//                             'ani_d01_large.mp4' )
//               AND media_type_id IN ( '80', '81', '82' )
//               AND deleted_at = To_timestamp(0)
//        GROUP  BY filename
//        ORDER  BY filename DESC)t
//       INNER JOIN media_animation md
//               ON md.media_datetime = t.media_datetime
//                  AND md.filename = t.filename
//                  AND deleted_at = To_timestamp(0)
//GROUP  BY md.media_path,
//          md.filename,
//          md.media_datetime
//ORDER BY md.filename
//`

var SQL_SelectPreWave = `
SELECT lm.media_path, 
       lm.filename, 
       lm.media_datetime 
FROM   ` + cacheTableName + ` lm 
WHERE  agency_id = 9 
       AND media_type_id = 170 
ORDER  BY lm.media_datetime 
`

//var SQL_SelectPreWaveAnimation = `
//SELECT media_path, a.filename, a.media_datetime
//FROM
//  (SELECT MAX(media_datetime) AS media_datetime,
//          filename
//   FROM media_animation
//   WHERE filename IN ('wave_168hr.gif')
//   GROUP BY filename) b
//JOIN media_animation a ON b.media_datetime = a.media_datetime
//AND media_type_id = 83
//`

func Gen_SQL_Radar() (string, []interface{}) {
	var setting = make([]map[string]interface{}, 0)
	//	var whereAgency string = ""
	json.Unmarshal(model_setting.GetSystemSettingJSON("Frontend.analyst.Radar.RadarTypeOrder"), &setting)
	var (
		similarTo      string = ""
		array_position string = ""
	)
	for i, v := range setting {
		if i != 0 {
			similarTo += "|"
			array_position += ","
		}
		similarTo += v["radar_type"].(string)
		array_position += "'" + v["radar_type"].(string) + "'"
	}
	var itf = []interface{}{}
	//	sql := `
	//SELECT     m.media_datetime,
	//           m.media_path,
	//           m.filename,
	//           lm.filename
	//FROM       ` + cacheTableName + ` m
	//INNER JOIN
	//           (
	//                     SELECT    substring(m.filename from 1 FOR 6)AS filename,
	//                               max(media_datetime)               AS media_datetime
	//                     FROM      ` + cacheTableName + ` m
	//                     LEFT JOIN lt_media_type mt
	//                     ON        m.media_type_id=mt.id
	//                     WHERE     m.media_type_id = 30
	//                     AND       m.filename similar TO '(` + similarTo + `)%'
	//	` // phb240|kkn240|srn240|svp240|hhn240|cmp240|krb240|pkt240|cri240|stp240|skn240|phs240|lmp240|njk
	//	// cri240|lmp240|nan240|nan120|phs240|pnb240|kkn240|srn240|ubn240|chn240|kkw240|svp240|svp120|ryg240|skm240|skm120|hhn240|cmp240|srt240|pkt240|pkt120|krb240|stp240|nrt240|oki240|pmi240|sat240|tkh240|pn240|njk240|nk240
	//
	//	//	for i, v := range agency_id {
	//	//		if i != 0 {
	//	//			whereAgency += " OR "
	//	//		}
	//	//		whereAgency += " m.agency_id = $" + strconv.Itoa(i+1)
	//	//		itf = append(itf, v)
	//	//	}
	//	if whereAgency != "" {
	//		sql += " AND ( " + whereAgency + ") "
	//	}
	//	sql += `
	//GROUP BY  substring(m.filename FROM 1 FOR 6)
	//                     ORDER BY  array_position(array[` + array_position + `], substring(m.filename FROM 1 FOR 6))
	//           ) lm
	//ON         substring(m.filename FROM 1 FOR 6) = lm.filename
	//AND        m.media_datetime = lm.media_datetime
	//	`
	//'phb240','kkn240','srn240','svp240','hhn240','cmp240','krb240','pkt240','cri240','stp240','skn240','phs240','lmp240','njk'
	// 'cri240','lmp240','nan240','nan120','phs240','pnb240','kkn240','srn240','ubn240','chn240','kkw240','svp240','svp120','ryg240','skm240','skm120','hhn240','cmp240','srt240','pkt240','pkt120','krb240','stp240','nrt240','oki240','pmi240','sat240','tkh240','pn240','njk240','nk240'
	sql := `
SELECT media_datetime 
       , media_path 
       , filename 
       , index 
FROM   CACHE.latest_media 
WHERE  index IN ( ` + array_position + ` ) 
ORDER BY array_position(array[` + array_position + `], INDEX)
	`
	return sql, itf
}

//	get รูปภาพเรดาร์ ล่าสุด รายจังหวัด
//	Parameters:
//		province_id []string รหัสจังหวัด
//	Return:
//		string []interface{}
// 	Author: 
//		permporn@haii.or.th

func Gen_SQL_Radar_By_Code(code []string) (string, []interface{}) {
	
	var (
		similarTo      string = ""
		array_position string = ""
	)	
	var itf = []interface{}{}
	
	if len(code) == 0 {
		var setting = make([]map[string]interface{}, 0)
		//	var whereAgency string = ""
		json.Unmarshal(model_setting.GetSystemSettingJSON("Frontend.analyst.Radar.RadarTypeOrder"), &setting)
		
		for i, v := range setting {
			if i != 0 {
				similarTo += "|"
				array_position += ","
			}
			similarTo += v["radar_type"].(string)
			array_position += "'" + v["radar_type"].(string) + "'"
		}
		//array_position = "'phb240','kkn240','srn240','svp240','hhn240','cmp240','krb240','pkt240','cri240','stp240','skn240','phs240','lmp240','njk''cri240','lmp240','nan240','nan120','phs240','pnb240','kkn240','srn240','ubn240','chn240','kkw240','svp240','svp120','ryg240','skm240','skm120','hhn240','cmp240','srt240','pkt240','pkt120','krb240','stp240','nrt240','oki240','pmi240','sat240','tkh240','pn240','njk240','nk240'"
	
	}else{
		for i, v := range code {
			if i != 0 {
				similarTo += "|"
				array_position += ","
			}
			array_position += "'" + v + "'"
		}
	}
	
	sql := `SELECT media_datetime , media_path , filename , index  FROM   CACHE.latest_media WHERE  index IN ( `+ array_position + ` ) ORDER BY array_position(array[` + array_position + `], INDEX) `
	
	return sql, itf
}
