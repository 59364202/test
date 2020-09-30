package metadata

import (
	model_server "haii.or.th/api/server/model"
	model_metadata_frequency "haii.or.th/api/thaiwater30/model/metadata_frequency"
	model_metadata_history "haii.or.th/api/thaiwater30/model/metadata_history"
	model_metadata_hydroinfo "haii.or.th/api/thaiwater30/model/metadata_hydroinfo"
	model_metadata_servicemethod "haii.or.th/api/thaiwater30/model/metadata_servicemethod"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
	"strconv"
)

//	update metadata
//	Parameters:
//		param
//			Struct_Metadata_Data_InputParam
//		userID
//			รหัสผู้ใช้งาน
//	Return:
//		Update Successful
func PutMetadata(param *Struct_Metadata_Data_InputParam, userID int64) (string, error) {
	//Check 'metadata_id' is not null.
	if param.MetadataID == "" {
		return "", rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
	}

	//Decrypt metadata_id
	strMetadataID, err := model_server.GetCipher().DecryptText(param.MetadataID)
	if err != nil {
		return "", rest.NewError(422, "Invalid metadata_id", errors.New("Invalid metadata_id."))
	}
	param.ID, err = strconv.ParseInt(strMetadataID, 10, 64)
	if err != nil {
		return "", rest.NewError(422, "metadata_id is not a number.", err)
	}

	//	log.Println(param.MetadataID)
	//	log.Println(param.ID)

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Update metadata table
	err = fncUpdateData(tx, param, userID)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Update metadata_hydroinfo table
	err = model_metadata_hydroinfo.FncDeleteDataByMetadata(tx, param.ID)
	if err != nil {
		return "", err
	}
	for _, intHydroID := range param.Hydroinfo {
		err = model_metadata_hydroinfo.FncInsertData(tx, param.ID, intHydroID, userID)
		if err != nil {
			return "", err
		}
	}

	//Update metadata_servicemethod table
	err = model_metadata_servicemethod.FncDeleteDataByMetadata(tx, param.ID)
	if err != nil {
		return "", err
	}
	for _, intServiceMethodID := range param.ServiceMethod {
		err = model_metadata_servicemethod.FncInsertData(tx, param.ID, intServiceMethodID, userID)
		if err != nil {
			return "", err
		}
	}

	//Update metadata_frequency table
	err = model_metadata_frequency.FncDeleteDataByMetadata(tx, param.ID)
	if err != nil {
		return "", err
	}
	for _, strFrequency := range param.Frequency {
		err = model_metadata_frequency.FncInsertData(tx, param.ID, strFrequency, userID)
		if err != nil {
			return "", err
		}
	}

	//Insert metadata_history table
	err = model_metadata_history.FncInsertData(tx, param.ID, param.HistoryDescription, userID)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	return "Update Successful", nil
}

//Update data to metadata table
func fncUpdateData(tx *pqx.Tx, param *Struct_Metadata_Data_InputParam, userId int64) error {
	var (
		jsonMetadataAgencyName  interface{} = nil
		jsonMetadataServiceName interface{} = nil
		jsonMetadataTag         interface{} = nil
		jsonMetadataDescription interface{} = nil

		err error
	)

	/*
		//Convert metadata_id type from string to int64
		intMetadataID, err := strconv.ParseInt(param.MetadataID, 10, 64)
		if err != nil {
			return 0, rest.NewError(422, "metadata_id is not a number.", err)
		}*/

	//Convert metadataagency_name to db-json type
	if param.MetadataAgencyName != nil {
		jsonMetadataAgencyName, err = param.MetadataAgencyName.MarshalJSON()
		if err != nil {
			return rest.NewError(422, "'metadataagency_name' not a json format.", err)
		}
	}
	//Convert metadataservice_name to db-json type
	if param.MetadataServiceName != nil {
		jsonMetadataServiceName, err = param.MetadataServiceName.MarshalJSON()
		if err != nil {
			return rest.NewError(422, "'metadataservice_name' not a json format.", err)
		}
	}
	//Convert metadata_tag to db-json type
	if param.MetadataTag != nil {
		jsonMetadataTag, err = param.MetadataTag.MarshalJSON()
		if err != nil {
			return rest.NewError(422, "'metadata_tag' not a json format.", err)
		}
	}
	//Convert metadata_description to db-json type
	if param.MetadataDescription != nil {
		jsonMetadataDescription, err = param.MetadataDescription.MarshalJSON()
		if err != nil {
			return rest.NewError(422, "'metadata_description' not a json format.", err)
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateMetadata)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	var arrParam = make([]interface{}, 0)
	arrParam = append(arrParam, param.ID)
	arrParam = append(arrParam, param.SubcategoryID)
	arrParam = append(arrParam, param.AgencyID)
	arrParam = append(arrParam, param.DataunitID)
	arrParam = append(arrParam, param.DataformatID)
	arrParam = append(arrParam, param.ConnectionFormat)
	arrParam = append(arrParam, param.MetadataContact)

	if param.MetadataAgencystoredate == "" {
		arrParam = append(arrParam, nil)
	} else {
		arrParam = append(arrParam, param.MetadataAgencystoredate)
	}

	if param.MetadataStartdatadate == "" {
		arrParam = append(arrParam, nil)
	} else {
		arrParam = append(arrParam, param.MetadataStartdatadate)
	}

	arrParam = append(arrParam, param.MetadataUpdatePlan)
	arrParam = append(arrParam, param.MetadataLaws)
	arrParam = append(arrParam, param.MetadataRemark)
	arrParam = append(arrParam, jsonMetadataAgencyName)
	arrParam = append(arrParam, jsonMetadataServiceName)
	arrParam = append(arrParam, jsonMetadataTag)
	arrParam = append(arrParam, jsonMetadataDescription)
	arrParam = append(arrParam, param.MetadataStatusID)
	arrParam = append(arrParam, userId)
	arrParam = append(arrParam, param.MetadataConvertFrequency)
	arrParam = append(arrParam, param.ImportCount)
	if param.MetadataReceiveDate == "" {
		arrParam = append(arrParam, nil)
	} else {
		arrParam = append(arrParam, param.MetadataReceiveDate)
	}

	//Execute insert statement with parameters and returning id
	_, err = statement.Exec(arrParam...)
	if err != nil {
		return err
	}

	return nil
}

//	update soft deleted dataset config
//	Parameters:
//		uid
//			รหัสผู้ใช้งาน
//		mid
//			รหัสบัญชีข้อมูล
//	Return:
//		nil, error
func PutMetadataOfflineDate(uid, mid int64) error {

	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql soft deleted dataset config
	q := sqlUpdateMetadataOfflineDate
	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	p := []interface{}{uid, mid}

	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}
	return nil
}
