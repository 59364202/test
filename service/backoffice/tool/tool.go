package tool

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/tool"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlePutData)
	dpt.Register(ServiceVersion, service.MethodPATCH, DataServiceName, srv.handlePatchData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "lastest_image_load":
		return srv.onLoadLastestImage(ctx)
	case "lastest_image":
		return srv.getLastestImage(ctx)
	case "image_type":
		return srv.getImageType(ctx)
	case "display_image":
		return srv.displayImage(ctx)
	case "display_setting_json":
		return srv.getDisplaySetting(ctx)
	case "setting_list":
		return srv.getSettingList(ctx)
	//=== Ignore Station ===//
	case "ignore_table":
		return srv.getIgnoreTable(ctx)
	case "ignore_history":
		return srv.getIgnoreHistory(ctx)
	case "ignore":
		return srv.getIgnoreStationLoad(ctx)
	case "ignore_rainfall_detail":
		return srv.getIgnoreRainfallDetail(ctx)		
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePutData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "display_setting_json":
		return srv.putDisplaySetting(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePatchData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "ignore":
		return srv.patchIgnoreStation(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}