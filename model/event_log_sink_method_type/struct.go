package event_log_sink_method_type

import ()

type Struct_EventLogSinkMethodType struct {
	Id          int64  `json:"id"` // example:`1` รหัส sink method เช่น 1 
	Description string `json:"description"` // example:`SMTP (User mail)` รายละเอียด sink method เช่น SMTP (User mail)
}

type Struct_EventLogSinkMethodTypeSwagger struct {
	Id          int64  `json:"id"` // example:`11` รหัส sink method เช่น 11 
}