package metadata

import (
	model "haii.or.th/api/thaiwater30/model/subcategory"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	"encoding/json"
)

// SubCategory //
type Param_getSubCategory struct {
	Category []int `json:"category"` // รหัสหมวดหมู่หลัก
}

type Struct_getSubCategory struct {
	Result string                      `json:"result"` // example:`OK`
	Data   []*model.SubCategory_struct `json:"data"`   // หมวดข้อมูลย่อย
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/subcategory
// @Summary			หมวดข้อมูลย่อยทั้งหมด
// @Method			GET
// @Parameter		category	query	[]int	required:false	example:`1` รหัสหมวดหมู่หลัก
// @Produces		json
// @Response		200	Struct_getSubCategory successful operation
func (srv *HttpService) getSubCategory(ctx service.RequestContext) error {
	p := &Param_getSubCategory{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model.GetSubCategory(p.Category)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Param_postSubCategory struct {
	Text       json.RawMessage `json:"subcategory_name"` // example:`{"th": "พายุ", "en": "Strom"}` ชื่อหมวดหมู่ย่อย
	CategoryId string          `json:"category_id"`      // example:`1` รหัสหมวดหมู่หลัก
}
type Struct_postSubCategory struct {
	Result string                   `json:"result"` // example:`OK`
	Data   model.SubCategory_struct `json:"data"`   // หมวดข้อมูลย่อย
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/subcategory
// @Summary			เพิ่มหมวดข้อมูลย่อย
// @Method			POST
// @Consumes		json
// @Parameter		-	body	Param_postSubCategory
// @Produces		json
// @Response		200	Struct_postSubCategory successful operation
func (srv *HttpService) postSubCategory(ctx service.RequestContext) error {
	p := &Param_postSubCategory{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	rs, err := model.PostSubCategory(ctx.GetUserID(), p.CategoryId, p.Text)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_putSubCategory struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/subcategory/{id}
// @Summary			แก้ไขหมวดข้อมูลย่อย
// @Method			PUT
// @Consumes		json
// @Parameter		id	path	string	example:`1` รหัสหมวดหมู่ย่อย example 1
// @Parameter		-	body	Param_postSubCategory
// @Produces		json
// @Response		200	Struct_putSubCategory successful operation
func (srv *HttpService) putSubCategory(ctx service.RequestContext) error {
	p := &Param_postSubCategory{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	rs, err := model.PutSubCategory(ctx.GetServiceParams("id"), ctx.GetUserID(), p.CategoryId, p.Text)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_deleteSubCategory struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/subcategory/{id}
// @Summary			ลบหมวดข้อมูลย่อย
// @Method			DELETE
// @Parameter		id	path	string	example:`1` รหัสหมวดหมู่ย่อย example 1
// @Produces		json
// @Response		200	Struct_deleteSubCategory successful operation
func (srv *HttpService) deleteSubCategory(ctx service.RequestContext) error {
	rs, err := model.DeleteSubCategory(ctx.GetServiceParams("id"), ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
