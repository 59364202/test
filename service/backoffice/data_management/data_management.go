package data_management

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/data_management"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	//dpt.Register(ServiceVersion, service.MethodPOST, DataServiceName, srv.handlePostData)
	//dpt.Register(ServiceVersion, service.MethodPATCH, DataServiceName, srv.handlePatchData)
	dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlePutData)
	dpt.Register(ServiceVersion, service.MethodDELETE, DataServiceName, srv.handleDeleteData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	//=== Lastest Data ===//
	case "lastest_data_load":
		return srv.onLoadLastestData(ctx)
	case "lastest_data":
		return srv.getLastestData(ctx)
	//=== Check Metadata ===//
	case "check_metadata_load":
		return srv.onLoadCheckMetadata(ctx)
	case "check_metadata":
		return srv.getCheckMetadata(ctx)
	case "check_metadata_agency":
		return srv.getAgencyByTable(ctx)
	//=== Check Error Data ===//
	case "check_error_data_load":
		return srv.onLoadCheckErrorData(ctx)
	case "check_error_data":
		return srv.getCheckErrorData(ctx)
	case "check_error_data_agency":
		return srv.getAgencyByStationTable(ctx)
	//=== XXXXXXXX ===//
	case "impdata_option_list":
		return srv.importDataOptionList(ctx)
	case "event_tracking_option_list":
		return srv.getSelectOptionEventTracking(ctx)
	case "event_tracking_option_list_invalid_data":
		return srv.getSelectOptionEventTrackingInvalidData(ctx)
	case "event_tracking":
		return srv.getEventTracking(ctx)
	case "event_tracking_solve":
		return srv.getSolveEventTracking(ctx)
	case "event_download_invalid_data":
		return srv.getDownloadEventInvalidData(ctx)
	case "event_send_invalid_data":
		return srv.getSendEventInvalidData(ctx)
	case "event_tracking_invalid_data":
		return srv.getTrackingEventInvalidData(ctx)
	case "event_file_csv":
		return srv.getFileInvalidData2(ctx)
	case "event_file_csv2":
		return srv.getFileInvalidData2(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePutData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "event_tracking_solve":
		return srv.putEventTracking(ctx)
	case "event_send_invalid_data":
		return srv.putSendEventInvalidData(ctx)
	case "event_tracking_invalid_data":
		return srv.putTrackingEventInvalidData(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "lastest_data":
		return srv.deleteLastestData(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
