package data_service

import (
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

type Struct_getApprove struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // คำขอที่ยังไม่ได้รับผลคำขอจากหน่วยงานเฉพาะอันที่อัพโหลดไฟล์ผลคำขอแล้ว
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/approve
// @Summary			คำขอที่ยังไม่ได้รับผลคำขอ
// @Description		คำขอที่ยังไม่ได้รับผลคำขอจากหน่วยงานเฉพาะอันที่อัพโหลดไฟล์ผลคำขอแล้ว
// @Method			GET
// @Parameter		agency_id query int64 required:false รหัสหน่วยงาน example 1
// @Parameter		detail_letterno query string required:false เลขที่เอกสารร้องขอข้อมูล example AB322
// @Produces		json
// @Response		200	Struct_getApprove successful operation
func (srv *HttpService) getApprove(ctx service.RequestContext) error {
	p := &model_order_detail.Param_OrderApprove{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	p.Service_Id = 3
	rs, err := model_order_detail.GetOrderDetailByStatus(p, ctx)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

type Struct_putApprove struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/approve
// @Summary			บันทึกผลคำขอ
// @Consumes		json
// @Method			PUT
// @Parameter		- body model_order_detail.Param_OrderApprove
// @Produces		json
// @Response		200	Struct_putApprove successful operation
func (srv *HttpService) putApprove(ctx service.RequestContext) error {
	p := []*model_order_detail.Pram_OrderApprove_Put{}
	if err := ctx.GetRequestParams(&p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_order_detail.UpdateSourceResult(p, ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1("OK"))
	}

	return nil
}
