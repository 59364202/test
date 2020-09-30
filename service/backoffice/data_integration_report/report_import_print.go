package data_integration_report

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"time"

	model_download_log "haii.or.th/api/thaiwater30/model/dataimport_download_log"
	model_event_log "haii.or.th/api/thaiwater30/model/event_log"
)

type Param_getReportImportPrint struct {
	StartDate string `json:"start_date"` // example: "2006-01-02 15:04" วันเวลาเริ่มต้น
	EndDate   string `json:"end_date"`   // example: "2006-01-02 15:04" วันเวลาสิ้นสุด
}

type Struct_ReportImportPrint struct {
	Data       *Struct_ReportDetail `json:"data"`         // รายงานตารางหน้าแรก
	ChartData  *Struct_ReportDetail `json:"chart_data"`   // รายงานกราฟ ตาราง วันเวลาที่เลือก
	ChartData1 *Struct_ReportDetail `json:"chart_data_1"` // รายงานกราฟ ตาราง วันเวลาที่เลือก ย้อนหลังไป 1 วัน
	ChartData2 *Struct_ReportDetail `json:"chart_data_2"` // รายงานกราฟ ตาราง วันเวลาที่เลือก ย้อนหลังไป 2 วัน
}
type Struct_ReportDetail struct {
	Data      *result.Result `json:"data"`       // ข้อมูล
	StartDate string         `json:"start_date"` // วันเวลาเริ่มต้น
	EndDate   string         `json:"end_date"`   // วันเวลาสิ้นสุด
}

type Struct_getReportImportPrint struct {
	Data       *Struct_getReportImportPrint_Data  `json:"data"`         // รายงานตารางหน้าแรก
	ChartData  *Struct_getReportImportPrint_Chart `json:"chart_data"`   // รายงานกราฟ ตาราง วันเวลาที่เลือก
	ChartData1 *Struct_getReportImportPrint_Chart `json:"chart_data_1"` // รายงานกราฟ ตาราง วันเวลาที่เลือก ย้อนหลังไป 1 วัน
	ChartData2 *Struct_getReportImportPrint_Chart `json:"chart_data_2"` // รายงานกราฟ ตาราง วันเวลาที่เลือก ย้อนหลังไป 2 วัน
}
type Struct_getReportImportPrint_Data struct {
	Data      *Struct_getReportImportPrint_Data `json:"data"`       // รายงานตารางหน้าแรก
	StartDate string                            `json:"start_date"` // example:`2006-01-02 15:04` วันเวลาเริ่มต้นที่เลือก
	EndDate   string                            `json:"end_date"`   // example:`2006-01-02 15:04` วันเวลาสิ้นสุดที่เลือก
}
type Struct_getReportImportPrint_Data_Data struct {
	Result string                                          `json:"result"` // example:`OK`
	Data   []*model_event_log.Struct_EventLogSummaryReport `json:"data"`   // รายงานตารางหน้าแรก
}
type Struct_getReportImportPrint_Chart struct {
	Data      *Struct_getReportImportPrint_Chart_Data `json:"data"`       // รายงานกราฟ ตาราง
	StartDate string                                  `json:"start_date"` // example:`2006-01-02 15:04` วันเวลาเริ่มต้น
	EndDate   string                                  `json:"end_date"`   // example:`2006-01-02 15:04` วันเวลาสิ้นสุด
}
type Struct_getReportImportPrint_Chart_Data struct {
	Result string                                           `json:"result"` // example:`OK`
	Data   []*model_download_log.Struct_DownloadLog_Summary `json:"data"`   // รายงานกราฟ ตาราง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_integration_report/report_import_print
// @Summary			รายงานการนำเข้าข้อมูล
// @Method			GET
// @Parameter		- query Param_getReportImportPrint
// @Produces		json
// @Response		200	Struct_getReportImportPrint successful operation
func (srv *HttpService) getReportImportPrint(ctx service.RequestContext) error {
	//Map parameters
	param := &model_event_log.Struct_EventLog_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return err
	}
	ctx.LogRequestParams(param)

	layout := "2006-01-02 15:04"
	layoutReturn := "2006-1-2 15:04"
	timeStartDate, err := time.Parse(layout, param.StartDate)
	if err != nil {
		return err
	}
	timeEndDate, err := time.Parse(layout, param.EndDate)
	if err != nil {
		return err
	}

	//Get List of Event Summary Report
	dataResult, err := model_event_log.GetEventLogSummaryReport(param)
	if err != nil {
		return err
	}

	dlParam := &model_download_log.Struct_DownloadLog_Inputparam{
		StartDate:        param.StartDate,
		EndDate:          param.EndDate,
		ConnectionFormat: "online",
	}
	dlDataResult, err := model_download_log.GetPercentDownload(dlParam) // วันที่เลือก
	if err != nil {
		return err
	}
	dateNoTime1 := timeStartDate.Add(-24 * time.Hour) // วันที่เลือก -1 วัน
	timeStartDate1 := time.Date(dateNoTime1.Year(), dateNoTime1.Month(), dateNoTime1.Day(), 0, 0, 0, 0, time.UTC)
	timeEndDate1 := time.Date(dateNoTime1.Year(), dateNoTime1.Month(), dateNoTime1.Day(), 23, 59, 0, 0, time.UTC)
	dlParam1 := &model_download_log.Struct_DownloadLog_Inputparam{
		StartDate:        timeStartDate1.Format(layout),
		EndDate:          timeEndDate1.Format(layout),
		ConnectionFormat: "online",
	}
	dlDataResult1, err := model_download_log.GetPercentDownload(dlParam1)
	if err != nil {
		return err
	}
	dateNoTime2 := timeStartDate.Add(-48 * time.Hour) // วันที่เลือก -2 วัน
	timeStartDate2 := time.Date(dateNoTime2.Year(), dateNoTime2.Month(), dateNoTime2.Day(), 0, 0, 0, 0, time.UTC)
	timeEndDate2 := time.Date(dateNoTime2.Year(), dateNoTime2.Month(), dateNoTime2.Day(), 23, 59, 0, 0, time.UTC)
	dlParam2 := &model_download_log.Struct_DownloadLog_Inputparam{
		StartDate:        timeStartDate2.Format(layout),
		EndDate:          timeEndDate2.Format(layout),
		ConnectionFormat: "online",
	}
	dlDataResult2, err := model_download_log.GetPercentDownload(dlParam2)
	if err != nil {
		return err
	}

	rs := &Struct_ReportImportPrint{
		Data:       &Struct_ReportDetail{Data: result.Result1(dataResult), StartDate: timeStartDate.Format(layoutReturn), EndDate: timeEndDate.Format(layoutReturn)},
		ChartData:  &Struct_ReportDetail{Data: result.Result1(dlDataResult), StartDate: timeStartDate.Format(layoutReturn), EndDate: timeEndDate.Format(layoutReturn)},
		ChartData1: &Struct_ReportDetail{Data: result.Result1(dlDataResult1), StartDate: timeStartDate1.Format(layoutReturn), EndDate: timeEndDate1.Format(layoutReturn)},
		ChartData2: &Struct_ReportDetail{Data: result.Result1(dlDataResult2), StartDate: timeStartDate2.Format(layoutReturn), EndDate: timeEndDate2.Format(layoutReturn)},
	}
	ctx.ReplyJSON(rs)

	return nil
}
