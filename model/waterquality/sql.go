package waterquality

var SQL_SelectWaterQuanlityLatest = `SELECT w.waterquality_id, w.waterquality_datetime, w.waterquality_do, 
       w.waterquality_ph, w.waterquality_temp, w.waterquality_turbid, 
       w.waterquality_bod, w.waterquality_tcb, w.waterquality_fcb, 
       w.waterquality_nh3n, w.waterquality_wqi, w.waterquality_ammonium, 
       w.waterquality_nitrate, w.waterquality_colorstatus, w.waterquality_status, 
       w.waterquality_salinity, w.waterquality_conductivity, w.waterquality_tds, 
       w.waterquality_chlorophyll, mws.waterquality_station_lat, mws.waterquality_station_long, 
       mws.waterquality_station_name, 
       lg.province_name, lg.province_code, lg.amphoe_name, lg.tumbon_name, a.agency_name, a.agency_shortname, 
       w.id AS data_id, mws.waterquality_station_oldcode AS oldcode,a.id,is_active,show_status
FROM waterquality w
INNER JOIN (
	SELECT waterquality_id , max(waterquality_datetime) as waterquality_datetime
	FROM waterquality
	WHERE deleted_at = '1970-01-01 07:00:00+07'
	GROUP BY waterquality_id
) tw 
  ON w.waterquality_id = tw.waterquality_id AND w.waterquality_datetime = tw.waterquality_datetime
INNER JOIN m_waterquality_station mws 
  ON w.waterquality_id = mws.id
INNER JOIN lt_geocode lg
  ON mws.geocode_id = lg.id
INNER JOIN agency a 
  ON mws.agency_id = a.id
WHERE w.deleted_at = '1970-01-01 07:00:00+07' AND mws.deleted_at = '1970-01-01 07:00:00+07' AND mws.is_ignore = 'false'`

func SQL_SelectWaterQualityGraph(field string) string {
	str := "SELECT waterquality_" + field + " , waterquality_datetime FROM waterquality WHERE waterquality_id = $1 AND waterquality_datetime <= $2 AND waterquality_datetime >= $3 "
	return str
}

var SQLSelectIgnoreData = ` SELECT NOW()
								, 'waterquality' AS data_category
								, d.id AS station_id
								, d.waterquality_station_oldcode
								, d.waterquality_station_name
								, g.province_name
								, agt.agency_shortname
								--, agt.agency_name
								, dd.id
								, dd.waterquality_datetime
								, '#{Remarks}' AS remark
								, #{UserID} AS user_created
								, #{UserID} AS user_updated
								, NOW()
								, NOW()
								, waterquality_salinity AS data_value
							FROM waterquality dd 	   
							LEFT JOIN m_waterquality_station d ON dd.waterquality_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE true `

// WHERE d.is_ignore <> $1 `

var (
	sqlSelectWaterQualityGraphCompare = "SELECT wq.waterquality_id,mwq.waterquality_station_name,wq.waterquality_datetime,mwq.agency_id"

	sqlSelectWaterQualityGraphCompareFROM = " FROM public.waterquality wq JOIN public.m_waterquality_station mwq ON wq.waterquality_id=mwq.id " +
		" WHERE wq.deleted_at=to_timestamp(0) AND (wq.waterquality_datetime BETWEEN $1 AND $2) " +
		" AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') "

	sqlSelectWaterQualityGraphCompareORDER = " ORDER BY mwq.waterquality_station_name->>'th'"

	sqlSelectWaterQualityGraphMultiParams = "SELECT wq.waterquality_id,mwq.waterquality_station_name,wq.waterquality_datetime,mwq.agency_id"

	sqlSelectWaterQualityGraphMultiParamsFrom = " FROM public.waterquality wq JOIN public.m_waterquality_station mwq ON wq.waterquality_id=mwq.id " +
		" WHERE wq.deleted_at=to_timestamp(0) AND (wq.waterquality_datetime BETWEEN $1 AND $2) AND wq.waterquality_id=$3 " +
		" AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') " +
		" ORDER BY mwq.waterquality_station_name->>'th'"

	sqlSelectWaterqualityCompare      = "SELECT m.waterquality_station_name,a.* FROM ( "
	sqlSelectWaterqualityFieldCompare = "SELECT gs.datetime "
	sqlSelectWaterqualityJoinCompare  = " FROM public.waterquality wq  " +
		"INNER JOIN m_waterquality_station m ON m.id = wq.waterquality_id AND "

	sqlSelectWaterqualityJoin2Compare = " RIGHT JOIN ( " +
		"select generate_series($1::date,$2, '10 min') as datetime " +
		")gs  " +
		"ON wq.waterquality_datetime=gs.datetime AND wq.deleted_at=to_timestamp(0) " +
		"ORDER BY gs.datetime ASC " +
		") a  " +
		"LEFT JOIN m_waterquality_station m ON "

	sqlSelectWaterqualityConditionCompare = " WHERE CASE WHEN m.agency_id = 14 THEN " +
		"( date_part('minute', a.datetime )::integer % 30) = 0 " +
		"ELSE( date_part('minute', a.datetime )::integer % 10) = 0 " +
		"END "

	sqlSelectWaterquality      = "SELECT m.waterquality_station_name, wq.* FROM ( "
	sqlSelectWaterqualityField = "SELECT gs.datetime "
	sqlSelectWaterqualityEnd   = `
	FROM   (SELECT * 
			FROM   waterquality wq 
			WHERE  waterquality_datetime BETWEEN 
				$1 AND $2 
				AND wq.waterquality_id = $3 
				AND wq.deleted_at = To_timestamp(0) 
				AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
			ORDER  BY wq.waterquality_datetime DESC) wq 
		RIGHT JOIN (SELECT Generate_series(  $1 :: date  , $2, '10 min') AS datetime) gs 
				ON wq.waterquality_datetime = gs.datetime) wq 
	LEFT JOIN m_waterquality_station m 
		ON m.id = $3 
	WHERE  CASE 
	WHEN m.agency_id = 14 THEN ( Date_part('minute', wq.datetime) :: INTEGER % 30  ) = 0 
	ELSE( Date_part('minute', wq.datetime) :: INTEGER % 10 ) = 0 
	END 
	ORDER BY datetime DESC
	`
	// sqlSelectWaterqualityEnd   = "FROM public.waterquality wq  " +
	// 	"INNER JOIN m_waterquality_station m ON m.id = wq.waterquality_id AND wq.waterquality_id=$3 " +
	// 	"RIGHT JOIN ( " +
	// 	"select generate_series($1::date,$2, '10 min') as datetime " +
	// 	")gs  " +
	// 	"ON wq.waterquality_datetime=gs.datetime AND wq.deleted_at=to_timestamp(0) " +
	// 	"ORDER BY gs.datetime ASC " +
	// 	") a  " +
	// 	"LEFT JOIN m_waterquality_station m ON m.id=$3 " +
	// 	"WHERE CASE WHEN m.agency_id = 14 THEN " +
	// 	"( date_part('minute', a.datetime )::integer % 30) = 0 " +
	// 	"ELSE( date_part('minute', a.datetime )::integer % 10) = 0 " +
	// 	"END "

	sqlSelectCanal = `
	SELECT m.canal_station_name, a.* 
	FROM   (SELECT gs.datetime, canal_waterlevel_value 
			FROM   (SELECT * 
					FROM   canal_waterlevel tw 
					WHERE  canal_waterlevel_datetime BETWEEN $1 AND $2 
						AND canal_station_id = $3 
						AND deleted_at = To_timestamp(0) 
						AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
					ORDER  BY canal_waterlevel_datetime ASC) wq 
				RIGHT JOIN (SELECT Generate_series($1 :: date, $2, '15 min') AS datetime) gs 
						ON wq.canal_waterlevel_datetime = gs.datetime) a 
		LEFT JOIN m_canal_station m ON m.id = $3 
	ORDER  BY datetime DESC 
	`
	// sqlSelectCanal = "SELECT m.canal_station_name,a.* FROM (SELECT gs.datetime,cw.canal_waterlevel_value " +
	// 	"FROM  (select generate_series($1::date,$2, '15 min') as datetime " +
	// 	") gs LEFT JOIN public.canal_waterlevel cw ON gs.datetime=cw.canal_waterlevel_datetime AND cw.deleted_at=to_timestamp(0) " +
	// 	"ORDER BY gs.datetime DESC )a  " +
	// 	"LEFT JOIN m_canal_station m ON m.id=$3"

	sqlSelectWaterlevel = "SELECT m.tele_station_name, a.* FROM ( SELECT gs.datetime"

	sqlSelectWaterlevelFrom = `
	FROM   (SELECT * 
			FROM   tele_waterlevel tw 
			WHERE  waterlevel_datetime BETWEEN $1 AND $2 
				AND tele_station_id = $3 
				AND deleted_at = To_timestamp(0) 
				AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
			ORDER  BY waterlevel_datetime ASC) wq 
		RIGHT JOIN (SELECT Generate_series($1 :: date, $2, '10 min') AS datetime) gs 
				ON wq.waterlevel_datetime = gs.datetime) a 
	LEFT JOIN m_tele_station m ON m.id = $3 
	WHERE  CASE 
	WHEN m.agency_id = 1 THEN ( Date_part('minute', a.datetime) :: INTEGER % 60 ) = 0 
	WHEN m.agency_id = 3 THEN ( Date_part('minute', a.datetime) :: INTEGER % 15 ) = 0 
	WHEN m.agency_id = 10 THEN ( Date_part('minute', a.datetime) :: INTEGER % 15 ) = 0 
	WHEN m.agency_id = 9 THEN ( Date_part('minute', a.datetime) :: INTEGER % 10 ) = 0 
	END 
	ORDER  BY datetime DESC 
	`
	// sqlSelectWaterlevelFrom = "FROM public.tele_waterlevel tw " +
	// 	"INNER JOIN m_tele_station m ON m.id = tw.tele_station_id AND tw.tele_station_id=$3 " +
	// 	"RIGHT JOIN ( select generate_series($1::date,$2, '5 min') as datetime " +
	// 	")gs ON tw.waterlevel_datetime=gs.datetime AND tw.deleted_at=to_timestamp(0) " +
	// 	"ORDER BY gs.datetime DESC) a " +
	// 	"LEFT JOIN m_tele_station m ON m.id=$3" +
	// 	"WHERE CASE WHEN m.agency_id = 1 THEN " +
	// 	"( date_part('minute', a.datetime )::integer % 60) = 0 " +
	// 	"WHEN m.agency_id = 3 THEN	" +
	// 	"( date_part('minute', a.datetime )::integer % 15) = 0 " +
	// 	"WHEN m.agency_id = 10 THEN	 " +
	// 	"( date_part('minute', a.datetime )::integer % 15) = 0 " +
	// 	"WHEN m.agency_id = 9 THEN	" +
	// 	"( date_part('minute', a.datetime )::integer % 10) = 0 " +
	// 	"END"

	sqlSelectWaterqualityCompare2     = "SELECT waterquality_id "
	sqlSelectWaterqualityCompareFrom2 = "FROM public.waterquality wq LEFT JOIN m_waterquality_station m ON m.id = wq.waterquality_id WHERE wq.waterquality_datetime=$1 AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') "

	sqlSelectWaterqualityMonitoring = `
	SELECT wq.waterquality_datetime 
			, mt.waterquality_station_name :: jsonb
			, wq.waterquality_salinity 
	FROM   cache.latest_waterquality wq 
	INNER JOIN m_waterquality_station mt ON mt.id = wq.waterquality_id
	WHERE  ( waterquality_id = $1 
			OR wq.waterquality_id = $2 
			OR wq.waterquality_id = $3 ) 
			AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') 
	ORDER BY mt.waterquality_station_name::jsonb, wq.waterquality_datetime DESC limit 3
	`
	// sqlSelectWaterqualityMonitoring = "SELECT DISTINCT ON (mt.waterquality_station_name::jsonb) wq.waterquality_datetime,mt.waterquality_station_name::jsonb, wq.waterquality_salinity " +
	// 	"FROM public.waterquality wq LEFT JOIN m_waterquality_station mt ON mt.id=wq.waterquality_id " +
	// 	"WHERE (wq.waterquality_id=$1 OR wq.waterquality_id=$2 OR wq.waterquality_id=$3) and wq.deleted_at=to_timestamp(0)" +
	// 	"GROUP BY wq.waterquality_datetime,mt.waterquality_station_name::jsonb,wq.waterquality_salinity ORDER BY mt.waterquality_station_name::jsonb,wq.waterquality_datetime DESC limit 3"
)
