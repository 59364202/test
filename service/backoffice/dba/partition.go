package dba

import (
	//	"encoding/json"
	//	model_setting "haii.or.th/api/server/model/setting"
	model_history "haii.or.th/api/thaiwater30/model/dbamodule_history"
	model "haii.or.th/api/thaiwater30/model/partition"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	"strconv"
	"time"
)

type Struct_TableList struct {
	TableName    string   `json:"table_name"`
	DateField    string   `json:"datetime_field_name"`
	UniqueFields []string `json:"patition_unique_field"`
}

type Struct_getPartitionTable struct {
	Result string   `json:"result"` // example:`OK`
	Data   []string `json:"data"`   // example:`["air","canal_waterlevel","dam_daily","dam_hourly","drought_area","flood_situation","floodforecast"]` พาร์ทิชั่น
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/partition_table
// @Summary			รายชื่อพาร์ทิชั่นทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getPartitionTable successful operation
func (srv *HttpService) getPartitionTable(ctx service.RequestContext) error {
	//Get all table name from public
	t, err := model.ListPartitionTable()
	if err != nil {
		return errors.Repack(err)
	}

	ctx.ReplyJSON(result.Result1(t))
	return nil
}

type Param_getPartitionHistory struct {
	TableName string `json:"table_name"` // example:air ชื่อตาราง
	Year      string `json:"year"`       // example: 2006 ปี
}

type Struct_getPartitionHistory struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_history.Struct_DBAModuleHistory `json:"data"`   // ประวัติการสร้างพาร์ทิชั่น
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/partition_history
// @Summary			ประวัติการสร้างพาร์ทิชั่น
// @Method			GET
// @Parameter		-	query	Param_getPartitionHistory
// @Produces		json
// @Response		200	Struct_getPartitionHistory successful operation
func (srv *HttpService) getPartitionHistory(ctx service.RequestContext) error {

	//Map parameters
	param := &model_history.Struct_DBAModuleHistory_InputParam{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Get List of History
	resultData, err := model_history.GetDBAModuleHistory(param)
	if err != nil {
		return errors.Repack(err)
	}

	ctx.ReplyJSON(result.Result1(resultData))
	return nil
}

type Struct_postPartition struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Create Successful`
}

type Param_postPartition struct {
	TableName string `json:"table_name"` // example:`air` ชื่อตาราง
	Year      string `json:"year"`       // example:`2006` ปี
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/partition
// @Summary			สร้างพาร์ทิชั่น
// @Method			POST
// @Consumes		json
// @Parameter		-	body Param_postPartition
// @Produces		json
// @Response		200	Struct_postPartition successful operation
func (srv *HttpService) postPartition(ctx service.RequestContext) error {
	//Map parameters
	param := &model.Struct_PartitionTable_InputParam{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Get List of History
	resultData, err := model.PostPartition(ctx.GetUserID(), param)
	if err != nil {
		return errors.Repack(err)
	}
	ctx.ReplyJSON(result.Result1(resultData))
	return nil
}

type Struct_deletePartition struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

type Param_deletePartition struct {
	TableName string `json:"table_name"` // ชื่อตาราง example air
	Year      string `json:"year"`       // ปี example 2006
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/dba/partition
// @Summary			ลบพาร์ทิชั่น
// @Method			DELETE
// @Parameter		-	query Param_deletePartition
// @Produces		json
// @Response		200	Struct_deletePartition successful operation
func (srv *HttpService) deletePartition(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_PartitionTable_InputParam{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Get List of History
	resultData, err := model.DeletePartitionTable(ctx.GetUserID(), param)
	if err != nil {
		return errors.Repack(err)
	}

	ctx.ReplyJSON(result.Result1(resultData))
	return nil
}

func postYearlyPartition(userID int64) error {
	//Get Table Data
	//	var mapTableName map[string]interface{}
	//	err := json.Unmarshal(model_setting.GetSystemSettingJSON("bof.DBA.Partition.TableList"), &mapTableName)
	//	if err != nil {
	//		return errors.Repack(err)
	//	}

	t, err := model.ListPartitionTable()
	if err != nil {
		return errors.Repack(err)
	}

	//Post Partition by List of Table
	for _, strTableName := range t {
		param := &model.Struct_PartitionTable_InputParam{}
		param.TableName = strTableName
		param.Year = strconv.Itoa(time.Now().Year() + 1)
		_, err := model.PostPartition(userID, param)
		if err != nil {
			return errors.Repack(err)
		}
	}

	//ctx.ReplyJSON(result.Result1("Created Partition Successful."))
	return nil
}
