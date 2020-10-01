// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataimport_config is a model for api.dataimport_download table. This table store dataimport config information.
package dataimport_config

import (
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// dataimport  soft deleted dataimport download
//  Parameters:
//		dataimport_download_id
//						dataimport download id for delete
//		uid
//						user id for deleted_by
//  Return:
//		Error
func DeleteDataimportDownload(dataimport_download_id string, uid int64) error {
	downloadID, _ := strconv.Atoi(dataimport_download_id)
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	//sql soft delted download config
	q := deleteDownloadConfig
	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	p := []interface{}{uid, downloadID}

	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	// soft deleted dataset by download id
	q = deleteDatasetConfigByDownloadID
	// prepare sql
	stmt, err = db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	p = []interface{}{uid, downloadID}

	// execute data
	res, err = stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	// sql soft deleted download and dataset from metadata
	q = deleteDownloadIDMetadata

	// prepare sql
	stmt, err = db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	p = []interface{}{uid, downloadID}

	// execute data
	res, err = stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}
	return nil
}

// dataimport soft delted dataimport dataset
//  Parameters:
//		dataimport_dataset_id
//						dataimport dataset id for delete
//		uid
//						user id for deleted_by
//  Return:
//		Error
func DeleteDataimportDataset(dataimport_dataset_id string, uid int64) error {

	datasetID, _ := strconv.Atoi(dataimport_dataset_id)
	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql soft deleted dataset config
	q := deleteDatasetConfig
	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	p := []interface{}{uid, datasetID}

	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	// sql soft deleted dataset from metadata
	q = deleteDatasetIDMetadata
	// prepare sql
	stmt, err = db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	p = []interface{}{uid, datasetID}

	// execute data
	res, err = stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}
	return nil
}

func DeleteConfigVariable(uid, vid int64) (*Struct_category2, error) {
	// q := `UPDATE api.config_variable SET category = $1, config_name = $2, variable_name = $3 , value = $4,updated_by = $5, updated_at = NOW() WHERE id = $6`
	q := ` UPDATE api.config_variable SET deleted_by = $2, deleted_at = NOW(), updated_by = $2, updated_at = NOW() WHERE id = $1 `

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(q)
	if err != nil {
		return nil, err
	}

	// ctime := time.Now()
	userid := uid

	_, err = stmt.Exec(userid, vid)
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx.Commit()
	if err != nil {
		return nil, errors.Repack(err)
	}

	return nil, nil
}
