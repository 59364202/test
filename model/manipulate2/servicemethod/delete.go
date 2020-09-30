package servicemethod

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlDeleteServiceMethodTranslate = ` DELETE FROM lt_servicemethod_translate WHERE method_id = $1 `

func DeleteServiceMethod(userId int64, serviceMethodId int64) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Update to Delete lt_servicemethod
	err = updateToDeleteServiceMethod(tx, serviceMethodId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Update to Delete lt_servicemethod_translate
	err = updateToDeleteServiceMethodTranslate(tx, serviceMethodId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx.Commit()

	return result.Result1("delete success"), nil
}

func deleteServiceMethodTranslate(tx *pqx.Tx, serviceMethodId int64, userId int64) error {
	statement, err := tx.Prepare(sqlDeleteServiceMethodTranslate)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(serviceMethodId)
	if err != nil {
		return err
	}

	return nil
}
