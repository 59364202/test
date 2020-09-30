package metadata

import (
	model "haii.or.th/api/thaiwater30/model/lt_department"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	"encoding/json"
)

// Department //
type Param_getDepartment struct {
	Ministry []int `json:"ministry"` // example:`1` รหัสกระทรวง
}

type Param_postDepartment struct {
	Text       json.RawMessage `json:"department_name"`      // example:`{"th":"สำนักงานปลัดสำนักนายกรัฐมนตรี"}` ชื่อ
	ShortText  json.RawMessage `json:"department_shortname"` // example:`{"th":""}` ชื่อย่อ
	Code       string          `json:"department_code"`      // example:`1001` รหัสกรม
	MinistryId string          `json:"ministry_id"`          // example:`1` รหัสกระทรวง
}

type Struct_getDepartment struct {
	Result string                     `json:"result"` // example:`OK`
	Data   []*model.Department_struct `json:"data"`   // กรมทั้งหมด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/department
// @Summary			หน่วยงานระดับกรมทั้งหมด
// @Method			GET
// @Parameter		-	query	Param_getDepartment
// @Produces		json
// @Response		200	Struct_getDepartment successful operation
func (srv *HttpService) getDepartment(ctx service.RequestContext) error {
	p := &Param_getDepartment{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model.GetDepartment(p.Ministry)
	if err != nil {
		return err
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_postDepartment struct {
	Result string                   `json:"result"` // example:`OK`
	Data   *model.Department_struct `json:"data"`   // กรม
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/department
// @Summary			เพิ่มหน่วยงานระดับกรม
// @Method			POST
// @Consumes		json
// @Parameter		-	body	Param_postDepartment
// @Produces		json
// @Response		200	Struct_postDepartment successful operation
func (srv *HttpService) postDepartment(ctx service.RequestContext) error {
	p := &Param_postDepartment{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model.PostDepartment(ctx.GetUserID(), p.Code, p.MinistryId, p.Text, p.ShortText)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putDepartment struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/department/{id}
// @Summary			แก้ไขหน่วยงานระดับกรม
// @Method			PUT
// @Consumes		json
// @Parameter		id	path	string	 รหัสหน่วยงานระดับกรม example 1
// @Parameter		-	body	Param_postDepartment
// @Produces		json
// @Response		200	Struct_putDepartment successful operation
func (srv *HttpService) putDepartment(ctx service.RequestContext) error {
	p := &Param_postDepartment{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	rs, err := model.PutDepartment(ctx.GetServiceParams("id"), p.Code, ctx.GetUserID(), p.MinistryId, p.Text, p.ShortText)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_deleteDepartment struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/department/{id}
// @Summary			ลบหน่วยงานระดับกรม
// @Method			DELETE
// @Parameter		id	path	string	 รหัสหน่วยงานระดับกรม example 1
// @Produces		json
// @Response		200	Struct_deleteDepartment successful operation
func (srv *HttpService) deleteDepartment(ctx service.RequestContext) error {
	rs, err := model.DeleteDepartment(ctx.GetServiceParams("id"), ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
