package event_management

import (
	model_target "haii.or.th/api/thaiwater30/model/event_log_sink_target"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_target_option
// @Method			GET
// @Summary			return sink target option
// @Produces		json
// @Response		200		ResultSinkTargetSelectOption successful operation
type ResultSinkTargetSelectOption struct {
	Result string                                 `json:"result"` // example:`OK`
	Data   model_target.TargetSelectOptionSwagger `json:"data"`
}

func (srv *HttpService) getSinkTargetSelectOption(ctx service.RequestContext) error {

	dataResult, err := model_target.GetSinkTargetSelectOption()
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_target/{id}
// @Method			GET
// @Parameter		id	path    string  sink target id ex. 1
// @Summary			return sink target information
// @Produces		json
// @Response		200		ResultTargetEdit successful operation
type ResultTargetEdit struct {
	Result string                  `json:"result"` // example:`OK`
	Data   model_target.TargetEdit `json:"data"`
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_target
// @Method			GET
// @Summary			return sink target list
// @Produces		json
// @Response		200		SinkTargetSwagger successful operation

type SinkTargetSwagger struct {
	Result string                       `json:"result"` //example:`OK`
	Data   []model_target.TargetDetails `json:"data"`
}

func (srv *HttpService) getSinkTarget(ctx service.RequestContext) error {
	var templateID string
	templateID = ctx.GetServiceParams("id")
	if templateID != "" {
		dataResult, err := model_target.GetTarget(templateID)
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	} else {
		dataResult, err := model_target.GetTargets()
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_target
// @Method			POST
// @Parameter		- form model_target.TargetInputAddSwagger
// @Summary			return sink target
// @Produces		json
// @Response		200		ResultSinkTargetID successful operation
// @Response		404			-		the request service name was not found
type ResultSinkTargetID struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`42`
}

func (srv *HttpService) postSinkTarget(ctx service.RequestContext) error {

	p := &model_target.TargetInputAdd{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	dataResult, err := model_target.AddTarget(ctx.GetUserID(), p.Condition, p.Method, p.Group, p.Lang, p.Color)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_target
// @Method			PUT
// @Parameter		- form model_target.TargetInputAdd
// @Summary			return sink target
// @Produces		json
// @Response		200		ResultSinkTargetID successful operation
// @Response		404			-		the request service name was not found
func (srv *HttpService) putSinkTarget(ctx service.RequestContext) error {

	p := &model_target.TargetInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	dataResult, err := model_target.UpdateTarget(ctx.GetUserID(), p.ID, p.Condition, p.Method, p.Group, p.Lang, p.Color)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_target
// @Method			DELETE
// @Parameter		- query TargetInputDel
// @Summary			return sink target
// @Produces		json
// @Response		200		ResultSinkTargetID successful operation
// @Response		404			-		the request service name was not found

type TargetInputDel struct {
	ID int64 `json:"id"` // example:`1` รหัสของ target เช่น 1
}

func (srv *HttpService) deletedSinkTarget(ctx service.RequestContext) error {
	p := &model_target.TargetInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	if p.ID != 0 {
		dataResult, err := model_target.DeleteTarget(p.ID, ctx.GetUserID())
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	} else {
		ctx.ReplyJSON(result.Result0(nil))
	}

	return nil
}
