package event_log_category

import (
	model_setting "haii.or.th/api/thaiwater30/model/setting"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// Deleted eventlog category
//  Parameters:
//		userId
//					user id update event code
//		params
//					Struct_EventLogCategory_InputParam
//  Return:
//		Delete Successful.
func DeleteEventLogCategory(userId int64, param *Struct_EventLogCategory_InputParam) (*result.Result, error) {
	//Check id is not null
	if param.ID == "" {
		return nil, errors.New("'id' is not null.")
	}

	//Convert id type from string to int64
	intID, err := strconv.ParseInt(param.ID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Check child table of lt_event_log_category
	isHasChild, err := checkEventLogCategoryChild(db, intID)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return result.Result0(""), nil
	}

	//Delete lt_event_log_category table
	err = deleteEventLogCategoryById(tx, intID)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Delete system_setting table
	objSettingParam := &model_setting.Struct_SystemSetting{}
	objSettingParam.Name = "bof.EventMgt.EventLogCategory.Color_" + strconv.FormatInt(intID, 10)
	_, err = model_setting.DeleteSystemSetting(userId, objSettingParam)
	if err != nil {
		defer tx.Rollback()
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return result.Result1("Delete Successful."), nil
}

//Check child table of lt_event_log_category
//  Parameters:
//		db
//			connection database
//		id
//			eventcode id
//  Return:
//		if is has child = true, if not = false
func checkEventLogCategoryChild(db *pqx.DB, id int64) (bool, error) {

	//Set default of return value
	var isHasChild bool = false

	//Query statement with parameters
	row, err := db.Query(sqlCheckEventLogCategoryChild, id)
	if err != nil {
		return isHasChild, err
	}

	//Check child
	for row.Next() {
		isHasChild = true
	}

	//Return result
	return isHasChild, nil
}

//Delete data at lt_event_log_category table
//  Parameters:
//		tx
//			transaction database
//		id
//			event code id
//  Return:
//		None
func deleteEventLogCategoryById(tx *pqx.Tx, id int64) error {

	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteEventLogCategory)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
