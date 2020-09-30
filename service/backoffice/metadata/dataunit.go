package metadata

import (
	"encoding/json"
	"fmt"
	model "haii.or.th/api/thaiwater30/model/lt_dataunit"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	"strconv"

	result "haii.or.th/api/thaiwater30/util/result"
)

// Data Unit //
//type DataUnitParam struct {
//	Id           string          `json:"id"`
//	DataunitName json.RawMessage `json:"dataunit_name,omitempty"`
//}

type Struct_getDataunit struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Dataunit_struct `json:"data"`   // หน่วยข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataunit
// @Summary			หน่วยข้อมูล
// @Method			GET
// @Produces		json
// @Response		200	Struct_getDataunit successful operation
func (srv *HttpService) getDataunit(ctx service.RequestContext) error {
	rs, err := model.GetDataunit("")
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Param_postDatainit struct {
	DataunitName json.RawMessage `json:"dataunit_name"` // example:`{"th":"กิโลเมตรต่อชั่วโมง","en":"Kilometer per hour","jp":""}` ชื่อหน่วยข้อมูล
}
type Struct_postDatainit struct {
	Result string                 `json:"result"` // example:`OK`
	Data   *model.Dataunit_struct `json:"data"`   // หน่วยข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataunit
// @Summary			เพิ่มหน่วยข้อมูล
// @Method			POST
// @Consumes		json
// @Parameter		-	body	Param_postDatainit
// @Produces		json
// @Response		200	Struct_postDatainit successful operation
func (srv *HttpService) postDatainit(ctx service.RequestContext) error {
	p := &Param_postDatainit{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	rs, err := model.PostDataunit(ctx.GetUserID(), p.DataunitName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataunit/{id}
// @Summary			แก้ไขหน่วยข้อมูล
// @Method			PUT
// @Consumes		json
// @Parameter		id	path	string	 รหัสหน่วยข้อมูล example 8
// @Parameter		-	body	Param_postDatainit
// @Produces		json
// @Response		200	Struct_postDatainit successful operation
func (srv *HttpService) putDataunit(ctx service.RequestContext) error {
	p := &Param_postDatainit{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get ID."))
	}

	dataunitId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}

	rs, err := model.PutDataunit(ctx.GetUserID(), dataunitId, p.DataunitName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataunit/{id}
// @Summary			ลบหน่วยข้อมูล
// @Method			DELETE
// @Parameter		id	path	string	 รหัสหน่วยข้อมูล example 8
// @Produces		json
// @Response		200	Struct_postDatainit successful operation
func (srv *HttpService) deleteDataunit(ctx service.RequestContext) error {

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get ID."))
	}

	dataunitId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}

	rs, err := model.DeleteDataunit(ctx.GetUserID(), dataunitId)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
