package data_service

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	model_order_header "haii.or.th/api/thaiwater30/model/order_header"
)

type Struct_getPrint struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // ละเอียดคำขอ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/print
// @Summary			รายละเอียดคำขอแยกตามหน่วยงาน
// @Parameter		order_header_id	query	int	example:`36` รหัสคำขอ
// @Method			GET
// @Produces		json
// @Response		200	Struct_getPrint successful operation
// @Response		404 string mismatch param name
func (srv *HttpService) getPrint(ctx service.RequestContext) error {
	p := &model_order_header.Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	if p.Order_Header_Id != 0 {
		rs, err := model_order_detail.GetOrderDetailByOrderHeaderId(p.Order_Header_Id, ctx)
		if err != nil {
			ctx.ReplyError(err)
			return nil
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		return rest.NewError(404, "mismatch param name", nil)
	}

	return nil
}
