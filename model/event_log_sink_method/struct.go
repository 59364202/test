package event_log_sink_method

import (
	"encoding/json"
	model_type "haii.or.th/api/thaiwater30/model/event_log_sink_method_type"
)

type Struct_EventLogSinkMethod struct {
	Event_log_sink_method_type *model_type.Struct_EventLogSinkMethodType `json:"event_log_sink_method_type"` // sink method type
	Id                         int64                                     `json:"id"` // example:`4` sink method id
	Description                string                                    `json:"description"` // example:`sinkmethod description` sink method description
	ConfigName                 string                                    `json:"config_name"` // example:`smtp4` // sink method config name
	Sink_params                json.RawMessage                           `json:"sink_params,omitempty"` // example:`{"system_setting_name":"thaiwater30.service.event_management.smtpserver.smtp1"}` sink method config
}

type Struct_EventLogSinkMethod_InputParamSwagger struct {
	Event_log_sink_method_type_id string          `json:"event_log_sink_method_type_id"` // example:`1` sink method type id ex. 1
	Description                   string          `json:"description"`// example:`description` sink method description
	Sink_params                   json.RawMessage `json:"sink_params,omitempty"` // example:`{"system_setting_name":"thaiwater30.service.event_management.smtpserver.smtp1"}` sink method config
}

type Struct_EventLogSinkMethod_InputParam struct {
	Event_log_sink_method_type_id string          `json:"event_log_sink_method_type_id"` // example:`3`sink method type id ex. 3
	Id                            string          `json:"id"` // example:`2`sink method id ex. 2
	Description                   string          `json:"description"` // example:`sinkmethod description` sink method description
	Sink_params                   json.RawMessage `json:"sink_params,omitempty"` // example:`{"system_setting_name":"thaiwater30.service.event_management.smtpserver.smtp1"}` sink method config
}

type Struct_EventLogSinkMethod_InputParamSwagger2 struct {
	Event_log_sink_method_type_id string          `json:"event_log_sink_method_type_id"` // example:`3`sink method type id ex. 3
	Description                   string          `json:"description"` // example:`sinkmethod description` sink method description
	Sink_params                   json.RawMessage `json:"sink_params,omitempty"` // example:`{"system_setting_name":"thaiwater30.service.event_management.smtpserver.smtp1"}` sink method config
}

type Struct_EventLogSinkMethodSystemSetting_InputParam struct {
	Event_log_sink_method_type_id string          `json:"event_log_sink_method_type_id"` // example:`1` sink method type id ex. 1
	Id                            string          `json:"id"` // example:`20` sink method id ex. 20
	Description                   string          `json:"description"` // example:`description` description sink method
	Name                          string          `json:"config_name"` // example:`config_name` sink method config name
	Sink_params                   json.RawMessage `json:"sink_params,omitempty"` // example:`{"system_setting_name":"thaiwater30.service.event_management.smtpserver.smtp1"}` sink method config
}

type Struct_EventLogSinkMethodSystemSetting_InputParamSwagger struct {
	Event_log_sink_method_type_id string          `json:"event_log_sink_method_type_id"` // example:`1` sink method type id ex. 1
	Id                            string          `json:"id"` // example:`20` sink method id ex. 20
	Description                   string          `json:"description"` // example:`description` description sink method
	Name                          string          `json:"config_name"` // example:`config_name` sink method config name
	Sink_params                   json.RawMessage `json:"sink_params,omitempty"` // example:`{"system_setting_name":"thaiwater30.service.event_management.smtpserver.smtp1"}` sink method config
}