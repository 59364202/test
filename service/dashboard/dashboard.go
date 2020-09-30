package dashboard

import (
	model_dashboard "haii.or.th/api/thaiwater30/model/dashboard"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/dashboard"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

// RegisterService Register service in this package
func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPATCH, DataServiceName, srv.handlePatchData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "monitor":
		return srv.getDashboard(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handlePatchData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "monitor":
		return srv.patchOfflineDate(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) getDashboard(ctx service.RequestContext) error {

	result, err := model_dashboard.GetDashboard()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)

	return nil
}

func (srv *HttpService) patchOfflineDate(ctx service.RequestContext) error {

	type MetadataID struct {
		ID int64 `json:"metadata_id"`
	}

	p := &MetadataID{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	result, err := model_dashboard.UpdateMetadataOfflineDate(p.ID)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}
