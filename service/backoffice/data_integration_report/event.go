package data_integration_report

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_key_access "haii.or.th/api/thaiwater30/model/agent"
	model_event_code "haii.or.th/api/thaiwater30/model/event_code"
	model "haii.or.th/api/thaiwater30/model/event_log"
	model_event_category "haii.or.th/api/thaiwater30/model/event_log_category"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_GetEventCategorySummary struct {
	EventCategorySummary         []*model_event_category.Struct_EventLogCategory       `json:"event_log_category_summary"`
	EventCategorySummaryByAgency []*model.Struct_EventLogSummary_GroupByAgencyCategory `json:"event_log_category_summary_by_agency"`
}
type Param_onLoadEvent struct {
	StartDate string `json:"start_date"` // example:`2017-01-02` วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // example:`2017-01-02` วันที่สิ้นสุด
}
type Struct_onLoadEvent struct {
	Result string                   `json:"result"` // example:`OK`
	Data   *Struct_onLoadEvent_Data `json:"data"`   // รายงานเหตุการณ์
}
type Struct_event struct {
	EventCategorySummary         []*model_event_category.Struct_EventLogCategory       `json:"event_log_category_summary"`           // เหตุการณ์ที่เกิดขึ้น
	EventCategorySummaryByAgency []*model.Struct_EventLogSummary_GroupByAgencyCategory `json:"event_log_category_summary_by_agency"` // เหตุการณ์ที่เกิดขึ้นแยกตามหน่วยงาน
}
type Struct_onLoadEvent_Data struct {
	AgencyList        []*model_agency.Struct_Agency      `json:"agency_list"`             // หน่วยงาน
	EventCategoryList []*model_event_category.Struct_ELC `json:"event_log_category_list"` // ประเภทเหตุการณ์
	DateRange         int64                              `json:"date_range"`              // example:`30` ช่วงวันที่ที่เลือกได้
	Struct_event
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_load
// @Summary			เริ่มต้นหน้ารายงานเหตุการณ์
// @Method			GET
// @Parameter		- query Param_onLoadEvent
// @Produces		json
// @Response		200	Struct_onLoadEvent successful operation
func (srv *HttpService) onLoadEvent(ctx service.RequestContext) error {

	dataResult := &Struct_onLoadEvent_Data{}

	//Map parameters
	param := &model.Struct_EventLog_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Agency
	resultAgency, err := model_agency.GetAllAgency()
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.AgencyList = resultAgency

	//Get List of Event Category
	paramEventCat := &model_event_category.Struct_EventLogCategory_InputParam{}
	resultEventCat, err := model_event_category.GetEventLogCategory(paramEventCat)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.EventCategoryList = resultEventCat

	//Get Date Range Setting
	dataResult.DateRange = model_setting.GetSystemSettingInt("bof.DataIntRpt.EventLog.DateRange")
	if (dataResult.DateRange) == 0 {
		dataResult.DateRange = model_setting.GetSystemSettingInt("setting.Default.DateRange")
	}

	//Get List of Summary Event By Agency and Category => for summary table
	resultEventSumData, err := model.GetEventLogSummaryGroupByAgencyCategory(param)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.EventCategorySummaryByAgency = resultEventSumData

	//Get List of Summary Event By Category => for graph
	resultEventSumGraph, err := model.GetEventLogSummaryGroupByCategory(param)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.EventCategorySummary = resultEventSumGraph

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Param_getEventLogCategorySummary struct {
	AgencyId  []int64 `json:"agency_id"`  // example:`9` รหัสหน่วยงาน
	StartDate string  `json:"start_date"` // example:`2017-06-29` วันที่เริ่มต้น
	EndDate   string  `json:"end_date"`   // example:`2017-06-30` วันที่สิ้นสุด
}
type Struct_getEventLogCategorySummary struct {
	Result string        `json:"result"` // example:`OK`
	Data   *Struct_event `json:"data"`   // รายงานเหตุการณ์
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_log_category_summary
// @Summary 		รายงานเหตุการณ์
// @Description 	เหตุการณ์ที่เกิดขึ้น, เหตุการณ์ที่เกิดขึ้นแยกตามหน่วยงาน
// @Method			GET
// @Parameter		- query Param_getEventLogCategorySummary
// @Produces		json
// @Response		200		Struct_GetEventCategorySummary successful operation
func (srv *HttpService) getEventLogCategorySummary(ctx service.RequestContext) error {

	dataResult := &Struct_GetEventCategorySummary{}

	//Map parameters
	param := &model.Struct_EventLog_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Summary Event By Agency and Category => for summary table
	resultEventSumData, err := model.GetEventLogSummaryGroupByAgencyCategory(param)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.EventCategorySummaryByAgency = resultEventSumData

	//Get List of Summary Event By Category => for graph
	resultEventSumGraph, err := model.GetEventLogSummaryGroupByCategory(param)
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.EventCategorySummary = resultEventSumGraph

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Param_getEventCodeSummary struct {
	AgencyID  string `json:"agency_id"`  // example:`14` รหัสหน่วยงาน
	StartDate string `json:"start_date"` // example:`2017-07-02` วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // example:`2017-07-02` วันที่สิ้นสุด
}
type Struct_getEventCodeSummary struct {
	Result string                                            `json:"result"` // example:`OK`
	Data   []*model_event_code.Struct_EventCode_SummaryEvent `json:"data"`   // รายงานเหตุการณ์
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_code_summary
// @Summary			รายงานเหตุการณ์ของหน่วยงาน
// @Method			GET
// @Parameter		- query Param_getEventCodeSummary
// @Produces		json
// @Response		200	Struct_getEventCodeSummary successful operation
func (srv *HttpService) getEventCodeSummary(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_EventLog_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Summary Event By Agency and Category => for summary table
	dataResult, err := model.GetEventLogSummaryGroupByCode(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Struct_getEventDetail struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model.Struct_EventLog `json:"data"`   // รายละเอียดเหตุการณ์
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_detail
// @Summary			รายละเอียดเหตุการณ์
// @Method			GET
// @Parameter		- query model.Struct_EventLog_InputParam
// @Produces		json
// @Response		200	Struct_getEventDetail successful operation
func (srv *HttpService) getEventDetail(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_EventLog_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Event
	dataResult, err := model.GetEventLogDetail(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Struct_getEventCodeList struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_event_code.Struct_EventCode `json:"data"`   // เหตุการณ์ย่อย
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_code_list
// @Summary			เหตุการณ์ย่อย
// @Method			GET
// @Parameter		event_log_category_id query string example:4 รหัสเหตุการณ์
// @Produces		json
// @Response		200	Struct_getEventCodeList successful operation
func (srv *HttpService) getEventCodeList(ctx service.RequestContext) error {

	//Map parameters
	param := &model_event_code.Struct_EventCode_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Event Code
	dataResult, err := model_event_code.GetEventCode(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Param_getEventSummaryReport struct {
	StartDate string `json:"start_date"` // example: 2006-01-02 วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // example: 2006-01-02 วันที่สิ้นสุด
}

type Struct_getEventSummaryReport struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model.Struct_EventLogSummaryReport `json:"data"`   // เหตุการณ์ย่อย
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_summary_report
// @Summary			รายงานภาพรวมเหตุการณ์
// @Method			GET
// @Parameter		- query Param_getEventSummaryReport
// @Produces		json
// @Response		200	Struct_getEventSummaryReport successful operation
func (srv *HttpService) getEventSummaryReport(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_EventLog_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Event Summary Report
	dataResult, err := model.GetEventLogSummaryReport(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Struct_getEventReportLoad struct {
	Agent interface{}                      `json:"agent"` // agent
	Event []*model_event_code.Struct_Event `json:"event"` // event
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_report_load
// @Summary			รายงานเหตุการณ์
// @Method			GET
// @Produces		json
// @Response		200	Struct_getEventReportLoad successful operation
func (srv *HttpService) getEventReportLoad(ctx service.RequestContext) error {
	_agent, err := model_key_access.GetKeyAccessTable()
	if err != nil {
		return errors.Repack(err)
	}
	agent := _agent.(*result.Result)

	//Get List of Event Summary Report
	event, err := model_event_code.GetEventCategoryEventCode()
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(&Struct_getEventReportLoad{Agent: agent.Data, Event: event}))
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/event_report
// @Summary			รายงานเหตุการณ์
// @Method			GET
// @Parameter		- query model.Struct_EventReport_Input
// @Produces		json
// @Response		200	Struct_getEventSummaryReport successful operation
func (srv *HttpService) getEventReport(ctx service.RequestContext) error {

	//Map parameters
	param := &model.Struct_EventReport_Input{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	//Get List of Event Summary Report
	dataResult, err := model.GetEventReport(param)
	if err != nil {
		return errors.Repack(err)
	}

	//Return Data
	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}
