package metadata

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	//	"strings"
	//	"time"
)

var sqlUpdateMetadataTranslate = sqlInsertMetadataTranslate +
	" ON CONFLICT (metadata_id , language_id ) " +
	" DO UPDATE SET metadataagency_name = $3 , metadataservice_name = $4 , metadata_tag = $5 , metadata_description = $6 , updated_by = $7 "

func PutMetadata(metadataIdString string, userId int64, p map[string]interface{}) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	metadataId, err := strconv.ParseInt(p["id"].(string), 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}
	var metadataagency_name = p["metadataagency_name"].(map[string]interface{})
	var metadataservice_name = p["metadataservice_name"].(map[string]interface{})
	var metadata_description = p["metadata_description"].(map[string]interface{})
	var metadata_tag = p["metadata_tag"].(map[string]interface{})
	var metadata_frequency = p["metadata_datafrequency"].(string)
	var history_description = p["history_description"].(string)
	var hydro = p["hydro"].(string)

	delete(p, "metadata_description")
	delete(p, "metadataagency_name")
	delete(p, "metadataservice_name")
	delete(p, "metadata_tag")
	delete(p, "metadata_datafrequency")
	delete(p, "history_description")
	delete(p, "hydro")
	delete(p, "id")

	err = updateMetadata(tx, metadataId, userId, p)
	if err != nil {
		return nil, err
	}

	err = deleteMetadataTranslate(tx, metadataId)
	if err != nil {
		return nil, err
	}

	err = updateMetadataTran(tx, userId, metadataId, metadataagency_name, metadataservice_name, metadata_description, metadata_tag)
	if err != nil {
		return nil, err
	}

	err = insertMetadataHistory(tx, userId, metadataId, history_description)
	if err != nil {
		return nil, err
	}

	err = deleteMetadataHydro(tx, metadataIdString)
	if err != nil {
		return nil, err
	}

	err = insertMetadataHydro(tx, userId, metadataId, hydro)
	if err != nil {
		return nil, err
	}

	err = deleteMetadataFrequency(tx, metadataIdString)
	if err != nil {
		return nil, err
	}

	err = insertMetadataFrequency(tx, userId, metadataId, metadata_frequency)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return result.Result1(metadataId), nil
}

func updateMetadata(tx *pqx.Tx, metadataId, userId int64, p map[string]interface{}) error {
	//	layout := "2006-01-02"
	var sqlSet = ""
	var param = make([]interface{}, 0)
	var countParam = 1
	for k, v := range p {
		sqlSet += k + " = $" + strconv.Itoa(countParam)
		if countParam < len(p) {
			sqlSet += " , "
		}

		param = append(param, v)
		countParam++
	}
	param = append(param, metadataId)
	statement, err := tx.Prepare(" UPDATE metadata SET " + sqlSet + " WHERE id = $" + strconv.Itoa(countParam))
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	_, err = statement.Exec(param...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func updateMetadataTran(tx *pqx.Tx, userId int64, metadataId int64, AgencyName, ServiceName, Description, Tag map[string]interface{}) error {
	statement, err := tx.Prepare(sqlUpdateMetadataTranslate)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	for k, v := range ServiceName {
		_, err = statement.Exec(metadataId, nil, AgencyName[k], v, Tag[k], Description[k], userId)
		if err != nil {
			return pqx.GetRESTError(err)
		}
	}

	return nil
}

func updateMetadataHistory(tx *pqx.Tx, userId int64, metadataId string, desc string) error {
	statement, err := tx.Prepare(sqlInsertMetadataHistory)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	_, err = statement.Exec(metadataId, desc, userId)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
