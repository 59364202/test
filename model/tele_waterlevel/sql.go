package tele_waterlevel

import "time"

//var SQL_SelectWaterLevel = `SELECT *
//FROM
//  (SELECT tw.tele_station_id, tw.waterlevel_datetime, tw.waterlevel_m, tw.waterlevel_msl,
//          ts.subbasin_id, ts.agency_id, ts.tele_station_name, ts.tele_station_lat, ts.tele_station_long, ts.tele_station_oldcode,
//          ts.ground_level, ts.min_bank, lg.tumbon_name, lg.amphoe_name, lg.province_code, lg.province_name,
//          ((tw.waterlevel_m - ts.ground_level) * 100) / (ts.min_bank - ts.ground_level)AS storage_percent,
//          b.basin_name
//   FROM v_tele_waterlevel_1h tw
//   INNER JOIN
//     (SELECT id, subbasin_id, agency_id, geocode_id, tele_station_name, tele_station_lat, tele_station_long, tele_station_oldcode,
//             ground_level,
//             CASE
//                 WHEN left_bank > right_bank THEN right_bank
//                 ELSE left_bank
//             END AS min_bank
//      FROM m_tele_station
//      WHERE deleted_at = '1970-01-01 07:00:00+07' ) ts ON tw.tele_station_id = ts.id
//   INNER JOIN lt_geocode lg ON ts.geocode_id = lg.id
//   INNER JOIN subbasin sb ON ts.subbasin_id = sb.id
//   INNER JOIN basin b ON sb.basin_id = b.id
//   ORDER BY b.id, lg.province_code, lg.amphoe_code, lg.tumbon_code, ts.id
//   ) q`
//var SQL_SelectWaterLevel = `SELECT tw.tele_station_id, tw.waterlevel_datetime, tw.waterlevel_m, tw.waterlevel_msl, ts.subbasin_id, ts.agency_id, ts.tele_station_name,
//          ts.tele_station_lat, ts.tele_station_long, ts.tele_station_oldcode, ts.ground_level, ts.min_bank, lg.tumbon_name,
//          lg.amphoe_name, lg.province_code, lg.province_name,
//           CASE
//			WHEN (ts.min_bank - ts.ground_level) != 0 THEN
// 			((tw.waterlevel_msl - ts.ground_level) * 100) / (ts.min_bank - ts.ground_level)
//			ELSE null
//		  END AS storage_percent,
//          b.basin_name,
//          lg.tumbon_code, lg.amphoe_code, b.id AS basin_id, tw.id AS data_id
//   FROM v_tele_waterlevel_1h tw
//   INNER JOIN
//     (SELECT id, subbasin_id, agency_id, geocode_id, tele_station_name, tele_station_lat, tele_station_long, tele_station_oldcode,
//             ground_level,
//             CASE
//                 WHEN left_bank > right_bank THEN right_bank
//                 ELSE left_bank
//             END AS min_bank
//      FROM m_tele_station
//      WHERE deleted_at = '1970-01-01 07:00:00+07' AND is_ignore = 'false' ) ts ON tw.tele_station_id = ts.id
//   INNER JOIN lt_geocode lg ON ts.geocode_id = lg.id
//   INNER JOIN subbasin sb ON ts.subbasin_id = sb.id
//   INNER JOIN basin b ON sb.basin_id = b.id`

//   LEFT JOIN
//     ( SELECT tele_station_id, sum(rainfall1h) AS raintoday
//      FROM rainfall_1h
//      WHERE date(rainfall_datetime_calc) = CURRENT_DATE
//      GROUP BY tele_station_id ) rtd ON tw.tele_station_id = rtd.tele_station_id
//   LEFT JOIN
//     (SELECT tele_station_id, sum(rainfall1h) AS rainyesterday
//      FROM rainfall_1h
//      WHERE date(rainfall_datetime_calc) = CURRENT_DATE - 1
//      GROUP BY tele_station_id ) ryd ON tw.tele_station_id = ryd.tele_station_id
//   WHERE tw.waterlevel_msl <> 999999

// var SQL_SelectTeleCanal = `
// SELECT     tw.id,
//            tw.tele_station_id,
//            tw.waterlevel_datetime,
//            tw.waterlevel_msl,
//            tw.pre_waterlevel_msl,
//            m.subbasin_id,
//            m.agency_id,
//            m.geocode_id,
//            m.tele_station_name,
//            m.tele_station_lat,
//            m.tele_station_long,
//            m.tele_station_oldcode,
//            ground_level,
//            CASE
//                       WHEN left_bank > right_bank THEN right_bank
//                       ELSE left_bank
//            END AS min_bank ,
//            m.sort_order,
//            'tele_waterlevel' AS table
// FROM       (
//                     SELECT   id,
//                              tele_station_id,
//                              waterlevel_msl AS waterlevel_msl,
//                              waterlevel_datetime,
//                              Rank() OVER ( partition BY tele_station_id ORDER BY waterlevel_datetime DESC) AS pos,
//                              Lead(waterlevel_msl) OVER ( partition BY tele_station_id )                    AS pre_waterlevel_msl
//                     FROM     tele_waterlevel
//                     WHERE    deleted_at = '1970-01-01 07:00:00+07') tw
// INNER JOIN m_tele_station m
// ON         tw.tele_station_id = m.id
// AND        m.is_ignore = 'false'
// WHERE      (
//                       m.agency_id = 3
//            OR         m.agency_id = 8
//            OR         m.agency_id = 10 )
// AND        tw.pos = 1
// UNION ALL
// SELECT     tw.id,
//            tw.tele_station_id,
//            tw.waterlevel_datetime,
//            tw.waterlevel_msl,
//            tw.pre_waterlevel_msl,
//            m.subbasin_id,
//            m.agency_id,
//            m.geocode_id,
//            m.tele_station_name,
//            m.tele_station_lat,
//            m.tele_station_long,
//            m.tele_station_oldcode,
//            ground_level,
//            CASE
//                       WHEN left_bank > right_bank THEN right_bank
//                       ELSE left_bank
//            END AS min_bank ,
//            m.sort_order,
//            'tele_waterlevel' AS table
// FROM       (
//                     SELECT   id,
//                              tele_station_id,
//                              waterlevel_msl AS waterlevel_msl,
//                              waterlevel_datetime,
//                              Rank() OVER ( partition BY tele_station_id ORDER BY waterlevel_datetime DESC) AS pos,
//                              Lead(waterlevel_msl) OVER ( partition BY tele_station_id )                    AS pre_waterlevel_msl
//                     FROM     tele_waterlevel
//                     WHERE    deleted_at = '1970-01-01 07:00:00+07'
//                     AND      date_part('minute'::text, waterlevel_datetime) = '0'::DOUBLE PRECISION ) tw
// INNER JOIN m_tele_station m
// ON         tw.tele_station_id = m.id
// AND        m.is_ignore = 'false'
// WHERE      m.agency_id = 9
// AND        tw.pos = 1
// UNION ALL
// SELECT     cw.id,
//            cw.canal_station_id,
//            cw.canal_waterlevel_datetime,
//            cw.waterlevel_msl,
//            cw.pre_waterlevel_msl,
//            m.subbasin_id,
//            m.agency_id,
//            m.geocode_id,
//            m.canal_station_name,
//            m.canal_station_lat,
//            m.canal_station_long,
//            m.canal_station_oldcode,
//            NULL AS ground_level,
//            NULL AS min_bank,
//            m.sort_order,
//            'canal_waterlevel ' AS TABLE
// FROM       (
//                     SELECT   id,
//                              canal_station_id,
//                              canal_waterlevel_value AS waterlevel_msl,
//                              canal_waterlevel_datetime,
//                              rank() OVER ( partition BY canal_station_id ORDER BY canal_waterlevel_datetime DESC) AS pos,
//                              lead(canal_waterlevel_value) OVER ( partition BY canal_station_id)                   AS pre_waterlevel_msl
//                     FROM     canal_waterlevel
//                     WHERE    deleted_at = '1970-01-01 07:00:00+07') cw
// INNER JOIN m_canal_station m
// ON         cw.canal_station_id = m.id
// AND        m.is_ignore = 'false'
// WHERE      cw.pos = 1
// `
// var SQL_SelectWaterLevel = `
// SELECT     ct.id,
//            ct.tele_station_id,
//            ct.waterlevel_datetime,
//            ct.waterlevel_m,
//            ct.waterlevel_msl,
//            ct.pre_waterlevel_msl,
//            ct.subbasin_id,
//            ct.agency_id,
//            ct.tele_station_name,
//            ct.tele_station_lat,
//            ct.tele_station_long,
//            ct.tele_station_oldcode,
//            ct.ground_level,
//            ct.min_bank,
//            CASE
//            		WHEN waterlevel_msl = 999999 THEN NULL
//                	WHEN (ct.min_bank - ct.ground_level) != 0 THEN ((ct.waterlevel_msl - ct.ground_level) * 100) / (ct.min_bank - ct.ground_level)
//                 ELSE NULL
//            END AS storage_percent,
//            ct.sort_order,
//            ct.table,
//            lg.tumbon_name,
//            lg.amphoe_name,
//            lg.province_code,
//            lg.province_name,
//            b.basin_name,
//            lg.tumbon_code,
//            lg.amphoe_code,
//            b.id AS basin_id,
//            agency.agency_shortname
// FROM       (
//                       SELECT     tw.id,
//                                  tw.tele_station_id,
//                                  tw.waterlevel_datetime,
//                                  tw.waterlevel_m,
//                                  tw.waterlevel_msl,
//                                  tw.pre_waterlevel_msl,
//                                  m.subbasin_id,
//                                  m.agency_id,
//                                  m.geocode_id,
//                                  m.tele_station_name,
//                                  m.tele_station_lat,
//                                  m.tele_station_long,
//                                  m.tele_station_oldcode,
//                                  ground_level,
//                                  CASE
//                                             WHEN left_bank > right_bank THEN right_bank
//                                             ELSE left_bank
//                                  END AS min_bank ,
//                                  m.sort_order,
//                                  'tele_waterlevel' AS table
//                       FROM       (
//                                           SELECT   id,
//                                                    tele_station_id,
//                                                    waterlevel_m,
//                                                    waterlevel_msl AS waterlevel_msl,
//                                                    waterlevel_datetime,
//                                                    Rank() OVER ( partition BY tele_station_id ORDER BY waterlevel_datetime DESC) AS pos,
//                                                    Lead(waterlevel_msl) OVER ( partition BY tele_station_id )                    AS pre_waterlevel_msl
//                                           FROM     tele_waterlevel
//                                           WHERE    deleted_at = '1970-01-01 07:00:00+07') tw
//                       INNER JOIN m_tele_station m
//                       ON         tw.tele_station_id = m.id
//                       AND        m.is_ignore = 'false'
//                       WHERE      (
//                                             m.agency_id = 3
//                                  OR         m.agency_id = 8
//                                  OR         m.agency_id = 10 )
//                       AND        tw.pos = 1
//                       UNION ALL
//                       SELECT     tw.id,
//                                  tw.tele_station_id,
//                                  tw.waterlevel_datetime,
//                                  tw.waterlevel_m,
//                                  tw.waterlevel_msl,
//                                  tw.pre_waterlevel_msl,
//                                  m.subbasin_id,
//                                  m.agency_id,
//                                  m.geocode_id,
//                                  m.tele_station_name,
//                                  m.tele_station_lat,
//                                  m.tele_station_long,
//                                  m.tele_station_oldcode,
//                                  ground_level,
//                                  CASE
//                                             WHEN left_bank > right_bank THEN right_bank
//                                             ELSE left_bank
//                                  END AS min_bank ,
//                                  m.sort_order,
//                                  'tele_waterlevel' AS table
//                       FROM       (
//                                           SELECT   id,
//                                                    tele_station_id,
//                                                    waterlevel_m,
//                                                    waterlevel_msl AS waterlevel_msl,
//                                                    waterlevel_datetime,
//                                                    Rank() OVER ( partition BY tele_station_id ORDER BY waterlevel_datetime DESC) AS pos,
//                                                    Lead(waterlevel_msl) OVER ( partition BY tele_station_id )                    AS pre_waterlevel_msl
//                                           FROM     tele_waterlevel
//                                           WHERE    deleted_at = '1970-01-01 07:00:00+07'
//                                           AND      date_part('minute'::text, waterlevel_datetime) = '0'::DOUBLE PRECISION ) tw
//                       INNER JOIN m_tele_station m
//                       ON         tw.tele_station_id = m.id
//                       AND        m.is_ignore = 'false'
//                       WHERE      m.agency_id = 9
//                       AND        tw.pos = 1
//                       UNION ALL
//                       SELECT     cw.id,
//                                  cw.canal_station_id,
//                                  cw.canal_waterlevel_datetime,
//                                  NULL AS waterlevel_m,
//                                  cw.waterlevel_msl,
//                                  cw.pre_waterlevel_msl,
//                                  m.subbasin_id,
//                                  m.agency_id,
//                                  m.geocode_id,
//                                  m.canal_station_name,
//                                  m.canal_station_lat,
//                                  m.canal_station_long,
//                                  m.canal_station_oldcode,
//                                  NULL AS ground_level,
//                                  NULL AS min_bank,
//                                  m.sort_order,
//                                  'canal_waterlevel ' AS TABLE
//                       FROM       (
//                                           SELECT   id,
//                                                    canal_station_id,
//                                                    canal_waterlevel_value AS waterlevel_msl,
//                                                    canal_waterlevel_datetime,
//                                                    rank() OVER ( partition BY canal_station_id ORDER BY canal_waterlevel_datetime DESC) AS pos,
//                                                    lead(canal_waterlevel_value) OVER ( partition BY canal_station_id)                   AS pre_waterlevel_msl
//                                           FROM     canal_waterlevel
//                                           WHERE    deleted_at = '1970-01-01 07:00:00+07') cw
//                       INNER JOIN m_canal_station m
//                       ON         cw.canal_station_id = m.id
//                       AND        m.is_ignore = 'false'
//                       WHERE      cw.pos = 1 ) ct
// INNER JOIN lt_geocode lg
// ON         ct.geocode_id = lg.id
// INNER JOIN subbasin sb
// ON         ct.subbasin_id = sb.id
// INNER JOIN basin b
// ON         sb.basin_id = b.id
// INNER JOIN agency
// ON		   ct.agency_id = agency.id
// `
// var SQL_SelectWaterLevel_OrderBy = " ORDER BY storage_percent DESC NULLS LAST "

// var SQL_SelectWaterLevelAnHourBefore = `SELECT *
// FROM
//   (SELECT tw.tele_station_id, tw.waterlevel_datetime, tw.waterlevel_m, tw.waterlevel_msl, ts.subbasin_id, ts.agency_id, ts.tele_station_name,
//           ts.tele_station_lat, ts.tele_station_long, ts.tele_station_oldcode, ts.ground_level, ts.min_bank, lg.tumbon_name,
//           lg.amphoe_name, lg.province_code, lg.province_name,
//           CASE
// 			WHEN (ts.min_bank - ts.ground_level) != 0 THEN
//  			((tw.waterlevel_msl - ts.ground_level) * 100) / (ts.min_bank - ts.ground_level)
// 			ELSE null
// 		  END AS storage_percent,
//           b.basin_name
//    FROM v_tele_waterlevel_1h tw
//    INNER JOIN
//      (SELECT id, subbasin_id, agency_id, geocode_id, tele_station_name, tele_station_lat, tele_station_long, tele_station_oldcode,
//              ground_level,
//              CASE
//                  WHEN left_bank > right_bank THEN right_bank
//                  ELSE left_bank
//              END AS min_bank
//       FROM m_tele_station
//       WHERE deleted_at = '1970-01-01 07:00:00+07' AND is_ignore = 'false' ) ts ON tw.tele_station_id = ts.id
//    INNER JOIN lt_geocode lg ON ts.geocode_id = lg.id
//    INNER JOIN subbasin sb ON ts.subbasin_id = sb.id
//    INNER JOIN basin b ON sb.basin_id = b.id
//    WHERE tw.waterlevel_datetime BETWEEN date_trunc('hour', Now() - interval '1 hour') AND date_trunc('hour', now())
//    ORDER BY b.id, lg.province_code, lg.amphoe_code, lg.tumbon_code, ts.id
//   ) q`

//var SQL_SelectWaterLevel_Where = " WHERE true "
//var SQL_SelectWaterLevel_Where = " WHERE storage_percent > 100 OR storage_percent <= 10 "

var arrWaterlevelLastestHeader = []string{
	"id",
	//"tele_station_id",
	"waterlevel_datetime",
	"waterlevel_m",
	"waterlevel_msl",
	"flow_rate",
	"discharge",
}

var sqlGetWaterlevelByStationAndDate = ` SELECT tw.id
											  , waterlevel_datetime
											  , waterlevel_m
											  , ROUND((waterlevel_msl + COALESCE(mt.offset,0))::numeric, 2) AS waterlevel_msl
											  , flow_rate
											  , discharge
									 FROM tele_waterlevel tw
								        INNER JOIN m_tele_station mt ON mt.id = tw.tele_station_id
									 WHERE tele_station_id = $1
									   AND waterlevel_datetime BETWEEN $2 AND $3
									   AND tw.deleted_at = to_timestamp(0) `

var sqlGetWaterlevelByStationAndDateOrderBy = ` ORDER BY waterlevel_datetime DESC `

var sqlUpdateToDeleteWaterlevel = ` UPDATE tele_waterlevel
								SET deleted_by = $1
								  , deleted_at = NOW()
								  , updated_by = $1
								  , updated_at = NOW() `

var sqlGetErrorWaterlevel = `  SELECT dd.id
									, tele_station_oldcode
									, waterlevel_datetime
									, tele_station_name
									, province_name
									, agency_name
									, agency_shortname
									, waterlevel_m
									, waterlevel_msl
									--, waterlevel_in
									--, waterlevel_out
									--, waterlevel_out2
									, flow_rate
									, discharge
									, d.id AS station_id
							FROM tele_waterlevel dd 	   
							LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE ((dd.waterlevel_m = -9999 OR dd.waterlevel_m = 999999)
								OR (dd.waterlevel_msl = -9999 OR dd.waterlevel_msl = 999999)
								--OR (dd.waterlevel_in = -9999 OR dd.waterlevel_in = 999999)
								--OR (dd.waterlevel_out = -9999 OR dd.waterlevel_out = 999999)
								--OR (dd.waterlevel_out2 = -9999 OR dd.waterlevel_out2 = 999999)
								OR (dd.flow_rate = -9999 OR dd.flow_rate = 999999)
								OR (dd.discharge = -9999 OR dd.discharge = 999999))
							  AND (dd.deleted_by IS NULL) `

var SQLSelectIgnoreData = ` SELECT NOW()
								, 'tele_waterlevel' AS data_category
								, d.id AS station_id
								, d.tele_station_oldcode
								, d.tele_station_name
								, g.province_name
								, agt.agency_shortname
								--, agt.agency_name
								, dd.id
								, dd.waterlevel_datetime
								, '#{Remarks}' AS remark
								, #{UserID} AS user_created
								, #{UserID} AS user_updated
								, NOW()
								, NOW()
								, waterlevel_msl AS data_value
							FROM tele_waterlevel dd 	   
							LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE true `

// WHERE d.is_ignore <> $1 `

// water level chart in index page, นักวิเคราะห์
//ระดับน้ำกรมชล. ข้อมูลมาทั้ง ม.รทก และม.รสม ใส่อยู่ใน field เดีวกันคือ waterlevel_m
//ข้อมูล ม.รสม ต้องแปลงเป็น ม.รทก โดย ดูที่ field station_type_msl 1 = ม.รทก., 0 = ม.รสม. โดยระดับน้ำ ม.รสม ต้องเอามาบวก offset ก่อน
//ข้อมูลแสดง เงื่อไนตาม agency_id
//รายชม.
//1 กรมเจ้าท่า
//8 กฟผ.
//12 กรมชล
//ข้อมูบแสดงราย 15 นาที
//3 กรมน้ำฯ
//10 กทม.
//ข้อมูลแสดงราย 10 นาที
//9 สสน.
var sqlSelectWaterlevelByStationAndDateAnalyst = `
SELECT M.ID,
	A.datetime,
	A.waterlevel_msl 
FROM
	(SELECT gs.datetime,	
		CASE 
			WHEN ((station_type_msl = 0) AND ("offset" IS NOT NULL) AND (agency_id = 12)) THEN 
				(tw.waterlevel_m + COALESCE("offset", (0)::real))
	    WHEN (((station_type_msl = 1) OR (station_type_msl IS NULL)) AND (agency_id = 12)) THEN 
		    tw.waterlevel_m
      ELSE tw.waterlevel_msl + COALESCE("offset", (0)::real)
    END AS waterlevel_msl
	FROM
		PUBLIC.tele_waterlevel tw
		INNER JOIN m_tele_station M ON M.ID = tw.tele_station_id AND tw.tele_station_id =$1
		RIGHT JOIN ( SELECT generate_series ($2::DATE,$3, '5 min' ) AS datetime ) gs 
		ON tw.waterlevel_datetime BETWEEN $2 AND $3 
		AND tw.waterlevel_datetime = gs.datetime  AND tw.deleted_at = to_timestamp( 0 ) 
		AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
	ORDER BY gs.datetime ASC 
)
	A LEFT JOIN m_tele_station M ON M.ID =$1
WHERE
CASE		
	WHEN M.agency_id = 1 OR M.agency_id = 8 OR M.agency_id = 12 THEN
		( date_part( 'minute', A.datetime ) :: INTEGER % 60 ) = 0 
	WHEN M.agency_id = 3 OR M.agency_id = 10 THEN
		( date_part( 'minute', A.datetime ) :: INTEGER % 15 ) = 0 
	WHEN M.agency_id = 9 THEN
		( date_part( 'minute', A.datetime ) :: INTEGER % 10 ) = 0 
END`

// watergate ปตร./ฝาย
var sqlSelectWatergateByStationAndDateAnalyst = `SELECT m.id,a.datetime,  
		case when a.watergate_in = 999999 then null else a.watergate_in end as watergate_in,
		case when a.watergate_out = 999999 then null else a.watergate_out end as watergate_out
	FROM (SELECT gs.datetime, tw.watergate_in, tw1.watergate_out 
	FROM public.tele_watergate tw  ` +
	"INNER JOIN m_tele_station m ON m.id = tw.tele_station_id AND tw.tele_station_id=$1 " +
	"RIGHT JOIN ( select generate_series($2::date,$3, '5 min') as datetime  " +
	")gs ON tw.watergate_datetime BETWEEN $2 AND $3 AND tw.watergate_datetime=gs.datetime AND tw.deleted_at=to_timestamp(0) " +
	"LEFT JOIN public.tele_watergate tw1 ON tw1.watergate_datetime BETWEEN $2 AND $3 AND tw1.tele_station_id=$1 AND tw1.watergate_datetime=gs.datetime " +
	" WHERE (tw.qc_status IS NULL  OR tw.qc_status ->> 'is_pass' = 'true') AND ( tw1.qc_status IS NULL  OR tw1.qc_status ->> 'is_pass' = 'true') " +
	"ORDER BY gs.datetime ASC) a  " +
	"LEFT JOIN m_tele_station m ON m.id=$1 " +
	"WHERE CASE WHEN m.agency_id = 1 THEN  " +
	"( date_part('minute', a.datetime )::integer % 60) = 0  " +
	"WHEN m.agency_id = 3 THEN	 " +
	"( date_part('minute', a.datetime )::integer % 15) = 0  " +
	"WHEN m.agency_id = 10 THEN	  " +
	"( date_part('minute', a.datetime )::integer % 15) = 0  " +
	"WHEN m.agency_id = 9 THEN	 " +
	"( date_part('minute', a.datetime )::integer % 10) = 0 " +
	"END "

// watergate
var sqlSelectWaterlevelInOutLatest = `SELECT tele_station_name, tw.watergate_datetime, 
case when watergate_in = 999999 then null else watergate_in end as watergate_in,
case when watergate_out = 999999 then null else watergate_out end as watergate_out
,tw.watergate_out2 
, m.id, m.tele_station_lat, m.tele_station_long, m.tele_station_oldcode , g.id, g.geocode, g.province_code, g.province_name, g.amphoe_code, g.amphoe_name 
, g.tumbon_code, g.tumbon_name, m.pump, m.floodgate, tw.pump_on, tw.floodgate_open, tw.floodgate_height
,m.agency_id,a.agency_name,a.agency_shortname,G.area_code,G.area_name,b.id AS basin_id,b.basin_name,watergate_out_datetime,m.subbasin_id 
FROM latest.tele_watergate tw LEFT JOIN public.m_tele_station m ON tw.tele_station_id=m.id  
LEFT JOIN public.lt_geocode g ON m.geocode_id=g.id 
LEFT JOIN agency a ON M.agency_id = a.id 
LEFT JOIN subbasin sb ON m.subbasin_id = sb.id 
LEFT JOIN basin b ON sb.basin_id = b.id 
WHERE tw.deleted_at=to_timestamp(0) AND m.deleted_at=to_timestamp(0) AND tele_station_name IS NOT NULL  
 AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true')
 ORDER BY watergate_datetime desc `

var sqlSelectWaterlevelLatestForFloodforecast = "SELECT data_id_current,datetime_current,value_current, m.id,m.tele_station_name,m.tele_station_lat, " +
	"m.tele_station_long,m.tele_station_oldcode,a.id,a.agency_name,a.agency_shortname " +
	"FROM cache.latest_waterlevel lw  " +
	"LEFT JOIN public.m_tele_station m ON m.id=lw.tele_station_id " +
	"LEFT JOIN public.agency a ON a.id=m.agency_id"

func SQL_WaterlevelBasinGraphAnalystAdvance(subbasin_id int64, datetime string) (string, interface{}) {
	dt, err := time.Parse("2006-01-02 15:04", datetime)
	if err != nil {
		dt, _ = time.Parse("2006-01-02", datetime)
	}
	// sql ย้อนหลัง 3row ถ้าไม่มีีไปข้างหน้า 3row
	mdSQLWhere := SQL_WaterlevelBasinGraphAnalystAdvance_MD(dt, 60)
	dwrSQLWhere := SQL_WaterlevelBasinGraphAnalystAdvance_MD(dt, 15)
	haiiSQLWhere := SQL_WaterlevelBasinGraphAnalystAdvance_MD(dt, 10)

	sql := `
  SELECT
      *
  FROM
      (
          SELECT
              m.id,
              m.tele_station_name,
              m.agency_id,
              ground_level,
              CASE WHEN left_bank > right_bank THEN right_bank ELSE left_bank END AS min_bank,
              m.sort_order,
              CASE WHEN t.waterlevel_msl IS NULL THEN CASE WHEN m.agency_id = 1 THEN
              /* MD ทุก 1 ชั่วโมง */
              ` + mdSQLWhere + `
              WHEN m.agency_id = 3
              OR m.agency_id = 10 THEN
              /* DWR, BMW ทุก 15 นาที */
              ` + dwrSQLWhere + `
              ELSE
              /* HAII ทุก 10 นาที */
              ` + haiiSQLWhere + `
              END ELSE t.waterlevel_msl END AS value,
              'tele_waterlevel' AS type
          FROM
              m_tele_station m
              LEFT JOIN (
                  SELECT
                      m.id,
                      tw.waterlevel_msl
                  FROM
                      m_tele_station m
                      INNER JOIN tele_waterlevel tw ON tw.tele_station_id = m.id
                      AND m.agency_id IN (1, 3, 9, 10)
                      AND m.subbasin_id = $1
                      AND tw.deleted_at = to_timestamp(0)
                  WHERE
                      tw.waterlevel_datetime = '` + datetime + `'
              ) t ON m.id = t.id
          WHERE
              m.agency_id IN (1, 3, 9, 10)
              AND m.subbasin_id = $1
          UNION ALL
          SELECT
              m.id,
              m.canal_station_name,
              m.agency_id,
              NULL,
              NULL,
              m.sort_order,
              t.canal_waterlevel_value,
              'canal_waterlevel' AS type
          FROM
              m_canal_station m
              LEFT JOIN (
                  SELECT
                      m.id,
                      cw.canal_waterlevel_value
                  FROM
                      m_canal_station m
                      INNER JOIN canal_waterlevel cw ON cw.canal_station_id = m.id
                      AND m.agency_id IN (1, 3, 9, 10)
                      AND m.subbasin_id = $1
                      AND cw.deleted_at = to_timestamp(0)
                  WHERE
                      cw.canal_waterlevel_datetime = '` + datetime + `'
              ) t ON m.id = t.id
          WHERE
              m.agency_id IN (1, 3, 9, 10)
              AND m.subbasin_id = $1
      ) a
  ORDER BY
      sort_order
  `

	return sql, subbasin_id
}

//SQL_WaterlevelBasinGraphAnalystAdvance_MD
//sql สำหรับหาค่า ย้อนหลัง 3row ถ้าไม่มีีไปข้างหน้า 3row ของ SQL_WaterlevelBasinGraphAnalystAdvance เฉพาะหน่วยงาน MD เท่านั้น
//MD ทุก 1 ชั่วโมง
//DWR, BMW ทุก 15 นาที
//HAII ทุก 10 นาที
//--
// dt วัน-เวลา,
// interval ความถี่ข้อมูล เช่น 60 (1 ชั่วโมง), 15(15 นาที), 30 (30 นาที)
func SQL_WaterlevelBasinGraphAnalystAdvance_MD(dt time.Time, interval int) string {
	layout := "2006-01-02 15:04"
	interP1 := dt.Add(time.Duration(interval*-1) * time.Minute).Format(layout) // ย้อนหลัง 1 row
	interP2 := dt.Add(time.Duration(interval*-2) * time.Minute).Format(layout) // ย้อนหลัง 2 row
	interP3 := dt.Add(time.Duration(interval*-3) * time.Minute).Format(layout) // ย้อนหลัง 3 row
	interN1 := dt.Add(time.Duration(interval*1) * time.Minute).Format(layout)  // ข้างหน้า 1 row
	interN2 := dt.Add(time.Duration(interval*2) * time.Minute).Format(layout)  // ข้างหน้า 2 row
	interN3 := dt.Add(time.Duration(interval*3) * time.Minute).Format(layout)  // ข้างหน้า 3 row

	sql := `
  (
  SELECT
    waterlevel_msl
  FROM
    (
          SELECT
            waterlevel_msl
          FROM(
                    SELECT
                            waterlevel_datetime,
                            waterlevel_msl
                    FROM
                            tele_waterlevel
                    WHERE
                            (
                                    waterlevel_datetime = '` + interP1 + `'
                                    OR waterlevel_datetime = '` + interP2 + `'
                                    OR waterlevel_datetime = '` + interP3 + `'
                            )
                            AND tele_station_id = m.id
                            AND deleted_at = To_timestamp(0)
                            AND (
                                    qc_status IS NULL
                                    OR qc_status ->> 'is_pass' = 'true'
                            )
                    order by
                            waterlevel_datetime desc
            ) r3
            UNION ALL
            SELECT
                    waterlevel_msl
            FROM
                    tele_waterlevel
            WHERE
                    (
                            waterlevel_datetime = '` + interN1 + `'
                            OR waterlevel_datetime = '` + interN2 + `'
                            OR waterlevel_datetime = '` + interN3 + `'
                    )
                    AND tele_station_id = m.id
                    AND deleted_at = To_timestamp(0)
                    AND (
                            qc_status IS NULL
                            OR qc_status ->> 'is_pass' = 'true'
                    )
    ) a
  WHERE
    waterlevel_msl IS NOT NULL
  LIMIT
    1
  )
  `
	return sql
}

// var SQL_WaterlevelBasinGraphAnalystAdvance = `
// SELECT   *
// FROM     (
//                 SELECT m.id,
//                        m.tele_station_name,
//                        m.agency_id,
//                        ground_level,
//                        CASE
//                               WHEN left_bank > right_bank THEN right_bank
//                               ELSE left_bank
//                        END AS min_bank ,
//                        m.sort_order,
//                        CASE
//                               WHEN t.waterlevel_msl IS NULL THEN
//                                      CASE
//                                             WHEN m.agency_id = 1 THEN /* MD ทุก 1 ชั่วโมง */
//                                                    (
//                                                               SELECT     waterlevel_msl
//                                                               FROM       tele_waterlevel itw
//                                                               RIGHT JOIN (
//                                                                          (
//                                                                                   SELECT   *
//                                                                                   FROM     generate_series($2::timestamp - interval '3 hour ', $2::timestamp - interval '1 hour ', '1 hour' )
//                                                                                   ORDER BY generate_series DESC)
//                                                                 UNION ALL
//                                                                 SELECT *
//                                                                 FROM   generate_series($2::timestamp + interval '1 hour ', $2::timestamp + interval '3 hour ', '1 hour' ) ) gs
//                                                                   ON     itw.waterlevel_datetime = gs.generate_series
//                                                                   AND    tele_station_id = m.id
//                                                                   AND    itw.deleted_at = to_timestamp(0)
//                                                                   WHERE  waterlevel_msl IS NOT NULL AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') limit 1 )
//                                         WHEN m.agency_id = 3
//                                         OR         m.agency_id = 10 THEN /* DWR, BMW ทุก 15 นาที */
//                                                    (
//                                                               SELECT     waterlevel_msl
//                                                               FROM       tele_waterlevel itw
//                                                               RIGHT JOIN (
//                                                                          (
//                                                                                   SELECT   *
//                                                                                   FROM     generate_series($2::timestamp - interval '45minute ', $2::timestamp - interval '15 minute ', '15 minute' )
//                                                                                   ORDER BY generate_series DESC)
//                                                                 UNION ALL
//                                                                 SELECT *
//                                                                 FROM   generate_series($2::timestamp + interval '15 minute ', $2::timestamp + interval '45 minute ', '15 minute' ) ) gs
//                                                                   ON     itw.waterlevel_datetime = gs.generate_series
//                                                                   AND    tele_station_id = m.id
//                                                                   AND    itw.deleted_at = to_timestamp(0)
//                                                                   WHERE  waterlevel_msl IS NOT NULL AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') limit 1 )
//                                         ELSE /* HAII ทุก 10 นาที */
//                                               (
//                                               SELECT     waterlevel_msl
//                                               FROM       tele_waterlevel itw
//                                               RIGHT JOIN (
//                                                          (
//                                                                   SELECT   *
//                                                                   FROM     generate_series($2::timestamp - interval '30 minute ', $2::timestamp - interval '10 minute ', '10 minute' )
//                                                                   ORDER BY generate_series DESC)
//                                                 UNION ALL
//                                                 SELECT *
//                                                 FROM   generate_series($2::timestamp + interval '10 minute ', $2::timestamp + interval '30 minute ', '10 minute' ) ) gs
//                                                   ON     itw.waterlevel_datetime = gs.generate_series
//                                                   AND    tele_station_id = m.id
//                                                   AND    itw.deleted_at = to_timestamp(0)
//                                                   WHERE  waterlevel_msl IS NOT NULL AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') limit 1 )
//                              END
//                   ELSE t.waterlevel_msl
//        END               AS value,
//        'tele_waterlevel' AS type
// FROM       d m
// LEFT JOIN
//            (
//                       SELECT     m.id,
//                                  tw.waterlevel_msl
//                       FROM       m_tele_station m
//                       INNER JOIN tele_waterlevel tw
//                       ON         tw.tele_station_id = m.id
//                       AND        m.agency_id IN ( 1, 3, 9, 10 )
//                       AND        m.subbasin_id = $1
//                       AND        tw.deleted_at = to_timestamp(0)
//                       WHERE      tw.waterlevel_datetime = $2::timestamp ) t
//     ON     m.id = t.id
//     WHERE  m.agency_id IN ( 1, 3, 9, 10 )
//     AND    m.subbasin_id = $1
//     UNION ALL
//     SELECT    m.id,
//               m.canal_station_name,
//               m.agency_id,
//               NULL,
//               NULL,
//               m.sort_order,
//               t.canal_waterlevel_value,
//               'canal_waterlevel' AS type
//     FROM      m_canal_station m
//     LEFT JOIN
//               (
//                          SELECT     m.id,
//                                     cw.canal_waterlevel_value
//                          FROM       m_canal_station m
//                          INNER JOIN canal_waterlevel cw
//                          ON         cw.canal_station_id = m.id
//                          AND        m.agency_id IN ( 1, 3, 9, 10 )
//                          AND        m.subbasin_id = $1
//                          AND        cw.deleted_at = to_timestamp(0)
//                          WHERE      cw.canal_waterlevel_datetime = $2::timestamp ) t
//     ON        m.id = t.id
//     WHERE     m.agency_id IN ( 1, 3, 9, 10 )
//     AND       m.subbasin_id = $1 ) a
// ORDER BY sort_order
// `

func SQL_WaterlevelBasinGraph24HAnalystAdvance() string {
	now := time.Now()
	layout := "2006-01-02 15:00"
	dt := now.Format(layout)
	dt23 := now.Add(-23 * time.Hour).Format(layout)
	dt24 := now.Add(-24 * time.Hour).Format(layout)
	s := `
SELECT   id, 
         tele_station_name, 
         tele_station_lat, 
         tele_station_long, 
         ground_level, 
         left_bank, 
         right_bank, 
         distance, 
         sort_order, 
         gs, 
         waterlevel_msl 
FROM     ( 
                   SELECT    * 
                   FROM      ( 
                                    SELECT m.id, 
                                           m.tele_station_name, 
                                           m.tele_station_lat, 
                                           m.tele_station_long, 
                                           m.ground_level, 
                                           m.left_bank, 
                                           m.right_bank, 
                                           m.distance, 
                                           m.sort_order 
                                    FROM   m_tele_station m 
                                    WHERE  m.agency_id IN (1, 
                                                           3, 
                                                           9, 
                                                           10) 
                                    AND    m.subbasin_id = $1 ) m  
                   NATURAL JOIN generate_series('` + dt23 + `'::timestamp, '` + dt + `', '1 hour') gs
                   LEFT    JOIN 
                             ( 
                                        SELECT     tele_station_id, 
                                                   waterlevel_datetime, 
                                                   waterlevel_msl 
                                        FROM       tele_waterlevel tw 
                                        INNER JOIN m_tele_station m 
                                        ON         m.id = tw.tele_station_id 
                                        AND        m.agency_id IN (1, 
                                                                   3, 
                                                                   9, 
                                                                   10) 
                                        AND        m.subbasin_id = $1 
                                        WHERE      tw.waterlevel_datetime BETWEEN '` + dt24 + `' AND '` + dt + `'
                                        AND        tw.deleted_at = to_timestamp(0) 
                                        AND        (qc_status IS NULL OR qc_status->>'is_pass' = 'true')) tw 
                   ON        m.id = tw.tele_station_id 
                   AND       gs = tw.waterlevel_datetime 
                   UNION ALL 
                   SELECT    * 
                   FROM      ( 
                                    SELECT m.id, 
                                           m.canal_station_name, 
                                           m.canal_station_lat, 
                                           m.canal_station_long, 
                                           NULL, 
                                           NULL, 
                                           NULL, 
                                           m.distance, 
                                           m.sort_order 
                                    FROM   m_canal_station m 
                                    WHERE  m.agency_id IN (1, 
                                                           3, 
                                                           9, 
                                                           10) 
                                    AND    m.subbasin_id = $1 ) m  
                   NATURAL JOIN generate_series('` + dt23 + `'::timestamp, '` + dt + `', '1 hour') gs
                   LEFT    JOIN 
                             ( 
                                        SELECT     canal_station_id, 
                                                   canal_waterlevel_datetime, 
                                                   canal_waterlevel_value 
                                        FROM       canal_waterlevel tw 
                                        INNER JOIN m_canal_station m 
                                        ON         m.id = tw.canal_station_id 
                                        AND        m.agency_id IN (1, 
                                                                   3, 
                                                                   9, 
                                                                   10) 
                                        AND        m.subbasin_id = $1 
                                        WHERE      tw.canal_waterlevel_datetime BETWEEN '` + dt24 + `' AND '` + dt + `'
                                        AND        tw.deleted_at = to_timestamp(0) 
                                        AND        (qc_status IS NULL OR qc_status->>'is_pass' = 'true')) tw 
                   ON        m.id = tw.canal_station_id 
                   AND       gs = tw.canal_waterlevel_datetime) a 
ORDER BY gs, 
         sort_order, 
         tele_station_name->>'th'
    `
	return s
}
