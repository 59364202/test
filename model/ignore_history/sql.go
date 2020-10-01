package ignore_history

var sqlSelectHistory = ` SELECT h.id
						  , ignore_datetime
						  , data_category
						  , station_id 
						  , station_oldcode
						  , station_name
						  , station_province
						  , agency_shortname
						  , '' AS agency_name
						  , data_id
						  , data_datetime
						  , remark
						  , u.id AS user_id
						  , u.full_name
						  , h.data_value
						FROM public.ignore_history h
						LEFT JOIN api.user u ON h.created_by = u.id `

var sqlSelectHistory_OrderBy = ` ORDER BY h.ignore_datetime DESC `

var sqlInsertHistory = ` INSERT INTO ignore_history (ignore_datetime, data_category
													, station_id, station_oldcode, station_name, station_province
													, agency_shortname, data_id
													, data_datetime, remark
													, created_by, updated_by, created_at, updated_at, data_value) `

// var sqlSelectIgnoreData = ` SELECT ih.id
// 								, ih.ignore_datetime
// 								, ih.data_category
// 								, ih.station_id
// 								, ih.station_oldcode
// 								, ih.station_name
// 								, ih.station_province
// 								, ih.agency_shortname
// 								, '' AS agency_name
// 								, ih.data_id
// 								, ih.data_datetime
// 								, ih.remark
// 								, 0 AS user_id
// 								, '' AS full_name
// 								, ih.data_value
// 							FROM ignore_history ih
// 							LEFT JOIN (SELECT id, rank() OVER (PARTITION BY station_id, data_category ORDER BY ignore_datetime DESC) AS pos FROM ignore_history) ih2 ON ih.id = ih2.id
// 							WHERE ih2.pos = 1
// 							AND ih.remark = 'Ignore station' `
var sqlSelectIgnoreData = ` SELECT ih.id
								, ih.ignore_datetime
								, ih.data_category
								, ih.station_id
								, ih.station_oldcode
								, ih.station_name
								, ih.station_province
								, ih.agency_shortname
								, '' AS agency_name
								, ih.data_id
								, ih.data_datetime
								, ih.remark
								, 0 AS user_id
								, '' AS full_name
								, ih.data_value
							FROM ignore ih
							WHERE is_ignore = true `
