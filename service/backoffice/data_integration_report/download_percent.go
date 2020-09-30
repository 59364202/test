package data_integration_report

import (
	model_setting "haii.or.th/api/server/model/setting"
	model "haii.or.th/api/thaiwater30/model/dataimport_download_log"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_OnLoadPercentDownload struct {
	DateRange           int64                                      `json:"date_range"`       // example:`30` ช่วงวันที่ที่เลือกได้
	PercentDownloadList []*model.Struct_DownloadLog_Summary_Agency `json:"percent_download"` // สัดส่วนการนำเข้าข้อมูล
}

type Param_onLoadPercentDownload struct {
	StartDate string `json:"start_date"` // example: 2006-01-02 วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // example: 2006-01-02 วันที่สิ้นสุด
}

type Struct_onLoadPercentDownload struct {
	Result string                        `json:"result"` // example:`OK`
	Data   *Struct_OnLoadPercentDownload `json:"data"`   // ช่วงวันที่, สัดส่วนการนำเข้าข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/download_percent_load
// @Summary			เริ่มต้นหน้าสัดส่วนการนำเข้าข้อมูล
// @Description     ช่วงวันที่, สัดส่วนการนำเข้าข้อมูล
// @Method			GET
// @Parameter		- query Param_onLoadPercentDownload
// @Produces		json
// @Response		200	Struct_onLoadPercentDownload successful operation
func (srv *HttpService) onLoadPercentDownload(ctx service.RequestContext) error {

	dataResult := &Struct_OnLoadPercentDownload{}

	//Get Date Range Setting
	dataResult.DateRange = model_setting.GetSystemSettingInt("bof.DataIntRpt.DwlPercent.DateRange")
	if (dataResult.DateRange) == 0 {
		dataResult.DateRange = model_setting.GetSystemSettingInt("setting.Default.DateRange")
	}

	//Map parameters
	param := &model.Struct_DownloadLog_Inputparam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of PercentDownload
	resultPercentDownload, err := model.GetPercentDownload(param)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.PercentDownloadList = resultPercentDownload

	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Struct_getPercentDownload struct {
	Result string                                     `json:"result"` // example:`OK`
	Data   []*model.Struct_DownloadLog_Summary_Agency `json:"data"`   // สัดส่วนการนำเข้าข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/download_percent
// @Summary			สัดส่วนการนำเข้าข้อมูล
// @Method			GET
// @Parameter		- query Param_onLoadPercentDownload
// @Produces		json
// @Response		200	Struct_getPercentDownload successful operation
func (srv *HttpService) getPercentDownload(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_DownloadLog_Inputparam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of PercentDownload
	dataResult, err := model.GetPercentDownload(param)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Param_getPercentDownloadDetail struct {
	Param_onLoadPercentDownload
	AgencyId string `json:"agency_id"` // example: 57 รหัสหน่วยงาน
}

type Struct_getPercentDownloadDetail struct {
	Result string                                       `json:"result"` // example:`OK`
	Data   []*model.Struct_DownloadLog_Summary_Metadata `json:"data"`   // สัดส่วนการนำเข้าข้อมูลแยกตามบัญชีข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/download_percent_detail
// @Summary			สัดส่วนการนำเข้าข้อมูลแยกตามบัญชีข้อมูล
// @Method			GET
// @Parameter		- query Param_getPercentDownloadDetail
// @Produces		json
// @Response		200	Struct_onLoadPercentDownload successful operation
func (srv *HttpService) getPercentDownloadDetail(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_DownloadLog_Inputparam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of PercentDownload
	dataResult, err := model.GetPercentDownloadDetail(param)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
