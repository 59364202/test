package agency

import (
	"haii.or.th/api/thaiwater30/util/result"
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	//	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model "haii.or.th/api/thaiwater30/model/dataimport_download_log"
)

type Struct_Agency_Summary struct {
	Result string                             `json:"result"` // example:`OK`
	Data   *Struct_Agency_ShoppingDetail_Data `json:"data"`   // สรุปข้อมูลหน่วยงาน
}

type Struct_getAgencyMetadataSummary_Result struct {
	Online  []*model.Struct_DownloadLog_Summary_Agency `json:"online"`  // online
	Offline []*model.Struct_DownloadLog_Summary_Agency `json:"offline"` // offline
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/agency/agency_summary
// @Summary 		ภาพรวมรายเดือนของคลัง
// @Description 	สรุปสัดส่วนการนำเข้าข้อมูล(%)แยกตามหน่วยงานของแต่ละเดือน
// @Parameter		month	query string required:true example:`06` เดือน
// @Parameter		year	query string required:true example:`2017` ปี
// @Method			GET
// @Produces		json
// @Response		200		Struct_Agency_Summary	successful operation
func (srv *HttpService) getAgencyMetadataSummary(ctx service.RequestContext) error {
	param := &model.Struct_DownloadLog_Inputparam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	rs := &Struct_getAgencyMetadataSummary_Result{}
	var err, err1, err2 error

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

	if err != nil {
		//return errors.Repack(err)
		ctx.ReplyError(err)
	} else {
		//Return Data
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
