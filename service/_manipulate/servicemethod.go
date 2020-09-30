package manipulate

import (
	"fmt"
	model "haii.or.th/api/thaiwater30/model/lt_servicemethod"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	"strconv"
)

// Service Method Parameter //
type ServiceMethodParam struct {
	Id                string            `json:"id"`
	ServiceMethodName map[string]string `json:"servicemethod_name"`
}

func (srv *HttpService) getServiceMethod(ctx service.RequestContext) error {
	//	p := &ServiceMethodParam{}
	//	if err := ctx.GetRequestParams(p); err != nil {
	//		return errors.Repack(err)
	//	}
	//	ctx.LogRequestParams(p)

	result, err := model.GetServiceMethod(ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) postServiceMethod(ctx service.RequestContext) error {
	p := &ServiceMethodParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model.PostServiceMethod(ctx.GetUserID(), p.ServiceMethodName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) putServiceMethod(ctx service.RequestContext) error {
	p := &ServiceMethodParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get ID."))
	}

	serviceMethodId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}

	result, err := model.PutServiceMethod(ctx.GetUserID(), serviceMethodId, p.ServiceMethodName)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) deleteServiceMethod(ctx service.RequestContext) error {
	//	p := &ServiceMethodParam{}
	//	if err := ctx.GetRequestParams(p); err != nil {
	//		return errors.Repack(err)
	//	}
	//
	//	ctx.LogRequestParams(p)

	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get ID."))
	}

	serviceMethodId, err := strconv.ParseInt(ctx.GetServiceParams("id"), 10, 64)
	if err != nil {
		ctx.ReplyError(err)
	}

	result, err := model.DeleteServiceMethod(ctx.GetUserID(), serviceMethodId)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}
