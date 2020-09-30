package event_management

import (
	model_sink_condition "haii.or.th/api/thaiwater30/model/event_log_sink_condition"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_condition_option
// @Method			GET
// @Summary			return email template
// @Produces		json
// @Response		200		ResultSinkConditionSelectOption successful operation
type ResultSinkConditionSelectOption struct {
	Result string                                                 `json:"result"` // example:`OK`
	Data   *model_sink_condition.SinkConditionSelectOptionSwagger `json:"data"`
}

func (srv *HttpService) getSinkConditionSelectOption(ctx service.RequestContext) error {

	dataResult, err := model_sink_condition.GetSinkSelectOption()
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_condition/{id}
// @Method			GET
// @Parameter		id	path    string  sink condition id เช่น 1
// @Summary			sink condition by id
// @Produces		json
// @Response		200		ResultSinkConditionID successful operation
type ResultSinkConditionID struct {
	Result string                             `json:"result"` // example:`OK`
	Data   model_sink_condition.SinkCondition `json:"data"`   // sink condition id
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_condition
// @Method			GET
// @Summary			sink condition list
// @Produces		json
// @Response		200		ResultSinkConditionList successful operation

type ResultSinkConditionList struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []model_sink_condition.SinkConditionList `json:"data"`   // sink condition List
}

func (srv *HttpService) getSinkCondition(ctx service.RequestContext) error {
	var sinkConditionID string
	sinkConditionID = ctx.GetServiceParams("id")
	type SinkCondition struct {
		ID []int64 `json:"id"`
	}
	p := &SinkCondition{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	if sinkConditionID != "" {
		dataResult, err := model_sink_condition.GetSinkCondition(sinkConditionID)
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	} else {
		dataResult, err := model_sink_condition.GetSinkConditions(p.ID)
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_condition
// @Method			POST
// @Consumes		json
// @Parameter		- body model_sink_condition.SinkConditionInputSwagger
// @Summary			add sink condition
// @Produces		json
// @Response		200		ResultRID successful operation
// @Response		404			-		the request service name was not found
type ResultRID struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`134`
}

func (srv *HttpService) postSinkCondition(ctx service.RequestContext) error {

	p := &model_sink_condition.SinkConditionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	dataResult, err := model_sink_condition.AddSinkCondition(ctx.GetUserID(), p.Name, p.Channel, p.Category, p.Code, p.Service, p.Agent, p.User, p.Template, p.PostStartInterval)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_condition
// @Method			PUT
// @Consumes		json
// @Parameter		- body model_sink_condition.SinkConditionInputSwagger2
// @Summary			update sink condition
// @Produces		json
// @Response		200		ResultRID successful operation
// @Response		404			-		the request service name was not found
func (srv *HttpService) putSinkCondition(ctx service.RequestContext) error {

	p := &model_sink_condition.SinkConditionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	dataResult, err := model_sink_condition.UpdateSinkCondition(ctx.GetUserID(), p.ID, p.Name, p.Channel, p.Category, p.Code, p.Service, p.Agent, p.User, p.Template, p.PostStartInterval)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_condition
// @Method			DELETE
// @Parameter		- query SinkConditionInputSwaggerDel
// @Summary			soft deleted sink condition
// @Produces		json
// @Response		200		ResultRID 	successful operation
// @Response		404			-		the request service name was not found

type SinkConditionInputSwaggerDel struct {
	ID int64 `json:"id"` // example:`1` รหัส condition เช่น 11
}

func (srv *HttpService) deletedSinkCondition(ctx service.RequestContext) error {
	p := &model_sink_condition.SinkConditionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	if p.ID != 0 {
		dataResult, err := model_sink_condition.DeleteSinkCondition(p.ID, ctx.GetUserID())
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
