package metadata

import (
	//	"encoding/json"
	"strconv"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	model "haii.or.th/api/thaiwater30/model/lt_servicemethod"
)

type Struct_getServicemethod struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model.Struct_LtServicemethod `json:"data"`   // วีธีการให้บริการข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/servicemethod
// @Summary			วีธีการให้บริการข้อมูลทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getServicemethod successful operation
func (srv *HttpService) getServicemethod(ctx service.RequestContext) error {

	rs, err := model.GetAllServiceMethod()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_postServiceMethody struct {
	Result string                       `json:"result"` // example:`OK`
	Data   model.Struct_LtServicemethod `json:"data"`   // วีธีการให้บริการข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/servicemethod
// @Summary			เพิ่มวีธีการให้บริการข้อมูล
// @Method			POST
// @Consumes		json
// @Parameter		-	body	model.Param_LtServicemethod{servicemethod_name}
// @Produces		json
// @Response		200	Struct_postServiceMethody successful operation
func (srv *HttpService) postServiceMethod(ctx service.RequestContext) error {
	p := &model.Struct_LtServicemethod{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	p.UserId = ctx.GetUserID()

	rs, err := model.PostServiceMethod(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putServiceMethod struct {
	Result string                       `json:"result"` // example:`OK`
	Data   model.Struct_LtServicemethod `json:"data"`   // วีธีการให้บริการข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/servicemethod/{id}
// @Summary			แก้ไขวีธีการให้บริการข้อมูล
// @Method			PUT
// @Consumes		json
// @Parameter		id	path	string   example:`1` รหัสวิธีการให้บริการข้อมูล
// @Parameter		-	body	model.Param_LtServicemethod{servicemethod_name}
// @Produces		json
// @Response		200	Struct_putServiceMethod successful operation
func (srv *HttpService) putServiceMethod(ctx service.RequestContext) error {
	p := &model.Struct_LtServicemethod{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(rest.NewError(422, "Can not get ID.", nil))
	}

	serviceMethodId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}
	p.ID = serviceMethodId
	p.UserId = ctx.GetUserID()
	p.ServiceMethodID = ctx.GetServiceParams("id")

	rs, err := model.PutServiceMethod(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteServiceMethod struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/servicemethod/{id}
// @Summary			ลบวีธีการให้บริการข้อมูล
// @Method			DELETE
// @Parameter		id	path	string 	example:`1` รหัสวิธีการให้บริการข้อมูล
// @Produces		json
// @Response		200	Struct_deleteServiceMethod successful operation
func (srv *HttpService) deleteServiceMethod(ctx service.RequestContext) error {
	p := &model.Struct_LtServicemethod{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(rest.NewError(422, "Can not get ID.", nil))
	}

	serviceMethodId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}
	p.UserId = ctx.GetUserID()
	p.ID = serviceMethodId

	result, err := model.DeleteServiceMethod(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}
