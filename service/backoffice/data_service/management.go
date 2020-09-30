package data_service

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	//	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	model_order_header "haii.or.th/api/thaiwater30/model/order_header"
	model_order_status "haii.or.th/api/thaiwater30/model/order_status"
)

type Management struct {
	Order_header *result.Result `json:"order_header,omitempty"`
	Status       *result.Result `json:"order_status,omitempty"`
	Agency       *result.Result `json:"agency,omitempty"`
	Order_detail *result.Result `json:"order_detail,omitempty"`
}

type Struct_getManagementInit struct {
	Order_header *Struct_getManagementInit_OrderHeader `json:"order_header"`           // รายการคำขอ
	Status       *Struct_getManagementInit_Status      `json:"order_status,omitempty"` // สถานะทั้งหมด
}
type Struct_getManagementInit_OrderHeader struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_header.Struct_OrderHeader `json:"data"`   // รายการคำขอ
}
type Struct_getManagementInit_Status struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_status.Struct_OrderStatus `json:"data"`   // สถานะทั้งหมด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/managementInit
// @Summary			เริ่มต้นการจัดการคำขอ
// @Method			GET
// @Produces		json
// @Response		200		Struct_getManagementInit successful operation
func (srv *HttpService) getManagementInit(ctx service.RequestContext) error {
	rs := &Management{}
	p := &model_order_header.Param{}
	rs_head, err := model_order_header.GetOrderHeaderByParamManagement(p)
	if err != nil {
		rs.Order_header = result.Result0(err)
	} else {
		rs.Order_header = result.Result1(rs_head)
	}

	rs_status, err := model_order_status.GetOrderStatusAll()
	if err != nil {
		rs.Status = result.Result0(err)
	} else {
		rs.Status = result.Result1(rs_status)
	}

	//	rs_agency, err := model_agency.GetAllAgency()
	//	if err != nil {
	//		rs.Agency = result.Result0(err)
	//	} else {
	//		rs.Agency = result.Result1(rs_agency)
	//	}
	ctx.ReplyJSON(rs)
	return nil
}

type Struct_getManagement struct {
	Result       string                                   `json:"result"`                 // example:`OK`
	Order_detail []*model_order_detail.Struct_OrderDetail `json:"order_detail,omitempty"` // รายละเอียดคำขอ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/management
// @Summary			ดูคำขอ
// @Method			GET
// @Parameter		datestart	query	string	required:false	example:`2006-01-02` วันที่เริ่มต้น ไม่ใส่ คือ ไม่จำกัดวันเรื่มต้น
// @Parameter		dateend	query	string required:false	example:`2006-01-02` วันที่สิ้นสุด  ไม่ใส่ คือ ไม่จำกัดวันสิ้นสุด
// @Parameter		user_id	query	int required:false example:`68` รหัสผู้ใช้  0 คือทั้งหมด
// @Parameter		agency_id	query	int required:false example:`9` รหัสหน่วยงาน  0 คือทั้งหมด
// @Produces		json
// @Response		200		Struct_getManagementInit_OrderHeader successful operation

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/management?
// @Summary			ดูรายละเอียดคำขอ
// @Method			GET
// @Parameter		order_header_id	query int example:`68` รหัสรายการคำขอ
// @Produces		json
// @Response		200		Struct_getManagement successful operation
func (srv *HttpService) getManagement(ctx service.RequestContext) error {
	p := &model_order_header.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs := &Management{}

	if p.Order_Header_Id == 0 {
		rs_head, err := model_order_header.GetOrderHeaderByParamManagement(p)
		if err != nil {
			rs.Order_header = result.Result0(err)
		} else {
			rs.Order_header = result.Result1(rs_head)
		}
	} else {
		rs_detail, err := model_order_detail.GetOrderDetailByOrderHeaderId(p.Order_Header_Id, ctx)
		if err != nil {
			rs.Order_detail = result.Result0(err)
		} else {
			rs.Order_detail = result.Result1(rs_detail)
		}
	}
	ctx.ReplyJSON(rs)
	return nil
}

type Struct_deleteManagement struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/management
// @Summary			ไม่อนุมัติคำขอ
// @Method			DELETE
// @Parameter		Order_Header_Id	query int min:1 example: 68  รหัสรายการคำขอ
// @Produces		json
// @Response		200		Struct_deleteManagement successful operation
func (srv *HttpService) deleteManagement(ctx service.RequestContext) error {
	p := &model_order_header.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	if p.Order_Header_Id != 0 {
		err := model_order_header.DeleteOrderHeaderById(p.Order_Header_Id, ctx.GetUserID())
		if err != nil {
			ctx.ReplyJSON(result.Result0(err.Error()))
		} else {
			ctx.ReplyJSON(result.Result1("OK"))
		}
	}

	return nil
}
