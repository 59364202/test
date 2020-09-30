package manipulate

import (
	"encoding/json"
	"fmt"
	model "haii.or.th/api/thaiwater30/model/manipulate/dataunit"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	"strconv"
)

// Data Unit //
type DataUnitParam struct {
	Id           string          `json:"id"`
	DataunitName json.RawMessage `json:"dataunit_name,omitempty"`
}

func (srv *HttpService) getDataunit(ctx service.RequestContext) error {
	//	p := &DataUnitParam{}
	//	if err := ctx.GetRequestParams(p); err != nil {
	//		return errors.Repack(err)
	//	}
	//	ctx.LogRequestParams(p)

	result, err := model.GetDataunit(ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) postDatainit(ctx service.RequestContext) error {
	p := &DataUnitParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model.PostDataunit(ctx.GetUserID(), p.DataunitName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) putDataunit(ctx service.RequestContext) error {
	p := &DataUnitParam{}
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

	result, err := model.PutDataunit(ctx.GetUserID(), dataunitId, p.DataunitName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) deleteDataunit(ctx service.RequestContext) error {
	//	p := &DataUnitParam{}
	//	if err := ctx.GetRequestParams(p); err != nil {
	//		return errors.Repack(err)
	//	}
	//
	//	ctx.LogRequestParams(p)

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get ID."))
	}

	dataunitId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}

	result, err := model.DeleteDataunit(ctx.GetUserID(), dataunitId)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}
