package event_management

import (
	model "haii.or.th/api/thaiwater30/model/event_log_category"
	//model_setting "haii.or.th/api/server/model/setting"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/event_load
// @Method			GET
// @Summary			load event cateogry log
// @Parameter		-	query model.Struct_EventLogCategory_InputParam{id,code}
// @Produces		json
// @Response		200		EventLoadSwagger successful operation
// @Response		404			-		the request service name was not found

type EventLoadSwagger struct {
	Result string             `json:"result"` //example:`OK`
	Data   []model.Struct_ELC `json:"data"`   // event information
}

func (srv *HttpService) getEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogCategory_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	dataResult, err := model.GetEventLogCategory(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/event
// @Method			POST
// @Summary			Add event cateogry log
// @Consumes		json
// @Parameter		-	body model.Struct_EventLogCategory_InputParamSwagger
// @Produces		json
// @Response		200		model.Struct_ELC successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) postEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogCategory_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Post Data
	dataResult, err := model.PostEventLogCategory(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/event
// @Method			PUT
// @Summary			Update event cateogry log
// @Consumes		json
// @Parameter		-	body model.Struct_EventLogCategory_InputParam
// @Produces		json
// @Response		200		model.Struct_ELC successful operation
// @Response		404			-		the request service name was not found
func (srv *HttpService) putEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogCategory_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Put Data
	dataResult, err := model.PutEventLogCategory(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/event
// @Method			DELETE
// @Summary			Delete event cateogry log
// @Parameter		-	query Struct_EventLogCategory_InputParamDel
// @Produces		json
// @Response		200		ResultdeleteEvent successful operation
// @Response		404			-		the request service name was not found

type Struct_EventLogCategory_InputParamDel struct {
	ID string `json:"id"` //example:1 eventlog category id
}
type ResultdeleteEvent struct {
	Result string      `json:"result"` //example:`ok`
	Data   interface{} `json:"data"`   // example:`Delete Successful.`
}

func (srv *HttpService) deleteEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventLogCategory_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Delete Data
	dataResult, err := model.DeleteEventLogCategory(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}
