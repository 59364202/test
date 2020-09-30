// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package partition is a model for manage partition table
package partition

import (
	model_partition_history "haii.or.th/api/thaiwater30/model/dbamodule_history"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"strconv"
	"strings"
)

//	Delete partition table
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		param
//			Struct_PartitionTable_InputParam
//	Return:
//		"Delete Successful"
func DeletePartitionTable(userId int64, param *Struct_PartitionTable_InputParam) (string, error) {

	tx, px, err := validatePartionTableParam(param)
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	cs, err := px.GetChildrenNames()
	if err != nil {
		return "", errors.Repack(err)
	}

	var dlist []string
	dy, _ := strconv.Atoi(param.Year)
	q := "SELECT "

	for _, s := range cs {
		_, y, _ := pqx.GetPartitionNameInfo(s)
		if y == dy {
			n := px.GetTableFullName(s)
			dlist = append(dlist, n)
			q += " CASE WHEN EXISTS (SELECT * FROM " + n +
				" LIMIT 1) THEN 1 ELSE 0 END +"
		}
	}
	q += "0"

	if len(dlist) == 0 {
		return "", rest.NewError(422, "ไม่พบตารางพาทิชันของ "+param.TableName+" ปี "+param.Year, nil)
	}

	//Check record in partition table
	var d int
	if err = tx.QueryRow(q).Scan(&d); err != nil {
		return "", pqx.GetRESTError(err)
	}
	if d > 0 {
		return "", rest.NewError(422, "พบข้อมูลในตารางพาทิชันของ '"+param.TableName+" ปี "+param.Year, nil)
	}

	//Delete partition table
	err = dropPartitionTable(tx, param.TableName, param.Year, dlist, userId)
	if err != nil {
		return "", err
	}

	//Commit transaction
	if err = tx.Commit(); err != nil {
		return "", err
	}

	//Return result
	return "Delete Successful", nil
}

//	Function for drop partition table
//	Parameters:
//		tx
//			transaction
//		tableName
//			ชื่อตาราง
//		year
//			ปี
//		arrPartitionTable
//			ชื่อตารางพาร์ทิชั่น
//		userId
//			รหัสผู้ใช้
//	Return:
//		nil, error
func dropPartitionTable(qs pqx.Querist, tableName string, year string, arrPartitionTable []string, userId int64) error {

	//Query statement with parameters
	_, err := qs.Query("DROP TABLE " + strings.Join(arrPartitionTable, ", ") + "; ")
	if err != nil {
		return pqx.GetRESTError(err)
	}

	//Insert delete log table
	paramHistory := &model_partition_history.Struct_DBAModuleHistory_InputParam{}
	paramHistory.TableName = tableName
	paramHistory.Year = year

	for _, strPartitionTable := range arrPartitionTable {
		paramHistory.Month = strPartitionTable[len(strPartitionTable)-2:]
		paramHistory.Remarks = "DROP " + strPartitionTable
		_, err = model_partition_history.PostDBAModuleHistory(userId, paramHistory)
		if err != nil {
			return err
		}
	}

	return nil
}
