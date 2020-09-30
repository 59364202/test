package manipulate

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

func (srv *HttpService) getDataformat(ctx service.RequestContext) error {
	//	//Map parameters
	p := &dataformatParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Dataformat
	var rs []*model.Dataformat_struct
	var err error
	//	pid := ctx.GetServiceParams("id")
	//	if pid == "" {
	//		rs, err = model.GetSelectOption()
	//	} else {
	rs, err = model.GetDataformat(ctx.GetServiceParams("id"), p.MetdataMethodID)
	//	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

func (srv *HttpService) postDataformat(ctx service.RequestContext) error {
	//Map parameters
	p := &dataformatParamPostPut{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Post Dataformat
	result, err := model.PostDataformat(ctx.GetUserID(), p.DataformatName, p.MetdataMethodID)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) putDataformat(ctx service.RequestContext) error {
	//Map parameters
	p := &dataformatParamPostPut{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Put Dataformat
	result, err := model.PutDataformat(ctx.GetUserID(), ctx.GetServiceParams("id"), p.DataformatName, p.MetdataMethodID)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) deleteDataformat(ctx service.RequestContext) error {
	//	//Map parameters
	//	p := &dataformatParam{}
	//	if err := ctx.GetRequestParams(p); err != nil {
	//		return errors.Repack(err)
	//	}
	//	ctx.LogRequestParams(p)

	//Delete Dataformat
	result, err := model.DeleteDataformat(ctx.GetUserID(), ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) getSelectOption(ctx service.RequestContext) error {

	result, err := model.GetSelectOption()
	if err != nil {
		ctx.ReplyJSON(err.Error())
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}
