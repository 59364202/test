package api

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/api"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

// RegisterService Register service in this package
func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPOST, DataServiceName, srv.handlePostData)
	dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlePutData)
	dpt.Register(ServiceVersion, service.MethodDELETE, DataServiceName, srv.handleDeleteData)
	dpt.Register(ServiceVersion, service.MethodPATCH, DataServiceName, srv.handlePatchData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "key_access":
		return srv.combineServiceKeyAccess(ctx)
	case "access_log":
		return srv.accessLog(ctx)
	case "service_name":
		return srv.serviceName(ctx)
	case "monitor_api_service":
		return srv.monitorApiService(ctx)
	case "monitor_api_service_onload":
		return srv.monitorApiServiceOnload(ctx)
	case "agent":
		return srv.agentName(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handlePostData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "key_access":
		return srv.newKeyAccess(ctx)
	case "agency":
		return srv.editAgency(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handlePutData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "key_access":
		return srv.editKeyAccess(ctx)
	case "gen_key":
		return srv.genKey(ctx)

	case "del_key":
		return srv.delKey(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "key_access":
		return srv.deletedKeyAccess(ctx)
	case "monitor_api_service":
		return srv.deleteMonitorApiService(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePatchData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "monitor_api_service":
		return srv.patchMonitorApiService(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
