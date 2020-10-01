package data_service

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/data_service"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPOST, DataServiceName, srv.handlePostData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "onload":
		return srv.getOnloadShopping(ctx)
	case "data_service":
		return srv.getShoppingTable(ctx)
	case "data_service_check":
		return srv.getCheckVerifyShopping(ctx)
	case "history":
		return srv.getShoppingHistory(ctx)
	case "popular_order_purpose":
		return srv.getDatapurpose(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
func (srv *HttpService) handlePostData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "data_service":
		return srv.postShopping(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
