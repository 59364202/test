// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for public.metadata table. This table store metadata.
package metadata

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"

	"database/sql"
	//	"log"
)

//	หา data_min_date, data_max_date ของ metadata ทุกตัวที่อยู่ใน shopping
func ReCacheDataDateRange() error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	// shopping ทั้งหมด
	shoppingTable, err := GetMetadataShoppingTable(&Param_Metadata{})
	if err != nil {
		return err
	}

	strSqlUpdate := "UPDATE metadata SET data_min_date = $1, data_max_date = $2 WHERE id = $3"

	// วนลูปหา min, max
	for _, s := range shoppingTable {
		st := GetTable(s.Table)
		if st.IsMaster {
			// เป็น master table ไม่ต้องมี data date range
			continue
		}
		var (
			metadata_fromdate   sql.NullString
			metadata_todate     sql.NullString
			dataimportDatasetId sql.NullInt64
			additionalDataset   sql.NullString
		)
		strSql := "SELECT data_min_date, data_max_date, dataimport_dataset_id, additional_dataset FROM metadata WHERE id = $1"
		err = db.QueryRow(strSql, s.Id).Scan(&metadata_fromdate, &metadata_todate, &dataimportDatasetId, &additionalDataset)
		if err != nil {
			return errors.Repack(err)
		}

		log.Log("start find data min, max date from metadata id ", s.Id)
		err = findFromDateToDate(s, "", dataimportDatasetId.Int64, additionalDataset.String, "")
		if err != nil {
			log.Log(err)
			continue
		}
		if s.FormDate == "" && s.ToDate == "" {
			log.Log("metadata id ", s.Id, " no formdate todate")
			continue
		}

		// update metadata
		_, err := db.Exec(strSqlUpdate, s.FormDate, s.ToDate, s.Id)
		if err != nil {
			log.Log("error recache date range metadata id ", s.Id, " error ", err)
			return errors.Repack(err)
		}
	}
	return nil
}

//	หา  data_max_date ของ metadata ทุกตัวที่อยู่ใน shopping หาจาก last_check
func ReCacheDataDateRange_Max() error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	// shopping ทั้งหมด
	shoppingTable, _ := GetMetadataShoppingTable(&Param_Metadata{})

	strSqlUpdate := "UPDATE metadata SET data_max_date = $1, data_last_check = NOW() WHERE id = $2"

	// วนลูปหา min, max
	for _, s := range shoppingTable {
		st := GetTable(s.Table)
		if st.IsMaster {
			// เป็น master table ไม่ต้องมี data date range
			continue
		}
		var (
			metadata_todate     sql.NullString
			dataimportDatasetId sql.NullInt64
			additionalDataset   sql.NullString
		)
		strSql := "SELECT data_last_check, dataimport_dataset_id, additional_dataset  FROM metadata WHERE id = $1"
		err = db.QueryRow(strSql, s.Id).Scan(&metadata_todate, &dataimportDatasetId, &additionalDataset)
		if err != nil {
			return errors.Repack(err)
		}

		log.Log("start find data max date from metadata id ", s.Id)
		err = findFromDateToDate(s, metadata_todate.String, dataimportDatasetId.Int64, additionalDataset.String, "max")
		if err != nil {
			log.Log(err)
			continue
		}
		if s.ToDate == "" {
			log.Log("metadata id ", s.Id, " no todate")
			continue
		}
		// update metadata
		_, err := db.Exec(strSqlUpdate, s.ToDate, s.Id)
		if err != nil {
			log.Log("error recache date range metadata max : id ", s.Id, " error ", err)
			return errors.Repack(err)
		}
	}
	return nil
}
