package servicemethod

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlUpdateServiceMethod = ` UPDATE lt_servicemethod
								  SET updated_by = $2 ,
									  updated_at = NOW()
								  WHERE id = $1 `

var sqlUpdateToDeleteServiceMethod = ` UPDATE lt_servicemethod
										  SET deleted_by = $2,
											  deleted_at = NOW(),
											  updated_by = $2 ,
											  updated_at = NOW()
										  WHERE id = $1 `

var sqlUpdateToDeleteServiceMethodTranslate = ` UPDATE lt_servicemethod_translate
												   SET deleted_by = $2,
													   deleted_at = NOW(),
													   updated_by = $2 ,
													   updated_at = NOW()
												   WHERE method_id = $1 `

func PutServiceMethod(userId int64, serviceMethodId int64, mapServiceMethodName map[string]string) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Update lt_servicemethod
	err = updateServiceMethod(tx, serviceMethodId, userId)
	if err != nil {
		return nil, err
	}

	//Delete lt_servicemethod_translate
	err = deleteServiceMethodTranslate(tx, serviceMethodId, userId)
	if err != nil {
		return nil, err
	}

	//Insert lt_servicemethod_translate
	err = insertServiceMethodTranslate(tx, serviceMethodId, mapServiceMethodName, userId)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return result.Result1(serviceMethodId), nil
}

func updateServiceMethod(tx *pqx.Tx, serviceMethodId int64, userId int64) error {
	//Update lt_servicemethod
	statement, err := tx.Prepare(sqlUpdateServiceMethod)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(serviceMethodId, userId)
	if err != nil {
		return err
	}

	return nil
}

func updateToDeleteServiceMethod(tx *pqx.Tx, serviceMethodId int64, userId int64) error {
	//Update lt_servicemethod to set 'Delete'
	statement, err := tx.Prepare(sqlUpdateToDeleteServiceMethod)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(serviceMethodId, userId)
	if err != nil {
		return err
	}

	return nil
}

func updateToDeleteServiceMethodTranslate(tx *pqx.Tx, serviceMethodId int64, userId int64) error {
	//Update lt_servicemethod_translate to set 'Delete'
	statement, err := tx.Prepare(sqlUpdateToDeleteServiceMethodTranslate)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(serviceMethodId, userId)
	if err != nil {
		return err
	}

	return nil
}
