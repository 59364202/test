// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package latest_media is a model for cache.latest_media This table store latest_media information.
package latest_media

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"

	"database/sql"
	"fmt"
	"strconv"
	"time"
)

//  update cache.latest_media
func UpdateMediadDataCache() error {
	var err error
	var agencys []int64 = []int64{0}
	for _, v := range agencys {
		err = UpdateMediadDataCacheByMediaId(v)
		if err != nil {
			return err
		}
	}
	return nil
}

//  update cache.latest_media ทีละหน่วยงาน
//	Parameters:
//		agency_id
//			รหัสหน่วยงานที่อยากให้อัพเดท (0 = ทั้งหมด)
func UpdateMediadDataCacheByMediaId(media_type_id int64) error {
	db, err := pqx.Open()
	if err != nil {
		return errors.Repack(err)
	}
	tx, err := db.Begin()
	if err != nil {
		//		return errors.Repack(err)
		return errors.New("err begin")
	}
	defer tx.Rollback()

	lx := pqx.NewAdvisoryLock(tx, cacheTableName)
	b, err := lx.Wait(time.Minute * 3)
	if err != nil {
		//		return errors.Repack(err)
		return errors.New("err NewAdvisoryLock")
	}
	if !b {
		return errors.Newf("Can not get %s update lock", cacheTableName)
	}
	defer lx.Unlock()
	q := "SELECT MAX(dataimport_log_id) FROM " + cacheTableName
	p := []interface{}{}
	if media_type_id != 0 {
		q += " WHERE media_type_id = $1"
		p = append(p, media_type_id)
	}
	var lastImport sql.NullInt64

	fmt.Println(q)
	if err = tx.QueryRow(q, p...).Scan(&lastImport); err != nil {
		return err
	}

	q = `
			SELECT id,
			       agency_id,
			       media_type_id,
			       media_datetime,
			       media_path,
			       media_desc,
			       filename,
			       refer_source,
			       remark,
			       dataimport_log_id
			FROM   PUBLIC.media
			WHERE  dataimport_log_id > $1
			       AND deleted_at = To_timestamp(0)
			ORDER  BY dataimport_log_id, media_datetime
		`
	//	q = `
	//	WITH a AS
	//	  ( SELECT agency_id,
	//	           media_type_id,
	//	           max(dataimport_log_id) AS log_id
	//	   FROM PUBLIC.media
	//	   WHERE dataimport_log_id > $1
	//	     AND deleted_at = To_timestamp(0)
	//	   GROUP BY agency_id,
	//	            media_type_id)
	//	SELECT id,
	//		   m.agency_id,
	//		   m.media_type_id,
	//		   media_datetime,
	//		   media_path,
	//		   media_desc,
	//		   filename,
	//		   refer_source,
	//		   remark,
	//		   m.dataimport_log_id
	//	FROM PUBLIC.media m
	//	INNER JOIN a ON m.agency_id = a.agency_id
	//	AND m.media_type_id = a.media_type_id
	//	AND m.dataimport_log_id = a.log_id
	//	ORDER BY media_datetime
	//	`

	fmt.Println(q)
	rows, err := db.Query(q, lastImport.Int64)
	if err != nil {
		return errors.New(" err tx.Query : " + tx.Commit().Error())
	}
	newdata := make(map[string]*Struct_UpdateCacheMedia)

	newdata_agency := make(map[string]bool)
	newdata_media_type := make(map[string]bool)

	for rows.Next() {
		d := new(Struct_UpdateCacheMedia)
		if err := rows.Scan(&d.DataId, &d.AgencyId, &d.MediaTypeId, &d.MediaDatetime, &d.MediaPath, &d.MediaDesc,
			&d.Filename, &d.ReferSource, &d.Remark, &d.DataimportLogId); err != nil {
			return errors.Repack(err)
		}
		d.SetIndex()
		if !d.Index.Valid {
			log.Locationf("no config agency_id : %d, media_type_id : %d", d.AgencyId.Int64, d.MediaTypeId.Int64)
			continue // no index no need to update
		}
		m_str := strconv.FormatInt(d.MediaTypeId.Int64, 10)
		a_str := strconv.FormatInt(d.AgencyId.Int64, 10)
		key := a_str + "." + m_str + "." + d.Index.String

		nd := newdata[key]
		if nd != nil {
			// key ซ้ำ เอาอันที่ media_dateteime มากสุด
			nd_t := pqx.NullStringToTime(nd.MediaDatetime)
			d_t := pqx.NullStringToTime(d.MediaDatetime)
			if d_t.Before(nd_t) {
				continue
			}
		}

		newdata[key] = d
		newdata_agency[a_str] = true
		newdata_media_type[m_str] = true
	}
	if err := rows.Close(); err != nil {
		return errors.New("err rows.Close() : " + err.Error())
	}
	if len(newdata) == 0 { // no new data
		return nil
	}

	q_agency := ""
	for a_str, _ := range newdata_agency {
		if q_agency != "" {
			q_agency += ","
		}
		q_agency += a_str
	}

	q_media_type := ""
	for m_str, _ := range newdata_media_type {
		if q_media_type != "" {
			q_media_type += ","
		}
		q_media_type += m_str
	}

	// Do we need to update old record?
	q = `
	SELECT	data_id, 
			agency_id, 
			media_type_id, 
			media_datetime, 
			media_path, 
			media_desc, 
			filename, 
			refer_source, 
			remark, 
			dataimport_log_id, 
			index  
	FROM	` + cacheTableName + `
	WHERE	agency_id IN (` + q_agency + `)
			AND media_type_id IN (` + q_media_type + `)
	`

	fmt.Println(q)
	rows, err = db.Query(q)
	olddata := make(map[string]*Struct_UpdateCacheMedia)
	for rows.Next() {
		d := new(Struct_UpdateCacheMedia)
		if err := rows.Scan(&d.DataId, &d.AgencyId, &d.MediaTypeId, &d.MediaDatetime, &d.MediaPath, &d.MediaDesc,
			&d.Filename, &d.ReferSource, &d.Remark, &d.DataimportLogId, &d.Index); err != nil {
			return errors.Repack(err)
		}

		m_str := strconv.FormatInt(d.MediaTypeId.Int64, 10)
		a_str := strconv.FormatInt(d.AgencyId.Int64, 10)
		key := a_str + "." + m_str + "." + d.Index.String
		olddata[key] = d
	}
	if err := rows.Close(); err != nil {
		return errors.Repack(err)
	}

	for k, d := range newdata {
		od := olddata[k]
		// No old data to check
		if od == nil {
			continue
		}
		// Data was not changed, no need to update
		if *od == *d {
			delete(newdata, k)
			continue
		}

		od_t := pqx.NullStringToTime(od.MediaDatetime)
		d_t := pqx.NullStringToTime(d.MediaDatetime)
		// Data was older but new dataimport_log_id
		if d_t.Before(od_t) {
			delete(newdata, k)
		}
	}

	q = `
	INSERT INTO ` + cacheTableName + ` 
	            ( 
	                        agency_id, 
	                        media_type_id, 
	                        media_datetime, 
	                        media_path, 
	                        media_desc, 
	                        filename, 
	                        refer_source, 
	                        remark, 
	                        data_id, 
	                        index, 
	                        dataimport_log_id 
	            ) 
	VALUES ($1,
	        $2,
	        $3,
	        $4,
	        $5,
	        $6,
	        $7,
	        $8,
	        $9,
	        $10,
	        $11) ON conflict (agency_id, media_type_id, index) DO
	UPDATE
	SET media_datetime = excluded.media_datetime,
	    media_path = excluded.media_path,
	    media_desc = excluded.media_desc,
	    filename = excluded.filename,
	    refer_source = excluded.refer_source,
	    remark = excluded.remark,
	    data_id = excluded.data_id,
	    dataimport_log_id = excluded.dataimport_log_id
	`

	fmt.Println(q)
	stmt, err := tx.Prepare(q)
	if err != nil {
		return errors.Repack(err)
	}
	for _, d := range newdata {
		if !d.Index.Valid || d.Index.String == "" {
			continue
		}
		if _, err := stmt.Exec(d.AgencyId, d.MediaTypeId, d.MediaDatetime, d.MediaPath, d.MediaDesc, d.Filename,
			d.ReferSource, d.Remark, d.DataId, d.Index, d.DataimportLogId); err != nil {
			log.Logf("error upsert data_id %s, data %s", d.DataId.Int64, d)
			return errors.Repack(err)
		}
	}
	stmt.Close()

	return errors.Repack(tx.Commit())
}

// Set media index ตาม config ในไฟล์ map.go
func (d *Struct_UpdateCacheMedia) SetIndex() {
	cfg := GetCfg(d.AgencyId.Int64, d.MediaTypeId.Int64)
	if cfg == nil { // no config no update
		return
	} else {
		filename := d.Filename.String
		index := ""
		switch cfg.Function {
		case "substr":
			param1 := GetInt(cfg.Param1)
			param2 := GetInt(cfg.Param2)
			index = filename[param1:param2]
			d.AddIndex(index)
			break
		case "custom":
			cfg.SetIndex(d)
		default:
			d.AddIndex(filename)
		}

	}

}

// Add index value to d.Index
//
//	Parameters:
//		s
//			index value
func (d *Struct_UpdateCacheMedia) AddIndex(s string) {
	d.Index.Valid = true
	d.Index.String = s
}
