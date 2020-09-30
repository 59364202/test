package agency

import (
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/agency"
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
	case "agency_shopping":
		return srv.getAgencyShopping(ctx)
	case "agency_shopping_detail":
		return srv.getAgencyShoppingDetail(ctx)
	case "agecncy_metadata":
		return srv.getMetaData(ctx)
	case "agency_summary":
		return srv.getAgencyMetadataSummary(ctx)
	case "agecncy_detail":
		return srv.getOrderResult(ctx)
	}

	return nil
}
