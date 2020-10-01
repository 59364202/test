package data_service

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
)

type Summarry struct {
	Order  *result.Result `json:"order,omitempty"`
	Agency *result.Result `json:"agency,omitempty"`
}

type Struct_getSummaryInit struct {
	Order  *Struct_getSummaryInit_Order  `json:"order"`  // รายละเอียดคำขอ
	Agency *Struct_getSummaryInit_Agency `json:"agency"` // หน่วยงานทั้งหมด
}
type Struct_getSummaryInit_Order struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // รายละเอียดคำขอ
}
type Struct_getSummaryInit_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงานทั้งหมด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/summaryInit
// @Summary			เริ่มต้นสรุปคำขอ
// @Method			GET
// @Produces		json
// @Response		200	Struct_getSummaryInit successful operation
func (srv *HttpService) getSummaryInit(ctx service.RequestContext) error {
	p := &model_order_detail.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	rs := &Summarry{}

	rs_order, err := model_order_detail.GetOrderDetailSummary(p)
	if err != nil {
		rs.Order = result.Result0(err)
	} else {
		rs.Order = result.Result1(rs_order)
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/summary
// @Summary			ดูสรุปคำขอ
// @Method			GET
// @Parameter		datestart	query	string	required:false	example:`2006-01-02` วันที่เริ่มต้น ไม่ใส่ คือ ไม่จำกัดวันเรื่มต้น
// @Parameter		dateend	query	string required:false	example:`2006-01-02` วันที่สิ้นสุด  ไม่ใส่ คือ ไม่จำกัดวันสิ้นสุด
// @Parameter		user_id	query	int required:false example:`68` รหัสผู้ใช้  0 คือทั้งหมด
// @Parameter		agency_id	query	int required:false example:`9` รหัสหน่วยงาน  0 คือทั้งหมด
// @Produces		json
// @Response		200	Struct_getSummaryInit_Order successful operation
func (srv *HttpService) getSummary(ctx service.RequestContext) error {
	p := &model_order_detail.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_order_detail.GetOrderDetailSummary(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/summary
// @Summary			อัพเดตวันหมดอายุ
// @Method			PUT
// @Parameter		id	query	int required:true example:`68` รหัส order_detail id
// @Parameter		expire_date	query	string	required:true	example:`1970-01-12 07:00:00+07` วันหมดอายุ
// @Response		200	Param_OrderExpireDate_Put successful operation
func (srv *HttpService) putExppiredate(ctx service.RequestContext) error {
	p := &model_order_detail.Param_OrderExpireDate_Put{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_order_detail.UpdateExpireDate(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1("OK"))
	}
	return nil
}
