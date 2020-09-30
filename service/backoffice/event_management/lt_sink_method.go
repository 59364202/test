package event_management

import (
	model_lt_sink_method_type "haii.or.th/api/thaiwater30/model/event_log_sink_method_type"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/lt_sink_method
// @Method			GET
// @Summary			return sink method type list
// @Produces		json
// @Response		200		LtSinkMethodSwagger successful operation
type LtSinkMethodSwagger struct {
	Result string                                                    `json:"result"` //example:`OK`
	Data   []model_lt_sink_method_type.Struct_EventLogSinkMethodType `json:"data"`
}

func (srv *HttpService) getLtSinkMethodType(ctx service.RequestContext) error {

	dataResult, err := model_lt_sink_method_type.GetEventLogSinkMethodType()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/lt_sink_method
// @Method			POST
// @Parameter		- form model_lt_sink_method_type.Struct_EventLogSinkMethodTypeSwagger
// @Summary			return sink method type list
// @Produces		json
// @Response		200		ResultLtSinkMethod successful operation
type ResultLtSinkMethod struct {
	Result string `json:"result"` // example:`OK`
	Data interface{} `json:"data"` // example:`134` lt sink method id
}
func (srv *HttpService) postLtSinkMethodType(ctx service.RequestContext) error {

	p := &model_lt_sink_method_type.Struct_EventLogSinkMethodType{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	dataResult, err := model_lt_sink_method_type.AddSinkMethodType(p.Description, ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/lt_sink_method
// @Method			PUT
// @Parameter		- form model_lt_sink_method_type.Struct_EventLogSinkMethodType
// @Summary			return sink method type list
// @Produces		json
// @Response		200		ResultLtSinkMethod successful operation
func (srv *HttpService) putLtSinkMethodType(ctx service.RequestContext) error {

	p := &model_lt_sink_method_type.Struct_EventLogSinkMethodType{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	dataResult, err := model_lt_sink_method_type.UpdateSinkMethodType(p.Id, p.Description, ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}
	return nil
}
