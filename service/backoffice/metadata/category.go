package metadata

import (
	model "haii.or.th/api/thaiwater30/model/lt_category"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	"encoding/json"
)

// Category //

type Struct_getCategory struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Struct_category `json:"data"`   // หมวดข้อมูลหลักทั้งหมด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/category
// @Summary			หมวดข้อมูลหลักทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getCategory successful operation
func (srv *HttpService) getCategory(ctx service.RequestContext) error {
	rs, err := model.GetAllCategory()
	if err != nil {
		return err
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Param_postCategory struct {
	Text json.RawMessage `json:"category_name"` // example:`{"th":"อุตุนิยมวิทยา","en":"meteorology","jp":"気象学"}` ชื่อหมวดข้อมูลหลัก
}

type Struct_postCategory struct {
	Result string                 `json:"result"` // example:`OK`
	Data   *model.Struct_category `json:"data"`   // หมวดข้อมูลหลัก
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/category
// @Summary			เพิ่มหมวดข้อมูลหลัก
// @Method			POST
// @Consumes 		json
// @Parameter		-	body	Param_postCategory
// @Produces		json
// @Response		200	Struct_postCategory successful operation
func (srv *HttpService) postCategory(ctx service.RequestContext) error {
	p := &Param_postCategory{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	rs, err := model.PostCategory(ctx.GetUserID(), p.Text)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_putCategory struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/category/{id}
// @Summary			แก้ไขหมวดข้อมูลหลัก
// @Method			PUT
// @Consumes 		json
// @Parameter		id path string รหัสหมวดหมู่ example 1
// @Parameter		-	body	Param_postCategory
// @Produces		json
// @Response		200	Struct_putCategory successful operation
func (srv *HttpService) putCategory(ctx service.RequestContext) error {
	p := &Param_postCategory{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	rs, err := model.PutCategory(ctx.GetServiceParams("id"), ctx.GetUserID(), p.Text)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_deleteCategory struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/category/{id}
// @Summary			ลบหมวดข้อมูลหลัก
// @Method			DELETE 
// @Parameter		id path string รหัสหมวดข้อมูลหลัก example 1
// @Produces		json
// @Response		200	Struct_deleteCategory successful operation
func (srv *HttpService) deleteCategory(ctx service.RequestContext) error {
	rs, err := model.DeleteCategory(ctx.GetServiceParams("id"), ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
