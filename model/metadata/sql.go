package metadata

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
)

var SQL_selectMetadata_Head = " SELECT a.* , a.jsontable -> a.table->>'partition_field' FROM ( "
var SQL_selectMetadata_Foot = " ) a "

// lt_metadata_method id 2 = Web Extract
var SQL_selectMetadata = `SELECT m.id, m.metadataservice_name, m.metadata_description, m.metadata_convertfrequency, 
	ls.id, ls.subcategory_name, 
	a.id, a.agency_name, 
	import_table -> 'tables' as jsontable, 
	CASE WHEN dd.import_table ->> 'tables' IS NULL THEN NULL 
	END as table 
	FROM metadata m 
	INNER JOIN lt_dataformat ldf ON m.dataformat_id = ldf.id
	INNER JOIN lt_metadata_method lmm ON ldf.metadata_method_id = lmm.id AND lmm.id <> 2
	INNER JOIN lt_subcategory ls ON m.subcategory_id = ls.id 
	INNER JOIN lt_category lc ON ls.category_id = lc.id 
	INNER JOIN agency a ON m.agency_id = a.id 
	INNER JOIN lt_department ld ON a.department_id = ld.id 
	INNER JOIN lt_ministry lm ON ld.ministry_id = lm.id 
	INNER JOIN api.dataimport_dataset dd ON m.dataimport_dataset_id = dd.id 
	WHERE m.metadatastatus_id = 1  AND m.deleted_at = to_timestamp(0) `
var SQL_selectMetadata_Orderby = " ORDER BY m.id ASC "

var SQL_selectPartitionFieldFromMetadata = `
SELECT a.table, a.jsontable -> a.table->>'partition_field', a.agency_id, a.field
FROM
  (SELECT import_table -> 'tables' AS jsontable,
                          CASE
                              WHEN dd.import_table ->> 'tables' IS NULL THEN NULL
                              ELSE json_object_keys(dd.import_table -> 'tables')
                          END AS TABLE,
           m.agency_id,
           dd.convert_setting#>'{configs,0,fields}' as field
   FROM metadata m
   LEFT JOIN api.dataimport_dataset dd ON m.dataimport_dataset_id = dd.id
   WHERE m.id = $1 )a
`

//	สร้าง string query ที่ใช้ในการหา วันที่ มากสุด, น้อยสุด
//	Parameters:
//		table
//			ชื่อตาราง
//		agency_id
//			รหัสหน่วยงาน
//		convert_setting
//			convert_setting ในรูปแบบ string
//		metadata_last_check
//			วันที่เช็ค ล่าสุด
//		find
//			ค่าที่ต้องการหา {min,max,} ไม่ใส่ถือว่าหา min, max
func SQL_genSqlSelectMaxMinDateFromTable(table string, agency_id int64, convert_setting string, metadata_last_check, sDatasetId, additionalDataset, find string, _todate, _fromdate sql.NullString) (string, []interface{}) {
	itf := make([]interface{}, 0)
	st := GetTable(table)
	if st == nil {
		return "", nil
	}
	if st.IsMaster {
		return "", nil
	} else {
		if st.PartitionField == "" {
			return "", nil
		}
	}

	strSql := "SELECT t." + st.PartitionField + " FROM " + table + " t "
	strWhere := " WHERE t.deleted_at = '1970-01-01 07:00:00+07' "
	if st.MasterId != "" && st.MasterTable != "" { // มี master table
		strSql += " INNER JOIN " + st.MasterTable + " m ON t." + st.MasterId + " = m.id AND m.agency_id = $1"
		itf = append(itf, agency_id)
	} else {
		strWhere += " AND t.agency_id = $1"
		itf = append(itf, agency_id)
	}

	if !st.IsMaster { // ไม่ใช่ master table ให้ join หา dataset id
		strSql += " LEFT JOIN api.dataimport_dataset_log ddl ON t.dataimport_log_id = ddl.id "
		strWhere += " AND ( ddl.dataimport_dataset_id = " + sDatasetId + " OR t.dataimport_log_id IS NULL "
		if additionalDataset != "" {
			strWhere += " OR ddl.dataimport_dataset_id IN (" + additionalDataset + ") "
		}
		strWhere += ") "
	}

	if IsMedia(table) { // เป็น media ต้องมี media_type_id
		itf = append(itf, agency_id)
		strWhere += " AND t.agency_id = $" + strconv.Itoa(len(itf))
		var _itfFields []map[string]interface{}
		var _mapField map[string]interface{}
		var media_type_id string = ""
		json.Unmarshal([]byte(convert_setting), &_itfFields)
	L:
		for _, iv := range _itfFields {
			_mapField = iv
			if _mapField["name"].(string) == "media_type_id" {
				break L
			}
		}
		_media_type_id := _mapField["transform_params"]
		if str, ok := _media_type_id.(string); ok {
			media_type_id = str
		} else {
			media_type_id = strconv.Itoa(_media_type_id.(int))
		}
		if media_type_id != "" {
			itf = append(itf, media_type_id)
			strWhere += " AND t.media_type_id = $" + strconv.Itoa(len(itf))
		}
	}

	var findSql string
	switch find {
	case "max":
		findSql = "max(" + st.PartitionField + ")"
		if metadata_last_check != "" {
			itf = append(itf, metadata_last_check)
			strWhere += " AND t." + st.PartitionField + " > $" + strconv.Itoa(len(itf))
		}
	case "min":
		findSql = "min(" + st.PartitionField + ")"
		if metadata_last_check != "" {
			itf = append(itf, metadata_last_check)
			strWhere += " AND t." + st.PartitionField + " <= $" + strconv.Itoa(len(itf))
		}
	default:
		findSql = "min(" + st.PartitionField + "), max(" + st.PartitionField + ")"
		if _todate.Valid {
			itf = append(itf, _todate.String[0:10])
			strWhere += " AND t." + st.PartitionField + " > $" + strconv.Itoa(len(itf))
		}
		if _fromdate.Valid {
			itf = append(itf, _fromdate.String[0:10])
			strWhere += " AND t." + st.PartitionField + " <= $" + strconv.Itoa(len(itf))
		}
	}
	strSql = "SELECT " + findSql + " FROM (" + strSql + strWhere + ") a"

	return strSql, itf
}

var SQL_selectTableFromDatainportDataset = "SELECT json_object_keys(lookup_table -> 'tables'), m.agency_id FROM api.dataimport_dataset dd INNER JOIN metadata m on dd.id = m.dataimport_dataset_id WHERE m.id = $1"
var SQL_selectAllProvinceFromTable = " SELECT  province_code, province_name::text FROM lt_geocode WHERE id in ( select distinct(geocode_id) from "
var SQL_selectAllProvinceFromTable_end = " ) GROUP BY province_code , province_name::text ORDER BY province_code "

var SQL_selectSummaryMetadataByAgency = ` SELECT m.agency_id
												, agt.agency_shortname::jsonb AS agency_shortname
												, agt.agency_name::jsonb AS agency_name
												, sum(m.integrate_status) AS integrate_status
												, count(*) AS total
										   FROM ( SELECT metadata.id, metadata.agency_id, (CASE WHEN metadata.metadatastatus_id = 1 THEN 1 ELSE 0 END) AS integrate_status
												  FROM metadata
												  WHERE metadata.deleted_at = to_timestamp(0)) m
										   LEFT JOIN agency agt ON agt.id = m.agency_id
										   WHERE agt.deleted_at = to_timestamp(0)
										   GROUP BY m.agency_id, (agt.agency_name::jsonb), (agt.agency_shortname::jsonb)
										   ORDER BY agency_name::jsonb->>'th' `

var SQL_selectSummaryMetadataImportedByAgency = " SELECT category_id " +
	"		 , category_name " +
	"		 , COUNT(category_id) AS count_metadata " +
	"	FROM (" + sqlViewMetadataByAgencyCategory + ") aa " +
	"	WHERE agency_id = $1 " +
	"	GROUP BY category_id, category_name "
var SQL_selectSummaryMetadataImportedByAgency_OrderBy = ` ORDER BY category_name->>'th' `

var SQL_selectSummaryMetadataImportedByCategory = " SELECT agency_id " +
	"   , agency_shortname " +
	"   , agency_name " +
	"   , COUNT(agency_id) AS count_metadata " +
	" FROM (" + sqlViewMetadataByAgencyCategory + ") aa " +
	" WHERE category_id = $1 " +
	" GROUP BY agency_id, agency_name, agency_shortname "
var SQL_selectSummaryMetadataImportedByCategory_OrderBy = ` ORDER BY agency_name->>'th' `

var SQL_selectMetadataImportedByAgencyCategory = " SELECT metadata_id " +
	"	 , metadataservice_name " +
	"	 , metadataagency_name " +
	"	 , total_record " +
	"	 , last_import_date " +
	" FROM (" + sqlViewMetadataByAgencyCategory + ") aa " +
	" WHERE agency_id = $1 AND category_id = $2 "
var SQL_selectMetadataImportedByAgencyCategory_OrderBy = ` ORDER BY metadataservice_name->>'th' `

var SQL_selectMetadataByAgencyAndStatus = `
SELECT id, metadataservice_name, metadata_convertfrequency, connection_format, dataimport_download_id, dataimport_dataset_id FROM metadata 
WHERE metadatastatus_id = $2 AND agency_id = $1 AND deleted_at = to_timestamp(0)
ORDER BY id`

var SQL_SelectMetadataImportTableByAgency = `SELECT m.id, m.agency_id, m.connection_format , dd.import_setting#>'{configs,0,imports,0}'->>'destination', 
case 
   when dd.lookup_table->> 'tables' != '{}' then
   json_object_keys(dd.lookup_table-> 'tables')
   else null
end as lookup_table ,
dd.convert_setting#>>'{configs,0,fields}'
FROM metadata m
INNER JOIN api.dataimport_dataset dd ON m.dataimport_dataset_id = dd.id
WHERE m.id = $1 AND m.deleted_at = '1970-01-01 07:00:00+07' `

//	genarate sql string สำหรับดึงข้อมูลที่นำเข้าล่าสุด จาก dataimport_log_id
//	Parameters:
//		connect_format
//			ประเภทการเชื่อมต่อ online, offline
//		table_name
//			ชื่อตาราง
//		agency_id
//			รหัสหน่วยงาน
//		media_type_id
//			รหัสประเภทข้อมูลสื่อ
//	Return:
//		sql string
func SQL_GenSQLSelectLatestImport(connect_format, table_name, agency_id, media_type_id string) string {
	st := GetTable(table_name)
	m_id := st.MasterId
	q := " SELECT "
	selectField := " tt.* " // เอา ทุก column ของ เทเบิ้ล
	if st.HasProvince {     // มี จังหวัด ต้องเอา จังหวัด อำเภอ ตำบล มาด้วย
		selectField += " , province_code, province_name, amphoe_code, amphoe_name, tumbon_code, tumbon_name "
	}

	if st.MasterTable == "" { // ไม่มี master table ไม่ต้อง join
		q += selectField + " FROM " + table_name + " tt "
		if st.HasProvince { // มีฟิล geocode_id ให้จอย geocode
			q += " LEFT JOIN lt_geocode lg ON tt.geocode_id = lg.id "
		}
		if st.HasBasin { // มีฟิล subbasin_id ให้จอย subbasin เพื่อใช้ในการหา basin
			q += " LEFT JOIN subbasin s ON tt.subbasin_id = s.id "
		}
		q += " WHERE tt.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
		if strings.ToUpper(connect_format) == "ONLINE" { // เป็น online ต้องเอาที่ dataimport_log_id ล่าสุด
			q += " AND tt.dataimport_log_id = ( SELECT max(dataimport_log_id) FROM " + table_name +
				" WHERE agency_id = '" + agency_id + "' AND deleted_at = to_timestamp(0) "
			if media_type_id != "" { // ถ้ามี media_type_id ต้องเอามาหาด้วย
				q += " AND media_type_id = '" + media_type_id + "'"
			}
			q += " )"
		}

	} else { // มีตัว master table ต้อง join master เพื่อเอา geocode, subbasin_id
		mt := GetTable(st.MasterTable)
		if mt.SelectColumn != "" { // master table มี SelectColumn
			selectField += " , " + mt.SelectColumn
		}

		q += selectField + " FROM " + st.MasterTable + " m " +
			" INNER JOIN " + table_name + " tt ON m.id = tt." + m_id

		if st.HasProvince {
			q += " LEFT JOIN lt_geocode lg ON m.geocode_id = lg.id "
		}
		if st.HasBasin {
			q += " LEFT JOIN subbasin s ON m.subbasin_id = s.id "
		}
		q += " WHERE m.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
		if connect_format == "ONLINE" { // เป็น online ต้องเอาที่ dataimport_log_id ล่าสุด
			q += " AND tt.dataimport_log_id = ( SELECT max(tt.dataimport_log_id) FROM " + st.MasterTable + " m " +
				" INNER JOIN " + table_name + " tt ON  m.id = tt." + m_id + " WHERE m.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) )"
		}
		if IsHAII(agency_id) {
			q += " AND " + st.WhereHAII
		}
	}
	if media_type_id != "" {
		q += " AND tt.media_type_id = " + media_type_id
	}

	if strings.ToUpper(connect_format) == "OFFLINE" { // เป็น offline เอา 5 record ล่าสุด
		q += " ORDER BY tt.updated_at DESC LIMIT 5 "
	} else {
		q += " ORDER BY  tt." + st.PartitionField + " DESC"
	}

	return q
}

//func SQL_GenSQLSelectFromMetadata_Online(table_name, lookup_table, agency_id string) string {
//
//	var strSql string
//	var strSql_Where string
//
//	m_id := GetColumnMasterId(table_name)
//	m_dt := GetColumnDateTime(table_name)
//	st := GetTable(table_name)
//
//	if lookup_table == "" || lookup_table == "lt_geocode" {
//		strSql = "SELECT tt.*" + " FROM " + table_name + " tt "
//		strSql_Where = " WHERE tt.agency_id = " + agency_id + " AND tt.deleted_at = '1970-01-01 07:00:00+07' " +
//			" AND tt.created_at = (SELECT max(created_at) FROM " + table_name + " WHERE agency_id = " + agency_id + " AND deleted_at = '1970-01-01 07:00:00+07') "
//	} else {
//
//		strSql = "SELECT " + GetSelectColumn(lookup_table) + ",tt.* FROM " + lookup_table + " m " +
//			" LEFT JOIN " +
//			" ( SELECT " + m_id + " , max(" + m_dt + ") as d FROM " + table_name + " WHERE deleted_at = '1970-01-01 07:00:00+07' GROUP BY " + m_id + ") m ON m.id = m." + m_id +
//			" LEFT JOIN " + table_name + " tt ON m.id = tt." + m_id + " AND m.d = tt." + m_dt + " AND tt.deleted_at = '1970-01-01 07:00:00+07' " +
//			" LEFT JOIN lt_geocode lg ON m.geocode_id = lg.id "
//		if st.HasBasin {
//			strSql += " LEFT JOIN subbasin s ON m.subbasin_id = s.id "
//		}
//		strSql_Where = " WHERE m.agency_id = " + agency_id
//	}
//
//	return strSql + strSql_Where
//}

func SQL_GenSQLSelectFromMetadata_Offline(table_name, lookup_table, agency_id string) string {
	var strSql string
	var strSql_Where string
	if lookup_table == "" || lookup_table == "lt_geocode" || GetColumnMasterId(table_name) == "" {
		strSql = "SELECT tt.*" + " FROM " + table_name + " tt "
		strSql_Where = " WHERE tt.agency_id = " + agency_id + " AND tt.deleted_at = '1970-01-01 07:00:00+07' "

	} else {
		strSql = "SELECT tt.* FROM " + table_name + " tt INNER JOIN " + lookup_table + " m ON tt." + GetColumnMasterId(table_name) + " = m.id "
		strSql_Where = " WHERE m.agency_id = " + agency_id + " AND tt.deleted_at = '1970-01-01 07:00:00+07' "

	}

	return strSql + strSql_Where + " ORDER BY tt.updated_at DESC LIMIT 5"
}
func SQL_GenSelectLastMetadataImportMedia(table_name, agency_id, media_type_id string) string {
	str := "SELECT tt.media_path , tt.filename FROM " + table_name + " tt  WHERE tt.agency_id = " + agency_id + " and tt.deleted_at = '1970-01-01 07:00:00+07' " +
		" AND tt.media_type_id = " + media_type_id + " AND tt.media_datetime = " +
		"( SELECT max(media_datetime) FROM " + table_name + " WHERE agency_id = " + agency_id + " AND deleted_at = '1970-01-01 07:00:00+07' AND media_type_id = " + media_type_id + " )"
	return str
}

var SQL_SelectLastMetadataImportMedia = " SELECT tt.media_path , tt.filename FROM media tt  WHERE tt.agency_id = $1 and tt.deleted_at = '1970-01-01 07:00:00+07' " +
	" AND tt.media_type_id = $2 AND tt.media_datetime = " +
	"( SELECT max(media_datetime) FROM media WHERE agency_id = $1 AND deleted_at = '1970-01-01 07:00:00+07' AND media_type_id = $2 )"

var SQL_SelectMetadataStatusByAgency = " SELECT oh.id , u.full_name, m.metadataservice_name, ls.servicemethod_name, od.created_at " +
	", od.detail_source_result_date, od.detail_source_result " +
	" FROM dataservice.order_detail od " +
	" INNER JOIN metadata m ON od.metadata_id = m.id " +
	" INNER JOIN dataservice.order_header oh ON od.order_header_id = oh.id " +
	" INNER JOIN lt_servicemethod ls ON od.service_id = ls.id " +
	" INNER JOIN api.user u ON oh.user_id = u.id " +
	" WHERE u.agency_id = $1 AND m.agency_id = $2 " +
	" ORDER BY oh.id desc, od.id desc"

var sqlGetMetadata = `	SELECT m.id
							 , m.subcategory_id
							 , m.agency_id
							 , m.dataunit_id
							 , m.dataformat_id
							 , m.connection_format
							 , m.metadata_channel
							 , m.metadata_convertfrequency
							 , m.metadata_contact
							 , m.metadata_agencystoredate
							 , m.metadata_startdatadate
							 , m.metadata_update_plan
							 , m.metadata_laws
							 , m.metadata_remark
							 , m.metadataagency_name
							 , m.metadataservice_name
							 , m.metadata_tag
							 , m.metadata_description
							 , m.metadatastatus_id
							 , df.metadata_method_id
							 , sc.category_id
							 , m.metadata_receive_date
						FROM metadata m
						LEFT JOIN lt_subcategory sc ON m.subcategory_id = sc.id
						LEFT JOIN lt_dataformat df ON m.dataformat_id = df.id
						WHERE m.deleted_at = to_timestamp(0) `

//var sqlGetMetadataGroupby = " GROUP BY m.id "
var sqlGetMetadataGroupby = ""
var sqlGetMetadataTable = `  SELECT m.id, m.metadataservice_name , m.metadataagency_name
								  , agt.id AS agency_id, agt.agency_shortname, agt.agency_name
								  , sc.id AS subcategory_id, sc.subcategory_name
								  , c.id AS category_id, c.category_name
								  , string_agg(mh.hydroinfo_id || '##' || CAST(h.hydroinfo_name AS TEXT), '|') AS hydroinfo
							 FROM metadata m
							 LEFT JOIN lt_subcategory sc ON m.subcategory_id = sc.id
							 LEFT JOIN lt_category c ON sc.category_id = c.id
							 LEFT JOIN agency agt ON m.agency_id = agt.id
							 LEFT JOIN metadata_hydroinfo mh ON m.id = mh.metadata_id
							 LEFT JOIN lt_hydroinfo h ON mh.hydroinfo_id = h.id
							 WHERE m.deleted_at = to_timestamp(0) AND agt.deleted_at = to_timestamp(0) AND sc.deleted_at = to_timestamp(0) AND c.deleted_at = to_timestamp(0) AND (h.deleted_at IS NULL OR h.deleted_at = to_timestamp(0)) `

var sqlGetMetadataTableGroupBy = ` GROUP BY m.id, agt.id, sc.id, c.id `

var sqlUpdateMetadata = ` UPDATE metadata
						   SET subcategory_id = $2
							   , agency_id = $3
							   , dataunit_id = $4
							   , dataformat_id = $5
							   , connection_format = $6
							   , metadata_contact = $7
							   , metadata_agencystoredate = $8
							   , metadata_startdatadate = $9
							   , metadata_update_plan = $10
							   , metadata_laws = $11
							   , metadata_remark = $12
							   , metadataagency_name = $13
							   , metadataservice_name = $14
							   , metadata_tag = $15
							   , metadata_description = $16
							   , metadatastatus_id = $17
							   , updated_by = $18
							   , updated_at = NOW()
							   , metadata_convertfrequency = $19
							   , import_count = $20
							   , metadata_receive_date = $21
							WHERE id = $1 `

var sqlUpdateToDeleteMetadata = ` UPDATE metadata
								   SET deleted_by = $2
									 , updated_by = $2
									 , deleted_at = NOW()
									 , updated_at = NOW()
								   WHERE id = $1 AND deleted_by IS NULL `

var sqlInsertMetadata = ` INSERT INTO metadata ( subcategory_id
												   , agency_id
												   , dataunit_id
												   , dataformat_id
												   , connection_format
												   , metadata_contact
												   , metadata_agencystoredate
												   , metadata_startdatadate
												   , metadata_update_plan
												   , metadata_laws
												   , metadata_remark
												   , metadataagency_name
												   , metadataservice_name
												   , metadata_tag
												   , metadata_description
												   , metadatastatus_id
												   , created_by
												   , updated_by
												   , metadata_convertfrequency
												   , import_count
												   , created_at
												   , updated_at
												   , metadata_receive_date
												   )
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $17, $18, $19, NOW(), NOW(),$20) RETURNING id `

var sqlUpdateMetadataOfflineDate = "UPDATE public.metadata SET updated_by=$1, updated_at=NOW(), metadata_offline_date=NOW() WHERE id=$2"
var sqlViewMetadataByAgencyCategory = ` SELECT  m.id AS metadata_id,
                                                c.id AS category_id,
                                                m.agency_id,
                                                c.category_name::jsonb AS category_name,
                                                agt.agency_shortname::jsonb AS agency_shortname,
                                                agt.agency_name::jsonb AS agency_name,
                                                m.metadataservice_name::jsonb AS metadataservice_name,
                                                m.metadataagency_name::jsonb AS metadataagency_name,
                                                sum(ddl.total_record) AS total_record,
                                                ddl.last_import_date
                                        FROM metadata m
                                        LEFT JOIN agency agt ON m.agency_id = agt.id
                                        LEFT JOIN lt_subcategory sc ON m.subcategory_id = sc.id
                                        LEFT JOIN lt_category c ON sc.category_id = c.id
                                        LEFT JOIN api.dataimport_dataset dd ON m.dataimport_dataset_id = dd.id
                                        LEFT JOIN ( SELECT  dataimport_dataset_log.dataimport_dataset_id,
                                                            max(dataimport_dataset_log.updated_at) AS last_import_date,
                                                            sum(dataimport_dataset_log.import_success_row) AS total_record
                                                    FROM api.dataimport_dataset_log
                                                    WHERE dataimport_dataset_log.import_success_row <> 0 AND dataimport_dataset_log.deleted_at IS NULL
                                                    GROUP BY dataimport_dataset_log.dataimport_dataset_id) ddl ON dd.id = ddl.dataimport_dataset_id
                                        WHERE m.metadatastatus_id = 1 AND m.deleted_at = to_timestamp(0) AND sc.deleted_at = to_timestamp(0) AND c.deleted_at = to_timestamp(0) AND dd.deleted_at IS NULL
                                        GROUP BY m.id, c.id, m.agency_id, (c.category_name::jsonb), (agt.agency_shortname::jsonb), (agt.agency_name::jsonb), (m.metadataservice_name::jsonb), (m.metadataagency_name::jsonb), ddl.last_import_date `
