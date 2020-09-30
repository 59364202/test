package agency

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	//	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
)

type Struct_Agency_Shopping struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_AgencyShopping `json:"data"`   // สรุปข้อมูลหน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/agency/agency_shopping
// @Summary 		สรุปข้อมูลหน่วยงานที่เชื่อมโยง
// @Description		สรุปข้อมูลหน่วยงานที่เชื่อมโยงโดยจะให้หน่วยงานที่ตัวเองสังกัดอยู่บนสุด
// @Method			GET
// @Produces		json
// @Response		200	Struct_Agency_Shopping successful operation
func (srv *HttpService) getAgencyShopping(ctx service.RequestContext) error {
	rs, err := model_agency.GetAgencyShopping(ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Param_Agency struct {
	Agency   int64 `json:"agency"`
	Metadata int64 `json:"metadata"`
	MAgency  int64 `json:"m_agency"`
}

type Struct_Agency_ShoppingDetail struct {
	Result string                            `json:"result"` // example:`OK`
	Data   Struct_Agency_ShoppingDetail_Data `json:"data"`   // สรุปข้อมูลหน่วยงาน
}
type Struct_Agency_ShoppingDetail_Data struct {
	Connect     []*model_metadata.Struct_M                              `json:"connect"`      // เชื่อมโยง
	WaitUpdate  []*model_metadata.Struct_M                              `json:"wait_update"`  // รอหน่วยงานปรับปรุงข้อมูล
	WaitConnect []*model_metadata.Struct_M                              `json:"wait_connect"` // รอการเชื่อมโยง
	Cancel      []*model_metadata.Struct_M                              `json:"cancel"`       // ยกเลิอก
	AgencyCount []*model_order_detail.Struct_CountOrderDetailByAgencyId `json:"agency_count"` // ขอใช้บริการข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/agency/agency_shopping_detail
// @Summary 		สรุปข้อมูลหน่วยงาน, จำนวนการที่ขอใช้บริการข้อมูล
// @Description		สรุปข้อมูลหน่วยงานแยกตามสถานะการเชื่อมโยง, จำนวนการที่ขอใช้บริการข้อมูล
// @Method			GET
// @Parameter		agency	query	int64 required:true example:`9`	รหัสหน่วยงานที่ต้องการดูรายละเอียด
// @Produces		json
// @Response		200		Struct_Agency_ShoppingDetail	successful operation
func (srv *HttpService) getAgencyShoppingDetail(ctx service.RequestContext) error {
	p := &Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	if p.Agency == 0 {
		ctx.ReplyError(rest.NewError(422, "no parameter", nil))
		return nil
	}

	rs := &Struct_Agency_ShoppingDetail_Data{}
	rs_connect, err := model_metadata.GetMetadataStatusConnect(p.Agency)
	if err == nil {
		rs.Connect = rs_connect
	}

	rs_waitupdate, err := model_metadata.GetMetadataStatusWaitUpdate(p.Agency)
	if err == nil {
		rs.WaitUpdate = rs_waitupdate
	}

	rs_waitconnect, err := model_metadata.GetMetadataStatusWaitConnect(p.Agency)
	if err == nil {
		rs.WaitConnect = rs_waitconnect
	}

	rs_cancel, err := model_metadata.GetMetadataStatusCancel(p.Agency)
	if err == nil {
		rs.Cancel = rs_cancel
	}

	rs_agencycount, err := model_order_detail.GetCountOrderDetailByAgencyId(p.Agency)
	if err == nil {
		rs.AgencyCount = rs_agencycount
	}

	ctx.ReplyJSON(rs)

	return nil
}

type Struct_Agency_Metadata struct {
	Result string                                        `json:"result"` // example:`OK`
	Data   *model_metadata.Struct_MetadataImportByAgency `json:"data"`   // ข้อมูลที่นำเข้าล่าสุด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/agency/agecncy_metadata
// @Summary 		ข้อมูลที่นำเข้าล่าสุด
// @Description		img, table, weather จะมีแค่อันใดอันหนึ่งเท่านั้น
// @Parameter		metadata	query int64	required:true example:`5` รหัสบัญชีข้อมูล
// @Method			GET
// @Produces		json
// @Response		200		Struct_Agency_Metadata	successful operation
func (srv *HttpService) getMetaData(ctx service.RequestContext) error {
	p := &Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	if p.Metadata == 0 {
		ctx.ReplyError(rest.NewError(422, "no parameter", nil))
		return nil
	}
	media_url := ctx.BuildURL(0, "thaiwater30/shared/file", true)
	rs, err := model_metadata.GetLastMetadataImportByAgency(p.Metadata, media_url)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_Agency_AgencyDetail struct {
	Result string                                  `json:"result"` // example:`OK`
	Data   []*model_metadata.Struct_MetadataStatus `json:"data"`   // ผู้ขอใช้บริการ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/agency/agecncy_detail
// @Summary 		รายละเอียดการขอใช้บริการข้อมูล
// @Description 	รายละเอียดการขอใช้บริการข้อมูล ดูตาม หน่วยงานที่ขอใช้บริการ, หน่วยงานของบัญชีข้อมูล
// @Parameter		agency	query int64	required:true example:`9` รหัสหน่วยงานที่ต้องการดูการขอใช้บริการ
// @Parameter		m_agency	query int64	required:true example:`9` รหัสหน่วยงานของบัญชีข้อมูลที่ให้บริการ
// @Method			GET
// @Produces		json
// @Response		200		Struct_Agency_AgencyDetail	successful operation
func (srv *HttpService) getOrderResult(ctx service.RequestContext) error {
	p := &Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	if p.Agency == 0 || p.MAgency == 0 {
		ctx.ReplyError(rest.NewError(422, "no parameter", nil))
		return nil
	}
	rs, err := model_metadata.GetOrderResult(p.Agency, p.MAgency)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
