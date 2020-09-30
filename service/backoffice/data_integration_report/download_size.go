package data_integration_report

import (
	model "haii.or.th/api/thaiwater30/model/dataimport_download_log"
	//model_agency "haii.or.th/api/thaiwater30/model/agency"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Param_getMonthlyDownloadSize struct {
	AgencyID string `json:"agency_id"` // example:57 รหัสหน่วยงาน
	Month    string `json:"month"`     // example: 1 เดือน
	Year     string `json:"year"`      // example: 2006 ปี
}

type Struct_getMonthlyDownloadSize struct {
	Result string                      `json:"result"` // example:`OK`
	Data   []*model.Struct_DownloadLog `json:"data"`   // พื้นที่ที่ใช้จัดเก็บข้อมูลนำเข้า
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/download_size
// @Summary			พื้นที่ที่ใช้จัดเก็บข้อมูลนำเข้า
// @Method			GET
// @Parameter		- query Param_getMonthlyDownloadSize
// @Produces		json
// @Response		200	Struct_getMonthlyDownloadSize successful operation
func (srv *HttpService) getMonthlyDownloadSize(ctx service.RequestContext) error {

	//Map parameters
	param := &Param_getMonthlyDownloadSize{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	p := &model.Struct_DownloadLog_Inputparam{AgencyID: param.AgencyID, Month: param.Month, Year: param.Year}

	//Get List of Over All
	dataResult, err := model.GetDownloadSizeByAgency(p)
	if err != nil {
		//return errors.Repack(err)
		ctx.ReplyError(err)
	} else {
		//Return Data
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
