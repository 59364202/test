package metadata_frequency

import (
//	result "haii.or.th/api/thaiwater30/util/result"
//	"haii.or.th/api/util/errors"
//	"haii.or.th/api/util/pqx"
//	"haii.or.th/api/util/rest"
)

//	edit data in 'metadata_frequency' table with metadata_id (delete old and insert new)
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//		arrDataFrequency
//			array ความถี่
//		userID
//			รหัสผู้ใช้
//	Return:
//		Array Struct_MetadataFrequency
//func PutMetadataFrequencyMethod(metadataID int64, arrDataFrequency []string, userID int64) (*result.Result, error) {
//	//Check 'metadata_id' is not null.
//	if metadataID == 0 {
//		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
//	}
//
//	//Check 'dataFrequency' is not null.
//	if len(arrDataFrequency) == 0 {
//		return nil, rest.NewError(422, "dataFrequency is not null.", errors.New("parameter 'dataFrequency' is not null."))
//	}
//
//	//Open database
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, errors.Repack(err)
//	}
//
//	//Begin Transaction
//	tx, err := db.Begin()
//	if err != nil {
//		return nil, errors.Repack(err)
//	}
//	defer tx.Rollback()
//
//	//Delete data in 'metadata_frequency' table with metadata_id
//	err = FncDeleteDataByMetadata(tx, metadataID)
//	if err != nil {
//		return nil, err
//	}
//
//	//Loop for insert to 'metadata_frequency' table with servicemethod's array.
//	for _, strDataFrequency := range arrDataFrequency {
//		err := FncInsertData(tx, metadataID, strDataFrequency, userID)
//		if err != nil {
//			return nil, pqx.GetRESTError(err)
//		}
//	}
//
//	//Commit transaction
//	tx.Commit()
//
//	//Return result
//	return result.Result1("Updated Successful"), nil
//}
