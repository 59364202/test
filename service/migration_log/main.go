package migration_log

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/migration_log"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlerPutData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "summary_data":
		return srv.getSummaryData(ctx)
	case "summary_master_data":
		return srv.getSummaryMasterData(ctx)
	case "data_by_table":
		return srv.getDataByTable(ctx)
	case "summary_image":
		return srv.getSummaryImage(ctx)
	case "image_by_media":
		return srv.getImgByMedia(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlerPutData(ctx service.RequestContext) error {
	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "summary_data":
		return srv.getRegenTable(ctx)
	case "summary_master_data":
		return srv.getRegenMasterTable(ctx)
	case "summary_image":
		return srv.getRegenTableImg(ctx)
	case "test1":
		return srv.getRegenTableImg1(ctx)
	case "test2":
		return srv.getRegenTableImg2(ctx)
	case "test3":
		return srv.getRegenTableImg3(ctx)	
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
