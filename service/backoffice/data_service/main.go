package data_service

import (
	//	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/data_service"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlerPutData)
	dpt.Register(ServiceVersion, service.MethodDELETE, DataServiceName, srv.handlerDeleteData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "to_agency":
		return srv.getToAgency(ctx)
	case "upload_result":
		return srv.getUploadResult(ctx)
	case "approve":
		return srv.getApprove(ctx)
	case "management":
		return srv.getManagement(ctx)
	case "managementInit":
		return srv.getManagementInit(ctx)
	case "summaryInit":
		return srv.getSummaryInit(ctx)
	case "summary":
		return srv.getSummary(ctx)
	case "print":
		return srv.getPrint(ctx)
	case "reupload":
		return srv.reUpload(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlerPutData(ctx service.RequestContext) error {
	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "to_agency":
		return srv.putToAgency(ctx)
	case "upload_result":
		return srv.putUploadResult(ctx)
	case "approve":
		return srv.putApprove(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
func (srv *HttpService) handlerDeleteData(ctx service.RequestContext) error {
	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "management":
		return srv.deleteManagement(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
