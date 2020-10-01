package data_service

import (
	//	model_lt_category "haii.or.th/api/thaiwater30/model/lt_category"
	//	model_lt_ministry "haii.or.th/api/thaiwater30/model/lt_ministry"
	"fmt"

	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	model_order_header "haii.or.th/api/thaiwater30/model/order_header"
	model_order_status "haii.or.th/api/thaiwater30/model/order_status"

	"haii.or.th/api/server/model/uac/getdata"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	"encoding/json"
	"time"
)

type Struct_Onload struct {
	Metadata *result.Result     `json:"metadata"`
	User     *result.ResultJson `json:"user"`
}

type Struct_Onload_Shopping struct {
	Metadata *Struct_Onload_Shopping_Metadata `json:"metadata"` // รายการบัญชีข้อมูล
	User     *Struct_Onload_Shopping_User     `json:"user"`     // รายชืื่อผู้ใช้
}
type Struct_Onload_Shopping_Metadata struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*model_metadata.Struct_Metadata `json:"data"`   // รายการบัญชีข้อมูล
}
type Struct_Onload_Shopping_User struct {
	Result string          `json:"result"` // example:`OK`
	Data   json.RawMessage `json:"data"`   // example:`{"id":68,"user_type_id":3,"account":"kantamat.cim@haii.or.th","full_name":"Kantamat Polsawang","is_active":true,"is_deleted":false,"account_expired_at":"0001-01-01T00:00:00Z","password_lifespan_days":0}` รายชืื่อผู้ใช้
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/onload
// @Summary			เริ่มต้นหน้าให้บริการข้อมูล
// @Description		รายชื่อ user และ บัญชีข้อมูลที่ให้บริการ
// @Method			GET
// @Produces		json
// @Response		200	Struct_Onload_Shopping successful operation
func (srv *HttpService) getOnloadShopping(ctx service.RequestContext) error {
	rs := &Struct_Onload{}
	p := &model_metadata.Param_Metadata{}

	rs_metadata, err := model_metadata.GetMetadataShoppingTable(p)
	if err != nil {
		rs.Metadata = result.Result0(err)
	} else {
		rs.Metadata = result.Result1(rs_metadata)
	}

	d, _, err := getdata.GetUserListJSON(0)
	if err == nil {
		rs.User = result.ResultJson1(d)
	}

	ctx.ReplyJSON(rs)

	return nil
}

type Struct_Shopping_Table struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*model_metadata.Struct_Metadata `json:"data"`   // รายการบัญชีข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/data_service
// @Summary			บัญชีข้อมูลที่ให้บริการ
// @Method			GET
// @Parameter		agency_id	query	int64	required:false example:`1` รหัสหน่วยงาน 0,ไม่ใส่ คือทั้งหมด
// @Produces		json
// @Response		200		Struct_Shopping_Table successful operation

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/data_service?
// @Summary			รายละเอียดของบัญชีข้อมูล
// @Method			GET
// @Parameter		id	query	int64 required:true example:`1`	รหัสบัญชีข้อมูล
// @Produces		json
// @Response		200		Struct_Shopping_Table successful operation
func (srv *HttpService) getShoppingTable(ctx service.RequestContext) error {
	//	rs, err := GetShoppingTable()
	p := &model_metadata.Param_Metadata{}
	if err := ctx.GetRequestParams(p); err != nil {
		return nil
	}
	ctx.LogRequestParams(p)
	var (
		rs  []*model_metadata.Struct_Metadata
		err error
	)
	if p.Id == 0 {
		rs, err = model_metadata.GetMetadataShoppingTable(p)
	} else {
		rs, err = model_metadata.GetMetadataShoppingDetail(p.Id)
	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_Shopping_Check struct {
	Result string `json:"result"` // example:`OK` เงื่อนไขที่เลือกผ่าน
	Data   string `json:"data"`   // example:`OK`
}
type lang struct {
	TH string `json:"th"`
	EN string `json:"en"`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/data_service_check
// @Summary			ตรวจสอบการขอบริการข้อมูล
// @Description		ตรวจสอบว่าถ้าเลือกการบริการเป็น download ต้องเลือกวันไม่เกิน 30 วัน, cd/dvdv ต้องเลือกวันไม่เกิน 6 เดือน
// @Method			GET
// @Parameter		- query model_order_detail.Param_OrderDetail
// @Produces		json
// @Response		200		Struct_Shopping_Check	successful operation
func (srv *HttpService) getCheckVerifyShopping(ctx service.RequestContext) error {
	p := &model_order_detail.Param_OrderDetail{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	if p.Service_Id == 2 || p.Service_Id == 3 {
		// 2: download ต้องไม่เกิน 30 วัน
		// 3: cd/dvd ต้องไม่เกิน 6 เดือน
		d, err := model_order_detail.GetMetadataByMetadata(p.Metadata_Id)
		if err != nil {
			ctx.ReplyJSON(result.Result0(err.Error()))
			return nil
		}
		var fromDateAddLimit time.Time
		var msg *lang
		from_date, err := time.Parse("2006-01-02", p.Detail_Fromdate)
		if err != nil {
			ctx.ReplyJSON(result.Result0(err.Error()))
			return nil
		}

		if p.Service_Id == 2 {
			fromDateAddLimit = from_date.AddDate(0, 0, 30)
			msg = &lang{"เกิน 30 วัน", "over 30 day"}
		} else if p.Service_Id == 3 {
			fromDateAddLimit = from_date.AddDate(0, 6, 0)
			msg = &lang{"เกิน 6 เดือน", "over 6 month"}
		}
		if !model_order_detail.IsMedia(d.Table_name.String) {

			to_date, err := time.Parse("2006-01-02", p.Detail_Todate)
			if err != nil {
				ctx.ReplyJSON(result.Result0(err.Error()))
				return nil
			}
			// if to_date.Sub(from_date).Hours()/24 > 30 {
			if !fromDateAddLimit.After(to_date) {
				ctx.ReplyJSON(result.Result0(msg))
				return nil
			} else {
				ctx.ReplyJSON(result.Result1("OK"))
				return nil
			}
		}
		ctx.ReplyJSON(result.Result1("OK"))
	} else {
		ctx.ReplyJSON(result.Result1("OK"))
	}
	return nil
}

type orderHistory struct {
	Status *result.Result `json:"status"`
	Order  *result.Result `json:"order"`
}
type Struct_OrderHistory struct {
	Status *Struct_OrderHistory_Status `json:"status"` // สถานะทั้งหมด
	Order  *Struct_OrderHistory_Order  `json:"order"`  // รายการที่เคยขอใช้บริการ
}
type Struct_OrderHistory_Status struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_status.Struct_OrderStatus `json:"data"`   // สถานะทั้งหมด
}
type Struct_OrderHistory_Order struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model_order_header.Struct_OH `json:"data"`   // รายการที่เคยขอใช้บริการ
}

type Struct_OrderHistory_Detail struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // รายละเอียดคำขอ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/history
// @Summary			ประวัติคำขอข้อมูล
// @Method			GET
// @Produces		json
// @Response		200		Struct_OrderHistory successful operation

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/history?
// @Summary			รายละเอียดคำขอข้อมูล
// @Parameter		id	query int required:true example:`45` orde detail id
// @Method			GET
// @Produces		json
// @Response		200		Struct_OrderHistory_Detail successful operation
func (srv *HttpService) getShoppingHistory(ctx service.RequestContext) error {
	setting.SetSystemDefault("api_service.media_url", ctx.BuildURL(0, "thaiwater30/shared/image", true))
	p := &model_order_header.Struct_OrderHeader{}
	if err := ctx.GetRequestParams(p); err != nil {
		return nil
	}
	ctx.LogRequestParams(p)

	fmt.Println("-- getShoppingHistory --", p.Id)

	if p.Id == 0 {
		rs := &orderHistory{}

		rs_order, err := model_order_header.GetOrderHeaderByUserId(ctx.GetUserID())
		if err != nil {
			rs.Order = result.Result0(err)
		} else {
			rs.Order = result.Result1(rs_order)
		}

		rs_status, err := model_order_status.GetOrderStatusAll()
		if err != nil {
			rs.Status = result.Result0(err)
		} else {
			rs.Status = result.Result1(rs_status)
		}
		ctx.ReplyJSON(rs)
		return nil
	}
	rs, err := model_order_detail.GetOrderDetailByOrderHeaderId(p.Id, ctx)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/data_purpose
// @Summary			วัตถุประสงค์
// @Parameter
// @Method			GET
// @Produces		json
// @Response		200 Struc_Order_Purpose successful operation
func (srv *HttpService) getDatapurpose(ctx service.RequestContext) error {
	rs, err := model_order_header.GetPopularOrderPurpose()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
