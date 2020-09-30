package image

import (
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/shared/api_service"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {
	service_id := ctx.GetServiceParams("id")
	switch service_id {
	case "exp":
		return srv.getDownloadExp(ctx)

	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
func (srv *HttpService) getDownloadExp(ctx service.RequestContext) error {
	ctx.ReplyJSON(result.Result1(setting.GetSystemSettingInt("service.api-service.download.ExpireDate")))
	return nil
}
