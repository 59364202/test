package rainfall

var arrRainfallHeaderByStationAndDate = []string{
	"id",
	//	"tele_station_id",
	"rainfall_date",
	"rainfall5m",
	"rainfall10m",
	"rainfall15m",
	"rainfall30m",
	"rainfall1h",
	"rainfall3h",
	"rainfall6h",
	"rainfall12h",
	"rainfall24h",
	"rainfall_acc",
}

var sqlGetRainfallByStationAndDate = ` SELECT id
											, rainfall_datetime
											, rainfall5m
											, rainfall10m
											, rainfall15m
											, rainfall30m
											, rainfall1h
											, rainfall3h
											, rainfall6h
											, rainfall12h
											, rainfall24h
											, rainfall_acc
									 FROM rainfall
									 WHERE rainfall_datetime BETWEEN $2 AND $3
									   AND tele_station_id = $1
									   AND deleted_at = '1970-01-01 07:00:00+07' `

var sqlGetRainfallByStationAndDateOrderBy = ` ORDER BY rainfall_datetime DESC `

var sqlUpdateToDeleteRainfall = ` UPDATE rainfall
								SET deleted_by = $1
								  , deleted_at = NOW()
								  , updated_by = $1
								  , updated_at = NOW() `

var sqlGetErrorRainfall = `  SELECT dd.id
								    , tele_station_oldcode
								    , rainfall_datetime
								    , tele_station_name
								    , province_name
								    , agency_name
								    , agency_shortname
									, rainfall5m
									, rainfall10m
									, rainfall15m
									, rainfall30m
									, rainfall1h
									, rainfall3h
									, rainfall6h
									, rainfall12h
									, rainfall24h
									, rainfall_acc, d.id AS station_id
							FROM rainfall dd 	   
							LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE ((dd.rainfall5m = -9999 OR dd.rainfall5m = 999999)
								OR (dd.rainfall10m = -9999 OR dd.rainfall10m = 999999)
								OR (dd.rainfall15m = -9999 OR dd.rainfall15m = 999999)
								OR (dd.rainfall30m = -9999 OR dd.rainfall30m = 999999)
								OR (dd.rainfall1h = -9999 OR dd.rainfall1h = 999999)
								OR (dd.rainfall3h = -9999 OR dd.rainfall3h = 999999)
								OR (dd.rainfall6h = -9999 OR dd.rainfall6h = 999999)
								OR (dd.rainfall12h = -9999 OR dd.rainfall12h = 999999)
								OR (dd.rainfall24h = -9999 OR dd.rainfall24h = 999999)
								OR (dd.rainfall_acc = -9999 OR dd.rainfall_acc = 999999)) AND (dd.deleted_by IS NULL) `

var (
	sqlAdvRainMonthlyStationGraph = "SELECT rainfall_datetime, rainfall_value FROM public.rainfall_monthly WHERE deleted_at=to_timestamp(0) AND tele_station_id=$1 AND (rainfall_datetime BETWEEN $2 AND $3) AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') "

	sqlAdvRainMonthlyGraphByArea = "SELECT rainfall_datetime, sum(rainfall_value) FROM public.rainfall_monthly rm LEFT JOIN m_tele_station mts ON rm.tele_station_id=mts.id LEFT JOIN lt_geocode lg ON mts.geocode_id=lg.id WHERE rm.deleted_at=to_timestamp(0) AND (rainfall_datetime BETWEEN $1 AND $2) AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true')	"

	sqlAdvRainBaselineAreaY48 = "SELECT month_id,sum(volume) FROM public.tr_avg_rainfall_y48 "

	sqlAdvRainBaselineAreaY30 = "SELECT month_id,sum(volume) FROM public.tr_avg_rainfall_y30 "

	sqlAdvRainYearlyGraph = "SELECT rainfall_datetime,rainfall_value FROM public.rainfall_yearly WHERE deleted_at=to_timestamp(0) AND tele_station_id=$1 AND (rainfall_datetime BETWEEN $2 AND $3) AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true')"

	sqlAdvRainNormal = "SELECT yt.volume, yf.volume FROM public.tr_regional_avg_rainfall_y30 yt LEFT JOIN public.tr_regional_avg_rainfall_y48 yf ON yt.reg_id=yf.reg_id WHERE yt.reg_id=(SELECT ge.area_code FROM public.m_tele_station mts LEFT JOIN public.lt_geocode ge ON mts.geocode_id=ge.id WHERE mts.deleted_at=to_timestamp(0) AND mts.id=$1)"
)
