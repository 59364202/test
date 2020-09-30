package event_management

import (
	model "haii.or.th/api/thaiwater30/model/event_log_sink_method"
	model_type "haii.or.th/api/thaiwater30/model/event_log_sink_method_type"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_OnLoadSinkMethod struct {
	SinkMethodType []*model_type.Struct_EventLogSinkMethodType `json:"sink_method_type"`
	SinkMethod     []*model.Struct_EventLogSinkMethod          `json:"sink_method"`
}

type ResultSinkMethodOnload struct {
	Result string                  `json:"result"` // example:`OK`
	Data   Struct_OnLoadSinkMethod `json:"data"`
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method_load
// @Method			GET
// @Summary			load sink method
// @Produces		json
// @Response		200		ResultSinkMethodOnload successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) onLoadSinkMethod(ctx service.RequestContext) error {
	dataResult := &Struct_OnLoadSinkMethod{}

	//Get List of SinkMethodType Data
	resultSinkMethodType, err := model_type.GetEventLogSinkMethodType()
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.SinkMethodType = resultSinkMethodType

	//Get List of SinkMethod Data
	p := &model.Struct_EventLogSinkMethod_InputParam{}
	resultSinkMethod, err := model.GetEventLogSinkMethod(p)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.SinkMethod = resultSinkMethod

	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method
// @Method			GET
// @Summary			get sink method
// @Parameter		-	query Struct_EventLogSinkMethod_InputParamGet
// @Produces		json
// @Response		200		SinkMethodSwagger successful operation
// @Response		404			-		the request service name was not found

type SinkMethodSwagger struct {
	Result string                            `json:"result"` //example:`OK`
	Data   []model.Struct_EventLogSinkMethod `json:"data"`
}

type Struct_EventLogSinkMethod_InputParamGet struct {
	Id                            string `json:"id"` // example:`3` รหัส sink method เช่น 3
	Event_log_sink_method_type_id string `json:"event_log_sink_method_type_id"` // example:`5` รหัส sink method type เช่น 5 
}

func (srv *HttpService) getSinkMethod(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogSinkMethod_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	dataResult, err := model.GetEventLogSinkMethod(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method
// @Method			POST
// @Summary			add sink method
// @Consumes		json
// @Parameter		-	body model.Struct_EventLogSinkMethod_InputParamSwagger2
// @Produces		json
// @Response		200		LtSinkMethodSwagger successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) postSinkMethod(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogSinkMethod_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Post Data
	dataResult, err := model.PostEventLogSinkMethod(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method_system_setting
// @Method			POST
// @Summary			update sink method
// @Consumes		json
// @Parameter		-	body model.Struct_EventLogSinkMethodSystemSetting_InputParamSwagger
// @Produces		json
// @Response		200		- successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) postSinkMethodSystemSetting(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogSinkMethodSystemSetting_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Post Data
	dataResult, err := model.PostEventLogSinkMethodSystemSetting(ctx.GetUserID(), p.Event_log_sink_method_type_id, p.Name, p.Description)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method_system_setting
// @Method			PUT
// @Summary			update sink method system_setting
// @Consumes		json
// @Parameter		-	body model.Struct_EventLogSinkMethodSystemSetting_InputParamSwagger
// @Produces		json
// @Response		200		ResultID successful operation
// @Response		404			-		the request service name was not found

type ResultID struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`31`
}

func (srv *HttpService) putSinkMethodSystemSetting(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogSinkMethodSystemSetting_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Put Data
	dataResult, err := model.PutEventLogSinkMethodSystemSetting(p.Id, ctx.GetUserID(), p.Event_log_sink_method_type_id, p.Name, p.Description)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method
// @Method			PUT
// @Summary			update sink method
// @Consumes		json
// @Parameter		-	body model.Struct_EventLogSinkMethod_InputParamSwagger
// @Produces		json
// @Response		200		LtSinkMethodSwagger successful operation
// @Response		404			-		the request service name was not found
func (srv *HttpService) putSinkMethod(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogSinkMethod_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Put Data
	dataResult, err := model.PutEventLogSinkMethod(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sink_method
// @Method			DELETE
// @Summary			soft  deleted sink method
// @Parameter		-	query Struct_EventLogSinkMethod_InputParamDEL
// @Produces		json
// @Response		200		- successful operation
// @Response		404			-		the request service name was not found

type Struct_EventLogSinkMethod_InputParamDEL struct {
	Id string `json:"id"` // example:`1` รหัส sink method เช่น 1
}

type ResultDelSinkMethod struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful.`
}

func (srv *HttpService) deleteSinkMethod(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogSinkMethod_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Delete Data
	dataResult, err := model.DeleteEventLogSinkMethod(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}
