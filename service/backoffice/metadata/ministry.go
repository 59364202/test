package metadata

import (
	model "haii.or.th/api/thaiwater30/model/lt_ministry"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	"encoding/json"
)

// Ministry //
type ministryParam struct {
	Code      string          `json:"ministry_code"`      // example:`1000` code
	Text      json.RawMessage `json:"ministry_name"`      // example:`{"th":"สำนักนายกรัฐมนตรี", "en":"Prime Minister's Office"}` ชื่อหน่วยงานระดับกระทรวง
	ShortText json.RawMessage `json:"ministry_shortname"` // example:`{"th":"นร."}` ชื่อย่อหน่วยงานระดับกระทรวง
}

type Struct_getMinistry struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Ministry_struct `json:"data"`   // กระทรวง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/ministry
// @Summary			หน่วยงานระดับกระทรวงทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getMinistry successful operation
func (srv *HttpService) getMinistry(ctx service.RequestContext) error {
	rs, err := model.GetMinistry()
	if err != nil {
		return err
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_postMinistry struct {
	Result string                 `json:"result"` // example:`OK`
	Data   *model.Ministry_struct `json:"data"`   // กระทรวง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/ministry
// @Summary			เพิ่มหน่วยงานระดับกระทรวง
// @Method			POST
// @Consumes		json
// @Parameter		-	body	ministryParam
// @Produces		json
// @Response		200	Struct_postMinistry successful operation
func (srv *HttpService) postMinistry(ctx service.RequestContext) error {
	p := &ministryParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model.PostMinistry(p.Text, p.ShortText, p.Code, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putMinistry struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/ministry/{id}
// @Summary			แก้ไขหน่วยงานระดับกระทรวง
// @Method			PUT
// @Consumes		json
// @Parameter		id	path	string	รหัสหน่วยงานระดับกระทรวง example 1
// @Parameter		-	body	ministryParam
// @Produces		json
// @Response		200	Struct_putMinistry successful operation
func (srv *HttpService) putMinistry(ctx service.RequestContext) error {
	p := &ministryParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model.PutMinistry(ctx.GetUserID(), ctx.GetServiceParams("id"), p.Code, p.ShortText, p.Text)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteMinistry struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/ministry/{id}
// @Summary			ลบหน่วยงานระดับกระทรวง
// @Method			DELETE
// @Parameter		id	path	string	example:`1` รหัสหน่วยงานระดับกระทรวง example 1
// @Produces		json
// @Response		200	Struct_deleteMinistry successful operation
func (srv *HttpService) deleteMinistry(ctx service.RequestContext) error {
	rs, err := model.DeleteMinistry(ctx.GetUserID(), ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
