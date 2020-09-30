package metadata

import (
	//	model_language "haii.or.th/api/server/model/language"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlCheckChild = " SELECT id FROM metadata_hydroinfo WHERE metadata_id = $1 AND deleted_by IS NULL LIMIT 1 "
var sqlUpdateToDeleteMetadata = " UPDATE metadata SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "

var sqlUpdateToDeleteMetadataTran = " UPDATE metadata_translate SET deleted_by = $2 , deleted_at = NOW() WHERE metadata_id = $1 "
var sqlDeleteMetadataTran = " DELETE FROM metadata_translate WHERE metadata_id = $1 "

var sqlUpdateToDeleteMetadataHydro = " UPDATE metadata_hydroinfo SET deleted_by = $2 , deleted_at = NOW() WHERE metadata_id = $1 "
var sqlDeleteMetadataHydro = " DELETE FROM metadata_hydroinfo WHERE metadata_id = $1 "

var sqlUpdateToDeleteMetadataFreq = " UPDATE metadata_frequency SET deleted_by = $2 , deleted_at = NOW() WHERE metadata_id = $1 "
var sqlDeleteMetadataFrequency = " DELETE FROM metadata_frequency WHERE metadata_id = $1 "

func DeleteMetadata(metadataId string, userId int64) (*result.Result, error) {
	//Check invalid input
	if metadataId == "" {
		return nil, errors.New("Null Input Parameter")
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Check Matadata's Child before delete
	//	var hasChild = false
	//	_result, err := db.Query(sqlCheckChild, metadataId)
	//	if err != nil {
	//		return nil, pqx.GetRESTError(err)
	//	}
	//	for _result.Next() {
	//		hasChild = true
	//	}
	//	if hasChild {
	//		return result.Result0(""), nil
	//	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	statment, err := tx.Prepare(sqlUpdateToDeleteMetadata)
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer statment.Close()

	_, err = statment.Exec(metadataId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	statmentTran, err := tx.Prepare(sqlUpdateToDeleteMetadataTran)
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer statmentTran.Close()

	_, err = statmentTran.Exec(metadataId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	statmentHydro, err := tx.Prepare(sqlUpdateToDeleteMetadataHydro)
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer statmentHydro.Close()

	_, err = statmentHydro.Exec(metadataId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	statmentFreq, err := tx.Prepare(sqlUpdateToDeleteMetadataFreq)
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer statmentFreq.Close()

	_, err = statmentFreq.Exec(metadataId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx.Commit()

	return result.Result1("delete success"), nil
}

func deleteMetadataTranslate(tx *pqx.Tx, metadataId int64) error {
	statement, err := tx.Prepare(sqlDeleteMetadataTran)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	_, err = statement.Exec(metadataId)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func deleteMetadataHydro(tx *pqx.Tx, metadataId string) error {
	statement, err := tx.Prepare(sqlDeleteMetadataHydro)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	_, err = statement.Exec(metadataId)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func deleteMetadataFrequency(tx *pqx.Tx, metadataId string) error {
	statement, err := tx.Prepare(sqlDeleteMetadataFrequency)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	_, err = statement.Exec(metadataId)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
