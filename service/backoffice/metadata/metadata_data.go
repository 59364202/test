package metadata

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_dataformat "haii.or.th/api/thaiwater30/model/dataformat"
	model_hydroinfo "haii.or.th/api/thaiwater30/model/hydroinfo"
	model_category "haii.or.th/api/thaiwater30/model/lt_category"
	model_dataunit "haii.or.th/api/thaiwater30/model/lt_dataunit"
	model_frequencyunit "haii.or.th/api/thaiwater30/model/lt_frequencyunit"
	model_servicemethod "haii.or.th/api/thaiwater30/model/lt_servicemethod"
	model_subcategory "haii.or.th/api/thaiwater30/model/lt_subcategory"
	model "haii.or.th/api/thaiwater30/model/metadata"
	model_metadata_history "haii.or.th/api/thaiwater30/model/metadata_history"
	model_metadata_method "haii.or.th/api/thaiwater30/model/metadata_method"
	model_metadata_status "haii.or.th/api/thaiwater30/model/metadata_status"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_FncGetMetadata struct {
	Data    []*model.Struct_Metadata_Data                    `json:"metadata"`
	Hydro   []int64                                          `json:"hydro"`
	History []*model_metadata_history.Struct_MetadataHistory `json:"history"`
}

type Struct_OnloadMetadata struct {
	Agency         *result.Result `json:"agency"`
	Category       *result.Result `json:"category"`
	Dataformat     *result.Result `json:"dataformat"`
	DataUnit       *result.Result `json:"dataunit"`
	FrequencyUnit  *result.Result `json:"frequencyunit"`
	Hydroinfo      *result.Result `json:"hydroinfo"`
	Metadata       *result.Result `json:"metadata"`
	MetadataMethod *result.Result `json:"metadata_method"`
	MetadataStatus *result.Result `json:"metadata_status"`
	ServiceMethod  *result.Result `json:"servicemethod"`
	SubCategory    *result.Result `json:"subcategory"`
}

type Struct_onLoadMetadata struct {
	Agency         *Struct_onLoadMetadata_Agency         `json:"agency"`          // หน่วยงาน
	Category       *Struct_onLoadMetadata_Category       `json:"category"`        // หมวดหมู่หลัก
	Dataformat     *Struct_onLoadMetadata_Dataformat     `json:"dataformat"`      // รูปแบบ
	DataUnit       *Struct_onLoadMetadata_Dataunit       `json:"dataunit"`        // หน่วยของข้อมูล
	FrequencyUnit  *Struct_onLoadMetadata_FrequencyUnit  `json:"frequencyunit"`   // หน่วยของความถี่การเชื่อมโยง
	Hydroinfo      *Struct_onLoadMetadata_Hydroinfo      `json:"hydroinfo"`       // กลุ่มข้อมูลด้านน้ำและภูมิอากาศ
	Metadata       *Struct_onLoadMetadata_Metadata       `json:"metadata"`        // บัญชีข้อมูล
	MetadataMethod *Struct_onLoadMetadata_MetadataMethod `json:"metadata_method"` // วิธีการได้มาซึ่งข้อมูล
	MetadataStatus *Struct_onLoadMetadata_MetadataStatus `json:"metadata_status"` // สถานะการเชื่อมโยงข้อมูล
	ServiceMethod  *Struct_onLoadMetadata_ServiceMethod  `json:"servicemethod"`   // การบริการข้อมูล
	SubCategory    *Struct_onLoadMetadata_SubCategory    `json:"subcategory"`     // หมวดหมู่ย่อย
}
type Struct_onLoadMetadata_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}
type Struct_onLoadMetadata_Category struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*model_category.Struct_category `json:"data"`   // หมวดหมู่หลัก
}
type Struct_onLoadMetadata_Dataformat struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model_dataformat.Dataformat_struct `json:"data"`   // รูปแบบ
}
type Struct_onLoadMetadata_Dataunit struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*model_dataunit.Dataunit_struct `json:"data"`   // หน่วยของข้อมูล
}
type Struct_onLoadMetadata_FrequencyUnit struct {
	Result string                                      `json:"result"` // example:`OK`
	Data   []*model_frequencyunit.FrequencyUnit_struct `json:"data"`   // หน่วยของความถี่การเชื่อมโยง
}
type Struct_onLoadMetadata_Hydroinfo struct {
	Result string                              `json:"result"` // example:`OK`
	Data   []*model_hydroinfo.Struct_Hydroinfo `json:"data"`   // กลุ่มข้อมูลด้านน้ำและภูมิอากาศ
}
type Struct_onLoadMetadata_Metadata struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Struct_Metadata `json:"data"`   // บัญชีข้อมูล
}
type Struct_onLoadMetadata_MetadataMethod struct {
	Result string                                        `json:"result"` // example:`OK`
	Data   []*model_metadata_method.MetadataMethodParams `json:"data"`   // วิธีการได้มาซึ่งข้อมูล
}
type Struct_onLoadMetadata_MetadataStatus struct {
	Result string                                        `json:"result"` // example:`OK`
	Data   []*model_metadata_status.MetadataStatusParams `json:"data"`   // สถานะการเชื่อมโยงข้อมูล
}
type Struct_onLoadMetadata_ServiceMethod struct {
	Result string                                        `json:"result"` // example:`OK`
	Data   []*model_servicemethod.Struct_LtServicemethod `json:"data"`   // การบริการข้อมูล
}
type Struct_onLoadMetadata_SubCategory struct {
	Result string                                  `json:"result"` // example:`OK`
	Data   []*model_subcategory.SubCategory_struct `json:"data"`   // หมวดหมู่ย่อย
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_load
// @Summary			เริ่มต้นหน้าบัญชีข้อมูล
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadMetadata successful operation
func (srv *HttpService) onLoadMetadata(ctx service.RequestContext) error {
	//Set result's variable
	objResultData := &Struct_OnloadMetadata{}

	//Get agency
	resultAgency, err := model_agency.GetAllAgency()
	if err != nil {
		objResultData.Agency = result.Result0(err)
	}
	objResultData.Agency = result.Result1(resultAgency)

	//Get hydroinfo
	resultHydroinfo, err := model_hydroinfo.GetAllHydroinfo()
	if err != nil {
		objResultData.Hydroinfo = result.Result0(err)
	}
	objResultData.Hydroinfo = result.Result1(resultHydroinfo)

	//Get category
	resultCategory, err := model_category.GetAllCategory()
	if err != nil {
		objResultData.Category = result.Result0(err)
	}
	objResultData.Category = result.Result1(resultCategory)

	//Get subcategory
	resultSubCategory, err := model_subcategory.GetSubCategory([]int{})
	if err != nil {
		objResultData.SubCategory = result.Result0(err)
	}
	objResultData.SubCategory = result.Result1(resultSubCategory)

	//Get MetadataMethod
	resultMetadataMethod, err := model_metadata_method.GetMetadataMethod()
	if err != nil {
		objResultData.MetadataMethod = result.Result0(err)
	}
	objResultData.MetadataMethod = result.Result1(resultMetadataMethod)

	//Get Dataformat
	resultDataformat, err := model_dataformat.GetDataformat("", []int64{})
	if err != nil {
		objResultData.Dataformat = result.Result0(err)
	}
	objResultData.Dataformat = result.Result1(resultDataformat)

	//Get MetadataStatus
	resultMetadataStatus, err := model_metadata_status.GetMetadataStatus()
	if err != nil {
		objResultData.MetadataStatus = result.Result0(err)
	}
	objResultData.MetadataStatus = result.Result1(resultMetadataStatus)

	//Get ServiceMethod
	resultServiceMethod, err := model_servicemethod.GetAllServiceMethod()
	if err != nil {
		objResultData.ServiceMethod = result.Result0(err)
	}
	objResultData.ServiceMethod = result.Result1(resultServiceMethod)

	//Get FrequencyUnit
	resultFrequencyUnit, err := model_frequencyunit.GetFrequencyUnit("")
	if err != nil {
		objResultData.FrequencyUnit = result.Result0(err)
	}
	objResultData.FrequencyUnit = result.Result1(resultFrequencyUnit)

	//Get Dataunit
	resultDataUnit, err := model_dataunit.GetDataunit("")
	if err != nil {
		objResultData.DataUnit = result.Result0(err)
	}
	objResultData.DataUnit = result.Result1(resultDataUnit)

	//Get metadata
	param := &model.Struct_Metadata_Table_InputParam{}
	resultMetadata, err := model.GetMetadataTable(param)
	if err != nil {
		objResultData.Metadata = result.Result0(err)
	}
	objResultData.Metadata = result.Result1(resultMetadata)

	//Return data
	ctx.ReplyJSON(objResultData)
	return nil
}

type Struct_getMetadataTable struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Struct_Metadata `json:"data"`   // บัญชีข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_table
// @Summary			บัญชีข้อมูลทั้งหมด
// @Method			GET
// @Parameter		subcategory_id	query	[]int64 required:false	รหัสหมวดหมู่ย่อย example 33
// @Parameter		agency_id	query	[]int64 required:false	รหัสหน่วยงาน example 12
// @Parameter		hydroinfo_id	query	[]int64 required:false	รหัสกลุ่มข้อมูลด้านน้ำและภูมิอากาศ example 2
// @Parameter		category_id	query	int64 required:false	รหัสหมวดหมู่หลัก example 3
// @Produces		json
// @Response		200	Struct_getMetadataTable successful operation
func (srv *HttpService) getMetadataTable(ctx service.RequestContext) error {
	p := &model.Struct_Metadata_Table_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	resultData, err := model.GetMetadataTable(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}

	return nil
}

type Struct_getMetadata struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model.Struct_Metadata_Data `json:"data"`   // บัญชีข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata
// @Summary			รายละเอียดบัญชีข้อมูล
// @Method			GET
// @Parameter		metadata_id	query	string 	example:`6ISkJmlyE41mdGl574O2j64ockelAp8NjbNPg3Q76FYeMKDpwO4iZpiXjjscylEBbJpnCVHbIRwzdDzI3Wecsg` รหัสบัญชีข้อมูลแบบเข้ารหัส
// @Produces		json
// @Response		200	Struct_getMetadata successful operation
func (srv *HttpService) getMetadata(ctx service.RequestContext) error {

	//Map input param
	p := &model.Struct_Metadata_Data_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get metadata
	resultMeatadata, err := model.GetMetadata([]string{p.MetadataID})
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(resultMeatadata))
	return nil
}

type Struct_postMetadata struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Insert Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata
// @Summary			เพิ่มบัญชีข้อมูล
// @Method			POST
// @Consumes		json
// @Parameter		-	body	model.Struct_Metadata_Data_Post_InputParam
// @Produces		json
// @Response		200	Struct_postMetadata successful operation
func (srv *HttpService) postMetadata(ctx service.RequestContext) error {
	//Map input param
	p := &model.Struct_Metadata_Data_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	resultData, err := model.PostMetadata(p, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}
	return nil
}

type Struct_putMetadata struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata
// @Summary			แก่ไขบัญชีข้อมูล
// @Method			PUT
// @Consumes		json
// @Parameter		-	body	model.Struct_Metadata_Data_InputParam
// @Produces		json
// @Response		200	Struct_putMetadata successful operation
func (srv *HttpService) putMetadata(ctx service.RequestContext) error {
	//Map input param
	p := &model.Struct_Metadata_Data_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	resultData, err := model.PutMetadata(p, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}
	return nil
}

type Struct_putMetadataOfflineDate struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/offlinedate
// @Summary			แก่ไขบัญชีข้อมูล
// @Method			PUT
// @Parameter		metadata_id	form	string	รหัสบัญชีข้อมูล example 1
// @Produces		json
// @Response		200	Struct_putMetadataOfflineDate successful operation
func (srv *HttpService) putMetadataOfflineDate(ctx service.RequestContext) error {
	//Map input param
	p := &model.MetadataOfflineDate{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	err := model.PutMetadataOfflineDate(p.MetadataID, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1("Update Successful"))
	}
	return nil
}

type Struct_deleteMetadata struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata
// @Summary			ลบบัญชีข้อมูล
// @Method			DELETE
// @Parameter		metadata_id	query	string	รหัสบัญชีข้อมูล example 1
// @Produces		json
// @Response		200	Struct_deleteMetadata successful operation
func (srv *HttpService) deleteMetadata(ctx service.RequestContext) error {
	//Map input param
	p := &model.Struct_Metadata_Data_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	resultData, err := model.DeleteMetadata(p.MetadataID, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}
	return nil
}
