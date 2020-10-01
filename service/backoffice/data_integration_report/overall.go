package data_integration_report

import (
	"fmt"

	model "haii.or.th/api/thaiwater30/model/dataimport_download_log"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Param_getOverAllPercentDownload struct {
	Month string `json:"month"` // required:false example:1 เดือน
	Year  string `json:"year"`  // example:2006 ปี

}
type Struct_getOverAllPercentDownload struct {
	Result string                                     `json:"result"` // example:`OK`
	Data   []*model.Struct_DownloadLog_Summary_Agency `json:"data"`   // ภาพรวมของคลังฯ
}

type Struct_getOverAllPercentDownload_Result struct {
	Online  []*model.Struct_DownloadLog_Summary_Agency `json:"online"`  // online
	Offline []*model.Struct_DownloadLog_Summary_Agency `json:"offline"` // offline
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/overall
// @Summary 		ภาพรวมของคลังฯ
// @Method			GET
// @Parameter		- query model.Struct_DownloadLog_Inputparam{month,year,connection_format}
// @Produces		json
// @Response		200	Struct_getOverAllPercentDownload successful operation
func (srv *HttpService) getOverAllPercentDownload(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_DownloadLog_Inputparam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	rs := &Struct_getOverAllPercentDownload_Result{}
	var err, err1, err2 error

	switch param.ConnectionFormat {
	case "online":
		rs.Online, err = model.GetOverAllPercentDownload(param)
	case "offline":
		rs.Offline, err = model.GetOverAllPercentDownload(param)
	default:
		param.ConnectionFormat = "online"
		rs.Online, err1 = model.GetOverAllPercentDownload(param)
		param.ConnectionFormat = "offline"
		rs.Offline, err2 = model.GetOverAllPercentDownload(param)
		if err1 != nil {
			err = err1
		}
		if err2 != nil {
			err = err2
		}
	}

	//Get List of Over All
	//	dataResult, err := model.GetOverAllPercentDownload(param)
	if err != nil {
		//return errors.Repack(err)
		ctx.ReplyError(err)
	} else {
		//Return Data
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/overall
// @Summary 		ภาพรวมของคลังฯ
// @Method			GET
// @Parameter		- query model.Struct_DownloadLog_Inputparam{month,year,connection_format}
// @Produces		json
// @Response		200	Struct_getOverAllPercentDownload successful operation
func (srv *HttpService) getOverAllMultiYear(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_DownloadLog_Inputparam_2{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	tt := make([]interface{}, 0)

	var rs1 interface{}
	var rs2 interface{}
	var err error
	var year string

	fmt.Println("-- month --",param.Month)

	for _, str := range param.Year_arr {
		param.ConnectionFormat = "online"
		year = str
		rs1, err = model.GetOverAllMultipleYear(year,param.Month,param.ConnectionFormat)
		tt = append(tt, rs1)
		param.ConnectionFormat = "offline"
		year = str
		rs2, err = model.GetOverAllMultipleYear(year,param.Month,param.ConnectionFormat)
		tt = append(tt, rs2)
	}

	if err != nil {
		//return errors.Repack(err)
		ctx.ReplyError(err)
	} else {
		//Return Data
		ctx.ReplyJSON(result.Result1(tt))
	}

	return nil
}

type Param_getYearlyComparePercentDownload struct {
	Year     string `json:"year"`      // example: "2006,2007"  ปี
	AgencyId string `json:"agency_id"` // example: 57 รหัสหน่วยงาน
}
type Struct_getYearlyComparePercentDownload struct {
	Result string                                    `json:"result"` // example:`OK`
	Data   []*model.Struct_DownloadLog_YearlyCompare `json:"data"`   // เปรียบเทียบรายปีของหน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/compare_yearly
// @Summary			เปรียบเทียบรายปีของหน่วยงาน
// @Method			GET
// @Parameter		- query model.Struct_DownloadLog_Inputparam{year,agency_id}
// @Produces		json
// @Response		200	Struct_getYearlyComparePercentDownload successful operation
func (srv *HttpService) getYearlyComparePercentDownload(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_DownloadLog_Inputparam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Over All
	dataResult, err := model.PercentDownloadYearlyCompare(param)
	if err != nil {
		//return errors.Repack(err)
		ctx.ReplyError(err)
	} else {
		//Return Data
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
