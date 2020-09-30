package data_management

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_shared "haii.or.th/api/thaiwater30/model/shared"
	//model_air_station "haii.or.th/api/thaiwater30/model/air_station"
	//model_canal_station "haii.or.th/api/thaiwater30/model/canal_station"
	model_station "haii.or.th/api/thaiwater30/model/station"
	//model_setting "haii.or.th/api/server/model/setting"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	//"log"
)

type Struct_CheckMetadata_InputParam struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	AgencyID   string `json:"agency_id"`
}

type Struct_CheckMetadata struct {
	ColumnName []string                                         `json:"column_name"` // example:`["id", "station_oldcode", "station_name", "agency_name", "lat", "long", "geocode"]` หัวตาราง
	Data       []*model_station.Struct_Station_ForCheckMetadata `json:"data"`        // ข้อมูล
}
type Struct_onLoadCheckMetadata struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_shared.Struct_MetadataTable `json:"data"`   // ตาราง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/check_metadata_load
// @Summary			เริ่มต้นหน้า ตรวจสอบบัญชีข้อมูลพื้นฐาน
// @Description		รายชื่อตาราง
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadCheckMetadata successful operation
func (srv *HttpService) onLoadCheckMetadata(ctx service.RequestContext) error {

	//Get Data
	dataResult, err := model_shared.GetMatadataTable()
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Struct_getCheckMetadata struct {
	Result string                `json:"result"` // example:`OK`
	Data   *Struct_CheckMetadata `json:"data"`   // ตาราง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/check_metadata
// @Summary			ตรวจสอบบัญชีข้อมูลพื้นฐาน
// @Parameter		-	query	model_station.Struct_Station_InputParam
// @Method			GET
// @Produces		json
// @Response		200	Struct_getCheckMetadata successful operation
func (srv *HttpService) getCheckMetadata(ctx service.RequestContext) error {
	//Map parameters
	param := &model_station.Struct_Station_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get Data
	resultData := &Struct_CheckMetadata{}
	resultData.ColumnName = model_station.GetColumnForCheckMetadata(param.TableName)
	resultStation, err := model_station.GetStationInfoForCheckMetadata(param)
	if err != nil {
		return errors.Repack(err)
	}
	resultData.Data = resultStation

	//Return Data
	ctx.ReplyJSON(result.Result1(resultData))
	return nil
}

type Struct_getAgencyByTable struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/check_metadata_agency
// @Summary			หน่วยงาน ตาม ตารางตรวจสอบบัญชีข้อมูลพื้นฐาน
// @Parameter		table_name	query	string example: m_air_station ชื่อตาราง
// @Method			GET
// @Produces		json
// @Response		200	Struct_getAgencyByTable successful operation
func (srv *HttpService) getAgencyByTable(ctx service.RequestContext) error {
	//Map parameters
	p := &Struct_CheckMetadata_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	dataResult, err := model_agency.GetAgencyInTable(p.TableName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
