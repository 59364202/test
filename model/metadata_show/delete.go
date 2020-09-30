// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
//
package metadata_show

import (
	//	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// soft deleted metadata_show by id
//  Parameters:
//		uid
//			user id
//		id
//			metadata_show id
//  Return:
//		"Delete Successful"
func Delete(uid int64, id string) (string, error) {

	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	//	inf, err := model.GetCipher().DecryptText(id)
	//	if err != nil {
	//		return "", errors.New("Can not deleted data invalid key")
	//	}
	//	mid, err := strconv.Atoi(inf)
	//	if err != nil {
	//		return "", err
	//	}

	mid, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}
	q := sqlDelete
	//	p := []interface{}{uid, mid}
	p := []interface{}{mid}

	stmt, err := db.Prepare(q)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	return "Delete Successful", nil
}
