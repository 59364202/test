package metadata

import (
	model_metadata_frequency "haii.or.th/api/thaiwater30/model/metadata_frequency"
	model_metadata_history "haii.or.th/api/thaiwater30/model/metadata_history"
	model_metadata_hydroinfo "haii.or.th/api/thaiwater30/model/metadata_hydroinfo"
	model_metadata_servicemethod "haii.or.th/api/thaiwater30/model/metadata_servicemethod"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
)

// 	insert metadata
//	Parameters:
//		param
//			Struct_Metadata_Data_InputParam
//		userID
//			รหัสผู้ใช้งาน
//	Return:
//		Insert Successful
func PostMetadata(param *Struct_Metadata_Data_InputParam, userID int64) (string, error) {

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

	//Insert metadata table
	intNewID, err := fncInsertData(tx, param, userID)
	if err != nil {
		return "", errors.Repack(err)
	}

	//	log.Println(intNewID)

	//Insert metadata_hydroinfo table
	for _, intHydroID := range param.Hydroinfo {
		err = model_metadata_hydroinfo.FncInsertData(tx, intNewID, intHydroID, userID)
		if err != nil {
			return "", err
		}
	}

	//Insert metadata_servicemethod table
	for _, intServiceMethodID := range param.ServiceMethod {
		err = model_metadata_servicemethod.FncInsertData(tx, intNewID, intServiceMethodID, userID)
		if err != nil {
			return "", err
		}
	}

	//Insert metadata_frequency table
	for _, strFrequency := range param.Frequency {
		err = model_metadata_frequency.FncInsertData(tx, intNewID, strFrequency, userID)
		if err != nil {
			return "", err
		}
	}

	//Insert metadata_history table
	param.HistoryDescription = "เพิ่มบัญชีข้อมูล"
	err = model_metadata_history.FncInsertData(tx, intNewID, param.HistoryDescription, userID)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	return "Insert Successful", nil
}

//  Insert to metadata table
//	Parameters:
//		tx
//			transaction
//		param
//			Struct_Metadata_Data_InputParam
//		userID
//			รหัสผู้ใช้งาน
//	Return:
//		new id
func fncInsertData(tx *pqx.Tx, param *Struct_Metadata_Data_InputParam, userId int64) (int64, error) {
	var (
		_id int64

		jsonMetadataAgencyName  interface{} = nil
		jsonMetadataServiceName interface{} = nil
		jsonMetadataTag         interface{} = nil
		jsonMetadataDescription interface{} = nil

		err error
	)

	//Convert metadataagency_name to db-json type
	if param.MetadataAgencyName != nil {
		jsonMetadataAgencyName, err = param.MetadataAgencyName.MarshalJSON()
		if err != nil {
			return 0, rest.NewError(422, "'metadataagency_name' not a json format.", err)
		}
	}
	//Convert metadataservice_name to db-json type
	if param.MetadataServiceName != nil {
		jsonMetadataServiceName, err = param.MetadataServiceName.MarshalJSON()
		if err != nil {
			return 0, rest.NewError(422, "'metadataservice_name' not a json format.", err)
		}
	}
	//Convert metadata_tag to db-json type
	if param.MetadataTag != nil {
		jsonMetadataTag, err = param.MetadataTag.MarshalJSON()
		if err != nil {
			return 0, rest.NewError(422, "'metadata_tag' not a json format.", err)
		}
	}
	//Convert metadata_description to db-json type
	if param.MetadataDescription != nil {
		jsonMetadataDescription, err = param.MetadataDescription.MarshalJSON()
		if err != nil {
			return 0, rest.NewError(422, "'metadata_description' not a json format.", err)
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertMetadata)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var arrParam = make([]interface{}, 0)
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
	err = statement.QueryRow(arrParam...).Scan(&_id)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return _id, nil
}
