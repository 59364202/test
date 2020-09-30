package servicemethod

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlInsertServiceMethod = ` INSERT INTO lt_servicemethod (created_by, updated_by) VALUES ($1, $1) RETURNING id `

var sqlInsertServiceMethodTranslate = ` INSERT INTO lt_servicemethod_translate (
											method_id,
											language_id,
											servicemethod_name,
											created_by,
											updated_by ) VALUES ($1, $2, $3, $4, $4) `

func PostServiceMethod(userId int64, mapServiceMethodName map[string]string) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Insert lt_servicemethod
	newId, err := insertServiceMethod(tx, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Insert lt_servicemethod_translate
	err = insertServiceMethodTranslate(tx, newId, mapServiceMethodName, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx.Commit()

	//data := &ServiceMethod_struct{Id: newId, ServiceMethodName: mapServiceMethodName}
	return result.Result1("Insert Successful"), nil
}

func insertServiceMethod(tx *pqx.Tx, userId int64) (int64, error) {
	var (
		_id int64
	)
	statment, err := tx.Prepare(sqlInsertServiceMethod)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	err = statment.QueryRow(userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}

func insertServiceMethodTranslate(tx *pqx.Tx, serviceMethodId int64, mapServiceMethodName map[string]string, userId int64) error {
	statementTrans, err := tx.Prepare(sqlInsertServiceMethodTranslate)
	if err != nil {
		return err
	}
	defer statementTrans.Close()

	//Loop mapServiceMethodName
	for _, methodName := range mapServiceMethodName {
		_, err = statementTrans.Exec(serviceMethodId, nil, methodName, userId)
		if err != nil {
			return errors.Repack(err)
		}
	}

	return nil
}
