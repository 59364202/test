package event_management

import (
	"encoding/json"
	model_setting "haii.or.th/api/server/model/setting"
	model "haii.or.th/api/thaiwater30/model/event_code"
	model_event "haii.or.th/api/thaiwater30/model/event_log_category"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_OnLoadSubEvent struct {
	Event           []*model_event.Struct_ELC `json:"event"`
	SubEvent        []*model.Struct_EventCode `json:"sub_event"`
	SubTypeCategory json.RawMessage           `json:"subtype_category"` // example:`{"en":"subevent category"}`
}

type ResultSubEventLoad struct {
	Result string                `json:"result"` // example:`ok`
	Data   Struct_OnLoadSubEvent `json:"data"`
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/subevent_load
// @Method			GET
// @Summary			load subevent
// @Produces		json
// @Response		200		ResultSubEventLoad successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) onLoadSubEvent(ctx service.RequestContext) error {
	dataResult := &Struct_OnLoadSubEvent{}

	//Get List of Event Data
	paramEvent := &model_event.Struct_EventLogCategory_InputParam{}
	resultEvent, err := model_event.GetEventLogCategory(paramEvent)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.Event = resultEvent

	//Get List of SubtypeCategory Data
	dataResult.SubTypeCategory = model_setting.GetSystemSettingJson("bof.EventMgt.EventCode.SubtypeCategory")

	//Get List of SubEvent Data
	p := &model.Struct_EventCode_InputParam{}
	resultSubEvent, err := model.GetEventCode(p)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.SubEvent = resultSubEvent

	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/subevent
// @Method			GET
// @Summary			load subevent
// @Parameter		-	query model.Struct_EventCode_InputParamGetSwagger
// @Produces		json
// @Response		200		SubEventeSwagger successful operation
// @Response		404			-		the request service name was not found

type SubEventeSwagger struct {
	Result string                   `json:"result"` //example:`OK`
	Data   []model.Struct_EventCode `json:"data"`
}

func (srv *HttpService) getSubEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventCode_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	dataResult, err := model.GetEventCode(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/subevent
// @Method			POST
// @Summary			post subevent
// @Consumes		json
// @Parameter		-	body model.Struct_EventCode_InputParamSwagger
// @Produces		json
// @Response		200		ResultSubEventList successful operation
// @Response		404			-		the request service name was not found
type ResultSubEventList struct {
	Result string `json:"result"` // example:`OK`
	Data []model.Struct_EventCode `json:"data"`
}


func (srv *HttpService) postSubEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventCode_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Post Data
	dataResult, err := model.PostEventCode(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/subevent
// @Method			PUT
// @Summary			put subevent
// @Consumes		json
// @Parameter		-	body model.Struct_EventCode_InputParam
// @Produces		json
// @Response		200		ResultSubEventList successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) putSubEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventCode_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Put Data
	dataResult, err := model.PutEventCode(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/subevent
// @Method			DELETE
// @Summary			Delete subevent
// @Parameter		-	query Struct_EventCode_InputParamSw
// @Produces		json
// @Response		200		ResultSubEventDel  successful operation
// @Response		404			-		the request service name was not found

type Struct_EventCode_InputParamSw struct {
	ID string `json:"id"` //example:3 sub event id
}

type ResultSubEventDel struct {
	Result string `json:"result"` // example:`OK`
	Data interface{} `json:"data"` // example:`Delete Successful.`
}

func (srv *HttpService) deleteSubEvent(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_EventCode_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Delete Data
	dataResult, err := model.DeleteEventCode(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(dataResult)
	}

	return nil
}
