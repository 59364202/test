package metadata

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//	"log"
	"strconv"
	"strings"
	//	"time"
)

var sqlInsertMetadataTranslate = ` INSERT INTO metadata_translate (
										metadata_id,
										language_id,
										metadataagency_name,
										metadataservice_name,
										metadata_tag,
										metadata_description,
										created_by) VALUES ($1, $2, $3, $4, $5, $6, $7) `
var sqlInsertMetadataHistory = ` INSERT INTO metadata_history(
									metadata_datetime,
									metadata_id,
									history_description,
									created_by) VALUES (NOW(), $1, $2, $3) `
var sqlInsertMetadataHydro = ` INSERT INTO metadata_hydroinfo(
									metadata_id,
									hydroinfo_id,
									created_by) VALUES ($1, $2, $3) `
var sqlInsertMetadataFreq = ` INSERT INTO metadata_frequency(
									metadata_id,
									datafrequency,
									created_by) VALUES ($1, $2, $3) `

func PostMetadata(userId int64, p map[string]interface{}) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	var metadataservice_name = p["metadataservice_name"].(map[string]interface{})
	var metadataagency_name = p["metadataagency_name"].(map[string]interface{})
	var metadata_description = p["metadata_description"].(map[string]interface{})
	var metadata_tag = p["metadata_tag"].(map[string]interface{})
	var metadata_frequency = p["metadata_datafrequency"].(string)
	var history_description = p["history_description"].(string)
	var hydro = p["hydro"].(string)

	delete(p, "metadata_description")
	delete(p, "metadataservice_name")
	delete(p, "metadataagency_name")
	delete(p, "metadata_tag")
	delete(p, "metadata_datafrequency")
	delete(p, "history_description")
	delete(p, "hydro")
	delete(p, "id")

	newId, err := inertMetadata(tx, userId, p)
	if err != nil {
		return nil, err
	}

	err = insertMetadataTran(tx, userId, newId, metadataagency_name, metadataservice_name, metadata_description, metadata_tag)
	if err != nil {
		return nil, err
	}

	err = insertMetadataHistory(tx, userId, newId, history_description)
	if err != nil {
		return nil, err
	}

	err = insertMetadataHydro(tx, userId, newId, hydro)
	if err != nil {
		return nil, err
	}

	err = insertMetadataFrequency(tx, userId, newId, metadata_frequency)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return result.Result1(newId), nil
}

func inertMetadata(tx *pqx.Tx, userId int64, p map[string]interface{}) (int64, error) {
	//	layout := "2006-01-02"
	var sqlColumn = ""
	var sqlValues = ""
	var param = make([]interface{}, 0)
	//	param = append(param, userId)
	var countParam = 1
	var newId int64
	for k, v := range p {
		sqlColumn += k
		sqlValues += "$" + strconv.Itoa(countParam)
		if countParam < len(p) {
			sqlColumn += ","
			sqlValues += ","
		} else {
			sqlColumn += ")"
			sqlValues += ")"
		}
		//		if k == "metadata_startdatadate" {
		//
		//			t, err := time.Parse(layout, v.(string))
		//			if err != nil {
		//				return 0, errors.Repack(err)
		//			}
		//			param = append(param, t)
		//		} else {
		//			param = append(param, v)
		//		}
		param = append(param, v)
		countParam++
	}
	statement, err := tx.Prepare("INSERT INTO metadata (" + sqlColumn + " VALUES (" + sqlValues + " RETURNING id ")
	//	log.Println("INSERT INTO metadata (" + sqlColumn + " VALUES (" + sqlValues + " RETURNING id ")
	if err != nil {
		return 0, errors.Repack(err)
	}
	defer statement.Close()

	err = statement.QueryRow(param...).Scan(&newId)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return newId, nil
}

func insertMetadataTran(tx *pqx.Tx, userId, metadataId int64, AgencyName, ServiceName, Description, Tag map[string]interface{}) error {
	statement, err := tx.Prepare(sqlInsertMetadataTranslate)
	//	log.Println("insertMetadataTran")
	//	log.Println(sqlInsertMetadataTranslate)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	for k, v := range ServiceName {
		_, err = statement.Exec(metadataId, nil, AgencyName[k], v, Tag[k], Description[k], userId)
		//log.Println(metadataId, model_language.GetID(k), AgencyName[k], v, Tag[k], Description[k], userId)
		if err != nil {
			return pqx.GetRESTError(err)
		}
	}

	return nil
}

func insertMetadataHistory(tx *pqx.Tx, userId, metadataId int64, desc string) error {
	statement, err := tx.Prepare(sqlInsertMetadataHistory)
	//	log.Println("insertMetadataHistory")
	//	log.Println(sqlInsertMetadataHistory)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	_, err = statement.Exec(metadataId, desc, userId)
	//	log.Println(metadataId, desc, userId)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func insertMetadataHydro(tx *pqx.Tx, userId, metadataId int64, hydros string) error {
	statement, err := tx.Prepare(sqlInsertMetadataHydro)
	//	log.Println("insertMetadataHydro")
	//	log.Println(sqlInsertMetadataHydro)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	arrayHydro := strings.Split(hydros, ",")
	for _, v := range arrayHydro {
		_, err = statement.Exec(metadataId, v, userId)
		//		log.Println(metadataId, v, userId)
		if err != nil {
			return pqx.GetRESTError(err)
		}
	}

	return nil
}

func insertMetadataFrequency(tx *pqx.Tx, userId, metadataId int64, strListOfFreq string) error {
	statement, err := tx.Prepare(sqlInsertMetadataFreq)
	//	log.Println("insertMetadataFrequency")
	//	log.Println(sqlInsertMetadataFreq)
	if err != nil {
		return errors.Repack(err)
	}
	defer statement.Close()

	arrMetadataFreq := strings.Split(strListOfFreq, ", ")
	for _, v := range arrMetadataFreq {
		_, err = statement.Exec(metadataId, v, userId)
		//		log.Println(metadataId, v, userId)
		if err != nil {
			return pqx.GetRESTError(err)
		}
	}
	return nil
}
