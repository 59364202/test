package rdl

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

type HttpService struct {
}

func HandlePostData(ctx service.RequestContext) error {
	srv := &HttpService{}
	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "ps":
		return srv.postPs(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func HandleDeleteData(ctx service.RequestContext) error {
	srv := &HttpService{}
	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "ps":
		return srv.deletePs(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
