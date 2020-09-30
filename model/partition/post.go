// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package partition is a model for manage partition table
package partition

import (
	"haii.or.th/api/server/model/setting"
	model_partition_history "haii.or.th/api/thaiwater30/model/dbamodule_history"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"strconv"
	"strings"
)

const DefaultSchema = "public"

//	validate parameter
//	Parameters:
//		param
//			Struct_PartitionTable_InputParam
//	Return:
//		transaction
//		PartitionTableMaker
func validatePartionTableParam(param *Struct_PartitionTable_InputParam) (*pqx.Tx, *pqx.PartitionTableMaker, error) {
	//Check input params
	if param.TableName == "" {
		return nil, nil, rest.NewError(422, "'table_name' must not empty.", errors.New("parameter 'table_name' must not empty."))
	}

	if y, _ := strconv.Atoi(param.Year); y <= 0 {
		return nil, nil, rest.NewError(422, "'year' must be a valid interger.", errors.New("parameter 'year' must be a valid interger."))
	}

	//Try to open database
	db, err := pqx.Open()
	if err != nil {
		return nil, nil, pqx.GetRESTError(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, pqx.GetRESTError(err)
	}

	schemaName, tableName := pqx.GetTableNameInfo(param.TableName, DefaultSchema)
	px, err := pqx.NewPartitionTableMaker(tx, schemaName, tableName, "")
	if err != nil {
		tx.Rollback()
		return nil, nil, pqx.GetRESTError(err)
	}
	return tx, px, nil
}

const SettingPartitionYearly = "server.model.dataimport.PartitionYearly"

//	create partition
//	Parameters:
//		userId
//			รหัสผู้ใข้
//		param
//			Struct_PartitionTable_InputParam
//	Return:
//		"Create Successful"
func PostPartition(userId int64, param *Struct_PartitionTable_InputParam) (string, error) {

	tx, px, err := validatePartionTableParam(param)
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	var createdP []string
	var pnames []string
	s := "," + setting.GetSystemSetting(SettingPartitionYearly) + ","
	if p := strings.Index(s, ","+px.GetMasterTableName()+","); p >= 0 {
		pnames = append(pnames, px.GetMasterTableName()+"_y"+param.Year)
	} else {
		for i := 1; i <= 12; i++ {
			s := px.GetMasterTableName() + "_y" + param.Year + "m"
			if i < 10 {
				s += "0"
			}
			s += strconv.Itoa(i)
			pnames = append(pnames, s)
		}
	}

	for _, pname := range pnames {
		isCreated, err := px.CreateTablesIfNotExisted(pname)
		if err != nil {
			return "", pqx.GetRESTError(err)
		}
		if isCreated {
			createdP = append(createdP, px.GetTableFullName(pname))
		}
	}

	//Commit Transaction
	if err = tx.Commit(); err != nil {
		return "", pqx.GetRESTError(err)
	}

	//Insert create log table
	paramHistory := &model_partition_history.Struct_DBAModuleHistory_InputParam{}
	paramHistory.TableName = param.TableName
	paramHistory.Year = param.Year
	for _, s := range createdP {
		paramHistory.Month = s[len(s)-2:]
		paramHistory.Remarks = "CREATE " + s
		_, err = model_partition_history.PostDBAModuleHistory(userId, paramHistory)
		if err != nil {
			return "", pqx.GetRESTError(err)
		}
	}

	//Return data
	return "Create Successful", nil
}
