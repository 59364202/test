// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"strconv"
	"time"
	"fmt"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	//	"time"
)

// ------------------------------ insert ------------------------------
var SQL_InserOrderDetail = `INSERT INTO dataservice.order_detail (metadata_id, order_header_id, detail_status_id, service_id, detail_frequency, detail_fromdate, detail_todate, 
	detail_remark, detail_province, detail_basin, e_id, detail_source_result, detail_source_result_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
var SQL_InserOrderDetail_service1 = `INSERT INTO dataservice.order_detail (metadata_id, order_header_id, detail_status_id, service_id, detail_frequency, detail_remark, 
	detail_province, detail_basin,detail_fromdate, detail_todate, e_id, detail_source_result, detail_source_result_date) VALUES ($1, $2, $3, 1, $4, $5, $6, $7, 
	null, null, $8, $9, $10)`
var SQL_InserOrderDetail_service4 = `INSERT INTO dataservice.order_detail (metadata_id, order_header_id, detail_status_id, service_id, detail_frequency, detail_remark, 
	detail_province, detail_basin, e_id, detail_source_result, detail_source_result_date) VALUES ($1, $2, $3, 4, $4, $5, $6, $7, $8, $9, $10)`

// ------------------------------ select ------------------------------

// var SQL_SelectOrderDetail = `SELECT m.id, m.metadataservice_name, ods.id, ods.detail_status, ls.id, ls.servicemethod_name, od.detail_province, od.detail_basin,
// 	lc.id, lc.category_name, a.id , a.agency_name, od.detail_fromdate, od.detail_todate, od.detail_letterdate, od.detail_letterno, od.order_header_id, ld.id, ld.department_name,
// 	lm.id, lm.ministry_name, od.id, od.detail_source_result, od.detail_source_result_date, ua.agency_name, au.id, au.office_name, au.full_name, ulm.ministry_name, au.contact_phone,
// 	oh.order_purpose, od.e_id, od.is_enabled, clog.count, clog.latest_access_time ,
// 	(SELECT string_agg(datafrequency,',') FROM metadata_frequency WHERE id = ANY(string_to_array(od.detail_frequency,',','')::int[])) as detail_frequency
// 	FROM dataservice.order_detail od
// 	INNER JOIN metadata m ON od.metadata_id = m.id
// 	INNER JOIN dataservice.order_detail_status ods ON od.detail_status_id = ods.id
// 	INNER JOIN lt_servicemethod ls ON od.service_id = ls.id
// 	INNER JOIN lt_subcategory lsub ON m.subcategory_id = lsub.id
// 	INNER JOIN lt_category lc ON lsub.category_id = lc.id
// 	INNER JOIN agency a ON m.agency_id = a.id
// 	INNER JOIN lt_department ld ON a.department_id = ld.id
// 	INNER JOIN lt_ministry lm ON ld.ministry_id = lm.id
// 	INNER JOIN dataservice.order_header oh ON od.order_header_id = oh.id
// 	INNER JOIN api.user au ON oh.user_id = au.id
// 	LEFT JOIN agency ua ON au.agency_id = ua.id
// 	LEFT JOIN lt_department uld ON au.department_id = uld.id
// 	LEFT JOIN lt_ministry ulm ON uld.ministry_id = ulm.id
// 	LEFT JOIN (
//    		SELECT od.id, aal.request_params ->>'id' , count(aal.request_params ->>'id') , max(aal.access_time) as latest_access_time
//    		FROM dataservice.order_detail od
//    		LEFT JOIN  api.access_log aal ON aal.service_id = 107 AND (aal.request_params ->>'id')::bigint = od.id
//    		GROUP BY od.id, aal.request_params ->>'id'
// 	) clog ON clog.id = od.id `
var SQL_SelectOrderDetail = `SELECT m.id, m.metadataservice_name, ods.id, ods.detail_status, ls.id, ls.servicemethod_name, od.detail_province, od.detail_basin, 
	lc.id, lc.category_name, a.id , a.agency_name, od.detail_fromdate, od.detail_todate, od.detail_letterdate, od.detail_letterno, od.order_header_id, ld.id, ld.department_name, 
	lm.id, lm.ministry_name, od.id, od.detail_source_result, od.detail_source_result_date, ua.agency_name, au.id, au.office_name, au.full_name, ulm.ministry_name, au.contact_phone, 
	oh.order_purpose, od.e_id, od.is_enabled, clog.count, clog.latest_access_time ,
	(SELECT string_agg(datafrequency,',') FROM metadata_frequency WHERE id = ANY(string_to_array(od.detail_frequency,',','')::int[])) as detail_frequency
	FROM dataservice.order_detail od 
	INNER JOIN metadata m ON od.metadata_id = m.id 
	INNER JOIN dataservice.order_detail_status ods ON od.detail_status_id = ods.id 
	INNER JOIN lt_servicemethod ls ON od.service_id = ls.id 
	INNER JOIN lt_subcategory lsub ON m.subcategory_id = lsub.id 
	INNER JOIN lt_category lc ON lsub.category_id = lc.id 
	INNER JOIN agency a ON m.agency_id = a.id 
	INNER JOIN lt_department ld ON a.department_id = ld.id 
	INNER JOIN lt_ministry lm ON ld.ministry_id = lm.id 
	INNER JOIN dataservice.order_header oh ON od.order_header_id = oh.id 
	INNER JOIN api.user au ON oh.user_id = au.id 
	LEFT JOIN agency ua ON au.agency_id = ua.id 
	LEFT JOIN lt_department uld ON au.department_id = uld.id 
	LEFT JOIN lt_ministry ulm ON uld.ministry_id = ulm.id 
	INNER JOIN (
   		SELECT od.id, aal.request_params ->>'id' , count(aal.request_params ->>'id') , max(aal.access_time) as latest_access_time
   		FROM dataservice.order_detail od
		   LEFT JOIN  api.access_log aal ON aal.service_id = 107 AND (aal.request_params ->>'id')::bigint = od.id
		   --WHERE
   		GROUP BY od.id, aal.request_params ->>'id'
	) clog ON clog.id = od.id `
var SQL_SelectOrderDetail_OrderBy = " ORDER BY od.id DESC "

var SQL_SelectOrderDetailGroupByAgency = ` SELECT od.order_header_id , a.id ,a.agency_name, od.detail_letterno, od.detail_letterpath, 
	 ld.id,ld.department_name, lm.id, lm.ministry_name, od.detail_fromdate, od.detail_todate, m.metadataservice_name 
	 FROM  dataservice.order_detail od  
	 INNER JOIN dataservice.order_header oh ON od.order_header_id = oh.id 
	 INNER JOIN metadata m ON od.metadata_id = m.id  
	 INNER JOIN agency a ON m.agency_id = a.id 
	 INNER JOIN lt_department ld ON a.department_id = ld.id 
	 INNER JOIN lt_ministry lm ON ld.ministry_id = lm.id 
	 INNER JOIN ( 
		 SELECT od.order_header_id, m.agency_id 
		 FROM dataservice.order_detail od 
		 INNER JOIN metadata m ON od.metadata_id = m.id 
		 INNER JOIN agency a ON m.agency_id = a.id 
		 WHERE od.service_id = 3 
		 GROUP BY od.order_header_id, m.agency_id 
	 ) t ON t.order_header_id = od.order_header_id AND t.agency_id = a.id 
	 WHERE od.service_id = 3 
	     AND od.detail_letterno IS NOT NULL 
	     AND od.detail_source_result_date IS NULL
	 ORDER BY od.order_header_id , m.id `

var SQL_SelectOrderDetailSummary = "SELECT od.id, od.order_header_id, m.metadataservice_name, u.full_name, ls.servicemethod_name, " +
	" oh.order_datetime, od.detail_source_result_date, od.detail_source_result, u.id , ag.agency_name " +
	" FROM  dataservice.order_detail od " +
	" INNER JOIN metadata m ON od.metadata_id = m.id " +
	" INNER JOIN dataservice.order_header oh ON od.order_header_id = oh.id " +
	" INNER JOIN api.user u ON oh.user_id = u.id " +
	" LEFT JOIN agency ag ON u.agency_id = ag.id " +
	" INNER JOIN lt_servicemethod ls ON od.service_id = ls.id "
var SQL_SelectOrderDetailSummary_OrderBy = " ORDER BY od.id DESC , od.order_header_id DESC "

var SQL_SelectCountOrderDetailByAgencyId = `SELECT a.id, a.agency_name, c.count_id
FROM agency a
INNER JOIN
  (SELECT COUNT(od.id) AS count_id, u.agency_id
   FROM dataservice.order_detail od
   INNER JOIN dataservice.order_header oh ON od.order_header_id = oh.id
   INNER JOIN metadata m ON od.metadata_id = m.id
   INNER JOIN api.user u ON u.id = oh.user_id
   WHERE m.agency_id = $1
   GROUP BY u.agency_id) c ON a.id = c.agency_id`

var SQL_SelectMetadata_ = `
SELECT     od.service_id, 
           od.detail_fromdate, 
           od.detail_todate, 
           m.id, 
           m.agency_id, 
           m.connection_format, 
           dd.import_setting#>'{configs,0,imports,0}'->>'destination' as table_name, 
           dd.convert_setting#>>'{configs,0,fields}', 
           CASE 
                      WHEN dd.import_table->> 'tables' != '{}' THEN dd.import_table->'tables'-> json_object_keys(dd.import_table-> 'tables') 
                      ELSE NULL 
           END AS import_f, 
           od.id, 
           od.detail_province, 
           od.detail_basin, 
           od.is_enabled, 
           od.created_at,
           ( 
                  SELECT string_agg(datafrequency,',') 
                  FROM   metadata_frequency 
                  WHERE  id = ANY(string_to_array(od.detail_frequency,',','')::int[])) AS detail_frequency,
		   dd.id,
		   m.additional_dataset
FROM       metadata m 
LEFT JOIN  dataservice.order_detail od 
ON         od.metadata_id = m.id 
INNER JOIN api.dataimport_dataset dd 
ON         m.dataimport_dataset_id = dd.id
`

var SQL_SelectMetadata = SQL_SelectMetadata_ + " WHERE od.id = $1 "
var SQL_SelectMetadata_FromEId = SQL_SelectMetadata_ + " WHERE od.e_id = $1 "
var SQL_SelectMetadata_FromMetadata = SQL_SelectMetadata_ + " WHERE m.id = $1 ORDER BY od.id DESC LIMIT 1"

//	genarate sql string สำหรับ data_service
//	Parameters:
//		service
//			รหัสประเภทบริการข้อมูล 1, 2, 3, 4
//		selectField
//			select field ดึงมาจาก import_setting
//		table_name
//			ชื่อตาราง
//		agency_id
//			รหัสหน่วยงาน
//		media_type_id
//			รหัสประเภทข้อมูลสื่อ
//		province
//			รหัสจังหวัด
//		basin
//			รหัสลุ่มน้ำ
//		fromdate
//			ข้อมูลตั้งแต่วันที่..
//		todate
//			ข้อมูลจนถึงวันที่
//	Retun:
//		sql string, parameter
func SQL_GenSQLSelectDataservice_All_bk(service, selectField, table_name, agency_id, media_type_id, province, basin, fromdate, todate, datasetId string, hasQC bool) (string, []interface{}) {
	var q string = ""
	itf := make([]interface{}, 0)
	m_id := model_metadata.GetColumnMasterId(table_name)
	st := model_metadata.GetTable(table_name)

	if st.MasterTable == "" { // ไม่มี master table ไม่ต้อง join
		//		q = "SELECT " + selectField + " FROM " + table_name + " tt "

		if st.HasProvince {
			q += " LEFT JOIN lt_geocode lg ON tt.geocode_id = lg.id "
		}
		if st.HasBasin {
			q += " LEFT JOIN subbasin s ON tt.subbasin_id = s.id "
		}
		if table_name == "agency" {
			q += " WHERE tt.id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
		} else {
			q += " LEFT JOIN api.dataimport_dataset_log ddl ON tt.dataimport_log_id = ddl.id " // ต้อง join dataimport_dataset_log เพื่อกรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
			q += " WHERE tt.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
			q += " AND (ddl.dataimport_dataset_id = " + datasetId + " OR tt.dataimport_log_id IS NULL) " // ต้อง กรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
		}

		if !st.IsMaster && st.PartitionField != "" { // ไม่เป็น master table
			switch service {
			case "1": // --datetime ย้อนหลัง 7 วัน
				// เป็นข้อมูลย้อนหลัง
				// เปลี่ยนเป็นเหมือน cd/dvd เผื่อให้รองรับการใส่ form_date, to_date จากการให้บริการข้อมูล
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				//				fd := time.Now().AddDate(0, 0, -8)
				//				td := time.Now().AddDate(0, 0, -1)
				//				itf = append(itf, fd.Format("2006-01-02"))
				//				itf = append(itf, td.Format("2006-01-02")+" 23:59 ")
				itf = append(itf, fromdate)
				itf = append(itf, todate+" 23:59 ")
				break
			case "4": // ล่าสุด หา dataimport_log_id ล่าสุดจาก table_name
				//				q += " AND tt.dataimport_log_id = ( SELECT max(dataimport_log_id) FROM " + table_name +
				//					" WHERE agency_id = '" + agency_id + "' AND deleted_at = to_timestamp(0) "
				//				if media_type_id != "" { // ถ้ามี media_type_id ต้องเอามาหาด้วย
				//					q += " AND media_type_id IN (" + media_type_id + ")"
				//				}
				//				q += " )"
				// ให้ไปดึงจาก latest
				if !IsMedia(table_name) {
					table_name = "latest." + table_name
				}
				break
			default: // download, cd/dvd
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				itf = append(itf, fromdate)
				itf = append(itf, todate+" 23:59 ")
				break
			}
		}

	} else { // มีตัว master table ต้อง join master เพื่อเอา geocode, subbasin_id
		//		q = "SELECT " + selectField + " FROM " + st.MasterTable + " m " +
		q = " INNER JOIN " + st.MasterTable + " m ON m.id = tt." + m_id
		q += " LEFT JOIN api.dataimport_dataset_log ddl ON tt.dataimport_log_id = ddl.id " // join dataimport_dataset_log เพื่อกรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
		if st.HasProvince {
			q += " LEFT JOIN lt_geocode lg ON m.geocode_id = lg.id "
		}
		if st.HasBasin {
			q += " LEFT JOIN subbasin s ON m.subbasin_id = s.id "
		}
		q += " WHERE m.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
		q += " AND (ddl.dataimport_dataset_id = " + datasetId + " OR tt.dataimport_log_id IS NULL) " // กรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
		if st.PartitionField != "" {
			switch service {
			case "1": // --datetime ย้อนหลัง 7 วัน
				// เป็นข้อมูลย้อนหลัง
				// เปลี่ยนเป็นเหมือน cd/dvd เผื่อให้รองรับการใส่ form_date, to_date จากการให้บริการข้อมูล
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				//				fd := time.Now().AddDate(0, 0, -8)
				//				td := time.Now().AddDate(0, 0, -1)
				//				itf = append(itf, fd.Format("2006-01-02"))
				//				itf = append(itf, td.Format("2006-01-02")+" 23:59 ")
				itf = append(itf, fromdate)
				itf = append(itf, todate+" 23:59 ")
				break
			case "4": // ล่าสุด หา dataimport_log_id ล่าสุดจาก table_name
				//			q += " AND tt.dataimport_log_id = ( SELECT max(tt.dataimport_log_id) FROM " + st.MasterTable + " m " +
				//				" INNER JOIN " + table_name + " tt ON  m.id = tt." + m_id + " WHERE m.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) )"
				// ให้ไปดึงจาก latest
				if !IsMedia(table_name) {
					table_name = "latest." + table_name
				}
				break
			default: // download, cd/dvd
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				itf = append(itf, fromdate)
				itf = append(itf, todate+" 23:59 ")
				break
			}
			//			if model_metadata.IsHAII(agency_id) && st.WhereHAII != "" {
			//				q += " AND " + st.WhereHAII
			//			}
		}

	}
	if media_type_id != "" {
		q += " AND tt.media_type_id IN (" + media_type_id + ") "
	}
	if province != "" {
		q += " AND lg.province_code::integer IN (" + province + ") "
	}
	if basin != "" {
		q += " AND s.basin_id IN (" + basin + ") "
	}

	// เพิ่มเงื่อนไข query ข้อมูล value <> 999999
	if st.Where != "" {
		q += " AND " + st.Where + " "
	}
	// เพิ่มเงื่อนไข query ข้อมูลถ้าใน dataset มี qc_status
	if hasQC {
		q += " AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') "
	}

	if st.PartitionField != "" {
		q += " ORDER BY  tt." + st.PartitionField + " DESC"
	}
	if st.IsMaster { // เป็น master table เรียงตาม id
		q += " ORDER BY tt.id "
	}
	q = "SELECT " + selectField + " FROM " + table_name + " tt " + q
	return q, itf
}

//	genarate sql string สำหรับ data_service
//	Parameters:
//		p
//			Strct_Data
//	Retun:
//		sql string, parameter
func SQL_GenSQLSelectDataservice_All(p *Strct_Data) (string, []interface{}) {
	var q string = ""
	itf := make([]interface{}, 0)
	var (
		table_name = p.Table_name.String
		agency_id  = p.Agency_id.String
		// datasetId         = datatype.MakeString(p.Dataset_id)
		service       = p.Service_id
		fromdate      = p.Detail_fromdate.String
		todate        = p.Detail_todate.String + " 23:59 "
		media_type_id = p.MediaTypeId
		province      = p.Province.String
		basin         = p.Basin.String
		hasQC         = p.HasQC
		selectField   = p.SelectFields
		metadata_id	= p.Metadata_id
		// additionalDataset = p.AdditionalDataset
	)
	
	m_id := model_metadata.GetColumnMasterId(table_name)
	
	st := model_metadata.GetTable(table_name)

	if st.MasterTable == "" { // ไม่มี master table ไม่ต้อง join
		//		q = "SELECT " + selectField + " FROM " + table_name + " tt "

		if st.HasProvince {
			q += " LEFT JOIN lt_geocode lg ON tt.geocode_id = lg.id "
		}
		if st.HasBasin {
			q += " LEFT JOIN subbasin s ON tt.subbasin_id = s.id "
		}
		if table_name == "agency" {
			q += " WHERE tt.id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) AND tt.deleted_at = to_timestamp(0)"
		}

		if !st.IsMaster && st.PartitionField != "" { // ไม่เป็น master table
			if table_name != "agency" {
				q += " LEFT JOIN api.dataimport_dataset_log ddl ON tt.dataimport_log_id = ddl.id "    // ต้อง join dataimport_dataset_log เพื่อกรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
				q += " WHERE tt.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) " // เอาข้อมูลเฉพาะในหน่วยงานเดียวกันเท่านั้น
				// q += " AND (ddl.dataimport_dataset_id = " + datasetId + " OR tt.dataimport_log_id IS NULL " // ต้อง กรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
				// if additionalDataset.Valid {
				// 	q += " OR ddl.dataimport_dataset_id IN (" + additionalDataset.String + ") "
				// }
				// q += ")"
			}

			switch service {
			case 1: // --datetime ย้อนหลัง 7 วัน
				// เป็นข้อมูลย้อนหลัง
				// เปลี่ยนเป็นเหมือน cd/dvd เผื่อให้รองรับการใส่ form_date, to_date จากการให้บริการข้อมูล
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				//				fd := time.Now().AddDate(0, 0, -8)
				//				td := time.Now().AddDate(0, 0, -1)
				//				itf = append(itf, fd.Format("2006-01-02"))
				//				itf = append(itf, td.Format("2006-01-02")+" 23:59 ")
				itf = append(itf, fromdate)
				itf = append(itf, todate)
				break
			case 4: // ล่าสุด หา dataimport_log_id ล่าสุดจาก table_name
				//				q += " AND tt.dataimport_log_id = ( SELECT max(dataimport_log_id) FROM " + table_name +
				//					" WHERE agency_id = '" + agency_id + "' AND deleted_at = to_timestamp(0) "
				//				if media_type_id != "" { // ถ้ามี media_type_id ต้องเอามาหาด้วย
				//					q += " AND media_type_id IN (" + media_type_id + ")"
				//				}
				//				q += " )"
				// ข้อมูลล่าสุด ให้ไปดึงจาก latest
				if !IsMedia(table_name) {
					table_name = "latest." + table_name
				}
				break
			default: // download, cd/dvd
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				itf = append(itf, fromdate)
				itf = append(itf, todate)
				break
			}
		} else if st.IsMaster {
			q += " WHERE tt.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
		}

	} else { // มีตัว master table ต้อง join master เพื่อเอา geocode, subbasin_id
		//		q = "SELECT " + selectField + " FROM " + st.MasterTable + " m " +
		q = " INNER JOIN " + st.MasterTable + " m ON m.id = tt." + m_id
		q += " LEFT JOIN api.dataimport_dataset_log ddl ON tt.dataimport_log_id = ddl.id " // join dataimport_dataset_log เพื่อกรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
		if st.HasProvince {
			q += " LEFT JOIN lt_geocode lg ON m.geocode_id = lg.id "
		}
		if st.HasBasin {
			q += " LEFT JOIN subbasin s ON m.subbasin_id = s.id "
		}
		q += " WHERE m.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) "
		// q += " AND (ddl.dataimport_dataset_id = " + datasetId + " OR tt.dataimport_log_id IS NULL " // กรองเอาเฉพาะข้อมูลที่อยู่ใน dataset เดียวกันเท่านั้น
		// if additionalDataset.Valid {
		// 	q += " OR ddl.dataimport_dataset_id IN (" + additionalDataset.String + ") "
		// }
		// q += ")"
		if st.PartitionField != "" {
			switch service {
			case 1: // --datetime ย้อนหลัง 7 วัน
				// เป็นข้อมูลย้อนหลัง
				// เปลี่ยนเป็นเหมือน cd/dvd เผื่อให้รองรับการใส่ form_date, to_date จากการให้บริการข้อมูล
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				//				fd := time.Now().AddDate(0, 0, -8)
				//				td := time.Now().AddDate(0, 0, -1)
				//				itf = append(itf, fd.Format("2006-01-02"))
				//				itf = append(itf, td.Format("2006-01-02")+" 23:59 ")
				itf = append(itf, fromdate)
				itf = append(itf, todate)
				break
			case 4: // ล่าสุด หา dataimport_log_id ล่าสุดจาก table_name
				//			q += " AND tt.dataimport_log_id = ( SELECT max(tt.dataimport_log_id) FROM " + st.MasterTable + " m " +
				//				" INNER JOIN " + table_name + " tt ON  m.id = tt." + m_id + " WHERE m.agency_id = '" + agency_id + "' AND tt.deleted_at = to_timestamp(0) )"
				// ให้ไปดึงจาก latest
				if !IsMedia(table_name) {
					table_name = "latest." + table_name
				}
				break
			default: // download, cd/dvd
				q += " AND tt." + st.PartitionField + " BETWEEN $1 AND $2 "
				itf = append(itf, fromdate)
				itf = append(itf, todate)
				break
			}
			//			if model_metadata.IsHAII(agency_id) && st.WhereHAII != "" {
			//				q += " AND " + st.WhereHAII
			//			}
		}

	}
	if media_type_id != "" {
		q += " AND tt.media_type_id IN (" + media_type_id + ") "
	}
	if province != "" {
		q += " AND nullif(lg.province_code,'')::integer IN (" + province + ") "
	}
	if basin != "" {
		q += " AND s.basin_id IN (" + basin + ") "
	}
	// เพิ่มเงื่อนไข query ข้อมูล value <> 999999
	if st.Where != "" {
		q += " AND " + st.Where + " "
	}
	if st.WhereHAII != "" {
		q += " AND " + st.WhereHAII + " "
	}
	// เพิ่มเงื่อนไข  query station ที่เเป็น hydro 1-8 
	if (metadata_id == 550 || metadata_id == 228) && st.WhereHydro != "" {
		q += " AND " + st.WhereHydro + " "
	}
	// เพิ่มเงื่อนไข query ข้อมูลถ้าใน dataset มี qc_status
	if hasQC {
		q += " AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') "
	}
	if st.PartitionField != "" {
		q += " ORDER BY  tt." + st.PartitionField + " DESC"
	}
	if st.IsMaster { // เป็น master table เรียงตาม id
		q += " ORDER BY tt.id "
	}
	q = "SELECT " + selectField + " FROM " + table_name + " tt " + q
	
	fmt.Println(q)
	
	return q, itf
}

var SQL_CountNoResult = `
SELECT count(od.id) 
FROM dataservice.order_detail od 
INNER JOIN dataservice.order_header oh ON oh.id = od.order_header_id
INNER JOIN (
SELECT id
FROM   dataservice.order_header 
WHERE  order_forexternal = true 
       AND id IN (SELECT DISTINCT( order_header_id ) 
                  FROM   dataservice.order_detail od 
                         INNER JOIN dataservice.order_header oh 
                                 ON od.order_header_id = oh.id 
                                    AND oh.order_forexternal = true 
                  WHERE  od.detail_letterpath IS NULL) 
       AND order_status_id IN (1,2)
) oh2 ON oh2.id = oh.id AND oh2.id = od.order_header_id
WHERE od.detail_source_result is NULL AND oh.order_status_id <> 3
`

//	genarate sql string สำหรับนับจำนวน order_detail ที่เป็น service download ที่ยังเปิดใช้งานอยู่และ ไฟล์ ยังไม่หมดอายุ
func Gen_SQL_CountOrderDetailDownloadOnlyEnable(date string) string {
	_date, _ := strconv.Atoi(date)
	now := time.Now().AddDate(0, 0, -1*_date)
	return `
SELECT count(od.id)
FROM dataservice.order_detail od
INNER JOIN dataservice.order_header oh ON oh.id = od.order_header_id
WHERE od.service_id = 2 AND od.is_enabled = true AND od.created_at > ` + now.Format("2006-01-02 15:04")
	// WHERE od.service_id = 2 AND od.is_enabled = true AND od.created_at > NOW() - interval '` + date + ` day '
}

var SQL_CountOrderDetailEnableWithoutDownload = `
SELECT count(od.id)
FROM dataservice.order_detail od
INNER JOIN dataservice.order_header oh ON oh.id = od.order_header_id
WHERE od.service_id <> 2 AND od.service_id <>3 AND od.is_enabled = true
`

// ------------------------------ update ------------------------------
var SQL_UpdateOrderDetail = "UPDATE dataservice.order_detail od "
var SQL_UpdateOrderDetail_From = " FROM public.metadata as m " +
	" WHERE od.metadata_id = m.id AND od.service_id = 3 AND od.order_header_id = $1 AND agency_id = $2 "

var SQL_UpdateOrderLetterno = SQL_UpdateOrderDetail +
	" SET detail_letterno = $3 , detail_letterdate = $4 , updated_by = $5 , detail_status_id = 1 " +
	SQL_UpdateOrderDetail_From

var SQL_UpdateOrderLetterPath = SQL_UpdateOrderDetail +
	" SET detail_letterpath = $3 , updated_by = $4 , detail_status_id = 2 " +
	SQL_UpdateOrderDetail_From

var SQL_UpdateOrderSourceResult = " UPDATE dataservice.order_detail SET detail_source_result_date = NOW() , detail_source_result = $2, updated_by = $3, detail_status_id = $4 WHERE id = $1 RETURNING order_header_id; "
var SQL_UpdateOrderHeaderStatusAfterSourceResult = ``
var SQL_UpdateOrderHeaderStatus = "UPDATE dataservice.order_header SET order_status_id = $3 , updated_at = NOW() , updated_by = $1 WHERE id = $2"

var SQL_UpdateEId = "UPDATE dataservice.order_detail od SET e_id = $1 WHERE id = $2"

var SQL_UpdateIsEnable = " UPDATE dataservice.order_detail SET is_enabled = $1 WHERE id = $2"

// ------------------------------ delete ------------------------------
var SQL_DeleteOrderDetailByOrderHeaderId = "DELETE FROM dataservice.order_detail WHERE order_header_id = $1"
var SQL_UpdateDetailToCancelByOrderHeaderId = " UPDATE dataservice.order_detail SET detail_source_result_date = NOW(), detail_source_result = 'DA' WHERE order_header_id = $1 "
