package data_integration_report

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/data_integration_report"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	//=== Event Log ===//
	case "event_load":
		return srv.onLoadEvent(ctx)
	case "event_log_category_summary":
		return srv.getEventLogCategorySummary(ctx)
	case "event_code_summary":
		return srv.getEventCodeSummary(ctx)
	case "event_detail":
		return srv.getEventDetail(ctx)
	case "event_code_list":
		return srv.getEventCodeList(ctx)
	case "event_summary_report":
		return srv.getEventSummaryReport(ctx)
	//=== Over All ===//
	case "overall":
		return srv.getOverAllPercentDownload(ctx)
	case "overall_multiple":
		return srv.getOverAllMultiYear(ctx)
	case "compare_yearly":
		return srv.getYearlyComparePercentDownload(ctx)
	//=== Download Size ===//
	case "download_size":
		return srv.getMonthlyDownloadSize(ctx)
	case "multi_download_size":
		return srv.getMultipleAgecncyAndMonthAndYear(ctx)
	//=== Download Percent ===//
	case "download_percent_load":
		return srv.onLoadPercentDownload(ctx)
	case "download_percent":
		return srv.getPercentDownload(ctx)
	case "download_percent_detail":
		return srv.getPercentDownloadDetail(ctx)
	//=== report_import_print ===/
	case "report_import_print":
		return srv.getReportImportPrint(ctx)
	case "event_report":
		return srv.getEventReport(ctx)
	case "event_report_load":
		return srv.getEventReportLoad(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
