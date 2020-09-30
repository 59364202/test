package metadata

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_category "haii.or.th/api/thaiwater30/model/lt_category"
	model "haii.or.th/api/thaiwater30/model/metadata"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_OnLoadSummaryMetadata struct {
	Summary_Metadata []*model.Struct_MetadataSummary   `json:"metadata"`
	Agency           []*model_agency.Struct_Agency     `json:"agency"`
	Category         []*model_category.Struct_category `json:"category"`
}

type Struct_onLoadSummaryMetadata struct {
	Summary_Metadata *Struct_onLoadSummaryMetadata_Summary_Metadata `json:"metadata"`
	Agency           *Struct_onLoadSummaryMetadata_Agency           `json:"agency"`   // หน่วยงาน
	Category         *Struct_onLoadSummaryMetadata_Category         `json:"category"` // หมวดหมู่หลัก
}
type Struct_onLoadSummaryMetadata_Summary_Metadata struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model.Struct_MetadataSummary `json:"data"`   // บัญชีข้อมูลที่เชื่อมโยง
}
type Struct_onLoadSummaryMetadata_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}
type Struct_onLoadSummaryMetadata_Category struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*model_category.Struct_category `json:"data"`   // หมวดหมู่หลัก
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/summary_load
// @Summary			เริ่มต้นหน้าบัญชีข้อมูลที่เชื่อมโยง
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadSummaryMetadata successful operation
func (srv *HttpService) onLoadSummaryMetadata(ctx service.RequestContext) error {

	dataResult := &Struct_OnLoadSummaryMetadata{}

	//Get List of Agency Data
	resultAgency, err := model_agency.GetAllAgency()
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.Agency = resultAgency

	//Get List of Category Data
	resultCategory, err := model_category.GetAllCategory()
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.Category = resultCategory

	//Get List of SummaryMetadata Data
	resultMetadata, err := model.GetSummaryMetadataGroupByAgency()
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.Summary_Metadata = resultMetadata

	ctx.ReplyJSON(result.Result1(dataResult))

	return nil
}

type Struct_getSummaryMetadata struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model.Struct_MetadataSummary `json:"data"`   // บัญชีข้อมูลที่เชื่อมโยง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/summary
// @Summary			บัญชีข้อมูลที่เชื่อมโยงทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getSummaryMetadata successful operation
func (srv *HttpService) getSummaryMetadata(ctx service.RequestContext) error {

	//Get List of SummaryMetadata Data
	dataResult, err := model.GetSummaryMetadataGroupByAgency()
	if err != nil {
		return errors.Repack(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Struct_getMetadataImported struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Struct_Metadata `json:"data"`   // บัญชีข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/summary_detail
// @Summary			รายละเอียดบัญชีข้อมูลที่เชื่อมโยงของหน่วยงาน
// @Method			GET
// @Parameter		-	query	model.Struct_Metadata_InputParam
// @Produces		json
// @Response		200	Struct_getMetadataImported successful operation
func (srv *HttpService) getMetadataImported(ctx service.RequestContext) error {
	//Map parameters
	param := &model.Struct_Metadata_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of SummaryMetadata Data
	dataResult, err := model.GetMetadataImported(param)

	if err != nil {
		return errors.Repack(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Struct_getSummaryMetadataImported struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model.Struct_MetadataSummary `json:"data"`   // รายละเอียดบัญชีข้อมูลที่เชื่อมโยง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/summary_by_agency
// @Summary			รายละเอียดบัญชีข้อมูลที่เชื่อมโยงทั้งหมด
// @Method			GET
// @Parameter		-	query	model.Struct_Metadata_InputParam{agency_id}
// @Produces		json
// @Response		200	Struct_getSummaryMetadataImported successful operation

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/summary_by_category
// @Summary			รายละเอียดบัญชีข้อมูลที่เชื่อมโยงทั้งหมด
// @Method			GET
// @Parameter		-	query	model.Struct_Metadata_InputParam{category_id}
// @Produces		json
// @Response		200	Struct_getSummaryMetadataImported successful operation
func (srv *HttpService) getSummaryMetadataImported(ctx service.RequestContext) error {
	//Map parameters
	param := &model.Struct_Metadata_InputParam{}

	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of SummaryMetadata Data
	dataResult, err := model.GetSummaryMetadataImported(param)

	if err != nil {
		return errors.Repack(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
