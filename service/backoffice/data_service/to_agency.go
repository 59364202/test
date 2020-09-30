package data_service

import (
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	model_order_header "haii.or.th/api/thaiwater30/model/order_header"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

type Struct_getToAgency_OrderHeader struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_header.Struct_OrderHeader `json:"data"`   // รายการตำขอสำหรับบุคคลภายนอกทั้งหมดที่ยังค้างคา
}
type Struct_getToAgency_OrderDetail struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // รายละเอียดคำขอ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/to_agency
// @Summary			คำขอข้อมูลสำหรับบุคคลภายนอก
// @Description		คำขอข้อมูลสำหรับบุคคลภายนอกที่ยังไม่ได้ดำเนินการ 
// @Parameter		date	query	string	required:false	example:`2006-01-02` วันที่คำขอ
// @Method			GET
// @Produces		json
// @Response		200	Struct_getToAgency_OrderHeader successful operation

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/to_agency?
// @Summary			รายละเอียดคำขอ
// @Parameter		order_header_id	query	int	example:`36` รหัสคำขอ
// @Method			GET
// @Produces		json
// @Response		200	Struct_getToAgency_OrderDetail successful operation
func (srv *HttpService) getToAgency(ctx service.RequestContext) error {
	p := &model_order_header.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	if p.Order_Header_Id != 0 {
		// มี order_header_id
		rs, err := model_order_detail.GetOrderDetailByOrderHeaderId(p.Order_Header_Id, ctx)
		if err != nil {
			ctx.ReplyError(err)
			return nil
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		// ไม่มี order_header_id โชว์ order header สำหรับบุคคลภายนอกทั้งหมดที่ยังค้างคา
		rs, err := model_order_header.GetOrderForExternal(p)
		if err != nil {
			ctx.ReplyError(err)
			return nil
		}
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/to_agency
// @Summary			บันทึกเอกสารคำขอหน่วยงาน
// @Consumes		json
// @Method			PUT
// @Parameter		- body model_order_detail.Param_OrderLetter
// @Produces		json
// @Response		200	Struct_deleteManagement	successful operation
func (srv *HttpService) putToAgency(ctx service.RequestContext) error {
	p := &model_order_detail.Param_OrderLetter{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	p.UserId = ctx.GetUserID()
	err := model_order_detail.UpdateLetterno(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1("OK"))
	}

	return nil
}
