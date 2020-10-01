package api

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/service"

	data "haii.or.th/api/server/model/eventlog"
	model_accessLog "haii.or.th/api/thaiwater30/model/accesslog"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
)

type mailData struct {
	Data []*model_order_detail.Struct_OrderDetail `json:"data"`
}
type Struct_MonitorApiService struct {
	OrderDetail *result.Result `json:"order_detail,omitempty"`
	Agency      *result.Result `json:"agency,omitempty"`
}

type Param_monitorApiService struct {
	Date_Start string `json:"datestart"` // required:false example:`2006-01-02` วันที่เริ่มต้น
	Date_End   string `json:"dateend"`   // required:false example:`2006-01-02` วันที่สิ้นสุด
	Agency_Id  int64  `json:"agency_id"` // required:false example:`9` รหัสหน่วยงาน
	User_Id    int64  `json:"user_id"`   // required:false example:`69` รหัสผู้ร้องขอ
}

type Struct_monitorApiService struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // api service ของการให้บริการข้อมูล
}

// @Service			thaiwater30/backoffice/api/monitor_api_service
// @Method			GET
// @Summary			monitor api service ของระบบให้บริการข้อมูล
// @Parameter		-	query Param_monitorApiService
// @Produces		json
// @Response		200	Struct_monitorApiService successful operation

type Struct_monitorApiService_ID struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_accessLog.ResultAccessLog `json:"data"`   // ประวัติการเข้าใช้ข้อมูล
}

// @Service			thaiwater30/backoffice/api/monitor_api_service?id={id}
// @Method			GET
// @Summary			ดูประวัติการเข้าใช้ข้อมูล
// @Parameter		id	path int64 example:`118` รหัส order detail
// @Produces		json
// @Response		200	Struct_monitorApiService_ID successful operation
func (srv *HttpService) monitorApiService(ctx service.RequestContext) error {
	p := &model_order_detail.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	if p.Id != 0 {
		rs, err := model_accessLog.GetOrderDetailLog(p.Id, p.Date_Start, p.Date_End)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		rs, err := model_order_detail.GetOrderDetail(p, ctx)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_monitorApiServiceOnload struct {
	Agency      *Struct_monitorApiServiceOnload_A `json:"agency"`       // หน่วยงานทั้งหมด
	OrderDetail *Struct_monitorApiService         `json:"order_detail"` // api service ของการให้บริการข้อมูล
}
type Struct_monitorApiServiceOnload_A struct {
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงานทั้งหมด
	Result string                        `json:"result"` // example:`OK`
}

// @Service			thaiwater30/backoffice/api/monitor_api_service_onload
// @Method			GET
// @Summary			เริ่มต้นหน้า monitor api service
// @Produces		json
// @Response		200	Struct_monitorApiServiceOnload successful operation
func (srv *HttpService) monitorApiServiceOnload(ctx service.RequestContext) error {
	rs := &Struct_MonitorApiService{}
	rs_order_detail, err := model_order_detail.GetOrderDetail(&model_order_detail.Param{}, ctx)
	if err != nil {
		rs.OrderDetail = result.Result0(err)
	} else {
		rs.OrderDetail = result.Result1(rs_order_detail)
	}
	rs_agency, err := model_agency.GetAllAgency()
	if err != nil {
		rs.Agency = result.Result0(err)
	} else {
		rs.Agency = result.Result1(rs_agency)
	}

	ctx.ReplyJSON(rs)
	return nil
}

type Struct_deleteMonitorApiService struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Disable Successful.`
}

// @Service			thaiwater30/backoffice/api/monitor_api_service
// @Method			DELETE
// @Summary			Disable service
// @Parameter		id	query int64  example:`118` รหัส order detail
// @Produces		json
// @Response		200	Struct_deleteMonitorApiService successful operation
// @Response		404			-		the request service name was not found
func (srv *HttpService) deleteMonitorApiService(ctx service.RequestContext) error {
	p := &model_order_detail.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	err := model_order_detail.DisableOrderDetail(p)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1("Disable Successful."))

	return nil
}

type Struct_patchMonitorApiService_eid struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`PaI_WOwIVzWLtNxAyB5fYVWT7RnDhF1iGMMysa9TU2di33YhsARIfSuI4FY_odv6UOyWxvLcyBgtmYyjNUpkXg`
}

// @Service			thaiwater30/backoffice/api/monitor_api_service?field=e_id
// @Method			PATCH
// @Summary			Re genarate eid
// @Description		Re genarate eid and send new eid to emil
// @Parameter		id	query int64 example:`118` รหัส order detail
// @Produces		json
// @Response		200	Struct_patchMonitorApiService_eid successful operation

type Struct_patchMonitorApiService_enable struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Enable Successful.`
}

// @Service			thaiwater30/backoffice/api/monitor_api_service?field=is_enabled
// @Method			PATCH
// @Summary			Enable service
// @Parameter		id	query int64  example:`118` รหัส order detail
// @Produces		json
// @Response		200	Struct_patchMonitorApiService_enable successful operation
func (srv *HttpService) patchMonitorApiService(ctx service.RequestContext) error {
	p := &model_order_detail.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	if p.Field == "e_id" {
		rs, err := model_order_detail.RegenerateKey(p.Id)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))

		mData := &mailData{}
		mData.Data, err = model_order_detail.GetOrderDetail(&model_order_detail.Param{Eid: rs}, ctx)
		if err != nil {
			return errors.Repack(err)
		}
		data.LogSystemEvent(ctx.GetServiceID(), ctx.GetAgentUserID(), mData.Data[0].User_Id, eventcode.EventDataServiceSendMailRegenerateKey, "PUT : regenerate key", mData)

	} else if p.Field == "is_enabled" {
		err := model_order_detail.EnableOrderDetail(p.Id)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1("Enable Successful."))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/agency_onload
// @Summary			รับค่าโลโก้ของหน่วยงานทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	successful operation
func (srv *HttpService) getAgencyLogo(ctx service.RequestContext) error {
	rs, err := model_agency.GetAgencyLogo()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	//Get Agency

	return nil
}
