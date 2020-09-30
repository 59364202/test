package metadata

import (
	"encoding/json"
	model "haii.or.th/api/thaiwater30/model/dataformat"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// Dataformat //
type dataformatParam struct {
	Id              string          `json:"id"`
	DataformatName  json.RawMessage `json:"dataformat_name"`
	MetdataMethodID []int64         `json:"metadata_method_id"`
}

type dataformatParamPostPut struct {
	Id              string          `json:"id"`
	DataformatName  json.RawMessage `json:"dataformat_name"`
	MetdataMethodID int64           `json:"metadata_method_id"`
}
type Param_getDataformat struct {
	MetdataMethodID []int64 `json:"metadata_method_id"` // รหัสวิธีการได้มาซึ่งข้อมูล example 3
}
type Struct_getDataformat struct {
	Result string                     `json:"result"` // example:`OK`
	Data   []*model.Dataformat_struct `json:"data"`   // รูปแบบข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataformat
// @Summary			รูปแบบข้อมูล
// @Method			GET
// @Parameter		- query Param_getDataformat
// @Produces		json
// @Response		200	Struct_getDataformat successful operation
func (srv *HttpService) getDataformat(ctx service.RequestContext) error {
	//	//Map parameters
	p := &Param_getDataformat{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Dataformat
	rs, err := model.GetDataformat("", p.MetdataMethodID)

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Param_postDataformat struct {
	DataformatName  json.RawMessage `json:"dataformat_name"`    // example:`{"th":"Realtime","en":"","jp":""}` ชื่อรูปแบบข้อมูล
	MetdataMethodID int64           `json:"metadata_method_id"` // example:`3` รหัสวิธีการได้มาซึ่งข้อมูล
}

type Struct_postDataformat struct {
	Result string                   `json:"result"` // example:`OK`
	Data   *model.Dataformat_struct `json:"data"`   // รูปแบบข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataformat
// @Summary			รูปแบบข้อมูล
// @Method			POST
// @Consumes 		json
// @Parameter		- body Param_postDataformat
// @Produces		json
// @Response		200	Struct_postDataformat successful operation
func (srv *HttpService) postDataformat(ctx service.RequestContext) error {
	//Map parameters
	p := &Param_postDataformat{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Post Dataformat
	rs, err := model.PostDataformat(ctx.GetUserID(), p.DataformatName, p.MetdataMethodID)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataformat/{id}
// @Summary			รูปแบบข้อมูล
// @Method			PUT
// @Consumes 		json
// @Parameter		id	path	string รหัสรูปแบบข้อมูล example 3
// @Parameter		- body Param_postDataformat
// @Produces		json
// @Response		200	Struct_postDataformat successful operation
func (srv *HttpService) putDataformat(ctx service.RequestContext) error {
	//Map parameters
	p := &Param_postDataformat{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Put Dataformat
	rs, err := model.PutDataformat(ctx.GetUserID(), ctx.GetServiceParams("id"), p.DataformatName, p.MetdataMethodID)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteDataformat struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/dataformat/{id}
// @Summary			รูปแบบข้อมูล
// @Method			DELETE
// @Parameter		id	path	string	example:3 รหัสรูปแบบข้อมูล
// @Produces		json
// @Response		200	Struct_deleteDataformat successful operation
func (srv *HttpService) deleteDataformat(ctx service.RequestContext) error {

	//Delete Dataformat
	rs, err := model.DeleteDataformat(ctx.GetUserID(), ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_getSelectOption struct {
	Data   []*model.Struct_Option `json:"data"`   // วิธีการได้มาของข้อมูล
	Result string                 `json:"result"` // example:`OK`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/select_option_dataformat
// @Summary			เริ่มต้นหน้า data_format
// @Description		วิธีการได้มาซึ่งข้อมูล
// @Method			GET
// @Produces		json
// @Response		200	Struct_getSelectOption successful operation
func (srv *HttpService) getSelectOption(ctx service.RequestContext) error {

	rs, err := model.GetSelectOption()
	if err != nil {
		ctx.ReplyJSON(err.Error())
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
