package weather_image

var (
	sqlSelectMediaImageStorm = "SELECT m.media_datetime,a.agency_name,m.media_path,m.filename FROM public.media m LEFT JOIN public.agency a ON m.agency_id=a.id WHERE m.deleted_at=to_timestamp(0) AND m.media_type_id=$1"

	sqlSelectAgencyByID = "SELECT id, agency_name FROM public.agency WHERE id=$1 "

	sqlSelectMediaImageLatest = "SELECT DISTINCT ON (media_type_id, agency_id) media_datetime,a.agency_name,m.media_path,m.filename,refer_source FROM public.media m LEFT JOIN public.agency a ON m.agency_id=a.id " +
		"WHERE m.deleted_at=to_timestamp(0) AND m.agency_id=$2 AND m.media_type_id=$1 "

	sqlSelectMediaImageDatetime = "SELECT m.media_datetime,a.agency_name,m.media_path,m.filename,m.media_desc,mt.id, mt.media_type_name, mt.media_subtype_name,mt.media_category FROM public.media m LEFT JOIN public.agency a ON m.agency_id=a.id LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id WHERE m.deleted_at=to_timestamp(0) AND m.agency_id=$1 AND m.media_type_id=$2 AND m.media_datetime BETWEEN $3 AND $4"

	sqlSelectMediaHistoryDatetime = "SELECT m.media_datetime,a.agency_name,m.media_path,m.filename FROM public.media m LEFT JOIN public.agency a ON m.agency_id=a.id WHERE m.deleted_at=to_timestamp(0) AND m.agency_id=$1 AND m.media_type_id=$2 AND m.media_datetime BETWEEN $3 AND $4"

	sqlHistoryWeatherTMD = "SELECT g.datetime,a.agency_name,m.media_path,m.filename,m.media_desc,mt.id, mt.media_type_name, mt.media_subtype_name,mt.media_category " +
		"FROM (SELECT generate_series($3::timestamp,$4, '6 hour') AS datetime) g  " +
		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $4 AND m.media_datetime=g.datetime AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
		" LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id " +
		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL)"

	sqlHistoryWeatherDaily = "SELECT g.datetime,a.agency_name,m.media_path,m.filename,m.media_desc,mt.id, mt.media_type_name, mt.media_subtype_name,mt.media_category " +
		"FROM (SELECT generate_series($3::timestamp,$4, '12 hour') AS datetime) g  " +
		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $4 AND m.media_datetime=g.datetime AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
		" LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id " +
		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL)"

		// 2019-12-03 thitiporn
		//ย้านไปใส่ใน get.go เนื่องจากมีเงื่อนไขเพิ่มเติม
		//	sqlHistoryWeatherDate = "SELECT g.datetime,a.agency_name,m.media_path,m.filename,m.media_desc,mt.id, mt.media_type_name, mt.media_subtype_name,mt.media_category " +
		//		"FROM (SELECT generate_series($3::timestamp,$4, '1 hour') AS datetime) g  " +
		//		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $4 " +
		//		"AND m.media_datetime=g.datetime " +
		//		"AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
		//		" LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id " +
		//		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL)"

	sqlHistoryWeatherYear = "SELECT g.datetime,a.agency_name,m.media_path,m.filename " +
		"FROM (SELECT generate_series($3::timestamp,$4, '7 day') AS datetime) g  " +
		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $4 AND m.media_datetime=g.datetime AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL) AND g.datetime != $3"

	sqlHistoryWeatherYear1 = "SELECT g.datetime,a.agency_name,m.media_path,m.filename,m.media_desc,mt.id, mt.media_type_name, mt.media_subtype_name,mt.media_category " +
		"FROM (SELECT generate_series( (select min(datetime) from " +
		"generate_series($3::timestamp,$6, '1 day') AS datetime " +
		"where EXTRACT(DOW FROM datetime) = 0) ,$4, '7 day') AS datetime) g  " +
		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $6 AND m.media_datetime=g.datetime  " +
		"AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
		" LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id " +
		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL) " +
		"AND date_part('year', g.datetime) = $5"

	sqlSelectHistory7Day = "SELECT g.datetime,a.agency_name,m.media_path,m.filename " +
		"FROM (SELECT generate_series($3::timestamp,$4, '1 day') AS datetime) g   " +
		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $4 AND m.media_datetime=g.datetime AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL)"
)
