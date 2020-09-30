package media

import ()

var sqlGetMedia = ` SELECT m.id
						, m.media_datetime
						, m.media_path
						, m.media_desc
						, m.filename
						, m.refer_source
						, m.media_type_id
						, mt.media_type_name
						, mt.media_subtype_name
						, m.agency_id
						, agt.agency_shortname
						, agt.agency_name
					FROM media m
					LEFT JOIN lt_media_type mt ON mt.id = m.media_type_id
					LEFT JOIN agency agt ON agt.id = m.agency_id
					WHERE m.deleted_at = '1970-01-01 07:00:00+07'
					  AND mt.deleted_at = '1970-01-01 07:00:00+07'
					  AND agt.deleted_at = '1970-01-01 07:00:00+07'
					  AND m.agency_id = $1
					  AND m.media_type_id = $2
					  AND m.media_datetime BETWEEN $3 AND $4 `

var sqlGetMediaOrderBy = ` ORDER BY m.media_datetime DESC, m.media_desc `

// var sqlMediaLatest = "SELECT DISTINCT ON (m.media_type_id,m.media_datetime) mt.media_type_name, mt.media_subtype_name, m.media_datetime, m.media_path, m.media_desc, m.filename " +
//	" FROM public.media m LEFT JOIN lt_media_type mt ON m.media_type_id=mt.id " +
//	" WHERE m.media_type_id = $1 AND m.deleted_at=to_timestamp(0) ORDER  BY m.media_datetime DESC NULLS LAST, m.media_type_id limit "

// add performance by query latest media from latest schema (by manorot, jan 2019)
var sqlMediaLatest = "SELECT m.media_type_id, mt.media_type_name, mt.media_subtype_name, mt.media_category, m.media_datetime, m.media_path, m.media_desc, m.filename " +
	" FROM latest.media m LEFT JOIN lt_media_type mt ON m.media_type_id=mt.id " +
	" WHERE m.media_type_id = $1 limit "
var sqlMediaLatestMonth = "SELECT DISTINCT ON (m.media_type_id,m.media_datetime) mt.media_type_name, mt.media_subtype_name, m.media_datetime, m.media_path, m.media_desc, m.filename " +
	" FROM public.media m LEFT JOIN lt_media_type mt ON m.media_type_id=mt.id " +
	" WHERE m.media_type_id = $1 AND m.deleted_at=to_timestamp(0) AND m.media_datetime BETWEEN $2 AND $3 ORDER  BY m.media_datetime DESC NULLS LAST, m.media_type_id limit "

var sqlMedia = "SELECT mt.media_type_name, mt.media_subtype_name, m.media_datetime, m.media_path, m.media_desc, m.filename " +
	" FROM public.media m LEFT JOIN lt_media_type mt ON m.media_type_id=mt.id WHERE (mt.deleted_at=to_timestamp(0) AND m.deleted_at=to_timestamp(0))"

var sqlMediaOther = "SELECT mt.media_type_name, mt.media_subtype_name, m.media_datetime, m.media_path, m.media_desc, m.filename " +
	" FROM public.media_other m LEFT JOIN lt_media_type mt ON m.media_type_id=mt.id WHERE (mt.deleted_at=to_timestamp(0) AND m.deleted_at=to_timestamp(0)) "

var sqlGetMediaHistory = "SELECT mt.id, mt.media_type_name, mt.media_subtype_name, m.media_datetime, m.media_path, m.media_desc, m.filename FROM public.media m LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id "

//var sqlGetMediaHistory2 = "SELECT mt.id, mt.media_type_name, mt.media_subtype_name, gs.date, m.media_path, m.media_desc, m.filename " +
//	"FROM (SELECT generate_series($1::timestamp,$2, '1 day') AS date) gs  " +
//	"LEFT JOIN public.media m ON gs.date=m.media_datetime AND m.media_type_id=$3 AND m.filename like '%0001.d02.jpg' AND m.deleted_at=to_timestamp(0) " +
//	"LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id AND mt.deleted_at=to_timestamp(0)"
var sqlGetMediaHistory2 = "SELECT mt.id, mt.media_type_name, mt.media_subtype_name, gs.date AS media_date, m.media_path, m.media_desc, m.filename " +
   "FROM (SELECT generate_series($1::timestamp,$2, '1 day') AS date) gs " +
   "LEFT JOIN public.media m ON gs.date=m.media_datetime AND m.media_type_id=$3 AND m.deleted_at=to_timestamp(0) " +
   "LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id AND mt.deleted_at=to_timestamp(0)"

// query top media other (pdf)
var sqlGetPdfHistory = "SELECT rank_filter.agency_id, rank_filter.media_type_id, rank_filter.media_path, rank_filter.filename, rank_filter.media_desc, rank_filter.media_datetime FROM (" +
	"SELECT *, rank() OVER " +
	"(PARTITION BY media_type_id ORDER BY media_datetime DESC) " +
	"FROM media_other" +
	"$where" +
") rank_filter"

func Gen_SQL_RadarHistory(radar_type, date, frequency string) (string, []interface{}) {
	var itf = []interface{}{}
	sql := `
SELECT gs, 
	m.media_path, 
    m.filename 
FROM   media m 
	INNER JOIN lt_media_type mt 
		ON m.media_type_id = mt.id 
		AND m.media_type_id = 30 
		AND m.deleted_at = To_timestamp(0) 
		AND m.filename LIKE $1
		AND m.media_datetime BETWEEN $2 AND $3
	RIGHT JOIN generate_series( $2::TIMESTAMP, $3, $4) gs 
		ON m.media_datetime = gs 
	ORDER  BY gs 
	`
	itf = append(itf, "%"+radar_type+"%")
	// itf = append(itf, date)
	itf = append(itf, date+" 00:00")
	itf = append(itf, date+" 23:59")
	itf = append(itf, frequency+" minute")

	return sql, itf
}
